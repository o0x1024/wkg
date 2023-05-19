import axios from "axios";
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";
import router from "../router"




const axiosInstance = axios.create({
    baseURL: process.env.VUE_APP_BASE_URL,
    headers: {
        "Content-Type": "application/json",
    },

});


axiosInstance.interceptors.request.use((req: any) => {
    if (localStorage.token) {
        const token = localStorage.getItem('token')
        // console.Glog('goke', token)
        req.headers['token'] = token;
    }
    return req;
},
    (error) => {
        return Promise.reject(error);
    }
);

axiosInstance.interceptors.response.use(
    (response) => {
        if (response.status == 403) {
            router.push("/login")
        }
        return response;
    },
    // 请求错误
    (error) => {
        if (error.response.status == 403) {
            // alert("token过期,请重新登录!");
            localStorage.removeItem("token");
            router.push("/login")
            // router.push("/login");
            // message.error("会话过期")
            return

        } else {
            // alert(error.response.data);
        }
        return Promise.reject(error);
    }
);

export default axiosInstance;