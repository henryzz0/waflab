import React from "react";
import {Button, Col, Popconfirm, Row, Table} from 'antd';
import moment from "moment";
import * as Setting from "./Setting";
import * as TestsetBackend from "./backend/TestsetBackend";

class TestsetListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testsets: null,
    };
  }

  componentWillMount() {
    this.getTestsets();
  }

  getTestsets() {
    TestsetBackend.getTestsets()
      .then((res) => {
        this.setState({
          testsets: res,
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

  renderTable(testsets) {
    const columns = [
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
        width: '120px',
        sorter: (a, b) => a.name.localeCompare(b.name),
        render: (text, record, index) => {
          return (
            <a href={`/testsets/${text}`}>{text}</a>
          )
        }
      },
      {
        title: 'Title',
        dataIndex: 'title',
        key: 'title',
        // width: '80px',
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
        title: 'Action',
        dataIndex: '',
        key: 'op',
        width: '160px',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{marginTop: '10px', marginBottom: '10px', marginRight: '10px'}} type="primary" onClick={() => Setting.goToLink(`/testsets/${record.name}`)}>Edit</Button>
              <Popconfirm
                title={`Sure to delete testset: ${record.name} ?`}
                onConfirm={() => this.deleteTestset(index)}
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
        <Table columns={columns} dataSource={testsets} rowKey="name" size="middle" bordered pagination={{pageSize: 100}}
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
        <Row style={{width: "100%"}}>
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
