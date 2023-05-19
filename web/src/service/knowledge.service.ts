import axiosInstance from './api';
import { message } from "ant-design-vue";


class KnowledgeService {

  getTopCategories() {
    return axiosInstance.get('/api/v1/knowledge/getTopCategories',);
  }

  getSecondCategories(key: string) {
    return axiosInstance.get('/api/v1/knowledge/getSecondCategories?key=' + key,);
  }

  getKnowledgeCategories(key: string) {
    return axiosInstance.get('/api/v1/knowledge/getKnowledgeCategories?key=' + key,);
  }

  GetKnowledgeByKey(key: string) {
    return axiosInstance.get('/api/v1/knowledge/getKnowledgeByKey?key=' + key,);
  }


  searchsharedKnowledgebyWord(pagenum: number, pagesize: number, word: string) {
    let data = { "page": pagenum, "pagesize": pagesize, "keyword": word }
    return axiosInstance.post('/api/v1/knowledge/searchsharedKnowledgebyWord', data)
  }

  getTopSelectOption() {
    return axiosInstance.get('/api/v1/knowledge/getTopSelectOption',);
  }

  getSecondSelectOption(key: string) {
    return axiosInstance.get('/api/v1/knowledge/getSecodSelectOption?key=' + key,);
  }


  getSummary() {
    return axiosInstance.get('/api/v1/knowledge/getSummary')
  }

  getTree() {
    return axiosInstance.get('/api/v1/knowledge/getTree')
  }

  async UploadImg(form: any) {
    const res = await axiosInstance.post('/api/v1/knowledge/uploadImage', form, {
      headers: {
        'Content-Type': 'multipart/form-data'
      }
    })
    return res
  }


  cancelShare(id: number) {
    return axiosInstance.get('/api/v1/knowledge/cancelShare?id=' + id)
  }

  cancelShareByKey(key: string) {
    return axiosInstance.get('/api/v1/knowledge/cancelShareByKey?key=' + key)
  }

  batchCancelShare(idList: string) {
    let data = { "idList": idList }
    return axiosInstance.post('/api/v1/knowledge/batchCancelShare', data)
  }


  DelTag(id: any) {
    return axiosInstance.get('/api/v1/knowledge/deltag?id=' + id)
  }

  AddTag(key: String) {
    return axiosInstance.get('/api/v1/knowledge/addTag?key=' + key)
  }



  SaveEditKnowledge(category: string, tags: string[], parentId: number, title: string, content: string, key: string) {
    const data = { "category": category, "tags": tags, "parentId": parentId, "title": title, "content": content, "key": key }
    return axiosInstance.post('/api/v1/knowledge/saveEditKnowledge', data,)
  }

  SaveNewKnowledge( tags: string[],  parentId: number, title: string, content: string) {
    const data = { "tags": tags, "parentId": parentId, "title": title, "content": content }
    return axiosInstance.post('/api/v1/knowledge/saveNewKnowledge', data,)
  }


  AddNode(parentId: number, topNode: string) {
    return axiosInstance.get('/api/v1/knowledge/addNode?nodeName=' + topNode + '&parentId=' + parentId)
  }

  AddRootNode(topNode: string) {
    return axiosInstance.get('/api/v1/knowledge/addNode?nodeName=' + topNode + '&parentId=0' )
  }
  
  DelTreeNode(isLeaf: boolean, id: number) {
    const data = { isLeaf: isLeaf, id: id }
    return axiosInstance.post('/api/v1/knowledge/delTreeNode', data,)
  }

  ModifyTreeNode(isLeaf:boolean,id: number, keyword: string) {
    const data = {"isLeaf":isLeaf, "id": id, "keyword": keyword }
    return axiosInstance.post('/api/v1/knowledge/modifyTreeNode', data,)

  }

  ShareKnowledge(key: string) {
    return axiosInstance.get('/api/v1/knowledge/shareknowledge?key=' + key,)
  }

  getShareknowledgeInBackend(pagenum: number, pagesize: number, keyword: string) {
    let data = { "page": pagenum, "pagesize": pagesize, "keyword": keyword }
    return axiosInstance.post('/api/v1/knowledge/getShareknowledgeInBackend', data)
  }
}

const knowledgeService = new KnowledgeService()

export default knowledgeService;