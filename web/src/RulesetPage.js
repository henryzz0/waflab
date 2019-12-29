import React from "react";
import {Button, Table, Tag, Typography} from "antd";
import * as Setting from "./Setting";

const {Text} = Typography;

class RulesetPage extends React.Component {
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
      },
      {
        title: 'Rule Count',
        dataIndex: 'count',
        key: 'count',
      },
      {
        title: 'Action',
        dataIndex: '',
        key: 'action',
        render: (text, record, index) => <Button type="primary" onClick={() => this.onClick.bind(this)(`/ruleset/${this.state.rulesetId}/rulefile/${record.id}`)}>View</Button>
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={rulefiles} size="small" bordered pagination={{pageSize: 100}} scroll={{y: 'calc(95vh - 170px)'}}
               title={() => <div><Text>Rulefiles for: </Text><Tag color="#108ee9">{title}</Tag></div>} />
      </div>
    );
  }

  render() {
    return (
      <div>
        {
          this.state.ruleset !== null ? this.renderTable(this.state.rulesetId, this.state.ruleset.rulefiles) : null
        }
      </div>
    );
  }

}

export default RulesetPage;
