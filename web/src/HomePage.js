import React from "react";
import {Button, Table, Tag, Typography} from "antd";
import * as Setting from "./Setting";
import {Link} from "react-router-dom";

const {Text} = Typography;

class HomePage extends React.Component {
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
        dataIndex: 'count',
        key: 'count',
      },
      {
        title: 'Action',
        dataIndex: '',
        key: 'action',
        render: (text, record, index) => <Button type="primary" onClick={() => this.onClick.bind(this)(`/ruleset/${record.id}`)}>View</Button>
      },
    ];

    return (
      <div>
        <Table columns={columns} dataSource={rulesets} size="small" bordered pagination={{pageSize: 100}} scroll={{y: 'calc(95vh - 450px)'}}
               title={() => 'Rulesets'} />
      </div>
    );
  }

  render() {
    return (
      <div>
        {
          this.renderTable(this.state.rulesets)
        }
      </div>
    );
  }

}

export default HomePage;
