/* @flow */
import React from 'react';
import Paper from 'material-ui/Paper';
import RaisedButton from 'material-ui/RaisedButton';
import {Card, CardActions, CardHeader, CardTitle, CardText} from 'material-ui/Card';
import CircularProgress from 'material-ui/CircularProgress';
import TextField from 'material-ui/TextField';
import Toggle from 'material-ui/Toggle';

import {getRandom} from './testApi.js';
import {createPost} from './postApi.js';

const paperStyle = {
	flex: 1,
	display: 'flex',
	margin: '40px',
	zDepth: 5
}

type State = {
	title: string,
	body: string,
	frontpage: boolean,
}

export default class NewPostPane extends React.Component {
	state: State = {title: '',body:'',frontpage: false};
	render() {
		const {title,body,frontpage} = this.state;
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
									onChange={(event) => this._onTitleChange(event)}
								/>
								<TextField
								 hintText="Body"
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
							<RaisedButton primary={true} style={{margin:'5px'}} label="Submit" onTouchTap={() => this._submit()}/>
						</div>
					</Card>
				</Paper>
			</div>
		);
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

	async _submit() {
		const {title, body,frontpage} = this.state;
		const post = {
			title,
			body,
			frontpage,
		}
		const resp = await createPost(post);
	}
}
