import React from "react";
import {Table, Tag, Typography} from "antd";
import * as Setting from "./Setting";

const {Text} = Typography;

class RulefilePage extends React.Component {
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

  renderTable(title, rulefiles) {
    const columns = [
      {
        title: 'Id',
        dataIndex: 'id',
        key: 'id',
      },
      {
        title: 'Type',
        dataIndex: 'type',
        key: 'type',
      },
      {
        title: 'No',
        dataIndex: 'no',
        key: 'no',
      },
      {
        title: 'Suffix',
        dataIndex: 'suffix',
        key: 'suffix',
      },
      {
        title: 'Rule Count',
        dataIndex: 'count',
        key: 'count',
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

export default RulefilePage;
