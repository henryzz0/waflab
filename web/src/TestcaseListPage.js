import React from "react";
import {Button, Col, Popconfirm, Row, Switch, Table} from 'antd';
import moment from "moment";
import * as Setting from "./Setting";
import * as TestcaseBackend from "./backend/TestcaseBackend";

class TestcaseListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testcases: null,
    };
  }

  componentWillMount() {
    this.getTestcases();
  }

  getTestcases() {
    TestcaseBackend.getTestcases()
      .then((res) => {
        this.setState({
          testcases: res,
        });
      });
  }

  newTestcase() {
    return {
      name: `testcase_${this.state.testcases.length}`,
      createdTime: moment().format(),
      title: `New Testcase - ${this.state.testcases.length}`,
      method: "GET",
      userAgent: navigator.userAgent,
      queryStrings: [],
      status: 200,
    }
  }

  addTestcase() {
    const newTestcase = this.newTestcase();
    TestcaseBackend.addTestcase(newTestcase)
      .then((res) => {
          Setting.showMessage("success", `Testcase added successfully`);
          this.setState({
            testcases: Setting.prependRow(this.state.testcases, newTestcase),
          });
        }
      )
      .catch(error => {
        Setting.showMessage("error", `Testcase failed to add: ${error}`);
      });
  }

  deleteTestcase(i) {
    TestcaseBackend.deleteTestcase(this.state.testcases[i])
      .then((res) => {
          Setting.showMessage("success", `Testcase deleted successfully`);
          this.setState({
            testcases: Setting.deleteRow(this.state.testcases, i),
          });
        }
      )
      .catch(error => {
        Setting.showMessage("error", `Testcase failed to delete: ${error}`);
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
        title: 'Description',
        dataIndex: 'desc',
        key: 'desc',
        width: '250px',
        sorter: (a, b) => a.desc.localeCompare(b.desc),
      },
      {
        title: 'Author',
        dataIndex: 'author',
        key: 'author',
        width: '120px',
        sorter: (a, b) => a.author.localeCompare(b.author),
      },
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
        width: '160px',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{marginTop: '10px', marginBottom: '10px', marginRight: '10px'}} type="primary" onClick={() => Setting.goToLink(`/testcases/${record.name}`)}>Edit</Button>
              <Popconfirm
                title={`Sure to delete testcase: ${record.name} ?`}
                onConfirm={() => this.deleteTestcase(index)}
              >
                <Button style={{marginBottom: '10px'}} type="danger">Delete</Button>
              </Popconfirm>
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
                   Testcases&nbsp;&nbsp;&nbsp;&nbsp;
                   <Button type="primary" size="small" onClick={this.addTestcase.bind(this)}>Add</Button>
                 </div>
               )}
               loading={testcases === null}
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

export default TestcaseListPage;
