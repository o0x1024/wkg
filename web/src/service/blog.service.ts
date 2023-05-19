import axiosInstance from './api';


class ShareKnowledgeService {

    getKnowledgeByKey(key: string) {
        return axiosInstance.get('/api/v2/blog/getShareknowledgeByKey?key=' + key,)
    }


    getKnowledgeList(pagenum: number, pagesize: number) {
        let data = { "page": pagenum, "pagesize": pagesize }
        return axiosInstance.post('/api/v2/blog/getShareknowledgeList', data)
    }

    searchsharedKnowledgebyWord(pagenum: number, pagesize: number, word: string) {
        let data = { "page": pagenum, "pagesize": pagesize, "keyword": word }
        return axiosInstance.post('/api/v2/blog/searchsharedKnowledgebyWord', data)
    }

    GetsharedKnowledgeListbyTag(pagenum: number, pagesize: number, word: string) {
        let data = { "page": pagenum, "pagesize": pagesize, "keyword": word }
        return axiosInstance.post('/api/v2/blog/getsharedKnowledgeListbyTag', data)
    }

    // getHotList() {
    //     return axiosInstance.get('/api/v2/getHotList')
    // }

    getTags(){
        return axiosInstance.get('/api/v2/blog/getTags',)
      }
    

}
const shareKnowledgeService = new ShareKnowledgeService()

export default shareKnowledgeService;