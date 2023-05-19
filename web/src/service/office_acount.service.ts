import axiosInstance from './api';

class OfficeAccoutService {

    GetWOAInfo(page: number, pagesize: number, type: string, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
        return axiosInstance.post('/api/v1/woa/getWOAInfo', data,)
    }


    getWOAInfoByCid(page: number, pagesize: number, keyword: string, cid: number) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword, "cid": cid }
        return axiosInstance.post('/api/v1/woa/getWOAInfoByCid', data,);
    }

    SearchWOAInfo(page: number, pagesize: number, type: string, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
        return axiosInstance.post('/api/v1/woa/searchWOAInfo', data,)
    }


    ReadFlagWOAInfoById(id: string) {
        const data = { "id": id }
        return axiosInstance.post('/api/v1/woa/readFlagWOAInfoById', data,);
    }

    ReadAllFlagWOAInfo() {
        return axiosInstance.get('/api/v1/woa/readAllFlagWOAInfo',)

    }

    DelWOAInfo(id: any) {
        return axiosInstance.get('/api/v1/woa/delWOAById?id=' + id)
    }



}
const woaService = new OfficeAccoutService()

export default woaService;