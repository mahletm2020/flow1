import axios from 'axios'

export const useAxios = ()  =>{
    const api = axios.create({
        baseURL:'htp://localhost:8080',
        // withcredentials:true.
    })
    return api
}