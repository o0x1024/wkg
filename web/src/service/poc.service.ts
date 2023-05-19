import axiosInstance from './api';

class PocService {


    SaveEditPoc(pocId: number, pocName: string|undefined, pocContent: string|undefined) {
        const data = { "pocId": pocId, "pocName": pocName, "pocContent": pocContent }
        return axiosInstance.post('/api/v1/poc/saveEditPoc', data);

    }

    SavePoc(pocName: string|undefined, pocContent: string|undefined) {
        const data = { "pocName": pocName, "pocContent": pocContent }
        return axiosInstance.post('/api/v1/poc/savePoc', data);

    }

    GetPocDetailById(id:any) {
        return axiosInstance.get('/api/v1/poc/getPocDetailById?id='+id);
    }

    GetPocInfo(page: number, pagesize: number, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/poc/getPocInfo?', data);
    }

    GetPocInfoByCId(page: number, pagesize: number, keyword: string, id: any) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword, "id": id }

        return axiosInstance.post('/api/v1/poc/getPocInfoByCId?' ,data);
    }

    GetPocList(page: number, pagesize: number, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/poc/getNoticeList', data,);
    }

    GetPocByKeyword(page: number, pagesize: number, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/poc/getNoticeByKeyword', data,);
    }


    DelPocInfo(id: any) {
        return axiosInstance.get('/api/v1/poc/delNoticeById?id=' + id)
    }



}
const podService = new PocService()

export default podService;