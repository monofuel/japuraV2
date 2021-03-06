/* @flow */
import React from 'react';
import Paper from 'material-ui/Paper';
import {Card, CardActions, CardHeader, CardTitle, CardText} from 'material-ui/Card';
import CircularProgress from 'material-ui/CircularProgress';

import {getRandom} from './testApi.js';

const paperStyle = {
	flex: 1,
	display: 'flex',
	margin: '40px',
	zDepth: 5
}

type State = {
	randomNumber: number,
}

export default class HomePane extends React.Component {
	state: State = {randomNumber: 0};
	render() {
		const {randomNumber} = this.state;
		return (
			<div style={{flex: 1, display: 'flex'}}>
				<Paper style={paperStyle}>
					<Card style={{flex: 1}}>
						<CardHeader
							title="Dynamic Map"
							avatar={<i className="fa fa-comment-o" aria-hidden="true"></i>}
							/>
						<CardText>
							<a href="https://japura.net:8123/">Fullscreen</a>
							<iframe src="https://japura.net:8123" style={{width:'100%', height:700}} id="dynframe"></iframe>
						</CardText>
					</Card>
				</Paper>
			</div>
		);
	}

	componentDidMount() {
		this._getRandomNumber();
	}
	async _getRandomNumber() {
		const x = await getRandom();
		this.setState({randomNumber: x});
	}
}
