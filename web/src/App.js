import React, { Component } from 'react'
import { observer } from 'mobx-react'
import DevTools from 'mobx-react-devtools'
import { Link } from 'react-router-dom'
import { Layout, Menu, Breadcrumb, Button } from 'antd'
const { Header, Content, Footer } = Layout
import 'antd/dist/antd.min.css'
import 'tachyons/css/tachyons.min.css'
import axios from 'axios'

const submit = () => {
  console.log('posting')
  axios
    .post('/users/create', {"role_id": 1}, { headers: {'X-API': 'v1'}})
    .then((resp) => { console.log(resp.data)})
    .catch((err) => { console.log(err) })
}

const App = (props) => {
  return (
  <Layout className='min-vh-100'>
    <Header>
      <Menu
        theme='dark'
        mode='horizontal'
        defaultSelectedKeys={['2']}
        style={{ lineHeight: '64px' }}
      >
        <Menu.Item key='Home'>Home</Menu.Item>
        <Link to='google'>Google</Link>
        <Link to='about'>About</Link>
      </Menu>
    </Header>
    <Content className='pa4'>
        <Button onClick={submit} className='' type='primary'>Click me</Button>
    </Content>
    <Footer className='tc'>
      Go Tempate ©2017 Created by John Nguyen
    </Footer>
  </Layout>
)}

export default observer(App)
