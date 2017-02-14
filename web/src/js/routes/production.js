import React from 'react'
import {
	Router,
	Route,
	IndexRoute,
	browserHistory
} from 'react-router'

const routes = (
	<div>
		<Router history={browserHistory}>
			<Route
				path="/"
				component={
					function Test() {
						return <div>hello world</div>
					}
				}
			/>
		</Router>
	</div>
)

export default routes
