/* @flow */

type PaneType = 'home' | 'item' | 'settings' | 'admin';

type Success = boolean

var isUserAdmin: boolean
var displayName: string

type Post = {
	key: String,
	title: String,
	body: String,
	timestamp: number,
	frontpage: bool,
	UserID: string
}

type KeyObject = {
	key: String
}
