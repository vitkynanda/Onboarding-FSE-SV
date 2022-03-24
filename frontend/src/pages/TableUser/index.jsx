import { DeleteOutlined, EditOutlined, EyeOutlined } from '@ant-design/icons';
import { Button, message, Drawer } from 'antd';
import { FormattedMessage } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import { ModalForm, ProFormSelect, ProFormText } from '@ant-design/pro-form';
import { getUserList } from '@/services/ant-design-pro/api';
import ProDescriptions from '@ant-design/pro-descriptions';
import ProTable from '@ant-design/pro-table';
import useTable from '@/hooks/use-table';

const TableUser = () => {
  const {
    handleRemove: handleRemoveUser,
    handleDetail: handleUserDetail,
    handleUpdate: handleUpdateUser,
    getRoleData,
    setCurrentRow,
    setShowDetail,
    handleModalVisible,
    actionRef,
    currentRow,
    showDetail,
    createModalVisible,
    intl,
    roleData,
  } = useTable('user');

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
        request={async () => {
          try {
            const res = await getUserList();
            if (res.statusCode === 200) {
              return res;
            }
          } catch (e) {
            message.error(e.data.error);
          }
        }}
        columns={columns}
      />

      <ModalForm
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
