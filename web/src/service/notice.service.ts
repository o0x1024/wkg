import axiosInstance from './api';

class NoticeService {


    SaveEditNotice(id: number, cid: number, title: string | undefined, content: string | undefined) {
        const data = { "cid": cid, "id": id, "title": title, "content": content }
        return axiosInstance.post('/api/v1/notice/saveEditNotice', data);

    }

    SaveNotice(cid: number, title: string | undefined, content: string | undefined) {
        const data = { "cid": cid, "title": title, "content": content }
        return axiosInstance.post('/api/v1/notice/saveNotice', data);

    }
    GetNoticeDetailById(id: any) {
        return axiosInstance.post('/api/v1/notice/getNoticeDetailById?id=' + id);
    }

    GetNoticeList(page: number, pagesize: number, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/notice/getNoticeList', data,);
    }

    GetNoticeByKeyword(page: number, pagesize: number, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "keyword": keyword }
        return axiosInstance.post('/api/v1/notice/getNoticeByKeyword', data,);
    }


    DelNoticeInfo(id: any) {
        return axiosInstance.get('/api/v1/notice/delNoticeById?id=' + id)
    }



}
const noticeService = new NoticeService()

export default noticeService;