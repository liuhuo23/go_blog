import request from '/@/utils/request'

let login = {}
systemInfo.login = function (data) {
    return request({
        url: '/ed/login',
        data,
        method: 'post'
    })
}
export default login