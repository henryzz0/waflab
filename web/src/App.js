import React, {Component} from 'react';
import {Switch, Route, withRouter} from 'react-router-dom';
import './App.css';
import {BackTop, Layout, Menu, Typography} from 'antd';
import * as Setting from "./Setting";
import RulesetListPage from "./RulesetListPage";
import RulefileListPage from "./RulefileListPage";
import RuleListPage from "./RuleListPage";
import TestsetListPage from "./TestsetListPage";
import TestsetEditPage from "./TestsetEditPage";
import TestcaseListPage from "./TestcaseListPage";
import TestcaseEditPage from "./TestcaseEditPage";
import TestsetTestcaseListPage from "./TestsetTestcaseListPage";

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

  getUrlPath() {
    // eslint-disable-next-line no-restricted-globals
    return location.pathname;
  }

  updateMenuKey() {
    const uri = this.getUrlPath();
    if (uri.includes('/rules/')) {
      this.setState({ selectedMenuKey: 2 });
    } else if (uri.includes('rulefiles')) {
      this.setState({ selectedMenuKey: 1 });
    } else if (uri.includes('testsets')) {
      this.setState({selectedMenuKey: 10});
    } else if (uri.includes('testcases')) {
      this.setState({selectedMenuKey: 11});
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
          Rule Sets
        </a>
      </Menu.Item>
    );
    if (this.getUrlPath().includes('rulefiles')) {
      res.push(
        <Menu.Item key="1">
          <a href="#">
            Rule Files
          </a>
        </Menu.Item>
      );
    }
    if (this.getUrlPath().includes('/rules/')) {
      res.push(
        <Menu.Item key="2">
          <a href="#">
            Rules
          </a>
        </Menu.Item>
      );
    }
    res.push(
      <Menu.Item key="10">
        <a href="/testsets">
          Testsets
        </a>
      </Menu.Item>
    );
    res.push(
      <Menu.Item key="11">
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
          <Route exact path="/" component={RulesetListPage}/>
          <Route exact path="/rulesets/:rulesetId/rulefiles/" component={RulefileListPage}/>
          <Route exact path="/rulesets/:rulesetId/rulefiles/:rulefileId/rules/" component={RuleListPage}/>
          <Route exact path="/testsets/" component={TestsetListPage}/>
          <Route exact path="/testsets/:testsetName" component={TestsetEditPage}/>
          <Route exact path="/testsets/:testsetName/testcases/" component={TestsetTestcaseListPage}/>
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
