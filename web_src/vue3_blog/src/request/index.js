import axios from "axios";

const service = axios.create({
    baseURL: '/1api', //所有的请求地址前缀部分
    timeout: 60000, //超时时间毫秒
    withCredentials: true, //异步请求携带cookie
    headers: {
        'Content-Type': 'application/json',
        'token': 'token',
        'X-Requested-With': 'XMLHttpRequest',
    },
})

//添加请求拦截器
service.interceptors.request.use(
    function (config) {
        //在发送请求之前
        return config
    },
    function (error) {
        // 对请求错误之前
        return Promise.reject(error)
    }

)
//添加响应拦截器
service.interceptors.response.use(
    function (response){
        let data = response.data
        if(data === null || data === undefined){
            data={
                code:-1,
                message:"请求无数据"
            }
        }
        return data
    },
    function (error) {
        // 超出 2xx 范围的状态码都会触发该函数
        return Promise.reject(error)
    }
)

export default service