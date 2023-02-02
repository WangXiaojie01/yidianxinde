import { Space, Tree } from 'antd';
import { EditOutlined, PlusOutlined, MinusOutlined } from '@ant-design/icons';
import { TreeNode } from 'antd/es/tree-select';
import { useState } from 'react';
const treeDatas = [
  {
    value: '0-0',
    key: '0-0',
    children: [
      {
        value: '0-0-0',
        key: '0-0-0',
        children: [
          {
            value: '0-0-0-0',
            key: '0-0-0-0',
          },
          {
            value: '0-0-0-1',
            key: '0-0-0-1',
          },
          {
            value: '0-0-0-2',
            key: '0-0-0-2',
            children: [
              {
                value: '0-0-0-2-1',
                key: '0-0-0-2-1'
              }
            ]
          },
        ],
      },
      {
        value: '0-0-1',
        key: '0-0-1',
        children: [
          {
            value: '0-0-1-0',
            key: '0-0-1-0',
          },
          {
            value: '0-0-1-1',
            key: '0-0-1-1',
          },
          {
            value: '0-0-1-2',
            key: '0-0-1-2',
          },
        ],
      },
      {
        value: '0-0-2',
        key: '0-0-2',
      },
    ],
  },
  {
    value: '0-1',
    key: '0-1',
    children: [
      {
        value: '0-1-0-0',
        key: '0-1-0-0',
      },
      {
        value: '0-1-0-1',
        key: '0-1-0-1',
      },
      {
        value: '0-1-0-2',
        key: '0-1-0-2',
      },
    ],
  },
  {
    value: '0-2',
    key: '0-2',
  },
];
const onDelete = (item) => {
  console.log("delete", item);
}
const ArticleTree = () => {
  const [expandedKeys, setExpandedKeys] = useState(['0-0-0', '0-0-1']);
  const [checkedKeys, setCheckedKeys] = useState(['0-0-0']);
  const [selectedKeys, setSelectedKeys] = useState([]);
  const [autoExpandParent, setAutoExpandParent] = useState(true);
  const onExpand = (expandedKeysValue) => {
    console.log('onExpand', expandedKeysValue);
    // if not set autoExpandParent to false, if children expanded, parent can not collapse.
    // or, you can remove all expanded children keys.
    setExpandedKeys(expandedKeysValue);
    setAutoExpandParent(false);
  };
  const onCheck = (checkedKeysValue) => {
    console.log('onCheck', checkedKeysValue);
    setCheckedKeys(checkedKeysValue);
  };
  const onSelect = (selectedKeysValue, info) => {
    console.log('onSelect', info);
    setSelectedKeys(selectedKeysValue);
  };
  const [data, setdata] = useState(treeDatas)
  const renderTreeNodes = (data) => {
    let dataNode = data.map((item) => {
      item.title = <div>
        <span>{item.value}</span>
        <span style={{margin: 20}}>
        <EditOutlined style={{margin: 5}}></EditOutlined>
        <MinusOutlined style={{margin: 5}} onClick={()=>onDelete(item)}></MinusOutlined>
        <PlusOutlined style={{margin: 5}}></PlusOutlined>
        </span>

      </div> ;
      if (item.children) {
        item.children = renderTreeNodes(item.children)
      }
      return item
    });
    return dataNode;
  };
  return data.length ? (
    <Tree
      checkable
      draggable
      selectable={false}
      expandedKeys={expandedKeys}
      selectedKeys={selectedKeys}
      checkedKeys={checkedKeys}
      autoExpandParent={autoExpandParent}
      onExpand={onExpand}
      onCheck={onCheck}
      onSelect={onSelect}
      treeData={renderTreeNodes(data)}
    />) : ('正在加载导航和文章...');
};

export default ArticleTree;
