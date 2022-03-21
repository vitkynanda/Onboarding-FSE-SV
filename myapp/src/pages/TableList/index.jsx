import { DeleteOutlined, EditOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, message, Input, Drawer } from 'antd';
import React, { useState, useRef } from 'react';
import { useIntl, FormattedMessage } from 'umi';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { ModalForm, ProFormText, ProFormSelect } from '@ant-design/pro-form';
// import ProDescriptions from '@ant-design/pro-descriptions';
// import UpdateForm from './components/UpdateForm';
import request from 'umi-request';
// import { rule, addRule, updateRule, removeRule } from '@/services/ant-design-pro/api';
/**
 * @en-US Add node
 * @zh-CN 添加节点
 * @param fields
 */

// const handleAdd = async (fields) => {
//   const hide = message.loading('正在添加');

//   try {
//     await addRule({ ...fields });
//     hide();
//     message.success('Added successfully');
//     return true;
//   } catch (error) {
//     hide();
//     message.error('Adding failed, please try again!');
//     return false;
//   }
// };
// /**
//  * @en-US Update node
//  * @zh-CN 更新节点
//  *
//  * @param fields
//  */

// const handleUpdate = async (fields) => {
//   const hide = message.loading('Configuring');

//   try {
//     await updateRule({
//       name: fields.name,
//       desc: fields.desc,
//       key: fields.key,
//     });
//     hide();
//     message.success('Configuration is successful');
//     return true;
//   } catch (error) {
//     hide();
//     message.error('Configuration failed, please try again!');
//     return false;
//   }
// };
// /**
//  *  Delete node
//  * @zh-CN 删除节点
//  *
//  * @param selectedRows
//  */

// const handleRemove = async (selectedRows) => {
//   const hide = message.loading('正在删除');
//   if (!selectedRows) return true;

//   try {
//     await removeRule({
//       key: selectedRows.map((row) => row.key),
//     });
//     hide();
//     message.success('Deleted successfully and will refresh soon');
//     return true;
//   } catch (error) {
//     hide();
//     message.error('Delete failed, please try again');
//     return false;
//   }
// };

const TableList = () => {
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
  const actionRef = useRef();
  const [currentRow, setCurrentRow] = useState();
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
            <Button
              onClick={() => {
                setUserData(row);
                handleModalVisible(true);
              }}
            >
              <EditOutlined />
            </Button>
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
    };

    const response = await request(`http://localhost:8001/users/${value.id}`, options);
    // const result = await response.json();
    if (response.status == 'ok') {
      handleModalVisible(false);

      if (actionRef.current) {
        actionRef.current.reload();
      }
    }
  };

  const removeUser = async (id, actionRef) => {
    // const hide = message.loading('Updating data');
    const options = {
      method: 'DELETE',
      headers: { Authorization: `${localStorage.getItem('token')}` },
      referrerPolicy: 'no-referrer',
    };

    try {
      const response = await request(`http://localhost:8001/users/${id}`, options);
      if (response.status === 'ok') {
        message.success('Deleted successfully and will refresh soon');
        actionRef.current.reload();
      } else {
        message.error(response.error);
      }
    } catch (err) {
      message.error(err.message);
    }
  };

  return (
    <PageContainer>
      <ProTable
        headerTitle={intl.formatMessage({
          id: 'pages.searchTable.title',
          defaultMessage: 'User Data',
        })}
        actionRef={actionRef}
        rowKey="key"
        search={{
          labelWidth: 120,
        }}
        // toolBarRender={() => [
        //   <Button
        //     type="primary"
        //     key="primary"
        //     onClick={() => {
        //       handleModalVisible(true);
        //     }}
        //   >
        //     <PlusOutlined /> <FormattedMessage id="pages.searchTable.new" defaultMessage="New" />
        //   </Button>,
        // ]}
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
              <FormattedMessage id="pages.searchTable.chosen" defaultMessage="Chosen" />{' '}
              <a
                style={{
                  fontWeight: 600,
                }}
              >
                {selectedRowsState.length}
              </a>{' '}
              <FormattedMessage id="pages.searchTable.item" defaultMessage="项" />
              &nbsp;&nbsp;
              <span>
                <FormattedMessage
                  id="pages.searchTable.totalServiceCalls"
                  defaultMessage="Total number of service calls"
                />{' '}
                {selectedRowsState.reduce((pre, item) => pre + item.callNo, 0)}{' '}
                <FormattedMessage id="pages.searchTable.tenThousand" defaultMessage="万" />
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
            <FormattedMessage
              id="pages.searchTable.batchDeletion"
              defaultMessage="Batch deletion"
            />
          </Button>
          <Button type="primary">
            <FormattedMessage
              id="pages.searchTable.batchApproval"
              defaultMessage="Batch approval"
            />
          </Button>
        </FooterToolbar>
      )}
      <ModalForm
        title={intl.formatMessage({
          id: 'pages.searchTable.createForm.newRul',
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
        <ProFormText
          rules={[
            {
              required: true,
              message: (
                <FormattedMessage
                  id="pages.searchTable.ruleName"
                  defaultMessage="Name is required"
                />
              ),
            },
          ]}
          value={userData?.name}
          width="md"
          name="name"
          label="Name"
        />
        <ProFormText
          rules={[
            {
              // required: true,
              message: (
                <FormattedMessage id="pages.searchTable.Email" defaultMessage="Email is required" />
              ),
            },
          ]}
          value={userData?.email}
          width="md"
          name="mail"
          label="Email"
        />
        <ProFormText
          rules={[
            {
              // required: true,
              message: (
                <FormattedMessage
                  id="pages.searchTable.personalNumber"
                  defaultMessage="Email is required"
                />
              ),
            },
          ]}
          value={userData?.personalNumber}
          width="md"
          name="personalNumber"
          label="Personal Number"
        />
        <ProFormText.Password
          rules={[
            {
              // required: true,
              message: (
                <FormattedMessage
                  id="pages.searchTable.password"
                  defaultMessage="Rule name is required"
                />
              ),
            },
          ]}
          value={userData?.password}
          width="md"
          name="pwd"
          label="Password"
        />

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
          value={userData?.active}
          width="xs"
          name="active"
          label="Active Status"
        />
      </ModalForm>
      {/* <UpdateForm
        onSubmit={async (value) => {
          const success = await handleUpdate(value);

          if (success) {
            handleUpdateModalVisible(false);
            setCurrentRow(undefined);

            if (actionRef.current) {
              actionRef.current.reload();
            }
          }
        }}
        onCancel={() => {
          handleUpdateModalVisible(false);

          if (!showDetail) {
            setCurrentRow(undefined);
          }
        }}
        updateModalVisible={updateModalVisible}
        values={currentRow || {}}
      />

      <Drawer
        width={600}
        visible={showDetail}
        onClose={() => {
          setCurrentRow(undefined);
          setShowDetail(false);
        }}
        closable={false}
      >
        {currentRow?.name && (
          <ProDescriptions
            column={2}
            title={currentRow?.name}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.name,
            }}
            columns={columns}
          />
        )}
      </Drawer> */}
    </PageContainer>
  );
};

export default TableList;
