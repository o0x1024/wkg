<template>
    <a-row type="flex" justify="start" :gutter="16">
        <a-col :span="3">
            <a-select
                ref="select"
                v-model:value="Groupval"
                :options="groupOptions.list"
                style="width: 120px"
                @change="selectHandleChange"
            ></a-select>
        </a-col>
    </a-row>

    <a-row type="flex" style="margin-top: 15px">
        <a-col :span="24" :order="4">
            <a-table
                sticky
                :row-selection="{
                    selectedRowKeys: selectedRowKeys,
                    onChange: onSelectChange,
                }"
                :columns="columns"
                :data-source="rolesList.list"
                :align="dbalign"
                :rowKey="(record: any, index: any) => index"
                @change="handleTableChange"
                :scroll="{ x: 1500, y: 1500 }"
                :pagination="Mypagination"
            >
                <template #bodyCell="{ column, text, record }">
                    <template v-if="column.key === 'operation'">
                        <a-row type="flex" justify="center">
                            <a-col :span="8">
                                <a-modal
                                    v-model:visible="visible"
                                    title="添加黑名单IP"
                                    @ok="handleOk(record.id)"
                                >
                                    <a-row type="flex" align="middle" :gutter="24">
                                        <a-col>团队</a-col>
                                        <a-col>
                                            <a-select
                                                ref="select"
                                                v-model:value="Groupval"
                                                :options="groupOptions.list"
                                                style="width: 120px"
                                            ></a-select>
                                        </a-col>
                                    </a-row>
                                </a-modal>
                                <a @click="OnSetGroup()">设置团队</a>
                            </a-col>
                            <a-col :span="1"></a-col>
                            <a-col :span="8">
                                <a @click="OnDetail(record.id)">查看详情</a>
                            </a-col>
                            <a-col :span="7">
                                <a-popconfirm
                                    title="你确定？"
                                    ok-text="Yes"
                                    cancel-text="No"
                                    @confirm="OnDelete(record.id)"
                                >
                                    <a>删除</a>
                                </a-popconfirm>
                            </a-col>
                        </a-row>
                    </template>
                    <template v-if="column.dataIndex === 'groupName'">
                        <div v-if="text">
                            <a-tag color="#6495ED">{{ text }}</a-tag>
                        </div>
                    </template>
                    <template v-if="column.dataIndex === 'online'">
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
    Id: string
    Version: string
    GroupId: number
    GroupName: string
    HostName: string
    ServerAddr: string
    ServerPort: string
    ServerPath: string
    ServerEnv: string
    Disk: string
    Network: string
    Memory: string
    Cpu: string
    Online: boolean
    Pid: string
    LastUpdateTime: string
}

type Key = string | number;

export default defineComponent({
    setup() {
        const rolesList: { list: Agent[] } = reactive({ list: [] });
        let total = ref(1);
        let curPage = ref(1);
        let pageSize = ref(10);
        let visible = ref(false)
        let Groupval = ref('')
        let columns = types.getAgentColumns();
        const groupOptions: { list: option[] } = reactive({ list: [] });
        const router = useRouter();


        const state = reactive<{
            selectedRowKeys: Key[];
            loading: boolean;
        }>({
            selectedRowKeys: [], // Check here to configure the default column
            loading: false,
        });


        const handleOk = (agentId: string) => {
            agentService.SetAgentGroup(agentId, Groupval.value)
            visible.value = false
            InitData()

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


        const handleTableChange = (pag: any) => {
            state.selectedRowKeys = [];
            curPage.value = pag.current;
            pageSize.value = pag.pageSize;
            agentService
                .getAgentList(curPage.value, pageSize.value)
                .then((res: any) => {
                    if (res.status == 400) {
                        alert(res.data.msg);
                    } else if (res.status == 200) {
                        rolesList.list = res.data.msg;
                        total.value = res.data.total;
                    }
                });
        };

        onMounted(() => {
            InitData();
        });

        const InitData = () => {
            agentService.GetGroupList().then((res: any) => {
                if (res.status == 200) {
                    groupOptions.list = res.data.data
                    Groupval.value = groupOptions.list[0].value
                    agentService
                        .GetAgentListByGroupId(curPage.value, pageSize.value, Groupval.value)
                        .then((res: any) => {
                            if (res.status == 400) {
                                alert(res.data.msg);
                            } else if (res.status == 200) {
                                rolesList.list = res.data.data;
                                total.value = res.data.total;
                            }
                        });
                } else {
                    message.error(res.data.msg)
                }
            })

        };

        const selectHandleChange = (value: string) => {
            agentService
                .GetAgentListByGroupId(curPage.value, pageSize.value, Groupval.value)
                .then((res: any) => {
                    if (res.status == 400) {
                        alert(res.data.msg);
                    } else if (res.status == 200) {
                        rolesList.list = res.data.data;
                        total.value = res.data.total;
                    }
                });
        };

        const OnDetail = (id: string) => {
            router.push({ name: "AgentDetail", params: { id: id } });

        }

        const OnSetGroup = () => {
            agentService.GetGroupList().then((res: any) => {
                if (res.status == 200) {
                    groupOptions.list = res.data.data
                    Groupval.value = groupOptions.list[0].value
                } else {
                    message.error(res.data.msg)
                }
            })
            visible.value = true
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
            dbalign: ref("center"),
            columns,
            handleOk,
            groupOptions,
            OnDetail,
            ...toRefs(state),
            Groupval,
            visible,
            selectHandleChange,
            Mypagination,
            OnSetGroup,
            OnDelete,
            onSelectChange,
            handleTableChange,
        };
    },
});
</script>
