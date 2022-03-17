import React, { useCallback } from 'react';
import { LogoutOutlined, SettingOutlined, UserOutlined } from '@ant-design/icons';
import { Avatar, Menu, Spin } from 'antd';
import { history, useModel } from 'umi';
import { stringify } from 'querystring';
import HeaderDropdown from '../HeaderDropdown';
import styles from './index.less';
import users from '@/helpers/auth';
// import { outLogin } from '@/services/ant-design-pro/api';

/**
 * 退出登录，并且将当前的 url 保存
 */
const loginOut = async () => {
  localStorage.removeItem('token');
  history.push('/user/login');
  // await outLogin();
  // const { query = {}, search, pathname } = history.location;
  // const { redirect } = query; // Note: There may be security issues, please note

  // if (window.location.pathname !== '/user/login' && !redirect) {
  //   history.replace({
  //     pathname: '/user/login',
  //     search: stringify({
  //       redirect: pathname + search,
  //     }),
  //   });
  // }
};

const AvatarDropdown = ({ menu }) => {
  const { initialState, setInitialState } = useModel('@@initialState');
  console.log(initialState);

  const onMenuClick = useCallback((event) => {
    const { key } = event;

    if (key === 'logout') {
      localStorage.removeItem('token');
      history.push('/user/login');
      return;
    }

    history.push(`/account/${key}`);
  }, []);
  const loading = (
    <span className={`${styles.action} ${styles.account}`}>
      <Spin
        size="small"
        style={{
          marginLeft: 8,
          marginRight: 8,
        }}
      />
    </span>
  );

  // if (!initialState) {
  //   return loading;
  // }

  // const { currentUser } = initialState;

  // if (!currentUser || !currentUser.name) {
  //   return loading;
  // }

  const menuHeaderDropdown = (
    <Menu className={styles.menu} selectedKeys={[]} onClick={onMenuClick}>
      {menu && (
        <Menu.Item key="center">
          <UserOutlined />
          个人中心
        </Menu.Item>
      )}
      {menu && (
        <Menu.Item key="settings">
          <SettingOutlined />
          个人设置
        </Menu.Item>
      )}
      {menu && <Menu.Divider />}

      <Menu.Item key="logout">
        <LogoutOutlined />
        Logout
      </Menu.Item>
    </Menu>
  );
  return (
    <HeaderDropdown overlay={menuHeaderDropdown}>
      <span className={`${styles.action} ${styles.account}`}>
        {/* <Avatar size="small" className={styles.avatar} src={currentUser?.avatar} alt="avatar" /> */}
        <span className={`${styles.name} anticon`}>{initialState.state.username}</span>
      </span>
    </HeaderDropdown>
  );
};

export default AvatarDropdown;
