import React from 'react';
import ReactDOM from 'react-dom';
import App from './App';

import {
  BrowserRouter as Router,
  StaticRouter, // for server rendering
  Route,
  Link
  // etc.
} from 'react-router-dom'

const Routes = () => {
  return (<Router>
    <div>
      <Route exact path="/" component={App}/>
      <Route path="/about" component={About}/>
      <Route path="/topics" component={App}/>
    </div>
  </Router>)
}

const About = () => {
  return (
    <div>
      Heyo
    </div>
  )
}

ReactDOM.render(
  <Routes/>,
  document.getElementById('root')
);
