<template>
  <div class="sharePanel">
    <a-row type="flex" justify="start" :gutter="16" style="padding: 10px;">
      <a-col :span="6">
        <a-input v-model:value="searchKey" placeholder="标题" />
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
        <a-button @click="OnCancelSelect" type="primary" :disabled="!IsSelect">
          <template #icon>
            <read-outlined />
          </template>
          批量取消分享
        </a-button>
      </a-col>
    </a-row>
    <a-row type="flex" style="margin-top: 10px;">
      <a-col :span="24" :order="4">
        <a-table
          :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }"
          :columns="columns"
          :rowKey="(record: any, index: any) => index"
          :data-source="rolesList.list"
          :pagination="Mypagination"
          :align="dbalign"
          @change="handleTableChange"
        >
          <template #bodyCell="{ column, text, record }">
            <template v-if="column.key === 'operation'">
              <a-row type="flex" justify="center">
                <a-col :span="8">
                  <a-button
                    size="small"
                    type="primary"
                    @click="OnCancelShare(record.id)"
                    danger
                  >取消分享</a-button>
                </a-col>
              </a-row>
            </template>
          </template>
        </a-table>
      </a-col>
    </a-row>
  </div>
</template>



<script lang="ts">


import { SearchOutlined, ReadOutlined, PlusCircleOutlined } from '@ant-design/icons-vue';
import { defineComponent, computed, toRefs, reactive, onMounted, ref } from "vue";
import knowledgeService from "../../service/knowledge.service"
import types from "../../common/types"
import { message } from 'ant-design-vue';

interface ShareData {
  id: number
  title: string
  content: string
  key: string
  updateTime: string
}

type Key = string | number;

export default defineComponent({
  components: { SearchOutlined, ReadOutlined },
  setup() {
    let total = ref(1)
    let curPage = ref(1)
    let pageSize = ref(10)
    let columns = types.getShareTableColumns()
    let delIdList = new Map();
    let IsSelect = ref(false)
    const rolesList: { list: ShareData[] } = reactive({ list: [] });
    let searchKey = ref('')

    const state = reactive<{
      selectedRowKeys: Key[];
      loading: boolean;
    }>({
      selectedRowKeys: [], // Check here to configure the default column
      loading: false,
    });

    const hasSelected = computed(() => state.selectedRowKeys.length > 0);

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
      knowledgeService.getShareknowledgeInBackend(curPage.value, pageSize.value, searchKey.value).then((res: any) => {
        if (res.data.code == 400) {
          alert(res.data.data)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    };


    const onSelectChange = (selectedRowKeys: Key[], selectedRows: Array<ShareData>) => {
      delIdList.clear()

      if (selectedRowKeys.length == 0) {
        IsSelect.value = false
      } else {
        IsSelect.value = true
      }

      for (let row of selectedRows) {
        delIdList.set(row.id, row.id)
      }
      state.selectedRowKeys = selectedRowKeys;
    };



    const OnCancelSelect = () => {
      let delist = ''
      for (let value of delIdList.values()) {
        delist += value + ','
      }

      knowledgeService.batchCancelShare(delist).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
          InitData()
        }
      })
      state.selectedRowKeys = []
    }



    const OnSearch = () => {
      knowledgeService.searchsharedKnowledgebyWord(curPage.value, pageSize.value, searchKey.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    }


    const OnCancelShare = (id: number) => {
      knowledgeService.cancelShare(id).then((res: any) => {
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
      knowledgeService.getShareknowledgeInBackend(curPage.value, pageSize.value, "").then((res: any) => {
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
      hasSelected,
      Mypagination,
      IsSelect,
      dbalign: ref('center'),
      OnCancelSelect,
      onSelectChange,
      handleTableChange,
      OnCancelShare,
      OnSearch,
      ...toRefs(state),
    };
  },
});
</script>

<style>
.ant-table-thead > tr > th {
  text-align: center;
}

.top {
  padding: 10px;
  background: #fff;
}

.sharePanel {
  background-color: #fff;
}
</style>