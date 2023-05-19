import axiosInstance from './api';
import { message } from "ant-design-vue";

class BugServer {
  [x: string]: any;
  getNewBug(page: number, pagesize: number) {
    const data = { "page": page, "pagesize": pagesize }
    return axiosInstance.post('/api/v1/bug/getNewBug', data);
  }


  searchNewBugReport(page: number, pagesize: number, keyword: string) {
    const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
    return axiosInstance.post('/api/v1/bug/searchNewBugReport', data);
  }

  searchOldBugReport(page: number, pagesize: number, keyword: string) {
    const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
    return axiosInstance.post('/api/v1/bug/searchOldBugReport', data);
  }

  getBugCollect(page: number, pagesize: number) {
    const data = { "page": page, "pagesize": pagesize }
    return axiosInstance.post('/api/v1/bug/getBugCollect', data);
  }

  delBugCollectById(id: number) {
    return axiosInstance.get('/api/v1/bug/delBugCollectById?id=' + id,);
  }

  searchBugCollectByKey(page: number, pagesize: number, keyword: string) {
    const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
    return axiosInstance.post('/api/v1/bug/searchBugCollectByKey', data);
  }

  collect(id: string) {
    return axiosInstance.get('/api/v1/bug/collect?id=' + id,).then((res: any) => {
      if (res.data.code == 400) {
        message.error(res.data.msg);
      } else if (res.data.code == 200) {
        message.error(res.data.msg);

      }
    });
  }

  BatchRead(id: string) {
    const data = { "id": id }
    return axiosInstance.post('/api/v1/bug/batchRead', data,);
  }

  getOldBug(page: number, pagesize: number) {
    const data = { "page": page, "pagesize": pagesize }
    return axiosInstance.post('/api/v1/bug/getOldBug', data);
  }

  read(id: any) {
    axiosInstance.get('/api/v1/bug/read?id=' + id,).then((res: any) => {
      if (res.data.code == 400) {
        message.error(res.data.msg);
      } else if (res.data.code == 200) {
        message.error(res.data.msg);

      }
    });

  }
}

const bugServer = new BugServer()

export default bugServer;