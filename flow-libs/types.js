/* @flow */

type PaneType = 'home' | 'item' | 'settings' | 'admin';

type Success = boolean

var isUserAdmin: boolean
var displayName: string

type Post = {
	key?: string,
	title: string,
	body: string,
	timestamp?: number,
	frontpage: bool,
	UserID?: string
}

type KeyObject = {
	key: string
}
