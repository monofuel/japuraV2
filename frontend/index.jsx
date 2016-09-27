/* @flow */

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import React from 'react';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import AppBar from 'material-ui/AppBar';
import Paper from 'material-ui/Paper';

import NavBar from './navBar.jsx';
import AccountBar from './accountBar.jsx'
import HomePane from './homePane.jsx';

const style ={
	display: 'flex',
	height: '100%',
	width: '100%',
	flexDirection: 'column'
}
const paperStyle = {
	display: 'flex',
	flex:1,
	zDepth: 0
}

type State = {
	pane: PaneType,
	sideBarOpen: boolean
}
export default class Index extends React.Component {
	state : State = {
		pane: 'home',
		sideBarOpen: false
	}
	render() {
		const {pane,sideBarOpen} = this.state;
		let paneElement = <div/>
		switch (pane) {
			case 'home':
				paneElement = <HomePane/>;
				break;
			default:
				paneElement = (<h1>Not implemented!</h1>)
		}

		return (
			// muiTheme={getMuiTheme(darkBaseTheme)}
			<MuiThemeProvider>
				<div id="indexRoot" style={style}>
					<AccountBar isOpen={sideBarOpen} switchPane={(pane: PaneType) => this._paneSwitch(pane)} onClose={() => this._toggleSideBar()}/>
					<AppBar title="webapp template" onLeftIconButtonTouchTap={() => this._toggleSideBar()}/>
					<Paper style={paperStyle}>
						<NavBar switchPane={(pane: PaneType) => this._paneSwitch(pane)}/>
						{paneElement}
					</Paper>
				</div>
			</MuiThemeProvider>
		);
	}

	_toggleSideBar() {
		const {sideBarOpen} = this.state;
		this.setState({sideBarOpen: !sideBarOpen});
	}

	_paneSwitch(pane: PaneType) {
		console.log('switching panes',pane);
		this.setState({pane});
	}
}
