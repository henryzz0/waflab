import React from "react";
import {DownOutlined, EditOutlined, DeleteOutlined, UpOutlined} from '@ant-design/icons';
import {Button, Col, Row, Select, Table, Tooltip} from 'antd';
import * as Setting from "./Setting";

const { Option } = Select;

class TestsetEditTestcaseTable extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
    };
  }

  updateTable(table) {
    this.props.onUpdateTable(table);
  }

  updateField(index, key, value) {
    let table = this.props.table;
    table[index][key] = value;
    this.updateTable(table);
  }

  addRow() {
    let table = this.props.table;
    let row = {name: "(Please select)"};
    if (table === undefined) {
      table = [];
    }
    if (table.length > 0) {
      const last = table.slice(-1)[0];
      row = Setting.deepCopy(last);
      row.id = last.id + 1;
    }
    table = Setting.addRow(table, row);
    this.updateTable(table);
  }

  deleteRow(i) {
    let table = this.props.table;
    table = Setting.deleteRow(table, i);
    this.updateTable(table);
  }

  upRow(i) {
    let table = this.props.table;
    table = Setting.swapRow(table, i - 1, i);
    this.updateTable(table);
  }

  downRow(i) {
    let table = this.props.table;
    table = Setting.swapRow(table, i, i + 1);
    this.updateTable(table);
  }

  isItemSelected(table, name) {
    for (let i = 0; i < table.length; i ++) {
      if (table[i].name === name) {
        return true;
      }
    }
    return false;
  }

  renderTable(table) {
    const columns = [
      {
        title: 'No.',
        dataIndex: 'id',
        key: 'id',
        width: '70px',
        render: (text, record, index) => {
          return index;
        }
      },
      {
        title: this.props.title,
        dataIndex: 'name',
        key: 'name',
        render: (text, record, index) => {
          return (
            <Select style={{width: '100%'}} value={text} onChange={value => {this.updateField(index, 'name', value);}}>
              {
                this.props.testcases?.filter((testcase) => !this.isItemSelected(table, testcase.name)).map((testcase, index) => <Option key={testcase.name} value={testcase.name}>{testcase.name}</Option>)
              }
            </Select>
          )
        }
      },
      {
        title: 'Action',
        key: 'action',
        width: '130px',
        render: (text, record, index) => {
          return (
            <div>
              <Tooltip placement="topLeft" title="Edit">
                <Button style={{marginRight: "5px"}} icon={<EditOutlined />} size="small" onClick={() => Setting.openLink(`/testcases/${record.name}`)} />
              </Tooltip>
              <Tooltip placement="bottomLeft" title="Up">
                <Button style={{marginRight: "5px"}} disabled={index === 0} icon={<UpOutlined />} size="small" onClick={() => this.upRow.bind(this)(index)} />
              </Tooltip>
              <Tooltip placement="topLeft" title="Down">
                <Button style={{marginRight: "5px"}} disabled={index === table.length - 1} icon={<DownOutlined />} size="small" onClick={() => this.downRow.bind(this)(index)} />
              </Tooltip>
              <Tooltip placement="topLeft" title="Delete">
                <Button icon={<DeleteOutlined />} size="small" onClick={() => this.deleteRow.bind(this)(index)} />
              </Tooltip>
            </div>
          );
        }
      },
    ];

    return (
      <Table columns={columns} dataSource={table} rowKey="name" size="middle" bordered pagination={{pageSize: 100}}
             title={() => (
               <div>
                 {this.props.title}&nbsp;&nbsp;&nbsp;&nbsp;
                 <Button type="primary" size="small" onClick={this.addRow.bind(this)}>Add</Button>
               </div>
             )}
      />
    );
  }

  render() {
    return (
      <div>
        <Row style={{marginTop: '20px'}} >
          <Col span={24}>
            {
              this.renderTable(this.props.table)
            }
          </Col>
        </Row>
      </div>
    )
  }
}

export default TestsetEditTestcaseTable;
