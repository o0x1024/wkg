import { AntdIconType } from '@ant-design/icons-vue/lib/components/AntdIcon';
import axiosInstance from './api';
import axios from "axios";
import { message } from "ant-design-vue";

class SrcService {
    getCompanyInfo(page: number, pagesize: number) {
        const data = { "page": page, "pagesize": pagesize }
        return axiosInstance.post('/api/v1/company/getCompanyInfo', data,)
    }

    searchCompanyInfo(page: number, pagesize: number, type: string, keyword: string) {
        const data = { "page": page, "pagesize": pagesize, "type": type, "keyword": keyword }
        return axiosInstance.post('/api/v1/company/getCompanyInfo', data,)
    }

    getCompanyInfoByCId(cid: any) {
        // console.log(cid)
        return axiosInstance.get('/api/v1/company/getCompanyByCId?cid=' + cid)
    }

    UpdateComapny(cid: any, ctype: any, companyName: any, rootDomain: any, srcUrl: any, keyWord: any, monitorStauts: any, monitorRate: any) {
        const data = {
            "id": cid.value, "projectType": ctype.value, "companyName": companyName.value, "domain": rootDomain.value,
            "srcUrl": srcUrl.value, "keyWord": keyWord.value, "monitorStatus": monitorStauts.value,
            "monitorRate": monitorRate.value, "vulnScanStatus": false, "vulnScanRate": 24
        }

        return axiosInstance.post('/api/v1/company/updateCompanyInfo', data,)
    }

    NewComapny(ctype: any, companyName: any, rootDomain: any, srcUrl: any, keyWord: any, monitorStauts: any, monitorRate: any) {
        const data = {
            "projectType": ctype.value, "companyName": companyName.value, "domain": rootDomain.value,
            "srcUrl": srcUrl.value, "keyWord": keyWord.value, "monitorStatus": monitorStauts.value,
            "monitorRate": monitorRate.value, "vulnScanStatus": false, "vulnScanRate": 24
        }

        return axiosInstance.post('/api/v1/company/newCompany', data,)
    }

    DelCompanyByCId(cid: any) {
        return axiosInstance.get('/api/v1/company/delCompanyByCId?cid=' + cid)
    }

    ScanCompanyDomain(id: any) {
        const data = { "id": id }
        return axiosInstance.post('/api/v1/company/scanCompanyInfo', data,)
    }

    GetSelectOption() {
        return axiosInstance.get('/api/v1/company/getSelectOption',)
    }

    Scan(id: any, scanType: string) {
        const data = { "id": id, "scanType": scanType }
        return axiosInstance.post('/api/v1/scan', data)

    }


    ExportByCid(cid: number) {
        const data = { "id": cid }
        return axiosInstance.post('/api/v1/company/export', data, { responseType: 'blob' })
    }

    InputData(companyid: string, dataType: string, content: string) {
        let contents = content.split("\n")
        let cid = parseInt(companyid)
        const data = { "companyId": cid, "dataType": dataType, "content": contents }
        axiosInstance.post('/api/v1/company/import', data,).then((res: any) => {
            if (res.status == 200) {
                message.success(res.data.msg)
            } else {
                message.error(res.data.msg)

            }
        })
    }


}
const srcSvc = new SrcService()

export default srcSvc;