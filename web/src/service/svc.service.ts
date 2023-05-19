import axiosInstance from './api';

class SvsService {

    GetServiceInfo(page: number, pagesize: number, type: string, keyword: string, assetOption:string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword,"assetOption" :assetOption}
        return axiosInstance.post('/api/v1/service/getServiceInfo', data,)
    }


    getServiceInfoByCid(page: number, pagesize: number, keyword: string, cid: number, assetOption:string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword, "cid": cid ,"assetOption" :assetOption}
        return axiosInstance.post('/api/v1/service/getServiceInfoByCid', data,);
    }

    SearchServiceInfo(page: number, pagesize: number, type: string, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
        return axiosInstance.post('/api/v1/service/searchServiceInfo', data,)
    }


    ReadFlagServiceInfoById(id: string) {
        const data = { "id": id }
        return axiosInstance.post('/api/v1/service/readFlagServiceInfoById', data,);
    }

    ReadAllFlagServiceInfo() {
        return axiosInstance.get('/api/v1/service/readAllFlagServiceInfo',)

    }

    DelServiceInfo(id: any) {
        return axiosInstance.get('/api/v1/service/delServiceById?id=' + id)
    }



}
const svcService = new SvsService()

export default svcService;