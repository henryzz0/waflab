import React from "react";
import {Table, Tag, Typography} from "antd";
import * as Setting from "./Setting";

const {Text} = Typography;

class HomePage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      classes: props,
      testsets: [],
    };
  }

  componentDidMount() {
    this.listTestSets();
  }

  listTestSets() {
    fetch(`${Setting.ServerUrl}/api/list-testsets`, {
      method: "GET",
      credentials: "include"
    })
      .then(res => res.json())
      .then((res) => {
          this.setState({
            testsets: res,
          });
        }
      );
  }

  renderTable(testsets) {
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
    ];

    return (
      <div>
        <Table columns={columns} dataSource={testsets} size="small" bordered pagination={{pageSize: 100}} scroll={{y: 'calc(95vh - 450px)'}}
               title={() => 'Testsets'} />
      </div>
    );
  }

  render() {
    return (
      <div>
        {
          this.renderTable(this.state.testsets)
        }
      </div>
    );
  }

}

export default HomePage;
