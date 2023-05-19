import axiosInstance from './api';

class DoaminService {

  getDomainInfo(page: number, pagesize: number, type: string, keyword: string,assetOption:string) {
    const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword ,"assetOption":assetOption}
    return axiosInstance.post('/api/v1/domain/getDomainInfo', data,);
  }

  getDomainInfoByKey(page: number, pagesize: number, type: string, keyword: string) {
    const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
    return axiosInstance.post('/api/v1/domain/getDomainInfoByKey', data,);
  }

  getDomainInfoByCid(page: number, pagesize: number, type: string, keyword: string, cid: number,assetOption:string) {
    const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword, "cid": cid,"assetOption":assetOption }
    return axiosInstance.post('/api/v1/domain/getDomainInfoByCid', data,);
  }



  DelDomainById(id: number) {
    const data = { "id": id }
    return axiosInstance.get('/api/v1/domain/delDomainById?id=' + id);
  }

  ReadFlagDomainInfoById(id: string) {
    const data = { "id": id }
    return axiosInstance.post('/api/v1/domain/readFlagDomainInfoById', data,);
  }

  ReadAllFlagDomainInfo() {
    return axiosInstance.get('/api/v1/domain/readAllFlagDomainInfo',);
  }

}

const domainService = new DoaminService()

export default domainService;