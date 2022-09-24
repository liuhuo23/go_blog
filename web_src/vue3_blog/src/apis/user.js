import httpRequest from '../request/index'

export function apiGetuserInfo(userid, username){
    return httpRequest({
        url: '/noapi/ping',
        method: 'post',
        data:{
            username: username,
            userid:userid
        }
    })
}