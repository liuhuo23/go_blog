import httpRequest from '@/request/index'

export function apiGetuserInfo(userid, username){
    return httpRequest({
        url: '',
        method: 'post',
        data:{
            username: username,
            userid:userid
        }
    })
}