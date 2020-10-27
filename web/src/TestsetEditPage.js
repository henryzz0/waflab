import React from "react";
import {Button, Card, Col, Input, Row} from 'antd';
import * as TestsetBackend from "./backend/TestsetBackend";
import * as Setting from "./Setting";

class TestsetEditPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testsetName: props.match.params.testsetName,
      testset: null,
      tasks: [],
      resources: [],
    };
  }

  componentWillMount() {
    this.getTestset();
  }

  getTestset() {
    TestsetBackend.getTestset(this.state.testsetName)
      .then((testset) => {
        this.setState({
          testset: testset,
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

  renderTestset() {
    return (
      <Card size="small" title={
        <div>
          Edit Testset&nbsp;&nbsp;&nbsp;&nbsp;
          <Button type="primary" onClick={this.submitTestsetEdit.bind(this)}>Save</Button>
        </div>
      } style={{marginLeft: '5px'}} type="inner">
        <Row style={{marginTop: '10px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            Name:
          </Col>
          <Col span={22} >
            <Input value={this.state.testset.name} onChange={e => {
              this.updateTestsetField('name', e.target.value);
            }} />
          </Col>
        </Row>
        <Row style={{marginTop: '20px'}} >
          <Col style={{marginTop: '5px'}} span={2}>
            Title:
          </Col>
          <Col span={22} >
            <Input value={this.state.testset.title} onChange={e => {
              this.updateTestsetField('title', e.target.value);
            }} />
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
        <Row style={{width: "100%"}}>
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
        <Row style={{margin: 10}}>
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
