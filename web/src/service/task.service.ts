import axiosInstance from './api';

class TaskService {
    NewTask(taskName: string,taskStatus:boolean, taskType: string, rate: number,companyRange: string,targetValue:string[],pocName:string,assetRange: string,domainCollectionType: string,portRange:string) {
        let cids:string = '';
        let pocs:string = '';
        targetValue.forEach(element => {
            cids = cids + element +","
        });
        if(taskType == "1"){
            domainCollectionType = "10000"
        }
        const data = { "taskName": taskName,"taskStatus":taskStatus, "taskType": taskType, "rate": rate, "companyRange": companyRange, "companyId": cids,"pocName":pocName, "assetRange": assetRange, "domainCollectionType": domainCollectionType, "portRange": portRange }
        return axiosInstance.post('/api/v1/task/new', data,)
    }


    UpdateTask(taskId:string,taskName: string,taskStatus:boolean, taskType: string, rate: number,companyRange: string,targetValue:string[],pocName:string,assetRange: string,domainCollectionType: string,portRange:string) {
        let cids:string = '';
        let pocs:string = '';
        targetValue.forEach(element => {
            cids = cids + element +","
        });
        if(taskType == "1"){
            domainCollectionType = "10000"
        }
        const data = {"taskId":taskId, "taskName": taskName,"taskStatus":taskStatus, "taskType": taskType, "rate": rate, "companyRange": companyRange, "companyId": cids,"pocName":pocName, "assetRange": assetRange, "domainCollectionType": domainCollectionType, "portRange": portRange }
        return axiosInstance.post('/api/v1/task/update', data,)
    }

    getCompIds() {
        return axiosInstance.get('/api/v1/task/getCompIds');
    }

    TaksInfo(taskid:string) {
        return axiosInstance.get('/api/v1//task/TaskInfo?taskId='+taskid)
    }

    getTaskList(page: number, pagesize: number) {
        const data = { "page": page, "pagesize": pagesize }
        return axiosInstance.post('/api/v1//task/getTaskList', data,)
    }

    DelTaskByTaskId(taskId: string) {
        return axiosInstance.get('/api/v1//task/delTaskByTaskId?taskId='+taskId)
    }

    clsTaskQueue() {
        return axiosInstance.get('/api/v1//task/clsTaskQueue')
    }

    

    stopTask(tid: string) {
        return axiosInstance.get('/api/v1//task/stopTask?taskId='+tid)
    }

    runTask(tid:string){
        return axiosInstance.get('/api/v1//task/runTask?taskId='+tid)

    }


    getPocs(){
        return axiosInstance.get('/api/v1/task/getPocs');
    }

}
const taskService = new TaskService()

export default taskService;