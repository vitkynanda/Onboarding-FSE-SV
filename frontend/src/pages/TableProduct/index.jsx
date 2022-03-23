import { DeleteOutlined, EditOutlined, EyeOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, message, Drawer } from 'antd';
import { useState, useRef } from 'react';
import { useIntl, FormattedMessage } from 'umi';
import { PageContainer, FooterToolbar } from '@ant-design/pro-layout';
import ProTable from '@ant-design/pro-table';
import { ModalForm, ProFormText } from '@ant-design/pro-form';
import ProDescriptions from '@ant-design/pro-descriptions';
import {
  getProductList,
  productDetail,
  updateProduct,
  createNewProduct,
  removeProduct,
} from '@/services/ant-design-pro/api';

const TableProduct = () => {
  const [createModalVisible, handleModalVisible] = useState(false);
  const [showDetail, setShowDetail] = useState(false);
  const [modalType, setModalType] = useState();
  const actionRef = useRef();
  const [currentRow, setCurrentRow] = useState();
  const [selectedRowsState, setSelectedRows] = useState([]);

  const handleRemoveProduct = async (id) => {
    try {
      const response = await removeProduct(id);
      if (response.statusCode === 200) {
        message.success('Deleted successfully');
        actionRef.current.reload();
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleUpdateProduct = async (value) => {
    const payload = {
      name: value.name,
      description: value.description,
    };

    try {
      const response = await updateProduct(value.id, payload);
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

  const handleProductDetail = async (id) => {
    try {
      const response = await productDetail(id);
      if (response.statusCode === 200) {
        setShowDetail(true);
        setCurrentRow(response.data);
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleCreateProduct = async (value) => {
    try {
      const response = await createNewProduct(value);
      if (response.statusCode === 201) {
        message.success('Product created successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const intl = useIntl();

  const columns = [
    {
      title: (
        <FormattedMessage
          id="pages.productTable.updateForm.ruleName.nameLabel"
          defaultMessage="ID"
        />
      ),
      dataIndex: 'id',
      tip: 'The id is the unique key',
      render: (dom) => {
        return <span>{dom.slice(0, 10)}...</span>;
      },
    },
    {
      title: (
        <FormattedMessage
          id="pages.productTable.updateForm.ruleName.nameLabel"
          defaultMessage="Product name"
        />
      ),
      dataIndex: 'name',
      tip: 'The user name is the unique key',
      render: (dom, rowData) => {
        return (
          <a
            onClick={() => {
              handleProductDetail(rowData.id);
            }}
          >
            {dom}
          </a>
        );
      },
    },

    {
      title: <FormattedMessage id="pages.productTable.titleStatus" defaultMessage="Description" />,
      dataIndex: 'role',
      hideInForm: true,
      render: (_, rowData) => {
        return <span>{rowData.description}</span>;
      },
    },
    {
      title: <FormattedMessage id="pages.productTable.titleStatus" defaultMessage="Status" />,
      dataIndex: 'status',
      hideInForm: true,
      render: (_, rowData) => {
        return <span>{rowData.status}</span>;
      },
    },

    {
      title: <FormattedMessage id="pages.productTable.titleOption" defaultMessage="Operating" />,
      dataIndex: 'option',
      valueType: 'option',
      render: (_, rowData) => {
        return (
          <div style={{ display: 'flex' }}>
            <div style={{ marginRight: 5 }}>
              <Button onClick={() => handleRemoveProduct(rowData.id)}>
                <DeleteOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  handleModalVisible(true);
                  setCurrentRow(rowData);
                  setModalType('edit');
                }}
              >
                <EditOutlined />
              </Button>
            </div>
            <div style={{ marginRight: 5 }}>
              <Button
                onClick={() => {
                  handleProductDetail(rowData.id);
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
          id: 'pages.productTable.title',
          defaultMessage: 'Product List',
        })}
        actionRef={actionRef}
        rowKey="key"
        search={{
          labelWidth: 120,
        }}
        toolBarRender={() => [
          <Button
            type="primary"
            key="primary"
            onClick={() => {
              setCurrentRow(undefined);
              setModalType('add');
              handleModalVisible(true);
            }}
          >
            <PlusOutlined /> <FormattedMessage id="pages.searchTable.new" defaultMessage="New" />
          </Button>,
        ]}
        request={getProductList}
        columns={columns}
        rowSelection={{
          onChange: (_, selectedRows) => {
            setSelectedRows(selectedRows);
          },
        }}
      />
      {/* {selectedRowsState?.length > 0 && (
        <FooterToolbar
          extra={
            <div>
              <FormattedMessage id="pages.productTable.chosen" defaultMessage="Chosen" />{' '}
              <a
                style={{
                  fontWeight: 600,
                }}
              >
                {selectedRowsState.length}
              </a>{' '}
              <FormattedMessage id="pages.productTable.item" defaultMessage="项" />
              &nbsp;&nbsp;
              <span>
                <FormattedMessage
                  id="pages.productTable.totalServiceCalls"
                  defaultMessage="Total number of service calls"
                />{' '}
                {selectedRowsState.reduce((pre, item) => pre + item.callNo, 0)}{' '}
                <FormattedMessage id="pages.productTable.tenThousand" defaultMessage="万" />
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
              id="pages.productTable.batchDeletion"
              defaultMessage="Batch deletion"
            />
          </Button>
          <Button type="primary">
            <FormattedMessage
              id="pages.productTable.batchApproval"
              defaultMessage="Batch approval"
            />
          </Button>
        </FooterToolbar>
      )} */}
      <ModalForm
        initialValues={{
          name: currentRow?.name,
          description: currentRow?.description,
        }}
        title={intl.formatMessage({
          id: 'pages.productTable.createForm.newRule',
          defaultMessage: modalType === 'edit' ? 'Edit Product' : 'Create Product',
        })}
        width="400px"
        visible={createModalVisible}
        onVisibleChange={handleModalVisible}
        onFinish={(value) =>
          modalType === 'edit'
            ? handleUpdateProduct({ ...value, id: currentRow.id })
            : handleCreateProduct(value)
        }
      >
        <ProFormText width="md" name="name" placeholder="Name" label="Name" />
        <ProFormText width="md" name="description" placeholder="Description" label="Description" />
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

export default TableProduct;
