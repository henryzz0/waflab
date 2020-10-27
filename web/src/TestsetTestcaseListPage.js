import React from "react";
import {Button, Col, Popconfirm, Row, Table, Tag} from 'antd';
import {SendOutlined, StopOutlined, CaretRightOutlined} from '@ant-design/icons';
import * as Setting from "./Setting";
import * as TestsetBackend from "./backend/TestsetBackend";
import * as TestcaseBackend from "./backend/TestcaseBackend";

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

  runTestcase(testcase) {
    Setting.showMessage("success", "Running: " + testcase.name);
  }

  renderTable(testcases) {
    const columns = [
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
        width: '100px',
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
        title: 'Action',
        dataIndex: '',
        key: 'op',
        width: '100px',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{marginTop: '10px', marginBottom: '10px', marginRight: '10px'}} type="primary" onClick={() => this.runTestcase(record)}>Run</Button>
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
