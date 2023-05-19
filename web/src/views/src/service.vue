<script lang="ts">
import { SearchOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { defineComponent, computed, toRefs, reactive, onMounted, ref } from "vue";
import svcService from "@/service/svc.service";
import types from "../../common/types"
import { message } from 'ant-design-vue';

interface ServiceData {
    id: number
    cid: number
    service: string
    host: string
    port: string
    product: string
    updateTime: string
    isNew: boolean
}


type Key = string | number;

export default defineComponent({
    components: { SearchOutlined, ReadOutlined },
    setup() {

        let total = ref(1)
        let searchKey = ref('')
        let curPage = ref(1)
        let columns = types.getServiceTableColumns()
        let assetOption = ref('0')
        let pageSize = ref(10)
        const rolesList: { list: ServiceData[] } = reactive({ list: [] });
        let selectVal = ref('service')
        let delIdList = new Map();


        const options = [
            {
                value: 'host',
                label: '主机/IP',
            },
            {
                value: 'port',
                label: '端口',
            },
            {
                value: 'service',
                label: '服务名',
            },
        ];

        const hasSelected = computed(() => state.selectedRowKeys.length > 0);
        const state = reactive<{
            selectedRowKeys: Key[];
            loading: boolean;
        }>({
            selectedRowKeys: [], // Check here to configure the default column
            loading: false,
        });


        const Mypagination = computed(() => ({
            total: total.value,
            current: curPage.value,
            pageSize: pageSize.value,
            showTotal: () => `总共 ${total.value} 项`,
            defaultPageSize: 10,
            pageSizeOptions: ['10', '20', '50', '100'], // 可不设置使用默认
            showSizeChanger: true, // 是否显示pagesize选择
            showQuickJumper: true, // 是否显示跳转窗

        }));



        const handleTableChange = (pag: any, filters: any, sorter: any) => {
            state.selectedRowKeys = []
            curPage.value = pag.current;
            pageSize.value = pag.pageSize;
            svcService.GetServiceInfo(curPage.value, pageSize.value, selectVal.value, searchKey.value, assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    alert(res.data.data)
                } else if (res.data.code == 200) {
                    rolesList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        };


        const onSelectChange = (selectedRowKeys: Key[], selectedRows: Array<ServiceData>) => {
            delIdList.clear()
            for (let row of selectedRows) {
                delIdList.set(row.id, row.id)
            }
            state.selectedRowKeys = selectedRowKeys;
        };

        const selectHandleChange = () => {
            console.log()
        }
        const OnScan = (record: ServiceData) => {
            console.log()
        }



        const OnAllHasRead = () => {

            let delist = ''
            for (let value of delIdList.values()) {
                delist += value + ','
            }

            svcService.ReadAllFlagServiceInfo().then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    message.success(res.data.msg)
                    InitData()
                }
            })
        }

        const OnHasRead = () => {
            let delist = ''
            for (let value of delIdList.values()) {
                delist += value + ','
            }
            svcService.ReadFlagServiceInfoById(delist).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    message.success(res.data.msg)
                    InitData()
                }
            })

        }

        const OnSearch = () => {
            curPage.value = 1
            pageSize.value = 10
            svcService.GetServiceInfo(curPage.value, pageSize.value, selectVal.value, searchKey.value, assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    rolesList.list = (res.data.data)
                    total.value = res.data.total

                }
            })
        }

        const onRadioChange = () => {
            curPage.value = 1
            pageSize.value = 10
            svcService.GetServiceInfo(curPage.value, pageSize.value, selectVal.value, searchKey.value, assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    rolesList.list = (res.data.data)
                    total.value = res.data.total

                }
            })
        }



        const OnDelete = (id: number) => {
            svcService.DelServiceInfo(id).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    message.success(res.data.msg)
                    InitData()
                }
            })
        }


        onMounted(() => {
            InitData()
        });

        const InitData = () => {
            svcService.GetServiceInfo(curPage.value, pageSize.value, searchKey.value, searchKey.value, assetOption.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    rolesList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }
        return {
            rolesList,
            searchKey,
            columns,
            selectVal,
            Mypagination,
            options,
            OnHasRead,
            dbalign: ref('center'),
            OnScan,
            onRadioChange,
            assetOption,
            hasSelected,
            OnAllHasRead,
            selectHandleChange,
            onSelectChange,
            handleTableChange,
            OnDelete,
            OnSearch,
            ...toRefs(state),
        };
    },
});
</script>

<style>
.ant-table-thead>tr>th {
    text-align: center;
}

.top {
    padding: 10px;
    background: #fff;
}
</style>