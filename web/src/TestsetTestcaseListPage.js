// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import React from "react";
import { Button, Col, Progress, Row, Switch, Table, Tag } from 'antd';
import * as Setting from "./Setting";
import * as TestsetBackend from "./backend/TestsetBackend";
import * as TestcaseBackend from "./backend/TestcaseBackend";
import * as ResultBackend from "./backend/ResultBackend";

class TestsetTestcaseListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testsetName: props.match.params.testsetName,
      testset: null,
      testcases: [],
    };
  }

  componentDidMount() {
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
    TestcaseBackend.getFilteredTestcases(this.state.testsetName)
      .then((res) => {
        this.setState({
          testcases: res,
        });
      });
  }

  setTestcaseValue(i, key, value) {
    let testcases = this.state.testcases;
    testcases[i][key] = value;
    this.setState({
      testcases: testcases,
    });
  }

  getResult(testcase, i) {
    this.setTestcaseValue(i, "state", "ongoing");
    ResultBackend.getResult(this.state.testset.name, testcase.name)
      .then((result) => {
        // Setting.showMessage("success", "Result: " + result.status);
        this.setTestcaseValue(i, "state", "finished");
        this.setTestcaseValue(i, "trueStatuses", result.statuses);
        this.setTestcaseValue(i, "response", result.response);
      })
      .catch(error => {
        Setting.showMessage("error", `failed to run: ${error}`);
        this.setTestcaseValue(i, "state", "error");
      });
  }

  getResults() {
    this.state.testcases.forEach((testcase, i) => {
      this.getResult(testcase, i);
    });
  }

  renderTable(testcases) {
    const columns = [
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
        width: '150px',
        sorter: (a, b) => a.name.localeCompare(b.name),
        render: (text, record, index) => {
          return (
            <a href={`/testcases/${text}`}>{text}</a>
          )
        }
      },
      // {
      //   title: 'Description',
      //   dataIndex: 'desc',
      //   key: 'desc',
      //   width: '250px',
      //   sorter: (a, b) => a.desc.localeCompare(b.desc),
      // },
      {
        title: 'Enabled',
        dataIndex: 'enabled',
        key: 'enabled',
        width: '80px',
        render: (text, record, index) => {
          return (
            <Switch disabled checked={text} />
          )
        }
      },
      {
        title: 'Test Count',
        dataIndex: 'testCount',
        key: 'testCount',
        width: '120px',
        sorter: (a, b) => a.testCount - b.testCount,
      },
      // {
      //   title: 'Created Time',
      //   dataIndex: 'createdTime',
      //   key: 'createdTime',
      //   width: '160px',
      //   sorter: (a, b) => a.createdTime.localeCompare(b.createdTime),
      //   render: (text, record, index) => {
      //     return Setting.getFormattedDate(text);
      //   }
      // },
      {
        title: 'Method',
        dataIndex: 'method',
        key: 'method',
        width: '100px',
        sorter: (a, b) => a.method.localeCompare(b.method),
        render: (text, record, index) => {
          return Setting.getMethodTag(text);
        }
      },
      // {
      //   title: 'User-Agent',
      //   dataIndex: 'userAgent',
      //   key: 'userAgent',
      //   // width: '100px',
      //   sorter: (a, b) => a.userAgent.localeCompare(b.userAgent),
      // },
      {
        title: 'Expected Status',
        dataIndex: 'statusLists',
        key: 'statusLists',
        width: '600px',
        // ellipsis: true,
        // sorter: (a, b) => a.statusLists - b.statusLists,
        render: (text, record, index) => {
          return Setting.getStatusTags(text);
        }
      },
      {
        title: 'True Status',
        dataIndex: 'trueStatuses',
        key: 'trueStatuses',
        width: '600px',
        // ellipsis: true,
        // sorter: (a, b) => a.trueStatuses - b.trueStatuses,
        render: (text, record, index) => {
          return Setting.getStatusTags(text);
        }
      },

      {
        title: 'Response',
        dataIndex: 'response',
        key: 'response',
        width: '100px',
        sorter: (a, b) => a.response.localeCompare(b.response),
        render: (text, record, index) => {
          if (record.trueStatus > 0 && text === "") {
            return "(Empty)";
          } else {
            return text;
          }
        }
      },
      {
        title: 'Progress',
        key: 'progress',
        width: '100px',
        // sorter: (a, b) => a.userAgent.localeCompare(b.userAgent),
        render: (text, record, index) => {
          if (record.state === undefined) {
            return (
              <Progress percent={0} size="small" />
            )
          } else if (record.state === "ongoing") {
            return (
              <Progress percent={50} size="small" />
            )
          } else if (record.state === "error") {
            return (
              <Progress percent={100} size="small" status="exception" />
            )
          } else {
            return (
              <Progress percent={100} size="small" />
            )
          }
        }
      },
      {
        title: 'Action',
        dataIndex: '',
        key: 'op',
        width: '100px',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{ marginTop: '10px', marginBottom: '10px', marginRight: '10px' }}
                loading={record.state === "ongoing"} type="primary" onClick={() => this.getResult(record, index)}>Run</Button>
            </div>
          )
        }
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={testcases} rowKey="name" size="middle" bordered pagination={{ pageSize: 1000 }}
          title={() => (
            <div>
              <Tag color="#108ee9">{this.state.testset === null ? "" : this.state.testset.name}</Tag> Testcases&nbsp;&nbsp;&nbsp;&nbsp;
              <Button type="primary" size="small" onClick={this.getResults.bind(this)}>Run All</Button>
            </div>
          )}
        />
      </div>
    );
  }

  render() {
    return (
      <div>
        <Row style={{ width: "100%" }}>
          <Col span={1}>
          </Col>
          <Col span={22}>
            {
              this.renderTable(this.state.testcases)
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
      </div>
    );
  }
}

export default TestsetTestcaseListPage;
