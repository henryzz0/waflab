import React from "react";
import {Icon, Table, Tag, Typography} from "antd";
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
    const expandedRowRender = (record, index, indent, expanded) => {
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
          width: 80,
        },
        {
          title: 'Text',
          dataIndex: 'text',
          key: 'text',
        },
      ];

      return <Table columns={columns} dataSource={record.chainRules} pagination={false} />;
    };

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
        width: 80,
      },
      {
        title: 'Text',
        dataIndex: 'text',
        key: 'text',
      },
    ];

    function expandIcon({ expanded, expandable, record, onExpand }) {
      if (!expandable || record.chainRules === null) return null;

      return (
        <a onClick={e => onExpand(record, e)}>
          {expanded ? <Icon type="minus-square" /> : <Icon type="plus-square" />}
        </a>
      );
    }

    return (
      <div>
        <Table columns={columns} dataSource={rules} size="small" bordered pagination={{pageSize: 100}} scroll={{y: 'calc(95vh - 170px)'}}
               expandIcon={expandIcon} expandedRowRender={expandedRowRender} title={() => <div><Text>Rules for: </Text><Tag color="#108ee9">{title}</Tag> => <Tag color="#108ee9">{title2}</Tag></div>} />
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
