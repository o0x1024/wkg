<template>
    <a-row type="flex" justify="start" :gutter="10">
        <a-col :span="4">
            <a-input v-model:value="searchKey" placeholder="标题" />
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
            <a-button @click="onAddNotice" type="primary">
                <template #icon>
                    <read-outlined />
                </template>
                新增笔记
            </a-button>
        </a-col>

    </a-row>
    <a-row type="flex" style="margin-top: 10px;">
        <a-col :span="24" :order="4">
            <a-table :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }" :columns="columns"
                :data-source="noticeList.list" :pagination="Mypagination" :align="dbalign" @change="handleTableChange">
                <template #bodyCell="{ column, text, record }">
                    <template v-if="column.dataIndex === 'title'">
                        <a-button type="link" @click="onView(record.id)">{{ text }}</a-button>


                    </template>
                    <template v-if="column.key === 'operation'">
                        <a-row  type="flex" justify="center">
                            <a-col>
                                <a-button size="small" type="primary" @click="onView(record.id)">查看</a-button>
                            </a-col>
                            <a-col>
                                <a-button size="small" type="primary" @click="onEdit(record.id)">编辑</a-button>
                            </a-col>
                            <a-col>
                                <a-button size="small" type="primary" @click="OnDelete(record.id)" danger>删除</a-button>
                            </a-col>
                        </a-row>
                    </template>
                </template>
            </a-table>
        </a-col>
    </a-row>
</template>



<script lang="ts" >
import { SearchOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { defineComponent, reactive, onMounted, ref, computed, toRefs } from "vue";
import noticeService from '@/service/notice.service';
import srcService from '@/service/src.service'
import types from "@/common/types"
import { useRouter } from "vue-router";
import MdEditor from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { message } from 'ant-design-vue';


interface ServiceData { }
interface option {
    value: string,
    label: string
}

interface Notice {
    id?: number
    cid?: number
    title?: string
    content?: string
    updateTime?: string
}


type Key = string | number;

export default defineComponent({
    components: { SearchOutlined, ReadOutlined, MdEditor },
    setup() {

        const noticeList: { list: Notice[] } = reactive({ list: [] });
        let total = ref(1)
        let curPage = ref(1)
        let viewNotice: Notice = reactive({})
        let pageSize = ref(10)
        let columns = types.getNoticeTableColumns()
        const rolesList: { list: ServiceData[] } = reactive({ list: [] });
        let searchKey = ref('')
        let cmpSelectVal = ref('')
        let curSelectVal = ref('')
        const router = useRouter();
        let drawerVisible = ref(false)
        const CmpSelectOptions: { list: option[] } = reactive({ list: [] })

        const hasSelected = computed(() => state.selectedRowKeys.length > 0);
        const state = reactive<{
            selectedRowKeys: Key[];
            loading: boolean;
        }>({
            selectedRowKeys: [], // Check here to configure the default column
            loading: false,
        });



        const onDrawerEdit = () => {
            drawerVisible.value = false
        }

        const onClose = () => {
            drawerVisible.value = false
        }


        const Mypagination = computed(() => ({
            total: total.value,
            current: curPage.value,
            pageSize: pageSize.value,
            showTotal: () => `总共 ${total.value} 项`,
            defaultPageSize: 10,
            pageSizeOptions: ['10', '20', '50', '100'], // å
        }))


        const handleTableChange = (pag: any, filters: any, sorter: any) => {
            state.selectedRowKeys = []
            curPage.value = pag.current;
            pageSize.value = pag.pageSize;
            
            noticeService.GetNoticeList(curPage.value, pageSize.value, "").then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    noticeList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        };

        const onSelectChange = (selectedRowKeys: Key[], selectedRows: Array<ServiceData>) => {

            // delIdList.clear()
            // for (let row of selectedRows) {
            //   delIdList.set(row.id, row.id)
            // }
            // state.selectedRowKeys = selectedRowKeys;
        };

        const onAddNotice = () => {
            router.push({ name: "newNotice", params: { cid: curSelectVal.value } });

        }

        const selectHandleChange = (value: any) => {
            curSelectVal.value = value
            noticeService.GetNoticeList(curPage.value, pageSize.value, "").then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    noticeList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }

        const onView = (id: number) => {
            router.push({ name: "viewNotice", params: { id: id } });

        }

        const onDrawerVisible = () => {

        }

        const onEdit = (id: number) => {
            router.push({ name: "editNotice", params: { id: id, cid: curSelectVal.value } });
        }


        const OnSearch = () => {
            curPage.value = 1
            pageSize.value = 10
            noticeService.GetNoticeByKeyword(curPage.value, pageSize.value, searchKey.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    noticeList.list = (res.data.data)
                    total.value = res.data.total
                }
            })
        }

        const OnDelete = (id: number) => {
            noticeService.DelNoticeInfo(id).then((res: any) => {
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

            noticeService.GetNoticeList(curPage.value, pageSize.value, "").then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    noticeList.list = (res.data.data)
                    total.value = res.data.total
                }
            })


        }

        const filterOption = (input: string, option: any) => {
            return option.label.indexOf(input.toLowerCase()) >= 0;
        };
        return {
            rolesList,
            searchKey,
            columns,
            onAddNotice,
            onClose,
            curSelectVal,
            onView,
            filterOption,
            cmpSelectVal,
            Mypagination,
            onDrawerVisible,
            drawerVisible,
            noticeList,
            viewNotice,
            onDrawerEdit,
            CmpSelectOptions,
            onEdit,
            dbalign: ref('center'),
            hasSelected,
            selectHandleChange,
            onSelectChange,
            handleTableChange,
            OnDelete,
            OnSearch,
            ...toRefs(state),
        };
    }
});



</script>

<style>
.top {
    padding: 10px;
    background: #fff;
}
</style>