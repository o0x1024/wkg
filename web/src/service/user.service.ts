import axiosInstance from './api';
import { message } from "ant-design-vue";


class UserService {
  [x: string]: any;


  GetUserById(id: number) {
    return axiosInstance.get('/api/v1/user/getUserById?id=' + id,);
  }

  GetUserList() {
    return axiosInstance.get('/api/v1/user/getUserOptions',);
  }

  SaveUser(username: string, password: string) {
    const data = { "username": username, "password": password }
    return axiosInstance.post('/api/v1/user/saveUser', data);
  }

  ResetPwd(uid: string) {
    return axiosInstance.get('/api/v1/user/resetpwd?id=' + uid,);
  }


  DeleteUserByid(id: number) {
    return axiosInstance.get('/api/v1/user/del?id=' + id,);
  }


  getUserList(page: number, pagesize: number) {
    const data = { "page": page, "pagesize": pagesize }
    return axiosInstance.post('/api/v1/user/list', data,);
  }

}

const userService = new UserService()

export default userService;