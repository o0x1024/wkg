import axiosInstance from './api';



class ToolService {
  [x: string]: any;



  SendQuestion(apiKey: string, secretKey: string,question:string) {
    const data = { apiKey: apiKey, secretKey: secretKey,question:question }
    return axiosInstance.post('/api/v2/tool/sendQuestion', data);
  }

}

const toolService = new ToolService()

export default toolService;