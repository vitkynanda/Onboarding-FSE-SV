import { ModalForm, ProFormText } from '@ant-design/pro-form';
import React from 'react';

const Modal = () => {
  return (
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
  );
};

export default index;
