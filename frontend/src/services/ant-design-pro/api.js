// @ts-ignore

/* eslint-disable */
import { request } from 'umi';

//SVI

export async function loginSv(payload) {
  return request(`http://localhost:8001/login`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/' },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function register(payload) {
  return request(`http://localhost:8001/users`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function getRole() {
  return request(`http://localhost:8001/roles`, {
    method: 'GET',
  });
}

export async function getUserList() {
  return request(`http://localhost:8001/users`, {
    method: 'GET',
  });
}

export async function userDetail(id) {
  return request(`http://localhost:8001/users/${id}`, {
    method: 'GET',
  });
}

export async function updateUser(id, payload) {
  return request(`http://localhost:8001/users/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', Authorization: localStorage.getItem('token') },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function removeUser(id) {
  return request(`http://localhost:8001/users/${id}`, {
    method: 'DELETE',
    headers: { 'Content-Type': 'application/json', Authorization: localStorage.getItem('token') },
    skipErrorHandler: true,
  });
}

export async function getProductList() {
  return request(`http://localhost:8001/products`, {
    method: 'GET',
  });
}

export async function productDetail(id) {
  return request(`http://localhost:8001/products/${id}`, {
    method: 'GET',
  });
}

export async function createNewProduct(payload) {
  return request(`http://localhost:8001/products`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json', Authorization: localStorage.getItem('token') },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function updateProduct(id, payload) {
  return request(`http://localhost:8001/products/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json', Authorization: localStorage.getItem('token') },
    body: JSON.stringify(payload),
    skipErrorHandler: true,
  });
}

export async function removeProduct(id) {
  return request(`http://localhost:8001/products/${id}`, {
    method: 'DELETE',
    headers: { 'Content-Type': 'application/json', Authorization: localStorage.getItem('token') },
    skipErrorHandler: true,
  });
}

/** 获取当前的用户 GET /api/currentUser */

export async function currentUser(options) {
  return request('/api/currentUser', {
    method: 'GET',
    ...(options || {}),
  });
}
/** 退出登录接口 POST /api/login/outLogin */

export async function outLogin(options) {
  return request('/api/login/outLogin', {
    method: 'POST',
    ...(options || {}),
  });
}
/** 登录接口 POST /api/login/account */

export async function login(body, options) {
  return request('/api/login/account', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}
/** 此处后端没有提供注释 GET /api/notices */

export async function getNotices(options) {
  return request('/api/notices', {
    method: 'GET',
    ...(options || {}),
  });
}
/** 获取规则列表 GET /api/rule */

export async function rule(params, options) {
  return request('/api/rule', {
    method: 'GET',
    params: { ...params },
    ...(options || {}),
  });
}
/** 新建规则 PUT /api/rule */

export async function updateRule(options) {
  return request('/api/rule', {
    method: 'PUT',
    ...(options || {}),
  });
}
/** 新建规则 POST /api/rule */

export async function addRule(options) {
  return request('/api/rule', {
    method: 'POST',
    ...(options || {}),
  });
}
/** 删除规则 DELETE /api/rule */

export async function removeRule(options) {
  return request('/api/rule', {
    method: 'DELETE',
    ...(options || {}),
  });
}
