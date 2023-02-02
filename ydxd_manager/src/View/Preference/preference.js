import { FileOutlined, SettingOutlined, EyeOutlined } from '@ant-design/icons';
import { Button, Layout, Form, theme, Input, Typography, Divider, Space } from 'antd';
import React, { useState } from 'react';

const { Text, Link } = Typography;
const { Header, Content, Footer, Sider } = Layout;


const PreferenceView = () => {
  const {
    token: { colorBgContainer },
  } = theme.useToken();
  return (
    <Content 
      style={{
        background: colorBgContainer,
        padding: 10,
      }}
    >
      <Layout style={{
        background: colorBgContainer,  
      }}>
        <Space>
          <Button type="primary" style={{width: 200,}}>更新前端代码</Button>
          <Button type="primary" style={{width: 200,}}>更新后端代码</Button>
        </Space>
        <Divider/>
        <Form>
          <Form.Item
            label="全局搜索的最大标题层级："
            name="global_level"
          >
            <Input />
          </Form.Item>

          <Form.Item
            label="建立localStorage间隔时间（分钟）："
            name="local_time"
          >
            <Input />
          </Form.Item>
          <Form.Item
            wrapperCol={{
              offset: 8,
              span: 16,
            }}
          >
            <Button type="primary" htmlType="submit">
              保存
            </Button>
          </Form.Item>
        </Form>
      </Layout>
    </Content>
  );
};
export default PreferenceView;
