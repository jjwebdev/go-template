import React from 'react'
import {
	Router,
	Route,
	IndexRoute,
	hashHistory
} from 'react-router'

import { Header, Icon, Image } from 'semantic-ui-react'
import {observer} from 'mobx-react'

import DevTools from 'mobx-react-devtools'

const Routes = (props) => {
	const { todoList } = props
	todoList.todos.map( todo => {
		console.log(todo.finished)
	})

	return (
		<div>
			<Router history={hashHistory}>
				<Route
					path="/"
					component={
						function Test() {
							return (
								<div>
									<Header as='h2' icon textAlign='center'>
										<Icon name='users' circular />
										<Header.Content>
											Friends
										</Header.Content>
									</Header>
									<Image centered size='large' src='http://semantic-ui.com/images/wireframe/centered-paragraph.png' />
								</div>
							)}
					}
				/>
			</Router>
			<DevTools />
		</div>
	)
}

export default observer(Routes)
