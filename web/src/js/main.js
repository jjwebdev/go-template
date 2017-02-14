import '../semantic/dist/semantic.css'
import '../semantic/dist/semantic'

import React from 'react'
import { render } from 'react-dom'
import Routes from './routes'

import { observable, computed } from 'mobx'

class Todo {
		id = Math.random()
		@observable title = ''
		@observable finished = false
}

class TodoList {
		@observable todos = [];
		@computed get unfinishedTodoCount() {
				return this.todos.filter(todo => !todo.finished).length
		}
}

const store = new TodoList()

store.todos.push(
		new Todo('Get Coffee'),
		new Todo('Write simpler code')
)

render(
	<Routes todoList={store}/>,
	document.getElementById('root')
)
