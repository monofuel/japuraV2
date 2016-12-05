/* @flow */
import React from 'react';
import _ from 'lodash';
import Paper from 'material-ui/Paper';
import {Card, CardActions, CardHeader, CardTitle, CardText} from 'material-ui/Card';
import IconMenu from 'material-ui/IconMenu';
import MenuItem from 'material-ui/MenuItem';
import IconButton from 'material-ui/IconButton';
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert';

import PostEditor from './postEditor.jsx';

type Props = {
	post: Post,
}

type State = {
	editing: boolean
}

const paperStyle = {
	flex: 1,
	display: 'flex',
	margin: '40px',
	zDepth: 5
}

export default class PostPaper extends React.Component {
	props: Props;
	state: State = {editing: false}
	render() {
		const {post} = this.props;
		const {editing} = this.state;

		const editMenu = (
				<IconMenu
					iconButtonElement={<IconButton><MoreVertIcon /></IconButton>}
					anchorOrigin={{horizontal: 'left', vertical: 'top'}}
					targetOrigin={{horizontal: 'left', vertical: 'top'}}
				>
					<MenuItem primaryText="Edit" onTouchTap={() => this._editPost()}/>
					<MenuItem primaryText="Delete" onTouchTap={() => this._deletePost()}/>
				</IconMenu>
		)

		if (editing) {
				return (
					<PostEditor existingPost={post} closeCallback={() => {this.setState({editing: false})}}/>
				)
		}
		return (
			<Paper style={paperStyle}>
				<Card style={{flex: 1}}>
					<CardHeader
						title={post.title}
						avatar={<i className="fa fa-comment-o" aria-hidden="true"></i>}
						/>
					<CardText style={{display: 'flex'}}>
						<div style={{flex:1}} dangerouslySetInnerHTML={{__html:post.body}}></div>
						{isUserAdmin ? editMenu : null}
					</CardText>
				</Card>
			</Paper>
		)
	}
	_editPost() {
		console.log("editing")
		this.setState({editing: true})

	}
	_deletePost() {
		console.log("deleting")
	}
}
