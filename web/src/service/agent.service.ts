import axiosInstance from './api';
import { message } from "ant-design-vue";


class AgentService {
  [x: string]: any;
  getAgentStatus(page: number, pagesize: number, type: string, keyword: string) {
    const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
    return axiosInstance.get('/api/v1/agent/getlist',);
  }

  async SetAgentGroup(agentId: string, groupId: string) {
    const res = await axiosInstance.get('/api/v1/agent/setAgentGroup?agentId=' + agentId + '&groupId=' + groupId,);
    if (res.status == 200) {
      message.success(res.data.msg)
    } else {
      message.error(res.data.msg)
    }
  }


  GetGroupList() {
    return axiosInstance.get('/api/v1/agent/getGroupOptions',);
  }



  DeleteAgentByid(id: string) {
    return axiosInstance.get('/api/v1/agent/delete?id=' + id,);
  }

  AgentDetail(id: string | string[]) {
    return axiosInstance.get('/api/v1/agent/AgentDetail?agent_id=' + id,);
  }

  GetAgentListByGroupId(page: number, pagesize: number, groupId: string) {
    const data = { "page": page, "pagesize": pagesize, "groupId": groupId }
    return axiosInstance.post('/api/v1/agent/getAgentListByGroupId', data,);
  }


  getAgentList(page: number, pagesize: number) {
    const data = { "page": page, "pagesize": pagesize }
    return axiosInstance.post('/api/v1/agent/list', data,);
  }

}

const agentService = new AgentService()

export default agentService;