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
      rulefileId: props.match.params.rulefileId,
      rulefile: null,
    };
  }

  componentDidMount() {
    this.listRules();
  }

  listRules() {
    fetch(`${Setting.ServerUrl}/api/list-rules?rulesetId=${this.state.rulesetId}&rulefileId=${this.state.rulefileId}`, {
      method: "GET",
      credentials: "include"
    })
      .then(res => res.json())
      .then((res) => {
          this.setState({
            rulefile: res,
          });
        }
      );
  }

  renderTable(title, title2, rules) {
    const columns = [
      {
        title: 'No',
        dataIndex: 'no',
        key: 'no',
        width: 50,
      },
      {
        title: 'Text',
        dataIndex: 'text',
        key: 'text',
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={rules} size="small" bordered pagination={{pageSize: 100}} scroll={{y: 'calc(95vh - 170px)'}}
               title={() => <div><Text>Rules for: </Text><Tag color="#108ee9">{title}</Tag> => <Tag color="#108ee9">{title2}</Tag></div>} />
      </div>
    );
  }

  render() {
    return (
      <div>
        {
          this.state.rulefile !== null ? this.renderTable(this.state.rulesetId, this.state.rulefileId, this.state.rulefile.rules) : null
        }
      </div>
    );
  }

}

export default RulefilePage;
