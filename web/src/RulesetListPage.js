// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

import React from "react";
import { Button, Col, Row, Table } from "antd";
import * as Setting from "./Setting";

class RulesetListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      rulesets: [],
    };
  }

  componentDidMount() {
    this.listRulesets();
  }

  listRulesets() {
    fetch(`${Setting.ServerUrl}/api/list-rulesets`, {
      method: "GET",
      credentials: "include"
    })
      .then(res => res.json())
      .then((res) => {
        this.setState({
          rulesets: res,
        });
      }
      );
  }

  onClick(link) {
    // this.props.history.push(link);
    const w = window.open('about:blank');
    w.location.href = link;
  }

  renderTable(rulesets) {
    const columns = [
      {
        title: 'Id',
        dataIndex: 'id',
        key: 'id',
      },
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
      },
      {
        title: 'Version',
        dataIndex: 'version',
        key: 'version',
      },
      {
        title: 'File Count',
        dataIndex: 'fileCount',
        key: 'fileCount',
      },
      {
        title: 'Rule Count',
        dataIndex: 'ruleCount',
        key: 'ruleCount',
      },
      {
        title: 'Action',
        key: 'action',
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{ marginTop: '10px', marginBottom: '10px', marginRight: '10px' }} type="primary" onClick={() => Setting.goToLink(`/rulesets/${record.id}/rulefiles/`)}>View</Button>
            </div>
          )
        }
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={rulesets} size="middle" bordered pagination={{ pageSize: 100 }}
          title={() => 'Rulesets'}
          loading={rulesets === null}
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
              this.renderTable(this.state.rulesets)
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
      </div>
    );
  }
}

export default RulesetListPage;
