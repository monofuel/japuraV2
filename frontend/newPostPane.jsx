/* @flow */
import React from 'react';
import Paper from 'material-ui/Paper';
import Button from 'material-ui/FlatButton';
import {Card, CardActions, CardHeader, CardTitle, CardText} from 'material-ui/Card';
import CircularProgress from 'material-ui/CircularProgress';
import TextField from 'material-ui/TextField';

import {getRandom} from './testApi.js';

const paperStyle = {
	flex: 1,
	display: 'flex',
	margin: '40px',
	zDepth: 5
}

type State = {
	title: string,
	body: string,
}

export default class NewPostPane extends React.Component {
	state: State = {title: '',body:''};
	render() {
		const {title,body} = this.state;
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
							<br></br>
							<p style={{margin:'5px'}}>{'authoring as ' + displayName}</p>
							<Button primary={true} style={{margin:'5px'}}>Submit</Button>
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
}
