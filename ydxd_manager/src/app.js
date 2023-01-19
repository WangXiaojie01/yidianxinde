import { FileOutlined, SettingOutlined, EyeOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme } from 'antd';
import React, { useState } from 'react';
const { Header, Content, Footer, Sider } = Layout;

function getItem(label, key, icon, children, content) {
  return {
    key,
    icon,
    children,
    label,
    content,
  };
}
const items = [
  getItem('文章管理', '0', <FileOutlined />, null, <div>11111</div>),
  getItem('界面管理', '1', <EyeOutlined />, null, <div>22222</div>),
  getItem('功能管理', '2', <SettingOutlined />, null, <div>33333</div>),
];

const App = () => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  const [selectedKey, setSelectedKey] = useState([]);
  const onSelectedKey = ({ item, key, keyPath, selectedKeys, domEvent }) => {
    setSelectedKey(selectedKeys);
    console.log("key is ", selectedKey)
  }
  const index = selectedKey.length > 0 ? parseInt(selectedKey[0]) : 0
  const content = items[index].content

  return (
    <Layout
      style={{
        minHeight: '100vh',
      }}
    >
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <div
          style={{
            height: 32,
            margin: 16,
            background: 'rgba(255, 255, 255, 0.2)',
          }}
        />
        <Menu theme="dark" defaultSelectedKeys={['0']} mode="inline" items={items} onSelect={onSelectedKey} />
      </Sider>
      <Layout className="site-layout">
        <Header
          style={{
            padding: 0,
            background: colorBgContainer,
            fontSize: 20,
            paddingLeft: 15,
          }}
        >一点心得后台管理</Header>
        <Content
          style={{
            margin: '0 16px',
          }}
        >
          {/*<Breadcrumb
            style={{
              margin: '16px 0',
            }}
            separator=">"
          >
            <Breadcrumb.Item>文章管理</Breadcrumb.Item>
            <Breadcrumb.Item>Bill</Breadcrumb.Item>
            <Breadcrumb.Item>Bill2</Breadcrumb.Item>
          </Breadcrumb>*/}
          {content}
        </Content>
        <Footer
          style={{
            textAlign: 'center',
          }}
        >
        Copyright ©2023 晓白齐齐
        </Footer>
      </Layout>
    </Layout>
  );
};
export default App;
