import axios from 'axios'

const service = axios.create({
    baseURL: '/',
    timeout: 9000
})
export default service