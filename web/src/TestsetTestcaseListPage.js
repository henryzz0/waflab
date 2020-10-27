import React from "react";
import {Button, Col, Progress, Row, Table, Tag} from 'antd';
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
    TestcaseBackend.getTestcases(this.state.testsetName)
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
    this.setTestcaseValue(i, "trueStatus", -1);
    ResultBackend.getResult(this.state.testset.name, testcase.name)
      .then((result) => {
        // Setting.showMessage("success", "Result: " + result.status);
        this.setTestcaseValue(i, "trueStatus", result.status);
        this.setTestcaseValue(i, "response", result.response);
      })
      .catch(error => {
        Setting.showMessage("error", `failed to run: ${error}`);
        this.setTestcaseValue(i, "trueStatus", -2);
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
      {
        title: 'Title',
        dataIndex: 'title',
        key: 'title',
        width: '150px',
        sorter: (a, b) => a.title.localeCompare(b.title),
      },
      {
        title: 'Created Time',
        dataIndex: 'createdTime',
        key: 'createdTime',
        width: '160px',
        sorter: (a, b) => a.createdTime.localeCompare(b.createdTime),
        render: (text, record, index) => {
          return Setting.getFormattedDate(text);
        }
      },
      {
        title: 'Method',
        dataIndex: 'method',
        key: 'method',
        width: '100px',
        sorter: (a, b) => a.method.localeCompare(b.method),
        render: (text, record, index) => {
          return Setting.getTags(text);
        }
      },
      {
        title: 'User-Agent',
        dataIndex: 'userAgent',
        key: 'userAgent',
        // width: '100px',
        sorter: (a, b) => a.userAgent.localeCompare(b.userAgent),
      },
      {
        title: 'Status',
        dataIndex: 'status',
        key: 'status',
        width: '80px',
        ellipsis: true,
        sorter: (a, b) => a.status - b.status,
        render: (text, record, index) => {
          return Setting.getStatusTag(text);
        }
      },
      {
        title: 'True Status',
        dataIndex: 'trueStatus',
        key: 'trueStatus',
        width: '120px',
        ellipsis: true,
        sorter: (a, b) => a.trueStatus - b.trueStatus,
        render: (text, record, index) => {
          return Setting.getStatusTag(text);
        }
      },

      {
        title: 'Response',
        dataIndex: 'response',
        key: 'response',
        width: '100px',
        sorter: (a, b) => a.response.localeCompare(b.response),
      },
      {
        title: 'Progress',
        key: 'progress',
        width: '100px',
        // sorter: (a, b) => a.userAgent.localeCompare(b.userAgent),
        render: (text, record, index) => {
          if (record.trueStatus === 0) {
            return (
              <Progress percent={0} size="small" />
            )
          } else if (record.trueStatus === -1) {
            return (
              <Progress percent={50} size="small" />
            )
          } else if (record.trueStatus === -2) {
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
              <Button style={{marginTop: '10px', marginBottom: '10px', marginRight: '10px'}}
                      loading={record.trueStatus === -1} type="primary" onClick={() => this.getResult(record, index)}>Run</Button>
            </div>
          )
        }
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={testcases} rowKey="name" size="middle" bordered pagination={{pageSize: 100}}
               title={() => (
                 <div>
                   Testcases for: <Tag color="#108ee9">{this.state.testset === null ? "" : this.state.testset.name}</Tag>
                 </div>
               )}
        />
      </div>
    );
  }

  render() {
    return (
      <div>
        <Row style={{width: "100%"}}>
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
