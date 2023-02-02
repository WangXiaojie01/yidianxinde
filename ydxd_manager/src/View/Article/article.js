import { FileOutlined, SettingOutlined, EyeOutlined } from '@ant-design/icons';
import { Breadcrumb, Layout, Menu, theme, List, Typography, Divider } from 'antd';
import React, { useState } from 'react';
import ArticleTree from './article_tree';

const { Text, Link } = Typography;
const { Header, Content, Footer, Sider } = Layout;


const ArticleView = () => {
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
        <Typography.Title
          level={5}
          style={{
            margin: 0,
            // background: 'green',
          }}
        >操作说明：</Typography.Title>
        <Divider style={{margin:0, padding: 0,}}/>
        <Text style={{padding: 3,}}>
          1. 通过勾选或取消勾选导航（或文章）前面的复选框以控制是否在网页上公开显示，勾选复选框公开展示，不勾选则不公开
        </Text>
        <Text style={{padding: 3}}>
          2. 通过拖拽导航（或文章）到对应的导航栏或对应的位置以调整导航（或文章）所属的一级导航、二级导航，或调整导航（或文章）的显示顺序
        </Text>
        <Text style={{padding: 3,}}>
          3. 通过点击对应文章可以查看对应的文章内容
        </Text>
        <Text style={{padding: 3,}}>
          4. 通过点击导航（或文章）该行后面的“添加”按钮，在对应的导航（或文章）之后添加同级导航（或文章）
        </Text>
        <Text style={{padding: 3,}}>
          5. 通过点击导航（或文章）该行后面的“删除”按钮，可删除对应的导航（或文章）
        </Text>
        <Text style={{padding: 3,}}>
          6. 通过点击导航（或文章）该行后面的“编辑”按钮，可删除编辑对应的导航（或文章）
        </Text>
      </Layout>
      <Divider/>
      <Layout style={{
        background: colorBgContainer,  
      }}>
        <Typography.Title
          level={5}
          style={{
            margin: 0,
            //background: 'green',
          }}
        >导航和文章列表：</Typography.Title>
        <Divider style={{margin:0, padding: 0,}}/>
        <ArticleTree/>
      </Layout>
    </Content>
  );
};
export default ArticleView;
