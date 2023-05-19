import axiosInstance from './api';

class IpsService {


    SaveEditIPs(IPsId: number, IPsName: string|undefined, IPsContent: string|undefined) {
        const data = { "IPsId": IPsId, "IPsName": IPsName, "IPsContent": IPsContent }
        return axiosInstance.post('/api/v1/IPs/saveEditIPs', data);

    }

    SaveIPs(IPsName: string|undefined, IPsContent: string|undefined) {
        const data = { "IPsName": IPsName, "IPsContent": IPsContent }
        return axiosInstance.post('/api/v1/IPs/saveIPs', data);

    }

    GetIPsDetailById(id:any) {
        return axiosInstance.get('/api/v1/IPs/getIPsDetailById?id='+id);
    }

    GetIPsInfo(page: number, pagesize: number, keyword: string,assetOpton:string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword ,"assetOpton":assetOpton}
        return axiosInstance.post('/api/v1/IPs/getIPsInfo', data);
    }

    ReadFlagIPInfoById(list:string) {
        const data = { "list": list}
        return axiosInstance.post('/api/v1/IPs/ReadFlagIPInfoById', data);
    }

    ReadAllFlagIPInfo(){
        return axiosInstance.get('/api/v1/IPs/readAllFlagIPInfo');

    }

    getIPsInfoByCId(page: number, pagesize: number, keyword: string, id: any,assetOption:string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword, "id": id ,"assetOption":assetOption}

        return axiosInstance.post('/api/v1/IPs/getIPsInfoByCId' ,data);
    }

    SearchIPInfo(page: number, pagesize: number, keyword: string){
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/IPs/searchIPInfo', data,);
    }

    GetIPsList(page: number, pagesize: number, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/IPs/getNoticeList', data,);
    }

    GetIPsByKeyword(page: number, pagesize: number, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/IPs/getNoticeByKeyword', data,);
    }


    DelIPsInfo(id: any) {
        return axiosInstance.get('/api/v1/IPs/delNoticeById?id=' + id)
    }



}
const ipsService = new IpsService()

export default ipsService;