import { DeleteOutlined, EditOutlined, EyeOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, message, Input, Drawer } from 'antd';
import React, { useState, useRef } from 'react';
import { useIntl, FormattedMessage } from 'umi';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { ModalForm, ProFormSelect, ProFormText, ProFormTextArea } from '@ant-design/pro-form';
import ProDescriptions from '@ant-design/pro-descriptions';
import UpdateForm from './components/UpdateForm';
import {
  rule,
  addRule,
  updateRule,
  removeRule,
  getUserList,
  removeUser,
  updateUser,
  getUserDetail,
  userDetail,
  getRole,
} from '@/services/ant-design-pro/api';
/**
 * @en-US Add node
 * @zh-CN 添加节点
 * @param fields
 */

const handleAdd = async (fields) => {
  const hide = message.loading('正在添加');

  try {
    await addRule({ ...fields });
    hide();
    message.success('Added successfully');
    return true;
  } catch (error) {
    hide();
    message.error('Adding failed, please try again!');
    return false;
  }
};
/**
 * @en-US Update node
 * @zh-CN 更新节点
 *
 * @param fields
 */

const handleUpdate = async (fields) => {
  const hide = message.loading('Configuring');

  try {
    await updateRule({
      name: fields.name,
      desc: fields.desc,
      key: fields.key,
    });
    hide();
    message.success('Configuration is successful');
    return true;
  } catch (error) {
    hide();
    message.error('Configuration failed, please try again!');
    return false;
  }
};
/**
 *  Delete node
 * @zh-CN 删除节点
 *
 * @param selectedRows
 */

const handleRemove = async (selectedRows) => {
  const hide = message.loading('正在删除');
  if (!selectedRows) return true;

  try {
    await removeRule({
      key: selectedRows.map((row) => row.key),
    });
    hide();
    message.success('Deleted successfully and will refresh soon');
    return true;
  } catch (error) {
    hide();
    message.error('Delete failed, please try again');
    return false;
  }
};

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
  const [roleData, setRoleData] = useState([]);
  const actionRef = useRef();
  const [currentRow, setCurrentRow] = useState();
  const [selectedRowsState, setSelectedRows] = useState([]);
  /**
   * @en-US International configuration
   * @zh-CN 国际化配置
   * */

  const handleRemoveUser = async (id) => {
    try {
      const response = await removeUser(id);
      if (response.statusCode === 200) {
        message.success('Deleted successfully');
        actionRef.current.reload();
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleUpdateUser = async (value) => {
    const payload = {
      personalNumber: value.personalNumber ?? '',
      password: value.pwd ?? '',
      email: value.email ?? '',
      name: value.name,
      active: value.active,
      role: {
        id: value.role ?? value.role_id,
      },
    };

    try {
      const response = await updateUser(value.id, payload);

      if (response.statusCode === 200) {
        message.success('Update user successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleUserDetail = async (id) => {
    try {
      const response = await userDetail(id);
      if (response.statusCode === 200) {
        setShowDetail(true);
        setCurrentRow(response.data);
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const getRoleData = async () => {
    try {
      const response = await getRole();
      if (response.statusCode === 200) {
        setRoleData(
          response.data.map((role) => {
            return { value: role.id, label: role.title };
          }),
        );
        handleModalVisible(true);
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const intl = useIntl();

  const columns = [
    {
      title: (
        <FormattedMessage id="pages.userTable.updateForm.ruleName.nameLabel" defaultMessage="ID" />
      ),
      dataIndex: 'id',
      tip: 'The id is the unique key',
      render: (dom, entity) => {
        return <span>{dom.slice(0, 10)}...</span>;
      },
    },
    {
      title: (
        <FormattedMessage
          id="pages.userTable.updateForm.ruleName.nameLabel"
          defaultMessage="User name"
        />
      ),
      dataIndex: 'name',
      tip: 'The user name is the unique key',
      render: (dom, entity) => {
        return (
          <a
            onClick={() => {
              handleUserDetail(entity.id);
            }}
          >
            {dom}
          </a>
        );
      },
    },

    {
      title: <FormattedMessage id="pages.userTable.titleStatus" defaultMessage="Role" />,
      dataIndex: 'role',
      hideInForm: true,
      render: (_, rowData) => {
        return <span>{rowData.role.title}</span>;
      },
    },
    {
      title: <FormattedMessage id="pages.userTable.titleStatus" defaultMessage="Status" />,
      dataIndex: 'active',
      hideInForm: true,
      render: (_, rowData) => {
        return <span>{rowData.active ? 'Active' : 'Inactive'}</span>;
      },
    },

    {
      title: <FormattedMessage id="pages.userTable.titleOption" defaultMessage="Operating" />,
      dataIndex: 'option',
      valueType: 'option',
      render: (_, rowData) => {
        return (
          <div style={{ display: 'flex' }}>
            <div style={{ marginRight: 5 }}>
              <Button onClick={() => handleRemoveUser(rowData.id)}>
                <DeleteOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  getRoleData(rowData);
                  setCurrentRow(rowData);
                }}
              >
                <EditOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  handleUserDetail(rowData.id);
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
  return (
    <PageContainer>
      <ProTable
        headerTitle={intl.formatMessage({
          id: 'pages.userTable.title',
          defaultMessage: 'User List',
        })}
        actionRef={actionRef}
        rowKey="key"
        search={{
          labelWidth: 120,
        }}
        request={getUserList}
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
              <FormattedMessage id="pages.userTable.chosen" defaultMessage="Chosen" />{' '}
              <a
                style={{
                  fontWeight: 600,
                }}
              >
                {selectedRowsState.length}
              </a>{' '}
              <FormattedMessage id="pages.userTable.item" defaultMessage="项" />
              &nbsp;&nbsp;
              <span>
                <FormattedMessage
                  id="pages.userTable.totalServiceCalls"
                  defaultMessage="Total number of service calls"
                />{' '}
                {selectedRowsState.reduce((pre, item) => pre + item.callNo, 0)}{' '}
                <FormattedMessage id="pages.userTable.tenThousand" defaultMessage="万" />
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
            <FormattedMessage id="pages.userTable.batchDeletion" defaultMessage="Batch deletion" />
          </Button>
          <Button type="primary">
            <FormattedMessage id="pages.userTable.batchApproval" defaultMessage="Batch approval" />
          </Button>
        </FooterToolbar>
      )}
      <ModalForm
        initialValues={{
          name: currentRow?.name,
          active: currentRow?.active,
        }}
        title={intl.formatMessage({
          id: 'pages.userTable.createForm.newRule',
          defaultMessage: 'Edit user',
        })}
        width="400px"
        visible={createModalVisible}
        onVisibleChange={handleModalVisible}
        onFinish={(value) =>
          handleUpdateUser({ ...value, id: currentRow.id, role_id: currentRow.role.id })
        }
      >
        <ProFormText
          rules={[
            {
              required: true,
              message: (
                <FormattedMessage
                  id="pages.userTable.ruleName"
                  defaultMessage="User name is required"
                />
              ),
            },
          ]}
          width="md"
          name="name"
          placeholder="Name"
          label="Name"
        />
        <ProFormText width="md" name="email" placeholder="Email" label="Email" />
        <ProFormText
          width="md"
          name="personalNumber"
          placeholder="Personal Number"
          label="Personal Number"
        />
        <ProFormText width="md" name="password" placeholder="Password" label="Password" />
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
        <ProFormSelect request={() => [...roleData]} width="xs" name="role" label="Role" />
      </ModalForm>
      <UpdateForm
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
            title={currentRow?.name.toUpperCase()}
            request={async () => ({
              data: currentRow || {},
            })}
            params={{
              id: currentRow?.name,
            }}
            columns={columns}
          />
        )}
      </Drawer>
    </PageContainer>
  );
};

export default TableUser;
