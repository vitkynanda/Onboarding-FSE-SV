import jwt_decode from 'jwt-decode';
const users = () => {
  const token = localStorage.getItem('token');
  try {
    if (token !== '') {
      const decode = jwt_decode(token);
      return {
        isLogin: true,
        userId: decode?.data?.user_id,
        role: decode?.data?.role,
      };
    } else {
      throw 'not login';
    }
  } catch (e) {
    return {
      isLogin: false,
      userId: null,
      role: null,
    };
  }
};

export default users;
