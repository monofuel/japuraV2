/* @flow */
import React from 'react';
import _ from 'lodash';
import Paper from 'material-ui/Paper';
import {Card, CardActions, CardHeader, CardTitle, CardText} from 'material-ui/Card';
import CircularProgress from 'material-ui/CircularProgress';
import PostPaper from './postPaper.jsx';

import {getLatestPosts} from './postApi.js';

const paperStyle = {
	flex: 1,
	display: 'flex',
	margin: '40px',
	zDepth: 5
}

type State = {
	posts: Array<Post>,
}
/* TODO
show timestamp
show author user
*/

export default class HomePane extends React.Component {
	state: State = {posts:[]};
	render() {
		const {posts} = this.state;
		return (
			<div style={{flex: 1, flexDirection:'column'}}>
				{ posts && posts.length == 0 ?
					<Paper style={paperStyle} key='placeholder'>
						<Card style={{flex:1}}>
							<CardHeader
								title='Welcome'
								avatar={<i className='fa fa-comment-o' aria-hidden='true'></i>}
								/>
							<CardText>
								<p>
									Loading content
								</p>
							</CardText>
						</Card>
					</Paper>
					:
					_.map(posts,(post) => {
						return (
							<PostPaper
								post={post}
								key={post.key}
								refreshCallback={() => {
									console.log("refreshing posts after update");
									this._getFrontpagePosts();
								}}
							/>
						)
					})
				}

			</div>
		);
	}

	componentDidMount() {
		this._getFrontpagePosts();
	}
	async _getFrontpagePosts() {
		const posts = await getLatestPosts(0,10);
		this.setState({posts: posts});
	}
}
