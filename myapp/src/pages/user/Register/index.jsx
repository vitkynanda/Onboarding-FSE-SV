import {
  AlipayCircleOutlined,
  LockOutlined,
  MobileOutlined,
  TaobaoCircleOutlined,
  UserOutlined,
  WeiboCircleOutlined,
} from '@ant-design/icons';
import { Alert, message, Tabs } from 'antd';
import React, { useState } from 'react';
import ProForm, {
  ProFormCaptcha,
  ProFormCheckbox,
  ProFormText,
  LoginForm,
} from '@ant-design/pro-form';
import { useIntl, history, FormattedMessage, SelectLang, useModel } from 'umi';
import Footer from '@/components/Footer';
import { registerSv } from '@/services/ant-design-pro/api';
// import { getFakeCaptcha } from '@/services/ant-design-pro/login';
import styles from './index.less';

const RegisterMessage = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

const Register = () => {
  const [userLoginState, setUserLoginState] = useState({});
  const [type, setType] = useState('account');
  const { initialState, setInitialState } = useModel('@@initialState');
  const intl = useIntl();

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();

    if (userInfo) {
      await setInitialState((s) => ({ ...s, currentUser: userInfo }));
    }
  };

  const handleSubmit = async (values) => {
    try {
    } catch (error) {}
  };

  const { status, type: loginType } = userLoginState;
  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.content}>
        <ProForm
          logo={<img alt="logo" src="/logo.svg" />}
          title="Ant Design"
          subTitle={intl.formatMessage({
            id: 'pages.layouts.userLayout.titles',
            defaultMessage: 'Register Account',
          })}
          initialValues={{
            autoLogin: true,
          }}
          actions={[
            <FormattedMessage
              key="loginWith"
              id="pages.register.loginWith"
              defaultMessage="其他登录方式"
            />,
          ]}
          onFinish={async (values) => {
            await handleSubmit(values);
          }}
        >
          <Tabs activeKey={type} onChange={setType}>
            <Tabs.TabPane
              key="account"
              tab={intl.formatMessage({
                id: 'pages.register.accountLogin.tabs',
                defaultMessage: 'Register Account',
              })}
            />
          </Tabs>

          {status === 'error' && loginType === 'account' && (
            <RegisterMessage
              content={intl.formatMessage({
                id: 'pages.register.accountLogin.errorMessage',
                defaultMessage: '(admin/ant.design)',
              })}
            />
          )}

          {type === 'account' && (
            <>
              <ProFormText
                name="name"
                placeholder={intl.formatMessage({
                  id: 'pages.register.name.placeholder',
                  defaultMessage: 'Username',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.name.required"
                        defaultMessage="Name requireed !"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText
                name="email"
                placeholder={intl.formatMessage({
                  id: 'pages.register.email.placeholder',
                  defaultMessage: 'Email',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.email.required"
                        defaultMessage="Email required"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText
                name="personalNumber"
                placeholder={intl.formatMessage({
                  id: 'pages.register.personalNumber.placeholder',
                  defaultMessage: 'Personal Number',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.personalNumber.required"
                        defaultMessage="Personal number required !"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                placeholder={intl.formatMessage({
                  id: 'pages.register.password.placeholder',
                  defaultMessage: 'Password',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.password.required"
                        defaultMessage="Password required !"
                      />
                    ),
                  },
                ]}
              />
            </>
          )}
        </ProForm>
      </div>
      <Footer />
    </div>
  );
};

export default Register;
