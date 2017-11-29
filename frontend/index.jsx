/* @flow */

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider';
import React from 'react';
import {green300} from 'material-ui/styles/colors';
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme';
import getMuiTheme from 'material-ui/styles/getMuiTheme';
import AppBar from 'material-ui/AppBar';
import Paper from 'material-ui/Paper';

import NavBar from './navBar.jsx';
import AccountBar from './accountBar.jsx'
import HomePane from './homePane.jsx';
import DynmapPane from './dynmapPane.jsx';
import PostEditor from './postEditor.jsx';
import AboutPane from './aboutPane.jsx';

const style ={
	flex: 1,
	flexDirection: 'column',
	minHeight: '100%'
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
			case 'dynmap':
				paneElement = <DynmapPane/>;
				break;
				case 'login':
					window.location = '/login';
					break;
			case 'ffxivmc':
				window.location = 'https://ffxivmc-1361.appspot.com/';
				break;
			case 'badmars':
				window.location = 'https://badmars.net';
				break;
			case 'EVE SoggySquad':
				window.location = 'https://soggysquad.space';
				break;
			case 'subreddit':
				window.location = 'https://www.reddit.com/r/JapuraGaming/';
				break;
			case 'newPost':
				paneElement = <PostEditor/>;
				break;
			case 'about':
				paneElement = <AboutPane/>;
				break;
			default:
				paneElement = (<h1>Not implemented!</h1>)
		}

		const theme = getMuiTheme({
			palette: {
			primary1Color: green300,
			},
			appBar: {
				height: 50,
			},
		});

		return (
			// muiTheme={getMuiTheme(darkBaseTheme)}
			<MuiThemeProvider muiTheme={theme}>
				<div id="indexRoot" style={style}>
					<AccountBar isOpen={sideBarOpen} switchPane={(pane: PaneType) => this._paneSwitch(pane)} onClose={() => this._toggleSideBar()}/>
					<AppBar title="Japura Gaming" onLeftIconButtonTouchTap={() => this._toggleSideBar()}/>
					<div id="paneRoot" style={{display: 'flex',flex:1}}>
						<NavBar switchPane={(pane: PaneType) => this._paneSwitch(pane)}/>
						{paneElement}
					</div>
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
