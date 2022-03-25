import {
  createNewProduct,
  updateProduct,
  updateUser,
  removeProduct,
  removeUser,
  userDetail,
  productDetail,
  getRole,
  publishProduct,
  checkProduct,
} from '@/services/ant-design-pro/api';
import { useIntl } from 'umi';
import { message } from 'antd';
import { useRef, useState } from 'react';

const useTable = (type) => {
  // State
  const intl = useIntl();
  const actionRef = useRef();
  const [createModalVisible, handleModalVisible] = useState(false);
  const [showDetail, setShowDetail] = useState(false);
  const [roleData, setRoleData] = useState([]);
  const [currentRow, setCurrentRow] = useState();
  const [modalType, setModalType] = useState();

  // General Function
  const handleRemove = async (id) => {
    try {
      const response = await (type === 'product' ? removeProduct(id) : removeUser(id));
      if (response.statusCode === 200) {
        message.success('Deleted successfully');
        actionRef.current.reload();
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleUpdate = async (value) => {
    const payload =
      type === 'product'
        ? {
            name: value.name,
            description: value.description,
          }
        : {
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
      const response = await (type === 'product'
        ? updateProduct(value.id, payload)
        : updateUser(value.id, payload));
      if (response.statusCode === 200) {
        message.success('Update user successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      message.error(error?.data?.error);
      handleModalVisible(false);
    }
  };

  const handleDetail = async (id) => {
    try {
      const response = await (type === 'product' ? productDetail(id) : userDetail(id));
      if (response.statusCode === 200) {
        setShowDetail(true);
        setCurrentRow(response.data);
      }
    } catch (error) {
      message.error(error?.data?.error);
    }
  };

  const handleCreate = async (value) => {
    try {
      const response = await (type === 'product'
        ? createNewProduct(value)
        : createNewProduct(value));
      if (response.statusCode === 201) {
        message.success('Product created successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      handleModalVisible(false);
      message.error(error?.data?.error);
    }
  };

  // Function User
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

  // Function Product
  const handlePublish = async (value) => {
    const payload = {
      name: value.name,
      description: value.description,
    };

    try {
      const response = await publishProduct(value.id, payload);
      if (response.statusCode === 200) {
        message.success('Update user successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      handleModalVisible(false);
      message.error(error?.data?.error);
    }
  };

  const handleCheck = async (value) => {
    const payload = {
      name: value.name ?? currentRow.name,
      description: value.description ?? currentRow.description,
    };

    try {
      const response = await checkProduct(value.id, payload);
      if (response.statusCode === 200) {
        message.success('Update user successfully');
        handleModalVisible(false);
        if (actionRef.current) {
          actionRef.current.reload();
        }
      }
    } catch (error) {
      handleModalVisible(false);
      message.error(error?.data?.error);
    }
  };

  return {
    handleRemove,
    handleUpdate,
    handleCreate,
    handleDetail,
    handleModalVisible,
    setModalType,
    setShowDetail,
    setCurrentRow,
    getRoleData,
    setRoleData,
    handlePublish,
    handleCheck,
    roleData,
    showDetail,
    currentRow,
    actionRef,
    modalType,
    createModalVisible,
    intl,
  };
};

export default useTable;
