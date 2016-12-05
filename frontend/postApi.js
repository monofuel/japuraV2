/* @flow */

import {makeRequest} from './apiHandler.js';

export async function getLatestPosts(page?: number, limit?: number): Promise<Array<Post>> {
	return  makeRequest('/v1/posts/','GET',{page,limit});
}

export async function createPost(post:Post): Promise<Post> {
	return makeRequest('/v1/posts/','POST',null,JSON.stringify(post))
}

export async function updatePost(post:Post,key: string): Promise<Post> {
	return makeRequest('/v1/posts/'+ key,'PUT',post)
}

export async function deletePost(key: string): Promise<void> {
	return makeRequest('/v1/posts/'+ key,'DELETE')
}
