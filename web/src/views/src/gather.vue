<template>
    <a-drawer v-model:visible="drawerVisible" class="custom-class" style="color: red" title="资产导入" width="520"
        placement="right">
        <!-- @after-visible-change="afterVisibleChange" -->
        <a-row :gutter="20">
            <a-col :span="10">
                <a-select ref="select" v-model:value="InputSelectVal" :options="InputSelectOptions.list"
                    style="width: 150px" @change="selectHandleChange"></a-select>
            </a-col>
            <a-col :span="12">
                <a-button style="width: 100px;" @click="OnImportData" type="primary">导入</a-button>
            </a-col>
        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col>
                <a-radio-group v-model:value="radioGroup" button-style="solid">
                    <a-radio-button value="domain">域名</a-radio-button>
                    <a-radio-button value="website">站点</a-radio-button>
                    <a-radio-button value="ip">ip</a-radio-button>
                    <a-radio-button value="service">服务</a-radio-button>
                    <a-radio-button value="miniProgram">小程序</a-radio-button>
                    <a-radio-button value="webchatAccout">公众号</a-radio-button>
                    <a-radio-button value="news">资讯</a-radio-button>
                </a-radio-group>
            </a-col>
        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col :span="24">
                <a-textarea v-model:value="InputText" :placeholder="radioGroup" :rows="30" />
            </a-col>
        </a-row>
    </a-drawer>


    <a-row type="flex" justify="start" :gutter="10">
        <a-col :span="4">
            <a-input v-model:value="searchKey" placeholder="关键字" />
        </a-col>
        <a-col v-if="typeSelectenable">
            <a-select ref="select" v-model:value="typeSelectVal" :options="typeOptions.list" style="width: 120px"
                @change="selectHandleChange"></a-select>
        </a-col>

        <a-col>
            <a-select ref="select" show-search :filter-option="filterOption" v-model:value="selectVal"
                :options="selectOptions.list" style="width: 150px" @change="selectHandleChange"></a-select>
        </a-col>
        <a-col>
            <a-button @click="OnSearch" type="primary">
                <template #icon>
                    <SearchOutlined />
                </template>
                搜索
            </a-button>
        </a-col>
        <a-col>
            <a-button @click="OnExport" type="primary">
                <template #icon>
                    <plus-circle-outlined />
                </template>
                资产导出
            </a-button>
        </a-col>
        <a-col>
            <a-button @click="OnImport" type="primary">
                <template #icon>
                    <plus-circle-outlined />
                </template>
                资产导入
            </a-button>
        </a-col>
        <a-col>
            <a-radio-group v-model:value="assetOption" @change="onRadioChange()">
                <a-radio-button value="0">全部</a-radio-button>
                <a-radio-button value="1">新增</a-radio-button>
            </a-radio-group>
        </a-col>
    </a-row>

    <a-row type="flex" style="margin-top: 10px;">
        <a-col :span="24" :order="4">
            <a-tabs v-model:activeKey="activeKey" @change="tableChange">
                <a-tab-pane :forceRender="false" key="1" tab="站点">
                    <a-table sticky :columns="websitecolumns" :data-source="WebSiteList.list" :align="dbalign"
                        :rowKey="(record: any, index: any) => index" @change="handleTableChange"
                        :scroll="{ x: 1500, y: 500 }" :pagination="Mypagination">
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'website'">
                                <a :href="text" target="_blank">{{ text }}</a>
                            </template>
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="2" tab="域名">
                    <a-table sticky :columns="domaincolumns" :data-source="DomainList.list" :align="dbalign"
                        :rowKey="(record: any, index: any) => index" @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1000 }" :pagination="Mypagination">
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                            <template v-if="column.key === 'operation'"></template>
                        </template>
                    </a-table>
                </a-tab-pane>

                <a-tab-pane :forceRender="false" key="3" tab="IP">
                    <a-table sticky :columns="ipcolumns" :data-source="IPsList.list" :align="dbalign"
                        :rowKey="(record: any, index: any) => index" @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }" :pagination="Mypagination">
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="4" tab="服务">
                    <a-table sticky :columns="serviceColumns" :data-source="ServiceList.list" :align="dbalign"
                        :rowKey="(record: any, index: any) => index" @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }" :pagination="Mypagination">
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="5" tab="小程序">
                    <a-table sticky :columns="MiniProgramcolumns" :data-source="MiniProgramList.list" :align="dbalign"
                        :rowKey="(record: any, index: any) => index" @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }" :pagination="Mypagination">
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                            <template v-if="column.key === 'operation'"></template>
                        </template>
                    </a-table>
                </a-tab-pane>
                <a-tab-pane :forceRender="false" key="6" tab="微信公众号">
                    <a-table sticky :columns="WOAColumns" :data-source="WOAList.list" :align="dbalign"
                        :rowKey="(record: any, index: any) => index" @change="handleTableChange"
                        :scroll="{ x: 1500, y: 1500 }" :pagination="Mypagination">
                        <template #bodyCell="{ column, text }">
                            <template v-if="column.dataIndex === 'isNew'">
                                <div v-if="text">
                                    <a-tag color="#87d068">Y</a-tag>
                                </div>
                                <div v-else>
                                    <a-tag color="#f50">N</a-tag>
                                </div>
                            </template>
                        </template>
                    </a-table>
                </a-tab-pane>
            </a-tabs>
        </a-col>
    </a-row>
</template>
<script lang="ts">
import { EditTwoTone, SearchOutlined, PlusCircleOutlined } from '@ant-design/icons-vue';
import { defineComponent, reactive, computed, onMounted, ref, toRefs } from "vue";
import srcService from '../../service/src.service';
import { useRouter } from 'vue-router';
import types from "../../common/types"
import { message } from 'ant-design-vue';
import domainService from '@/service/domain.service';
import websiteService from '@/service/website.service';
import ipsSvc from '@/service/ips.service';
import myoption from '../../common/options'
import axios from "axios";
import svcService from '@/service/svc.service';
import miniService from '@/service/mini_program.service';
import woaService from '@/service/office_acount.service';

interface Company {
    id: number;
    projectType: string;
    companyName: string;
    domain: string;
    SrcUrl: string;
    keyWord: string;
    monitorStatus: boolean;
    monitorRate: number;
    lastUpdateTime: string;
}
interface option {
    value: string,
    label: string
}

type Key = string | number;

export default defineComponent({
    components: { PlusCircleOutlined, SearchOutlined },
    setup() {
        const DomainList: { list: Company[] } = reactive({ list: [] });
        const WebSiteList: { list: Company[] } = reactive({ list: [] });
        const IPsList: { list: Company[] } = reactive({ list: [] });
        const ServiceList: { list: Company[] } = reactive({ list: [] });
        const MiniProgramList: { list: Company[] } = reactive({ list: [] });
        const NewsList: { list: Company[] } = reactive({ list: [] });
        const WOAList: { list: Company[] } = reactive({ list: [] });

        let radioGroup = ref('domain')
        let typeSelectVal = ref('')
        let typeSelectenable = ref(true)
        let searchKey = ref('')
        let total = ref(1)
        let assetOption = ref('0')
        let InputText = ref('')
        let drawerVisible = ref(false)

        let curPage = ref(1)
        let pageSize = ref(10)
        let domaincolumns = types.getGatherDomainTableColumns()
        let websitecolumns = types.getGatherWebsiteTableColumns()
        let ipcolumns = types.getGatherIPsTableColumns()
        let serviceColumns = types.getGatherServiceTableColumns()

        let MiniProgramcolumns = types.getMiniTableColumns()
        let WOAColumns = types.getWechatOfficeAccountTableColumns()
        // let serviceColumns = types.getGatherServiceTableColumns()
        let columns = [{}]
        let selectVal = ref('')
        let InputSelectVal = ref('')
        let activeKey = ref('1')
        // const router = useRouter();
        const selectOptions: { list: option[] } = reactive({ list: [] })
        const InputSelectOptions: { list: option[] } = reactive({ list: [] })

        const typeOptions: { list: option[] } = reactive({ list: [] })
        const websiteOptions = [
            {
                value: 'ips',
                label: 'ips',
            },
            {
                value: 'website',
                label: '网站',
            },
            {
                value: 'title',
                label: '标题',
            },
            {
                value: 'favicon',
                label: 'favicon',
            }
        ];


        const OnExport = () => {
            srcService.ExportByCid(parseInt(selectVal.value)).then((res: any) => {
                if (res.status == 200) {
                    //获取文件名，需要在后端进行配置
                    let filename = res.headers['content-disposition'].split('filename=')[1].split('; filename')[0]

                    let type = res.headers['content-type'].split(';')[0]
                    let blob = new Blob([res.data], { type: type })
                    const a = document.createElement('a')

                    // 创建URL
                    const blobUrl = window.URL.createObjectURL(blob)
                    a.download = filename
                    a.href = blobUrl
                    console.log(a)
                    document.body.appendChild(a)
                    // 下载文件
                    a.click()
                    // 释放内存
                    URL.revokeObjectURL(blobUrl)
                    document.body.removeChild(a)
                } else {
                    message.error(res.data.msg)
                }
            })
        }

        const OnImportData = () => {
            srcService.InputData(InputSelectVal.value, radioGroup.value, InputText.value)
            InputText.value = ''
            InitData()
            drawerVisible.value = false
        }

        const OnImport = () => {
            drawerVisible.value = true
            srcService.GetSelectOption().then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    InputSelectOptions.list = (res.data.data)
                    InputSelectVal.value = selectOptions.list[0].value
                }
            })

        }


        onMounted(() => {
            InitData()

        })
        const InitData = () => {
            typeOptions.list = websiteOptions
            typeSelectVal.value = websiteOptions[2].value
            srcService.GetSelectOption().then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    selectOptions.list = (res.data.data)
                    selectVal.value = selectOptions.list[0].value
                    // console.Glog(selectOptions)
                    websiteService.GetWebSiteInfoByCid(curPage.value, pageSize.value, typeSelectVal.value, "", parseInt(selectVal.value), assetOption.value).then((res: any) => {
                        if (res.data.code == 400) {
                            message.error(res.data.msg)
                        } else if (res.data.code == 200) {
                            WebSiteList.list = (res.data.data)
                            total.value = res.data.total
                        }
                    })
                }
            })

        }

        const getDomainInfoByCid = () => {
            domainService.getDomainInfoByCid(curPage.value, pageSize.value, "domain", searchKey.value, parseInt(selectVal.value), assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    DomainList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }

        const getWebsiteInfoByCid = () => {
            websiteService.GetWebSiteInfoByCid(curPage.value, pageSize.value, typeSelectVal.value, searchKey.value, parseInt(selectVal.value), assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    WebSiteList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }

        const getIPsInfoByCid = () => {
            ipsSvc.getIPsInfoByCId(curPage.value, pageSize.value, searchKey.value, parseInt(selectVal.value), assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    IPsList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }

        const getServiceInfoByCid = () => {
            svcService.getServiceInfoByCid(curPage.value, pageSize.value, searchKey.value, parseInt(selectVal.value), assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    ServiceList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }

        const getMiniProgramByCid = () => {
            miniService.getMiniInfoByCid(curPage.value, pageSize.value, searchKey.value, parseInt(selectVal.value)).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    MiniProgramList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }
        const getWOAByCid = () => {
            woaService.getWOAInfoByCid(curPage.value, pageSize.value, searchKey.value, parseInt(selectVal.value)).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    WOAList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }

        const onRadioChange = () => {
            console.log(1111)
            UpdateData()
        }

        const UpdateData = () => {
            switch (activeKey.value) {
                case '1'://域名
                    getWebsiteInfoByCid()
                    break;
                case '2'://website
                    getDomainInfoByCid()
                    break;
                case '3'://IPS
                    getIPsInfoByCid()
                    break;
                case '4':
                    getServiceInfoByCid()
                    break;
                case '5':
                    getMiniProgramByCid()
                    break;
                case '6':
                    getWOAByCid()
                    break;
                case '7':
                // getNoticeByCid()
            }
        }


        //tabls切换事件
        const tableChange = (activeKey: string) => {
            curPage.value = 1
            pageSize.value = 10
            switch (activeKey) {
                case '1':
                    typeOptions.list = websiteOptions
                    typeSelectVal.value = websiteOptions[2].value
                    typeSelectenable.value = true
                    break;
                case '2':
                    typeSelectenable.value = false
                    break;
                default:
                    typeSelectenable.value = false

            }
            UpdateData()
        }

        const OnSearch = () => {
            curPage.value = 1
            pageSize.value = 10
            UpdateData()
        }

        //table 翻页切换事件
        const handleTableChange = (pag: any, filters: any, sorter: any) => {
            curPage.value = pag.current;
            pageSize.value = pag.pageSize;
            UpdateData()
            // srcService.getCompanyInfo(curPage.value, pageSize.value).then((res: any) => {
            //     if (res.data.code == 400) {
            //         alert(res.data.data)
            //     } else if (res.data.code == 200) {
            //         // rolesList.list = (res.data.data)
            //         total.value = res.data.total
            //     }
            // })
        }

        //select切换事件
        const selectHandleChange = () => {
            curPage.value = 1
            pageSize.value = 10
            UpdateData()
        }


        const filterOption = (input: string, option: any) => {
            return option.label.indexOf(input.toLowerCase()) >= 0;
        };

        const Mypagination = computed(() => ({
            total: total.value,
            current: curPage.value,
            pageSize: pageSize.value,
            showTotal: () => `总共 ${total.value} 项`,
            defaultPageSize: 10,
            pageSizeOptions: ['5', '10', '20', '50'], // 可不设置使用默认
            showSizeChanger: true, // 是否显示pagesize选择
            showQuickJumper: true, // 是否显示跳转窗
        }));

        return {
            MiniProgramcolumns,
            WOAColumns,
            DomainList,
            WebSiteList,
            assetOption,
            InputSelectVal,
            IPsList,
            ServiceList,
            onRadioChange,
            MiniProgramList,
            NewsList,
            InputSelectOptions,
            WOAList,
            searchKey,
            radioGroup,
            InputText,
            dbalign: ref('center'),
            domaincolumns,
            websitecolumns,
            ipcolumns,
            columns,
            serviceColumns,
            typeSelectenable,
            drawerVisible,
            typeOptions,
            typeSelectVal,
            activeKey,
            selectOptions,
            selectVal,
            Mypagination,
            OnSearch,
            filterOption,
            OnExport,
            OnImport,
            OnImportData,
            tableChange,
            selectHandleChange,
            handleTableChange,
        };
    },
});
</script>

<style>
.top {
    padding: 10px;
    background: #fff;
}
</style>