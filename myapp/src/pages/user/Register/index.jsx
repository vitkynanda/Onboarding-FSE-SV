import { LockOutlined, UserOutlined } from '@ant-design/icons';
import { Alert, message, Tabs } from 'antd';
import React, { useState } from 'react';
import { ProFormCheckbox, ProFormText, LoginForm } from '@ant-design/pro-form';
import { useIntl, history, FormattedMessage, SelectLang, useModel } from 'umi';
import Footer from '@/components/Footer';
import { svRegister } from '@/services/ant-design-pro/api';
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
  const [userRegisterState, setUserRegisterState] = useState({});
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
      // 登录
      const msg = await svRegister({ ...values, type });

      if (msg.status === 'ok') {
        const defaultRegisterSuccessMessage = intl.formatMessage({
          id: 'pages.register.success',
          defaultMessage: '登录成功！',
        });
        message.success(defaultRegisterSuccessMessage);
        await fetchUserInfo();
        /** 此方法会跳转到 redirect 参数所在的位置 */

        if (!history) return;
        const { query } = history.location;
        const { redirect } = query;
        history.push(redirect || '/');
        return;
      }

      console.log(msg); // 如果失败去设置用户错误信息

      setUserRegisterState(msg);
    } catch (error) {
      const defaultRegisterFailureMessage = intl.formatMessage({
        id: 'pages.register.failure',
        defaultMessage: '登录失败，请重试！',
      });
      message.error(defaultRegisterFailureMessage);
    }
  };

  const { status, type: RegisterType } = userRegisterState;
  return (
    <div className={styles.container}>
      <div className={styles.lang} data-lang>
        {SelectLang && <SelectLang />}
      </div>
      <div className={styles.content}>
        <LoginForm
          logo={<img alt="logo" src="/logo.svg" />}
          title="Sharing Vision"
          subTitle={intl.formatMessage({
            id: 'pages.layouts.userLayout.title',
          })}
          onFinish={async (values) => {
            console.log(values);
            await handleSubmit(values);
          }}
        >
          <Tabs activeKey={type} onChange={setType}>
            <Tabs.TabPane
              key="account"
              tab={intl.formatMessage({
                id: 'pages.register.accountRegister.tab',
                defaultMessage: '账户密码登录',
              })}
            />
          </Tabs>

          {status === 'error' && RegisterType === 'account' && (
            <RegisterMessage
              content={intl.formatMessage({
                id: 'pages.register.accountRegister.errorMessage',
                defaultMessage: '账户或密码错误(admin/ant.design)',
              })}
            />
          )}
          {type === 'account' && (
            <>
              <ProFormText
                name="username"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: '用户名: admin or user',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="请输入用户名!"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText
                name="username"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: '用户名: admin or user',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="请输入用户名!"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText
                name="username"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: '用户名: admin or user',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="请输入用户名!"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText
                name="username"
                fieldProps={{
                  size: 'large',
                  prefix: <UserOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.username.placeholder',
                  defaultMessage: '用户名: admin or user',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.username.required"
                        defaultMessage="请输入用户名!"
                      />
                    ),
                  },
                ]}
              />
              <ProFormText.Password
                name="password"
                fieldProps={{
                  size: 'large',
                  prefix: <LockOutlined className={styles.prefixIcon} />,
                }}
                placeholder={intl.formatMessage({
                  id: 'pages.register.password.placeholder',
                  defaultMessage: '密码: ant.design',
                })}
                rules={[
                  {
                    required: true,
                    message: (
                      <FormattedMessage
                        id="pages.register.password.required"
                        defaultMessage="请输入密码！"
                      />
                    ),
                  },
                ]}
              />
            </>
          )}

          {status === 'error' && RegisterType === 'mobile' && (
            <RegisterMessage content="验证码错误" />
          )}

          <div
            style={{
              marginBottom: 24,
            }}
          >
            <ProFormCheckbox noStyle name="autoRegister">
              <FormattedMessage id="pages.register.rememberMe" defaultMessage="自动登录" />
            </ProFormCheckbox>
            <a
              style={{
                float: 'right',
              }}
            >
              <FormattedMessage id="pages.register.forgotPassword" defaultMessage="忘记密码" />
            </a>
          </div>
        </LoginForm>
      </div>
      <Footer />
    </div>
  );
};

export default Register;
