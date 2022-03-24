import { DeleteOutlined, EditOutlined, EyeOutlined, PlusOutlined } from '@ant-design/icons';
import { Button, message, Drawer } from 'antd';
import { FormattedMessage } from 'umi';
import { PageContainer } from '@ant-design/pro-layout';
import { getProductList } from '@/services/ant-design-pro/api';
import { ModalForm, ProFormText } from '@ant-design/pro-form';
import ProDescriptions from '@ant-design/pro-descriptions';
import ProTable from '@ant-design/pro-table';
import useTable from '@/hooks/use-table';

const TableProduct = () => {
  const {
    handleRemove: handleRemoveProduct,
    handleUpdate: handleUpdateProduct,
    handleCreate: handleCreateProduct,
    handleDetail: handleProductDetail,
    handlePublish,
    handleCheck,
    handleModalVisible,
    setCurrentRow,
    setShowDetail,
    setModalType,
    showDetail,
    currentRow,
    actionRef,
    createModalVisible,
    modalType,
    intl,
  } = useTable('product');

  const columns = [
    {
      title: <FormattedMessage id="pages.productTable.updateForm.product" defaultMessage="ID" />,
      dataIndex: 'id',
      tip: 'The id is the unique key',
      render: (dom) => {
        return <span>{dom.slice(0, 10)}...</span>;
      },
    },
    {
      title: (
        <FormattedMessage
          id="pages.productTable.updateForm.product"
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
              setCurrentRow({});
              setModalType('create');
              handleModalVisible(true);
            }}
          >
            <PlusOutlined /> <FormattedMessage id="pages.searchTable.new" defaultMessage="New" />
          </Button>,
        ]}
        request={async () => {
          try {
            const res = await getProductList();
            if (res.statusCode === 200) {
              return res;
            }
          } catch (error) {
            message.error(error.data.error);
          }
        }}
        columns={columns}
      />

      <ModalForm
        title={intl.formatMessage({
          id: 'pages.productTable.createForm.user',
          defaultMessage: modalType === 'edit' ? 'Edit Product' : 'Create Product',
        })}
        width="400px"
        visible={createModalVisible}
        onVisibleChange={handleModalVisible}
        onFinish={(value) => {
          switch (modalType) {
            case 'edit':
              handleUpdateProduct({ ...value, id: currentRow.id });
              break;
            case 'publish':
              handlePublish({ ...value, id: currentRow.id });
              break;
            case 'check':
              handleCheck({ ...value, id: currentRow.id });
              break;
            case 'create':
              handleCreateProduct(value);
          }
        }}
        submitter={{
          render: (props) => {
            return modalType === 'edit' ? (
              <div id="actions" style={{ display: 'flex' }}>
                <button
                  onClick={() => {
                    setModalType('check');
                    props.form?.submit?.();
                  }}
                >
                  Checked
                </button>
                <button
                  onClick={() => {
                    setModalType('publish');
                    props.form?.submit?.();
                  }}
                >
                  Published
                </button>
                <button onClick={() => props.form?.submit?.()}>Submit</button>
              </div>
            ) : (
              <div>
                <button type="submit" onClick={() => props.form?.submit?.()}>
                  Submit
                </button>
              </div>
            );
          },
        }}
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
