import React, {Component} from 'react';
import './App.css';
import * as Setting from "./Setting";
import {Layout, Menu, Typography} from 'antd';

import {Switch, Route} from 'react-router-dom'
import HomePage from "./HomePage";
import RulesetPage from "./RulesetPage";
import RulefilePage from "./RulefilePage";

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
    if (uri.includes('rulefile')) {
      this.setState({ selectedMenuKey: 3 });
    } else if (uri.includes('ruleset')) {
      this.setState({ selectedMenuKey: 2 });
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
            <Text style={{marginRight: '30px'}}>WAF Lab</Text>

            <Menu.Item key="1">
              <a href="/">
                Home
              </a>
            </Menu.Item>
            <Menu.Item key="2">
              <a href="/ruleset">
                Ruleset
              </a>
            </Menu.Item>
            <Menu.Item key="3">
              <a href="/rulefile">
                Rulefile
              </a>
            </Menu.Item>
          </Menu>
        </Header>
        <Switch>
          <Route exact path="/" component={HomePage}/>
          <Route path="/ruleset/:rulesetId/rulefile/:rulefileId" component={RulefilePage}/>
          <Route path="/ruleset/:rulesetId" component={RulesetPage}/>
        </Switch>
        {/*<Footer style={{ textAlign: 'center' }}>WAF Lab</Footer>*/}
      </div>
    );
  }
}

export default App;
