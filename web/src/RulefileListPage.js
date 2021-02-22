// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
import React from "react";
import {Button, Col, Row, Table, Tag, Typography} from "antd";
import * as Setting from "./Setting";

const {Text} = Typography;

class RulefileListPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      rulesetId: props.match.params.rulesetId,
      ruleset: null,
    };
  }

  componentDidMount() {
    this.listRulefiles();
  }

  listRulefiles() {
    fetch(`${Setting.ServerUrl}/api/list-rulefiles?rulesetId=${this.state.rulesetId}`, {
      method: "GET",
      credentials: "include"
    })
      .then(res => res.json())
      .then((res) => {
          this.setState({
            ruleset: res,
          });
        }
      );
  }

  onClick(link) {
    // this.props.history.push(link);
    const w = window.open('about:blank');
    w.location.href = link;
  }

  renderTable(title, rulefiles) {
    const columns = [
      {
        title: 'No',
        dataIndex: 'no',
        key: 'no',
        width: 60,
      },
      {
        title: 'Id',
        dataIndex: 'id',
        key: 'id',
        width: 400,
      },
      {
        title: 'Type',
        dataIndex: 'type',
        key: 'type',
        width: 100,
      },
      {
        title: 'Name',
        dataIndex: 'name',
        key: 'name',
        width: 100,
      },
      {
        title: 'Description',
        dataIndex: 'desc',
        key: 'desc',
        width: 300,
      },
      {
        title: 'Rule Count',
        dataIndex: 'count',
        key: 'count',
        width: 220,
        render: (text, record, index) => {
          return `${record.count} (${record.pl1Count} + ${record.pl2Count} + ${record.pl3Count} + ${record.pl4Count})`;
        }
      },
      {
        title: 'Test Count',
        dataIndex: 'testCount',
        key: 'testCount',
        width: 220,
        render: (text, record, index) => {
          return `${record.testCount} (${record.pl1TestCount} + ${record.pl2TestCount} + ${record.pl3TestCount} + ${record.pl4TestCount})`;
        }
      },
      {
        title: 'Action',
        key: 'action',
        width: 100,
        render: (text, record, index) => {
          return (
            <div>
              <Button style={{marginTop: '10px', marginBottom: '10px', marginRight: '10px'}} type="primary" onClick={() => Setting.goToLink(`/rulesets/${this.state.rulesetId}/rulefiles/${record.id}/rules/`)}>View</Button>
            </div>
          )
        }
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={rulefiles} size="middle" bordered pagination={{pageSize: 100}}
               title={() => <div><Text>Rulefiles for: </Text><Tag color="#108ee9">{title}</Tag></div>}
               loading={rulefiles === null}
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
              this.state.ruleset !== null ? this.renderTable(this.state.rulesetId, this.state.ruleset.rulefiles) : null
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
      </div>
    );
  }
}

export default RulefileListPage;
