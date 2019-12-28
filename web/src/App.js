import React, {Component} from 'react';
import './App.css';
import * as Setting from "./Setting";
import {Layout, Menu, Typography} from 'antd';

import {Switch, Route} from 'react-router-dom'
import HomePage from "./HomePage";

const { Header, Footer } = Layout;
const { Text } = Typography;

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      selectedMenuKey: 1,
    };
  }

  componentWillMount() {
    // eslint-disable-next-line no-restricted-globals
    const uri = location.pathname;
    if (uri.includes('page1')) {
      this.setState({ selectedMenuKey: 2 });
    } else if (uri.includes('pagw2')) {
      this.setState({ selectedMenuKey: 3 });
    } else {
      this.setState({ selectedMenuKey: 1 });
    }
  }

  render() {
    Setting.initServerUrl();

    return (
      <div className="layout">
        <Header style={{ padding: '0', marginBottom: '8px'}}>
          <div className="logo" />
          <Menu
            // theme="dark"
            mode="horizontal"
            defaultSelectedKeys={[`${this.state.selectedMenuKey}`]}
            style={{ lineHeight: '64px' }}
            inlineCollapsed={false}
          >
            <Text>WAF Lab</Text>

            <Menu.Item key="1">
              <a href="/">
                Home
              </a>
            </Menu.Item>
            <Menu.Item key="2">
              <a href="/page1">
                Page1
              </a>
            </Menu.Item>
            <Menu.Item key="3">
              <a href="/page2">
                Page2
              </a>
            </Menu.Item>
          </Menu>
        </Header>
        <Switch>
          <Route exact path="/" component={HomePage}/>
          {/*<Route path="/task/" component={TaskPage}/>*/}
        </Switch>
        {/*<Footer style={{ textAlign: 'center' }}>WAF Lab</Footer>*/}
      </div>
    );
  }
}

export default App;
