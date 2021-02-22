// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import React from "react";
import { Button, Card, Col, Input, Row } from 'antd';
import { LinkOutlined } from "@ant-design/icons";
import * as TestsetBackend from "./backend/TestsetBackend";
import * as TestcaseBackend from "./backend/TestcaseBackend";
import * as Setting from "./Setting";
import TestsetEditTestcaseTable from "./TestsetEditTestcaseTable";

class TestsetEditPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testsetName: props.match.params.testsetName,
      testset: null,
      testcases: [],
    };
  }

  componentWillMount() {
    this.getTestset();
    this.getTestcases();
  }

  getTestset() {
    TestsetBackend.getTestset(this.state.testsetName)
      .then((testset) => {
        this.setState({
          testset: testset,
        });
      });
  }

  getTestcases() {
    TestcaseBackend.getTestcases()
      .then((res) => {
        this.setState({
          testcases: res,
        });
      });
  }

  parseTestsetField(key, value) {
    // if ([].includes(key)) {
    //   value = Setting.myParseInt(value);
    // }
    return value;
  }

  updateTestsetField(key, value) {
    value = this.parseTestsetField(key, value);

    let testset = this.state.testset;
    testset[key] = value;
    this.setState({
      testset: testset,
    });
  }

  onUpdateTestsetField(key, value) {
    this.updateTestsetField(key, value);
  }

  renderTestset() {
    return (
      <Card size="small" title={
        <div>
          Edit Testset&nbsp;&nbsp;&nbsp;&nbsp;
          <Button type="primary" onClick={this.submitTestsetEdit.bind(this)}>Save</Button>
        </div>
      } style={{ marginLeft: '5px' }} type="inner">
        <Row style={{ marginTop: '10px' }} >
          <Col style={{ marginTop: '5px' }} span={2}>
            Name:
          </Col>
          <Col span={22} >
            <Input value={this.state.testset.name} onChange={e => {
              this.updateTestsetField('name', e.target.value);
            }} />
          </Col>
        </Row>
        <Row style={{ marginTop: '20px' }} >
          <Col style={{ marginTop: '5px' }} span={2}>
            Description:
          </Col>
          <Col span={22} >
            <Input value={this.state.testset.desc} onChange={e => {
              this.updateTestsetField('desc', e.target.value);
            }} />
          </Col>
        </Row>
        <Row style={{ marginTop: '20px' }} >
          <Col style={{ marginTop: '5px' }} span={2}>
            Target Url:
          </Col>
          <Col span={22} >
            <Input prefix={<LinkOutlined />} value={this.state.testset.targetUrl} onChange={e => {
              this.updateTestsetField('targetUrl', e.target.value);
            }} />
          </Col>
        </Row>
        <Row style={{ marginTop: '20px' }} >
          <Col style={{ marginTop: '5px' }} span={2}>
            Testcases:
          </Col>
          <Col span={14} >
            <TestsetEditTestcaseTable
              title="Testcases"
              table={this.state.testset.testcases}
              testcases={this.state.testcases}
              onUpdateTable={(value) => { return this.onUpdateTestsetField("testcases", value) }}
            />
          </Col>
        </Row>
      </Card>
    )
  }

  submitTestsetEdit() {
    let testset = Setting.deepCopy(this.state.testset);
    TestsetBackend.updateTestset(this.state.testsetName, testset)
      .then((res) => {
        if (res) {
          Setting.showMessage("success", `Successfully saved`);
          this.setState({
            testsetName: this.state.testset.name,
          });
          this.props.history.push(`/testsets/${this.state.testset.name}`);
        } else {
          Setting.showMessage("error", `failed to save: server side failure`);
          this.updateTestsetField('name', this.state.testsetName);
        }
      })
      .catch(error => {
        Setting.showMessage("error", `failed to save: ${error}`);
      });
  }

  render() {
    return (
      <div>
        <Row style={{ width: "100%" }}>
          <Col span={1}>
          </Col>
          <Col span={22}>
            {
              this.state.testset !== null ? this.renderTestset() : null
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
        <Row style={{ margin: 10 }}>
          <Col span={2}>
          </Col>
          <Col span={18}>
            <Button type="primary" size="large" onClick={this.submitTestsetEdit.bind(this)}>Save</Button>
          </Col>
        </Row>
      </div>
    );
  }
}

export default TestsetEditPage;
