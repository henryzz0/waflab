// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
import React from "react";
import {DownOutlined, DeleteOutlined, UpOutlined} from '@ant-design/icons';
import {Button, Col, Input, Popconfirm, Row, Table, Tooltip} from 'antd';
import * as Setting from "./Setting";

class TestcaseHeaderTable extends React.Component {
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
    // value = this.parseField(key, value);

    let table = this.props.table;
    table[index][key] = value;
    this.updateTable(table);
  }

  addRow() {
    let table = this.props.table;
    let row = {id: 1, key: "", value: ""};

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

  renderTable(template, table) {
    const columns = [
      {
        title: 'Key',
        dataIndex: 'key',
        key: 'key',
        width: '150px',
        render: (text, record, index) => {
          return (
            <Input value={text} onChange={e => {
              this.updateField(index, "key", e.target.value);
            }} />
          )
        }
      },
      {
        title: 'Value',
        dataIndex: 'value',
        value: 'value',
        // width: '150px',
        render: (text, record, index) => {
          return (
            <Input value={text} onChange={e => {
              this.updateField(index, "value", e.target.value);
            }} />
          )
        }
      },
      {
        title: 'Action',
        key: 'action',
        width: '120px',
        render: (text, record, index) => {
          return (
            <div>
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
      <Table columns={columns} dataSource={table} rowKey="id" size="middle" bordered pagination={{pageSize: 100}}
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
              this.renderTable(this.props.template, this.props.table)
            }
          </Col>
        </Row>
      </div>
    )
  }
}

export default TestcaseHeaderTable;
