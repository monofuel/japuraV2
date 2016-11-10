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

export default class AboutPane extends React.Component {
	state: State = {randomNumber: 0};
	render() {
		const {randomNumber} = this.state;
		return (
			<div style={{flex: 1, display: 'flex'}}>
				<Paper style={paperStyle}>
					<Card style={{flex: 1}}>
						<CardHeader
							title="About Japura.net"
							avatar={<i className="fa fa-comment-o" aria-hidden="true"></i>}
							/>
						<CardText>
							<a href="http://steamcommunity.com/groups/Japura</p>">Join our steam group!</a>

			        <p>Our proud donor wall!</p>

			        <p>diamond tier donors:
			          PolarBearWP,
			          Zyluxic,
			          Asynthic
			        </p>

			        <p>regular tier donors:
			        wentam,
			        lonewanderer666,
			        rob0423,
			        djan,
			        stitches_14,
			        videogamezombie,
			        force2reckon,
			        tinyt,
			        cantaroji,
			        jimagine,
			        123abe2,
			        dgjjl98u76t5,
			        darkone93,
			        jonathan_c7,
			        elpwner,
			        tmotom,
			        smgpenguin,
			        gemhunter76,
			        apexpredator29,
			        tomatoheadache,
			        infernalboy13,
			        reptilians13,
			        nayr,
			        stupidpandabear,
			        danzmanz,
			        phyc,
			        staticgrazer,
			        treeman892,
			        goodcaster1,
			        kibarob,
			        spartankiller76,
			        Zinkey246,
			        wildfox,
			        bluebee02,
			        firestrike33,
			        spriteslam,
			        thekiller419,
			        plescia1,
			        applecube16,
			        (Super) kooljoe,
			        ChunkNinja,
			        TheSoulCamel,
			        MonkeePilot,
			        The_8th_Doctor,
			        <em>Jazperz</em>,
			        freakmaster2012,
			        RPatt98,
			        kaiden_patt,
			        patt_cole,
			        matthew_szawaluk,
			        DerekJ64,
			        blackholetnt,
			        Jubbah,
			        soccerstriker,
			        purplecreeper92,
			        soccer_rex,
			        WAAT,
			        DarkAzure,
			        Elyceellen</p>

			        <p>The best way to help support japura (aside from donating) and to help youself is to vote for the server on all the server lists!</p>

			        <ul style={{listStyle: 'none'}}>
								<li><a href='https://mcserverstatus.com/viewserver/16533'><img src='http://mcserverstatus.com/16533_1.png'></img></a></li>
				        <li><a href="http://www.minecraftserverfinder.com/2297-japura-survival-city"><img src="http://www.minecraftserverfinder.com/rank.php?id=2297" /></a></li>
				        <li><a href="http://minecraft-server-list.com/server/41077/"><img src="http://minecraft-server-list.com/server/logo/41077.png" alt="Minecraft Server Japura Survival City"  /></a></li>
				        <li><a href="http://minecraft-mp.com/server-s108081" target="_blank"><img src="http://minecraft-mp.com/banner-108081.png" border="0"/></a> (Very cool site, has a handy ping map to see latency around the globe)</li>
				        <li><a href="http://topg.org/Minecraft/in-423225">TopG Minecraft</a></li>
				        <li><a href="http://www.serverpact.com/vote-30228">serverpact.com</a> (very poor site. password cleartext offender, displays database errors, forms reset. so bad i had to mention it here)</li>
				        <li><a href="http://minecraft-server.net/details/japura/">minecraft-server.net</a></li>
				        <li><a href="http://mineservers.com/server/2230/view/japura-net">mineservers.com</a></li>
				        <li><a href="https://minestatus.net/22991-japura-survival-city">minestatus.net</a></li>
				        <li><a href="https://minecraftlist.org/server/561">minecraftlist.org</a></li>
				        <li><a href="http://minecraftservers.org/server/25872">minecraftservers.org</a></li></ul>
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
