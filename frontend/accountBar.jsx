/* @flow */

import React from 'react';
import Drawer from 'material-ui/Drawer';
import MenuItem from 'material-ui/MenuItem';
import Divider from 'material-ui/Divider';
import List from 'material-ui/List/List';
import ListItem from 'material-ui/List/ListItem';
import Avatar from 'material-ui/Avatar';

import {
deepOrange300,
purple500,
} from 'material-ui/styles/colors';

const style = {
  width: '200px',
	zDepth: 5
};

type Props = {
	isOpen: boolean,
	switchPane: (pane: PaneType) => void,
	onClose: () => void
}
type State = {}
export default class AccountBar extends React.Component {
	state : State
	render() {
		const {switchPane,isOpen,onClose} = this.props;
		return (
			<Drawer open={isOpen}>
				<List>
					<MenuItem onTouchTap={() => this._dostuff()}>Login</MenuItem>
				</List>
				<Divider/>
				<MenuItem onTouchTap={onClose}>Close</MenuItem>
			</Drawer>
		);
	}
	_dostuff() {
		alert("stuff")
	}
}
