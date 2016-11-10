/* @flow */
import React from 'react';
import Paper from 'material-ui/Paper';

export default class adsense extends React.Component {
	render() {
		return ( <ins className="adsbygoogle"
     style={{display:"block"}}
     data-ad-client="ca-pub-9414560385755205"
     data-ad-slot="9674151252"
     data-ad-format="auto"></ins>);
		}

		componentDidMount() {
			let adsbygoogle
			(adsbygoogle = window.adsbygoogle || []).push({});
		}
}
