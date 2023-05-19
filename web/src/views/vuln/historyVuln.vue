/* eslint-disable @typescript-eslint/no-explicit-any */
<template>
    <a-row type="flex" justify="start" :gutter="16">
        <a-col :span="6">
            <a-input v-model:value="searchKey" placeholder />
        </a-col>

        <a-col :span="2">
            <a-button @click="OnSearch" type="primary">
                <template #icon>
                    <SearchOutlined />
                </template>
                搜索
            </a-button>
        </a-col>
        <a-col :span="1">
            <a-button @click="OnBatchRead" type="primary">
                <template #icon>
                    <plus-circle-outlined />
                </template>
                批量已阅
            </a-button>
        </a-col>
    </a-row>

    <a-row type="flex" style="margin-top: 15px">
        <a-col :span="24" :order="4">
            <a-table sticky :row-selection="{
                selectedRowKeys: selectedRowKeys,
                onChange: onSelectChange,
            }" :columns="columns" :data-source="rolesList.list" :align="dbalign"
                :rowKey="(record: any, index: any) => index" @change="handleTableChange" :scroll="{ x: 1500, y: 1500 }"
                :pagination="Mypagination">
                <template #bodyCell="{ column, text, record }">
                    <template v-if="column.key === 'operation'">
                        <a-row type="flex" justify="center">
                            <a-col :span="6">
                                <a-button size="small" type="primary" @click="OnRead(record.id)">已阅</a-button>
                            </a-col>

                            <a-col :span="6">
                                <a-button size="small" type="primary" @click="OnCollectx(record.id)">收藏</a-button>
                            </a-col>

                            <a-col :span="6">
                                <a-popconfirm title="你确定？" ok-text="Yes" cancel-text="No"
                                    @confirm="OnDelete(record.id)">
                                    <a-button size="small" type="primary" danger>删除</a-button>
                                </a-popconfirm>
                            </a-col>
                        </a-row>
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
        </a-col>
    </a-row>
</template>
<script lang="ts">
import { SearchOutlined, PlusCircleOutlined } from "@ant-design/icons-vue";
import {
    defineComponent,
    reactive,
    computed,
    onMounted,
    ref,
    toRefs,
} from "vue";
import bugService from "../../service/bug.service";
import { useRouter } from "vue-router";
import types from "../../common/types";
import { message } from "ant-design-vue"

interface bug {
    id: number;
    detail: string;
    updateTime: string;
    isNew: boolean;
}

type Key = string | number;

export default defineComponent({
    components: { SearchOutlined, PlusCircleOutlined },
    setup() {
        const rolesList: { list: bug[] } = reactive({ list: [] });
        let searchKey = ref("");
        let total = ref(1);
        let curPage = ref(1);
        let pageSize = ref(10);
        let columns = types.getBugReportTableColumns();
        let selectVal = ref("companyName");
        let batch = ref(false);
        let ScanVisible = ref(false)
        const router = useRouter();
        let curCid = 0;


        const state = reactive<{
            selectedRowKeys: Key[],
            loading: boolean,
            checkedList: string[],
        }>({
            selectedRowKeys: [], // Check here to configure the default column
            loading: false,
            checkedList: ['站点扫描', '域名扫描'],
        });

        const hasSelected = computed(() => state.selectedRowKeys.length > 0);

        const selectHandleChange = () => {
            let a= 1
        };

        const OnDelete = (cid: number) => {
            bugService.DelCompanyByid(cid).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    message.success(res.data.msg);
                    InitData();
                }
            });
        };

        const onSelectChange = (selectedRowKeys: Key[]) => {
            state.selectedRowKeys = selectedRowKeys;
        };


        const OnCollectx = (id: any) => {
            bugService.collect(id)
        }
        const OnSearch = () => {
            bugService
                .searchOldBugReport(
                    curPage.value,
                    pageSize.value,
                    searchKey.value
                )
                .then((res: any) => {
                    if (res.data.code == 400) {
                        message.error(res.data.msg);
                    } else if (res.data.code == 200) {
                        rolesList.list = (res.data.data);
                        total.value = res.data.total;
                    }
                });
        };

        const OnRead = (id: any) => {
            bugService.read(id)

            InitData();
        }

        const OnBatchRead = () => {
            console.log(1)
        }

        const handleTableChange = (pag: any) => {
            state.selectedRowKeys = [];
            curPage.value = pag.current;
            pageSize.value = pag.pageSize;
            bugService
                .searchOldBugReport(curPage.value, pageSize.value, searchKey.value)
                .then((res: any) => {
                    if (res.data.code == 400) {
                        message.error(res.data.msg);
                    } else if (res.data.code == 200) {
                        rolesList.list = (res.data.data);
                        total.value = res.data.total;
                    }
                });
        };

        onMounted(() => {
            InitData();
        });

        const InitData = () => {
            bugService
                .getOldBug(curPage.value, pageSize.value)
                .then((res: any) => {
                    if (res.data.code == 400) {
                        message.error(res.data.msg);
                    } else if (res.data.code == 200) {
                        rolesList.list = (res.data.data);
                        total.value = res.data.total;
                    }
                });
        };

        const Mypagination = computed(() => ({
            total: total.value,
            current: curPage.value,
            pageSize: pageSize.value,
            showTotal: () => `总共 ${total.value} 项`,
            defaultPageSize: 10,
            pageSizeOptions: ["5", "10", "20", "50"], // 可不设置使用默认
            showSizeChanger: true, // 是否显示pagesize选择
            showQuickJumper: true, // 是否显示跳转窗
        }));

        return {
            rolesList,
            searchKey,
            dbalign: ref("center"),
            columns,
            ...toRefs(state),
            OnSearch,
            OnRead,
            OnCollectx,
            selectVal,
            OnBatchRead,
            Mypagination,
            selectHandleChange,
            OnDelete,
            onSelectChange,
            handleTableChange,
        };
    },
});
</script>

<!-- <style>
.ant-table-thead > tr > th {
  text-align: center;
}
</style> -->
