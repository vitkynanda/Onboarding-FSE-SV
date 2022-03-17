import jwt_decode from 'jwt-decode';
const users = () => {
  const token = localStorage.getItem('token');
  const name = localStorage.getItem('username');

  try {
    if (token !== '') {
      const decode = jwt_decode(`${token}`);
      console.log(decode);
      return {
        isLogin: true,
        userId: decode?.user_id,
        role: decode?.role,
        name,
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
