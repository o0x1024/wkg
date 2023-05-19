import axiosInstance from './api';

class MiniService {

    GetMiniInfo(page: number, pagesize: number, type: string, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
        return axiosInstance.post('/api/v1/mini/getMiniInfo', data,)
    }


    getMiniInfoByCid(page: number, pagesize: number, keyword: string, cid: number) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword, "cid": cid }
        return axiosInstance.post('/api/v1/mini/getMiniInfoByCid', data,);
    }

    SearchMiniInfo(page: number, pagesize: number, type: string, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
        return axiosInstance.post('/api/v1/mini/searchMiniInfo', data,)
    }


    ReadFlagMiniInfoById(id: string) {
        const data = { "id": id }
        return axiosInstance.post('/api/v1/mini/readFlagMiniInfoById', data,);
    }

    ReadAllFlagMiniInfo() {
        return axiosInstance.get('/api/v1/mini/readAllFlagMiniInfo',)

    }

    DelMiniInfo(id: any) {
        return axiosInstance.get('/api/v1/mini/delMiniById?id=' + id)
    }



}
const miniService = new MiniService()

export default miniService;