<template>
  <a-row type="flex" justify="start" :gutter="10">
    <a-col :span="4">
      <a-input v-model:value="searchKey" placeholder="关键字" />
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
      <a-button :disabled="!hasSelected" @click="OnHasRead" type="primary">
        <template #icon>
          <read-outlined />
        </template>
        已阅
      </a-button>
    </a-col>
    <a-col>
      <a-button @click="OnAllHasRead" type="primary">
        <template #icon>
          <read-outlined />
        </template>
        全部已阅
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
      <a-table :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }" :columns="columns"
        :rowKey="(record: any, index: any) => index" :data-source="rolesList.list" :pagination="Mypagination"
        :align="dbalign" @change="handleTableChange">
        <template #bodyCell="{ column, text, record }">
          <template v-if="column.key === 'operation'">
            <a-row>
              <a-col :span="8">
                <a-button size="small" type="primary" @click="OnScan(record.ip)">扫描</a-button>
              </a-col>
              <a-col :span="8">
                <a-button size="small" type="primary" @click="OnDelete(record.id)" danger>删除</a-button>
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
import { SearchOutlined, ReadOutlined, PlusCircleOutlined } from '@ant-design/icons-vue';
import { defineComponent, computed, toRefs, reactive, onMounted, ref } from "vue";
import ipsSvc from "@/service/ips.service";
import types from "../../common/types"
import { message } from 'ant-design-vue';

interface ipsData {
  id: number
  cid: number
  ip: string
  os: string
  indomains: string
  geo: string
  updateTime: string
  isNew: boolean
}


type Key = string | number;

export default defineComponent({
  components: { SearchOutlined, ReadOutlined },
  setup() {
    let total = ref(1)
    let curPage = ref(1)
    let pageSize = ref(10)
    let assetOption = ref('0')
    let columns = types.getIPsTableColumns()
    let delIdList = new Map();
    const rolesList: { list: ipsData[] } = reactive({ list: [] });
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
      curPage.value = pag.current;
      pageSize.value = pag.pageSize;
      ipsSvc.GetIPsInfo(curPage.value, pageSize.value, searchKey.value, assetOption.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    };


    const onSelectChange = (selectedRowKeys: Key[], selectedRows: Array<ipsData>) => {
      delIdList.clear()
      for (let row of selectedRows) {
        delIdList.set(row.id, row.id)
      }
      state.selectedRowKeys = selectedRowKeys;
    };

    const selectHandleChange = () => {
      console.log()
    }
    const OnScan = (ip: string) => {
      console.log()
    }


    const OnAllHasRead = () => {

      ipsSvc.ReadAllFlagIPInfo().then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
        }
      })
    }

    const OnHasRead = () => {
      let delist = ''
      for (let value of delIdList.values()) {
        delist += value + ','
      }

      ipsSvc.ReadFlagIPInfoById(delist).then((res: any) => {
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
      ipsSvc.SearchIPInfo(curPage.value, pageSize.value, searchKey.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    }


    const OnDelete = (id: number) => {
      ipsSvc.DelIPsInfo(id).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
          InitData()
        }
      })
    }

    const onRadioChange = () => {
      ipsSvc.GetIPsInfo(curPage.value, pageSize.value, searchKey.value, assetOption.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    }
    

    onMounted(() => {
      InitData()
    });

    const InitData = () => {
      ipsSvc.GetIPsInfo(curPage.value, pageSize.value, searchKey.value, assetOption.value).then((res: any) => {
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
      dbalign: ref('center'),
      OnHasRead,
      OnAllHasRead,
      OnScan,
      onRadioChange,
      selectHandleChange,
      onSelectChange,
      handleTableChange,
      OnDelete,
      OnSearch,
      assetOption,
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