import * as React from 'react';
import { Container, Button } from 'semantic-ui-react'
import axios from 'axios'

const submit = (e) => {
  axios.post('http://localhost:8080/users/sign_in', null, {headers: {'X-API': 'v1'}})
    .then((response) => { console.log(response) })
    .catch((err) => {console.log(err)})
}

export const App = (props) => {
  return (
    <Container>
      <Button onClick={submit} primary content="Call da server"/>
      { DevTool(props) }
    </Container>
  )
}

const DevTool = (props) => {
  if (process.env.NODE_ENV !== 'production') {
    const DevTools = require('mobx-react-devtools').default;
    return (<DevTools />)
  }
}
