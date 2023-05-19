import axiosInstance from './api';

class WebsiteService {
  GetWebSiteInfo(page: number, pagesize: number, type: string, keyword: string,assetOption:string) {
    const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword,"assetOption":assetOption }
    return axiosInstance.post('/api/v1/website/getWebSiteInfo', data,)
  }

  GetWebSiteInfoByCid(page: number, pagesize: number, type: string, keyword: string, cid: number,assetOption:string) {
    const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword, "cid": cid ,"assetOption":assetOption}
    return axiosInstance.post('/api/v1/website/getWebSiteInfoByCid', data,)
  }

  DelWebSiteInfo(id: any) {
    return axiosInstance.get('/api/v1/website/delWebSiteById?id=' + id)
  }

  ReadFlagWebSiteInfoById(id: string) {
    const data = { "id": id }
    return axiosInstance.post('/api/v1/website/readFlagWebSitefoById', data,);
  }



  ReadAllFlagWebSiteInfo() {
    return axiosInstance.get('/api/v1/website/readAllFlagWebSiteInfo',)

  }

  ScanNew() {
    return axiosInstance.get('/api/v1/website/scanNew',)
  }



}

const websiteService = new WebsiteService()

export default websiteService;