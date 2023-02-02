import { FileOutlined, SettingOutlined, EyeOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme, Divider, Typography, AutoComplete } from 'antd';
import React, { useState } from 'react';
import ArticleView from './View/Article/article';
import SettingView from './View/Setting/setting';
import PreferenceView from './View/Preference/preference';
const { Header, Content, Footer, Sider } = Layout;
const { Text, Link } = Typography;


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
  getItem('文章管理', '0', <FileOutlined />, null, <ArticleView/>),
  getItem('界面管理', '1', <EyeOutlined />, null, <PreferenceView/>),
  getItem('功能管理', '2', <SettingOutlined />, null, <SettingView/>),
];

const App = () => {
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer },
  } = theme.useToken();

  const [selectedKey, setSelectedKey] = useState([]);
  const onSelectedKey = ({ item, key, keyPath, selectedKeys, domEvent }) => {
    setSelectedKey(selectedKeys);
  }
  const index = selectedKey.length > 0 ? parseInt(selectedKey[0]) : 0
  const content = items[index].content
  const title = collapsed ? "" : "一点心得后台管理"

  return (
    <Layout
      style={{
        minHeight: '100vh',
        // margin: 0,
        // padding: 0,
        // backgroundColor: "red",
      }}
    >
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}
       style={{margin: 0, padding: 0, /*backgroundColor: 'blue'*/}}>
        <div
          style={{
            height: 40,
            fontSize: 22,
            textAlign: 'center',
            padding: 5,
            background: 'rgba(255, 255, 255, 0.2)',
          }}
        >{title}</div>
        <Menu theme="dark" defaultSelectedKeys={['0']} mode="inline" items={items} onSelect={onSelectedKey} />
      </Sider>
      <Layout className="site-layout"
        style={{
          padding: 0,
          margin: 0,
        //  background: 'green',
          fontSize: 20,
          // paddingLeft: 15,
        }}
      >
        {content}
        <Divider style={{margin:0, padding: 0,}}/>
        <Footer
          style={{
            textAlign: 'center',
            background: colorBgContainer,
          }}
        >Copyright ©2023 晓白齐齐</Footer>
      </Layout>
    </Layout>
  );
};
export default App;
