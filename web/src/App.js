import React, {Component} from 'react';
import {Switch, Route, withRouter} from 'react-router-dom';
import './App.css';
import {BackTop, Layout, Menu, Typography} from 'antd';
import * as Setting from "./Setting";
import HomePage from "./HomePage";
import RulesetPage from "./RulesetPage";
import RulefilePage from "./RulefilePage";
import TestcaseListPage from "./TestcaseListPage";
import TestcaseEditPage from "./TestcaseEditPage";

const { Header, Footer } = Layout;
const { Text } = Typography;

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      selectedMenuKey: 0,
    };

    Setting.initServerUrl();
  }

  componentWillMount() {
    this.updateMenuKey();
  }

  updateMenuKey() {
    // eslint-disable-next-line no-restricted-globals
    const uri = location.pathname;
    if (uri.includes('ruleset')) {
      this.setState({ selectedMenuKey: 1 });
    } else if (uri.includes('rulefile')) {
      this.setState({ selectedMenuKey: 2 });
    } else if (uri.includes('testcases')) {
      this.setState({selectedMenuKey: 3});
    } else {
      this.setState({ selectedMenuKey: 0 });
    }
  }

  renderAccount() {
    return null;
  }

  renderMenu() {
    let res = [];

    res.push(
      <Menu.Item key="0">
        <a href="/">
          Home
        </a>
      </Menu.Item>
    );
    res.push(
      <Menu.Item key="1">
        <a href="/ruleset">
          Ruleset
        </a>
      </Menu.Item>
    );
    res.push(
      <Menu.Item key="2">
        <a href="/rulefile">
          Rulefile
        </a>
      </Menu.Item>
    );
    res.push(
      <Menu.Item key="3">
        <a href="/testcases">
          Testcases
        </a>
      </Menu.Item>
    );

    return res;
  }

  renderContent() {
    return (
      <div>
        <Header style={{ padding: '0', marginBottom: '2px'}}>
          <a className="logo" href={"/"} />
          {/*<Text style={{marginRight: '30px'}}>WAF Lab</Text>*/}
          <Menu
            // theme="dark"
            mode="horizontal"
            defaultSelectedKeys={[`${this.state.selectedMenuKey}`]}
            style={{ lineHeight: '64px' }}
          >
            {
              this.renderMenu()
            }
            {
              this.renderAccount()
            }
          </Menu>
        </Header>
        <Switch>
          <Route exact path="/" component={HomePage}/>
          <Route path="/ruleset/:rulesetId/rulefile/:rulefileId" component={RulefilePage}/>
          <Route path="/ruleset/:rulesetId" component={RulesetPage}/>
          <Route exact path="/testcases/" component={TestcaseListPage}/>
          <Route exact path="/testcases/:testcaseName" component={TestcaseEditPage}/>
        </Switch>
      </div>
    )
  }

  renderFooter() {
    // How to keep your footer where it belongs ?
    // https://www.freecodecamp.org/neyarnws/how-to-keep-your-footer-where-it-belongs-59c6aa05c59c/

    return (
      <Footer id="footer" style={
        {
          borderTop: '1px solid #e8e8e8',
          backgroundColor: 'white',
          textAlign: 'center',
        }
      }>
        Made with <span style={{color: 'rgb(255, 255, 255)'}}>❤️</span> by <a style={{fontWeight: "bold", color: "black"}} target="_blank" href="https://microsoftapc.sharepoint.com/teams/BotDetection">WAFLab</a>
      </Footer>
    )
  }

  render() {
    return (
      <div id="parent-area">
        <BackTop />
        <div id="content-wrap">
          {
            this.renderContent()
          }
        </div>
        {
          this.renderFooter()
        }
      </div>
    );
  }
}

export default withRouter(App);
