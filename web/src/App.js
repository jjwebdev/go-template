import React, { Component } from 'react';
import { observer } from 'mobx-react'
import DevTools from 'mobx-react-devtools';
import { Link } from 'react-router-dom'
import { Layout, Menu, Breadcrumb } from 'antd';
const { Header, Content, Footer } = Layout;
import 'antd/dist/antd.min.css'
import 'tachyons/css/tachyons.min.css'
const App = (props) => {
  return (
  <Layout className="min-vh-100">
    <Header>
      <Menu
        theme="dark"
        mode="horizontal"
        defaultSelectedKeys={['2']}
        style={{ lineHeight: '64px' }}
      >
        <Menu.Item key="Home">Home</Menu.Item>
        <Link to='google'>Google</Link>
        <Link to='about'>About</Link>
      </Menu>
    </Header>
    <Content className='pa4'>
      <div style={{ background: '#fff', padding: 24, minHeight: 280 }}>Content</div>
    </Content>
    <Footer className='tc'>
      Ant Design ©2016 Created by Ant UED
    </Footer>
  </Layout>
)
}

export default observer(App);
