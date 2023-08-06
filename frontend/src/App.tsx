import { useEffect, useState } from "react";
import {
  SettingOutlined,
  CameraOutlined,
  ProfileOutlined,
} from "@ant-design/icons";
import "./App.css";
import { Layout, Menu, Button } from "antd";
import Snap from "./page/snap/snap";
const { Header, Content } = Layout;
const menu = [
  {
    label: "snap",
    key: "snap",
    icon: <CameraOutlined />,
  },
  {
    label: "profile",
    key: "profile",
    icon: <ProfileOutlined />,
  },
  {
    label: "settings",
    key: "settings",
    icon: <SettingOutlined />,
  },
];

const App = () => {
  const [current, setCurrent] = useState("snap");
  const onClick = (e: any) => {
    console.log("click ", e);
    setCurrent(e.key);
  };
  const renderContent = () => {
    switch (current) {
      case "snap":
        return <Snap />;
      case "profile":
        return <div>PROFILE</div>;
      case "settings":
        return <div>SETTINGS</div>;
      default:
        return <div>SNAP</div>;
    }
  };
  return (
    <Layout className="layout">
      <Menu
        onClick={onClick}
        selectedKeys={[current]}
        mode="horizontal"
        items={menu}
      ></Menu>
      <Content className="content">{renderContent()}</Content>
    </Layout>
  );
};
export default App;
