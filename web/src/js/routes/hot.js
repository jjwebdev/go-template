import React from 'react'
import {
	Router,
	Route,
	IndexRoute,
	hashHistory
} from 'react-router'

import { Button } from 'semantic-ui-react'

import DevTools from 'mobx-react-devtools'

const routes = (
	<div>
		<Router history={hashHistory}>
			<Route
				path="/"
				component={
					function Test() {
						return (<div>
							<Button primary>
								Press here
							</Button>
							hello world
						</div>)
					}
				}
			/>
		</Router>
		<DevTools />
	</div>
)

export default routes
