import axiosInstance from './api';

class DashService {


    GetDashboard() {
        return axiosInstance.get('/api/v1/getDashBoard');

    }



}
const dashService = new DashService()

export default dashService;