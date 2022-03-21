import { DeleteOutlined, EditOutlined, EyeOutlined } from '@ant-design/icons';
import { Button, message, Input, Drawer } from 'antd';
import React, { useState, useRef } from 'react';
import { useIntl, FormattedMessage } from 'umi';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { ModalForm, ProFormText, ProFormSelect } from '@ant-design/pro-form';
import request from 'umi-request';

/*
 * @en-US Add node
 * @zh-CN 添加节点
 * @param fields
 */

const TableUser = () => {
  /**
   * @en-US Pop-up window of new window
   * @zh-CN 新建窗口的弹窗
   *  */
  const [createModalVisible, handleModalVisible] = useState(false);
  /**
   * @en-US The pop-up window of the distribution update window
   * @zh-CN 分布更新窗口的弹窗
   * */

  const [updateModalVisible, handleUpdateModalVisible] = useState(false);
  const [showDetail, setShowDetail] = useState(false);
  const [currentRow, setCurrentRow] = useState();
  const actionRef = useRef();
  const [selectedRowsState, setSelectedRows] = useState([]);
  const [userData, setUserData] = useState({});
  /**
   * @en-US International configuration
   * @zh-CN 国际化配置
   * */

  const intl = useIntl();
  const columns = [
    {
      title: 'ID',
      dataIndex: 'id',
      sorter: true,
      copyable: true,
      tip: 'The rule id is the unique key',
      render: (_, row) => {
        return <p title={row.id}>{row.id.slice(0, 8)}...</p>;
      },
    },
    {
      title: 'Name',
      dataIndex: 'name',
      sorter: true,
      tip: 'User Name',
      render: (_, row) => {
        return <p>{row.name}</p>;
      },
    },
    {
      title: 'Role',
      dataIndex: 'role',
      sorter: true,
      tip: 'User Role',
      render: (_, row) => {
        return <p>{row.role.title}</p>;
      },
    },
    {
      title: 'Action',
      dataIndex: 'action',
      sorter: true,
      tip: 'User Role',
      render: (_, row) => {
        return (
          <div style={{ display: 'flex' }}>
            <div style={{ marginRight: 5 }}>
              <Button onClick={() => removeUser(row.id, actionRef)}>
                <DeleteOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  setUserData(row);
                  handleModalVisible(true);
                }}
              >
                <EditOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  setUserData(row);
                  handleModalVisible(true);
                }}
              >
                <EyeOutlined />
              </Button>
            </div>
          </div>
        );
      },
    },
  ];

  const handleEditUser = async (value) => {
    const options = {
      method: 'PUT',
      body: JSON.stringify({
        personalNumber: value.personalNumber ?? '',
        password: value.pwd ?? '',
        email: value.mail ?? '',
        name: value.name,
        active: value.active,
        role: {
          id: value.role_id,
        },
      }),
      headers: { Authorization: `${localStorage.getItem('token')}` },
    };

    try {
      const response = await request(`http://localhost:8001/users/${value.id}`, options);
      if (response.status === 'ok') {
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      message.error(error.data.error);
    }
  };

  const removeUser = async (id, actionRef) => {
    const options = {
      method: 'DELETE',
      headers: { Authorization: `${localStorage.getItem('token')}` },
    };

    try {
      const response = await request(`http://localhost:8001/users/${id}`, options);
      if (response.status === 'ok') {
        message.success('Deleted successfully');
        actionRef.current.reload();
      }
    } catch (err) {
      message.error(err.data.error);
    }
  };

  return (
    <PageContainer>
      <ProTable
        headerTitle={intl.formatMessage({
          id: 'pages.usersTable.title',
          defaultMessage: 'User Data',
        })}
        actionRef={actionRef}
        rowKey="key"
        search={{
          labelWidth: 120,
        }}
        request={async (params = {}) => {
          const response = await request('http://localhost:8001/users', {
            params,
          });

          return response;
        }}
        columns={columns}
        rowSelection={{
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows);
          },
        }}
      />
      {selectedRowsState?.length > 0 && (
        <FooterToolbar
          extra={
            <div>
              <FormattedMessage id="pages.usersTable.chosen" defaultMessage="Chosen" />{' '}
              <a
                style={{
                  fontWeight: 600,
                }}
              >
                {selectedRowsState.length}
              </a>{' '}
              <FormattedMessage id="pages.usersTable.item" defaultMessage="项" />
              &nbsp;&nbsp;
              <span>
                <FormattedMessage
                  id="pages.usersTable.totalServiceCalls"
                  defaultMessage="Total number of service calls"
                />{' '}
                {selectedRowsState.reduce((pre, item) => pre + item.callNo, 0)}{' '}
                <FormattedMessage id="pages.usersTable.tenThousand" defaultMessage="万" />
              </span>
            </div>
          }
        >
          <Button
            onClick={async () => {
              await handleRemove(selectedRowsState);
              setSelectedRows([]);
              actionRef.current?.reloadAndRest?.();
            }}
          >
            <FormattedMessage id="pages.usersTable.batchDeletion" defaultMessage="Batch deletion" />
          </Button>
          <Button type="primary">
            <FormattedMessage id="pages.usersTable.batchApproval" defaultMessage="Batch approval" />
          </Button>
        </FooterToolbar>
      )}
      <ModalForm
        title={intl.formatMessage({
          id: 'pages.usersTable.createForm.newRul',
          defaultMessage: 'Edit User',
        })}
        initialValues={{
          name: userData.name,
          active: userData.active,
        }}
        width="400px"
        visible={createModalVisible}
        onVisibleChange={handleModalVisible}
        onFinish={(value) => {
          handleEditUser({ ...value, id: userData.id, role_id: userData.role.id });
        }}
      >
        <ProFormText width="md" name="name" label="Name" autoComplete="off" />
        <ProFormText width="md" name="email" label="Email" autoComplete="off" />
        <ProFormText width="md" name="personalNumber" label="Personal Number" autoComplete="off" />
        <ProFormText.Password width="md" name="pwd" label="Password" autoComplete="off" />
        <ProFormSelect
          request={async () => [
            {
              value: true,
              label: 'Active',
            },
            {
              value: false,
              label: 'Inactive',
            },
          ]}
          width="xs"
          name="active"
          label="Active Status"
        />
      </ModalForm>
    </PageContainer>
  );
};

export default TableUser;
