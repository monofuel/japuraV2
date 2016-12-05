/* @flow */
import React from 'react';

import Paper from 'material-ui/Paper';
import RaisedButton from 'material-ui/RaisedButton';
import {Card, CardActions, CardHeader, CardTitle, CardText} from 'material-ui/Card';
import TextField from 'material-ui/TextField';
import Toggle from 'material-ui/Toggle';

import {createPost,updatePost} from './postApi.js';

const paperStyle = {
	flex: 1,
	display: 'flex',
	margin: '40px',
	zDepth: 5
}


type Props = {
	existingPost?: Post,
	closeCallback?: () => void,
}
type State = {
	title: string,
	body: string,
	frontpage: boolean
}

export default class PostEditor extends React.Component {
	props: Props;
	state: State;
	constructor(props: Props) {
		super(props);
		const {existingPost} = props;
		if (existingPost) {
			console.log("setting default state",existingPost);
			this.state = {
				title: existingPost.title,
				body: existingPost.body,
				frontpage: existingPost.frontpage,
			}
		} else {
			this.state = {
				title: '',
				body: '',
				frontpage: false,
			}
		}
	}
	render() {
		const {existingPost} = this.props;
		const {title,body,frontpage} = this.state;
		const newPost = !existingPost;

		return (
			<div style={{flex: 1,  flexDirection:'column'}}>
				<Paper style={paperStyle}>
					<Card style={{flex: 1}}>
						<CardHeader
							title="Create New Post"
							avatar={<i className="fa fa-comment-o" aria-hidden="true"></i>}
							/>
						<CardText>
							<div>
								<TextField
									hintText="Post Title"
									defaultValue={title}
									onChange={(event) => this._onTitleChange(event)}
								/>
								<TextField
								 hintText="Body"
								 defaultValue={body}
								 onChange={(event) => this._onBodyChange(event)}
								 multiLine={true}
								 fullWidth={true}
							 	/>
					 		</div>
						</CardText>
					</Card>
				</Paper>
				<Paper style={paperStyle}>
					<Card style={{flex: 1}}>
						<CardHeader
							title={title}
							avatar={<i className="fa fa-comment-o" aria-hidden="true"></i>}
							/>
						<CardText>
							<div dangerouslySetInnerHTML={{__html:body}}></div>
						</CardText>
						<div>
							<br/>
							<p style={{margin:'5px'}}>{'authoring as ' + displayName}</p>
							<br/>
								<Toggle
									label="Front Page"
									onToggle={() => this._frontpageToggle()}
									toggled={frontpage}
								/>
							<br/>
							{
								newPost ?
								<RaisedButton primary={true} style={{margin:'5px'}} label="Submit" onTouchTap={() => this._newPost()}/>
								:
								<div>
									<RaisedButton primary={true} style={{margin:'5px'}} label="Update" onTouchTap={() => this._updatePost()}/>
									<RaisedButton secondary={true} style={{margin:'5px'}} label="Cancel" onTouchTap={() => this._close()}/>
								</div>
							}
						</div>
					</Card>
				</Paper>
			</div>
		)
	}

	_onTitleChange(event: Object) {
		const title = event.target.value;
		this.setState({title});
	}
	_onBodyChange(event: Object) {
		const body = event.target.value;
		this.setState({body});
	}

	_frontpageToggle() {
		const {frontpage} = this.state;
		this.setState({frontpage: !frontpage});
	}

	async _newPost() {
		const {closeCallback} = this.props;
		const {title, body,frontpage} = this.state;
		const post = {
			title,
			body,
			frontpage,
		}
		const resp = await createPost(post);

		if (closeCallback) {
			closeCallback();
		}
	}
	async _updatePost() {
		const {existingPost,closeCallback} = this.props;
		const {title, body,frontpage} = this.state;
		const post: Post = Object.assign({},existingPost);

		post.title = title;
		post.body = body;
		post.frontpage = frontpage;
		if (!post.key) {
			throw new Error("key missing from post");

		}

		const resp = await updatePost(post,post.key);

		if (closeCallback) {
			closeCallback();
		}
	}
	_close() {
		const {closeCallback} = this.props;
		if (!closeCallback) {
			console.log("close callback called when none set");
			return;
		}
		closeCallback();
	}
}
