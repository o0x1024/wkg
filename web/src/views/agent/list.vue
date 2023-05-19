<template>
    <a-row type="flex" justify="start" :gutter="16">
        <a-col :span="3">
            <a-input v-model:value="searchKey" placeholder="agent id" />
        </a-col>

        <a-col :span="2">
            <a-button @click="OnSearch" type="primary">
                <template #icon>
                    <SearchOutlined />
                </template>
                搜索
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
                            <a-col :span="1"></a-col>
                            <a-col :span="8">
                                <a @click="OnDetail(record.agent_id)">查看详情</a>
                            </a-col>
                            <a-col :span="7">
                                <a-popconfirm title="你确定？" ok-text="Yes" cancel-text="No"
                                    @confirm="OnDelete(record.agent_id)">
                                    <a>删除</a>
                                </a-popconfirm>
                            </a-col>
                        </a-row>
                    </template>
                    <template v-if="column.dataIndex === 'source_usage'">
                        <a-row justify="center" :gutter="20">
                            <a-col>
                                cpu:{{ record.cpu_usage }}
                            </a-col>
                            <a-col>
                                mem:{{ record.mem_usage }}
                            </a-col>
                        </a-row>
                    </template>

                    <template v-if="column.dataIndex === 'status'">
                        <div v-if="text == 'running'">
                            <a-row justify="center">
                                <a-col>
                                    <div
                                        style="width: 10px;height: 10px; margin-top: 6px; background-color: green; border-radius: 50%;">
                                    </div>
                                </a-col>
                                <a-col>
                                    运行中
                                </a-col>
                            </a-row>
                        </div>
                        <div v-else>
                            <a-row justify="center">
                                <a-col>
                                    <div
                                        style="width: 10px;height: 10px; margin-top: 6px; background-color: red; border-radius: 50%;">
                                    </div>
                                </a-col>
                                <a-col>
                                    离线
                                </a-col>
                            </a-row>
                        </div>
                    </template>
                </template>
            </a-table>
        </a-col>
    </a-row>
</template>
<script lang="ts">
import { SearchOutlined } from "@ant-design/icons-vue";
import {
    defineComponent,
    reactive,
    computed,
    onMounted,
    ref,
    toRefs,
} from "vue";
import types from "../../common/types"
import agentService from "../../service/agent.service"
import { message } from "ant-design-vue";
import { useRouter } from "vue-router";



interface option {
    value: string,
    label: string
}


interface Agent {
    agent_id: string
    host_name: string
    platform: string
    status: string
    cpu_usage: string
    mem_usage: string
    update_time: string
}

type Key = string | number;

export default defineComponent({
    components: { SearchOutlined },
    setup() {
        const rolesList: { list: Agent[] } = reactive({ list: [] });
        let searchKey = ref("");
        let total = ref(1);
        let curPage = ref(1);
        let pageSize = ref(10);
        let visible = ref(false)
        let columns = types.getAgentColumns();
        let selectVal = ref("projectName");
        const router = useRouter();


        const state = reactive<{
            selectedRowKeys: Key[];
            loading: boolean;
        }>({
            selectedRowKeys: [], // Check here to configure the default column
            loading: false,
        });

        const hasSelected = computed(() => state.selectedRowKeys.length > 0);
        const options = [
            {
                value: "hostName",
                label: "主机名",
            },
            {
                value: "projectName",
                label: "应用名称",
            },
            {
                value: "serverAddr",
                label: "服务器地址",
            },
        ];

        const selectHandleChange = () => {
            console.log(1);
        };




        const OnDelete = (cid: string) => {
            agentService.DeleteAgentByid(cid).then((res: any) => {
                if (res.status == 400) {
                    message.error(res.data.msg);
                } else if (res.status == 200) {
                    message.success(res.data.msg);
                    InitData();
                }
            });
        };

        const onSelectChange = (selectedRowKeys: Key[]) => {
            state.selectedRowKeys = selectedRowKeys;
        };

        const OnSearch = () => {
            agentService
                .SearchAgentInfo(
                    curPage.value,
                    pageSize.value,
                    selectVal.value,
                    searchKey.value
                )
                .then((res: any) => {
                    if (res.status == 400) {
                        alert(res.data.msg);
                    } else if (res.status == 200) {
                        rolesList.list = res.data.msg;
                        total.value = res.data.total;
                    }
                });
        };

        const handleTableChange = (pag: any, filters: any, sorter: any) => {
            state.selectedRowKeys = [];
            curPage.value = pag.current;
            pageSize.value = pag.pageSize;
            agentService
                .getAgentList(curPage.value, pageSize.value)
                .then((res: any) => {
                    if (res.status == 400) {
                        alert(res.data.msg);
                    } else if (res.status == 200) {
                        rolesList.list = (res.data.msg);
                        total.value = res.data.total;
                    }
                });
        };

        onMounted(() => {
            InitData();
        });

        const InitData = () => {
            agentService
                .getAgentList(curPage.value, pageSize.value)
                .then((res: any) => {
                    if (res.status == 400) {
                        alert(res.data.msg);
                    } else if (res.status == 200) {
                        rolesList.list = res.data.data;
                        total.value = res.data.total;
                    }
                });
        };

        const OnDetail = (agentId: string) => {
            router.push({ name: "AgentDetail", params: { agentId: agentId } });

        }




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
            options,
            OnDetail,
            ...toRefs(state),
            OnSearch,
            selectVal,
            visible,
            Mypagination,
            selectHandleChange,
            OnDelete,
            onSelectChange,
            handleTableChange,
        };
    },
});
</script>

