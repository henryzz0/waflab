import React from "react";
import {Col, Row, Table, Tag, Typography} from "antd";
import {MinusSquareOutlined, PlusSquareOutlined} from "@ant-design/icons";
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
          title: 'Type',
          dataIndex: 'type',
          key: 'type',
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
        title: 'Type',
        dataIndex: 'type',
        key: 'type',
        width: 80,
      },
      {
        title: 'Paranoia Level',
        dataIndex: 'paranoiaLevel',
        key: 'paranoiaLevel',
        width: 80,
      },
      {
        title: 'Test Count',
        dataIndex: 'testCount',
        key: 'testCount',
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
          {expanded ? <MinusSquareOutlined /> : <PlusSquareOutlined />}
        </a>
      );
    }

    const plColors = ["pl1", "pl2", "pl3", "pl4"];

    return (
      <div>
        <Table columns={columns} dataSource={rules} size="middle" bordered pagination={{pageSize: 100}}
               expandIcon={expandIcon} expandedRowRender={expandedRowRender}
               title={() => <div><Text>Rules for: </Text><Tag color="#108ee9">{title}</Tag> => <Tag color="#108ee9">{title2}</Tag></div>}
               rowClassName={(record, index) => { return plColors[record.paranoiaLevel - 1] }}
               loading={rules === null}
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
              this.state.rulefile !== null ? this.renderTable(this.state.rulesetId, this.state.rulefileId, this.state.rulefile.rules) : null
            }
          </Col>
          <Col span={1}>
          </Col>
        </Row>
      </div>
    );
  }
}

export default RulefilePage;
