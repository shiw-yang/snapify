import { Button, Space } from "antd";
import style from "./snap.less";

const Snap = () => {
  const onClick = () => {
    console.log("click");
  };
  return (
    <Space direction="vertical">
      <Button type="primary" size="large" onClick={onClick}>
        SNAP
      </Button>
    </Space>
  );
};

export default Snap;
