// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import React from "react";
import { Button, Col, List, Popconfirm, Row, Table, Tooltip } from 'antd';
import { EditOutlined } from "@ant-design/icons";
import moment from "moment";
import * as Setting from "./Setting";
import * as TestsetBackend from "./backend/TestsetBackend";
import * as TestcaseBackend from "./backend/TestcaseBackend";

class TestsetListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testsets: null,
      testcaseMap: {},
    };
  }

  componentWillMount() {
    this.getTestsets();
    this.getTestcases();
  }

  getTestsets() {
    TestsetBackend.getTestsets()
      .then((res) => {
        this.setState({
          testsets: res,
        });
      });
  }

  getTestcases() {
    TestcaseBackend.getTestcases()
      .then((testcases) => {
        let testcaseMap = {};

        testcases.forEach((testcase, i) => {
          testcaseMap[testcase.name] = testcase;
        });

        this.setState({
          testcaseMap: testcaseMap,
        });
      });
  }

  newTestset() {
    return {
      name: `testset_${this.state.testsets.length}`,
      createdTime: moment().format(),
      title: `New Testset - ${this.state.testsets.length}`,
      targetUrl: "http://localhost:9000/api/test",
      testcases: [],
    }
  }

  addTestset() {
    const newTestset = this.newTestset();
    TestsetBackend.addTestset(newTestset)
      .then((res) => {
        Setting.showMessage("success", `Testset added successfully`);
        this.setState({
          testsets: Setting.prependRow(this.state.testsets, newTestset),
        });
      }
      )
      .catch(error => {
        Setting.showMessage("error", `Testset failed to add: ${error}`);
      });
  }

  deleteTestset(i) {
    TestsetBackend.deleteTestset(this.state.testsets[i])
      .then((res) => {
        Setting.showMessage("success", `Testset deleted successfully`);
        this.setState({
          testsets: Setting.deleteRow(this.state.testsets, i),
        });
      }
      )
      .catch(error => {
        Setting.showMessage("error", `Testset failed to delete: ${error}`);
      });
  }

  renderTestcaseLink(record, i) {
    const testcaseName = record.testcases[i];
    return <a target="_blank" href={`/testcases/${testcaseName}`}>{`${i}. ${testcaseName}`}</a>
  }

  renderTable(testsets) {
    const columns = [
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
        width: '140px',
        sorter: (a, b) => a.name.localeCompare(b.name),
        render: (text, record, index) => {
          return (
            <a href={`/testsets/${text}`}>{text}</a>
          )
        }
      },
      // {
      //   title: 'Description',
      //   dataIndex: 'desc',
      //   key: 'desc',
      //   width: '220px',
      //   sorter: (a, b) => a.desc.localeCompare(b.desc),
      // },
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
        title: 'Baseline Url',
        dataIndex: 'baselineUrl',
        key: 'baselineUrl',
        width: '250px',
        ellipsis: true,
        sorter: (a, b) => a.baselineUrl.localeCompare(b.baselineUrl),
        render: (text, record, index) => {
          return (
            <a target="_blank" href={text}>{text}</a>
          )
        }
      },
      {
        title: 'Target Url',
        dataIndex: 'targetUrl',
        key: 'targetUrl',
        width: '250px',
        ellipsis: true,
        sorter: (a, b) => a.targetUrl.localeCompare(b.targetUrl),
        render: (text, record, index) => {
          return (
            <a target="_blank" href={text}>{text}</a>
          )
        }
      },
      {
        title: 'Testcases',
        dataIndex: 'testcases',
        key: 'testcases',
        width: '600px',
        render: (text, record, index) => {
          const testcases = text;
          if (testcases.length === 0) {
            return "(None)";
          }

          const half = Math.floor((testcases.length + 1) / 2);

          return (
            <div>
              <Row>
                <Col span={12}>
                  <List
                    size="small"
                    dataSource={testcases.slice(0, half)}
                    renderItem={(row, i) => {
                      return (
                        <List.Item>
                          <div style={{ display: "inline" }}>
                            <Tooltip placement="topLeft" title="Edit">
                              <Button style={{ marginRight: "5px" }} icon={<EditOutlined />} size="small" onClick={() => Setting.openLink(`/testcases/${row}`)} />
                            </Tooltip>
                            {
                              this.renderTestcaseLink(record, i)
                            }
                          </div>
                        </List.Item>
                      )
                    }}
                  />
                </Col>
                <Col span={12}>
                  <List
                    size="small"
                    dataSource={testcases.slice(half)}
                    renderItem={(row, i) => {
                      return (
                        <List.Item>
                          <div style={{ display: "inline" }}>
                            <Tooltip placement="topLeft" title="Edit">
                              <Button style={{ marginRight: "5px" }} icon={<EditOutlined />} size="small" onClick={() => Setting.openLink(`/testcases/${row}`)} />
                            </Tooltip>
                            {
                              this.renderTestcaseLink(record, i + half)
                            }
                          </div>
                        </List.Item>
                      )
                    }}
                  />
                </Col>
              </Row>
            </div>
          )
        },
      },
      {
        title: 'Action',
        dataIndex: '',
        key: 'op',
        width: '240px',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{ marginTop: '10px', marginBottom: '10px', marginRight: '10px' }} onClick={() => Setting.goToLink(`/testsets/${record.name}/testcases`)}>Run</Button>
              <Button style={{ marginTop: '10px', marginBottom: '10px', marginRight: '10px' }} type="primary" onClick={() => Setting.goToLink(`/testsets/${record.name}`)}>Edit</Button>
              <Popconfirm
                title={`Sure to delete testset: ${record.name} ?`}
                onConfirm={() => this.deleteTestset(index)}
              >
                <Button style={{ marginBottom: '10px' }} type="danger">Delete</Button>
              </Popconfirm>
            </div>
          )
        }
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={testsets} rowKey="name" size="middle" bordered pagination={{ pageSize: 1000 }}
          title={() => (
            <div>
              Testsets&nbsp;&nbsp;&nbsp;&nbsp;
              <Button type="primary" size="small" onClick={this.addTestset.bind(this)}>Add</Button>
            </div>
          )}
          loading={testsets === null}
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
              this.renderTable(this.state.testsets)
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
      </div>
    );
  }
}

export default TestsetListPage;
