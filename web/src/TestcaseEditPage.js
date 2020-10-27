import React from "react";
import {AutoComplete, Button, Card, Col, Input, Row, Select} from 'antd';
import * as TestcaseBackend from "./backend/TestcaseBackend";
import * as Setting from "./Setting";
import TestcaseHeaderTable from "./TestcaseHeaderTable";

const { Option } = AutoComplete;

class TestcaseEditPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testcaseName: props.match.params.testcaseName,
      testcase: null,
      tasks: [],
      resources: [],
    };
  }

  componentWillMount() {
    this.getTestcase();
  }

  getTestcase() {
    TestcaseBackend.getTestcase(this.state.testcaseName)
      .then((testcase) => {
        this.setState({
          testcase: testcase,
        });
      });
  }

  parseTestcaseField(key, value) {
    if (["status"].includes(key)) {
      value = Setting.myParseInt(value);
    }
    return value;
  }

  updateTestcaseField(key, value) {
    value = this.parseTestcaseField(key, value);

    let testcase = this.state.testcase;
    testcase[key] = value;
    this.setState({
      testcase: testcase,
    });
  }

  onUpdateTestcaseField(key, value) {
    this.updateTestcaseField(key, value);
  }

  renderTestcase() {
    return (
      <Card size="small" title={
        <div>
          Edit Testcase&nbsp;&nbsp;&nbsp;&nbsp;
          <Button type="primary" onClick={this.submitTestcaseEdit.bind(this)}>Save</Button>
        </div>
      } style={{marginLeft: '5px'}} type="inner">
        <Row style={{marginTop: '10px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            Name:
          </Col>
          <Col span={22} >
            <Input value={this.state.testcase.name} onChange={e => {
              this.updateTestcaseField('name', e.target.value);
            }} />
          </Col>
        </Row>
        <Row style={{marginTop: '20px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            Title:
          </Col>
          <Col span={22} >
            <Input value={this.state.testcase.title} onChange={e => {
              this.updateTestcaseField('title', e.target.value);
            }} />
          </Col>
        </Row>
        <Row style={{marginTop: '20px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            Method:
          </Col>
          <Col span={22} >
            <Select style={{width: '200px'}} value={this.state.testcase.method} onChange={(value => {this.updateTestcaseField('method', value);})}>
              {
                [
                  "GET",
                  "POST",
                  "PUT",
                  "DELETE",
                ].map((item, index) => <Option key={index} value={item}>{item}</Option>)
              }
            </Select>
          </Col>
        </Row>
        <Row style={{marginTop: '20px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            User-Agent:
          </Col>
          <Col span={22} >
            <Input value={this.state.testcase.userAgent} onChange={e => {
              this.updateTestcaseField('userAgent', e.target.value);
            }} />
          </Col>
        </Row>
        <Row style={{marginTop: '20px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            Query Strings:
          </Col>
          <Col span={22} >
            <TestcaseHeaderTable
              title="Query Strings"
              table={this.state.testcase.queryStrings}
              onUpdateTable={(value) => { return this.onUpdateTestcaseField("queryStrings", value)}}
            />
          </Col>
        </Row>
        <Row style={{marginTop: '20px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            Status:
          </Col>
          <Col span={22} >
            <Select style={{width: '200px'}} value={this.state.testcase.status} onChange={(value => {this.updateTestcaseField('status', value);})}>
              {
                [
                  200,
                  403,
                ].map((item, index) => <Option key={index} value={item}>{item}</Option>)
              }
            </Select>
          </Col>
        </Row>
      </Card>
    )
  }

  submitTestcaseEdit() {
    let testcase = Setting.deepCopy(this.state.testcase);
    TestcaseBackend.updateTestcase(this.state.testcaseName, testcase)
      .then((res) => {
        if (res) {
          Setting.showMessage("success", `Successfully saved`);
          this.setState({
            testcaseName: this.state.testcase.name,
          });
          this.props.history.push(`/testcases/${this.state.testcase.name}`);
        } else {
          Setting.showMessage("error", `failed to save: server side failure`);
          this.updateTestcaseField('name', this.state.testcaseName);
        }
      })
      .catch(error => {
        Setting.showMessage("error", `failed to save: ${error}`);
      });
  }

  render() {
    return (
      <div>
        <Row style={{width: "100%"}}>
          <Col span={1}>
          </Col>
          <Col span={22}>
            {
              this.state.testcase !== null ? this.renderTestcase() : null
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
        <Row style={{margin: 10}}>
          <Col span={2}>
          </Col>
          <Col span={18}>
            <Button type="primary" size="large" onClick={this.submitTestcaseEdit.bind(this)}>Save</Button>
          </Col>
        </Row>
      </div>
    );
  }
}

export default TestcaseEditPage;
