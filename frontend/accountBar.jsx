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
		let list;
		if (displayName) {
			if (isUserAdmin) {
				list = (
					<List>
						<MenuItem onTouchTap={() => {switchPane('newPost');onClose()}}>New Post</MenuItem>
						<MenuItem onTouchTap={() => {switchPane('editPost');onClose()}}>Edit Post</MenuItem>
					</List>
				)
			} else {
				list = (
					<List>
						<MenuItem onTouchTap={() => {switchPane('logout');onClose()}}>Log Out</MenuItem>
					</List>
				)
			}
		} else {
			list = (
				<List>
					<MenuItem onTouchTap={() => {switchPane('login');onClose()}}>Login</MenuItem>
				</List>
			)
		}
		return (
			<Drawer open={isOpen}>
				{list}
				<Divider/>
				<MenuItem onTouchTap={onClose}>Close</MenuItem>
			</Drawer>
		);
	}
	_dostuff() {
		alert("stuff")
	}
}
