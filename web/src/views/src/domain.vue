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

    <!-- <a-col :span="3">
      <a-button @click="OnExport" type="primary">
        <template #icon>
          <plus-circle-outlined />
        </template>
        资产导出
      </a-button>
    </a-col>-->
  </a-row>
  <a-row type="flex" style="margin-top: 10px;">
    <a-col :span="24" :order="4">
      <a-table :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }" :columns="columns"
        :rowKey="(record: any, index: any) => index" :data-source="rolesList.list" :pagination="Mypagination"
        :align="dbalign" @change="handleTableChange">
        <template #bodyCell="{ column, text, record }">
          <template v-if="column.key === 'operation'">
            <a-row type="flex" justify="center">
              <a-col :span="8">
                <a-button size="small" type="primary" @click="OnScan(record.domain)">扫描</a-button>
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
import domainService from "../../service/domain.service";
import types from "../../common/types"
import { message } from 'ant-design-vue';

interface domainData {
  id: number;
  cid: number;
  domain: string;
  title: string;
  type: string;
  ip: string;
  source: string;
  updateTime: string;
  isNew: boolean;
}

type Key = string | number;

export default defineComponent({
  components: { SearchOutlined, ReadOutlined },
  setup() {
    let total = ref(1)
    let curPage = ref(1)
    let assetOption = ref('0')
    let pageSize = ref(10)
    let columns = types.getDomainTableColumns()
    let delIdList = new Map();
    const rolesList: { list: domainData[] } = reactive({ list: [] });
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
      domainService.getDomainInfo(curPage.value, pageSize.value, "domain", searchKey.value, assetOption.value).then((res: any) => {
        if (res.data.code == 400) {
          alert(res.data.data)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    };


    const onSelectChange = (selectedRowKeys: Key[], selectedRows: Array<domainData>) => {
      delIdList.clear()
      for (let row of selectedRows) {
        delIdList.set(row.id, row.id)
      }
      state.selectedRowKeys = selectedRowKeys;
    };



    const OnAllHasRead = () => {

      domainService.ReadAllFlagDomainInfo().then((res: any) => {
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
      domainService.ReadFlagDomainInfoById(delist).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
          InitData()
        }
      })
    }

    const OnSearch = () => {
      domainService.getDomainInfoByKey(curPage.value, pageSize.value, "domain", searchKey.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    }

    const OnScan = (domain: string) => {
      domainService.getDomainInfoByKey(curPage.value, pageSize.value, "domain", searchKey.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    }



    const OnDelete = (id: number) => {
      domainService.DelDomainById(id).then((res: any) => {
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
      domainService.getDomainInfo(curPage.value, pageSize.value, "domain", searchKey.value, assetOption.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          rolesList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    }


    const onRadioChange = () => {
      domainService.getDomainInfo(curPage.value, pageSize.value, "domain", searchKey.value, assetOption.value).then((res: any) => {
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
      onRadioChange,
      hasSelected,
      Mypagination,
      dbalign: ref('center'),
      OnHasRead,
      OnAllHasRead,
      onSelectChange,
      handleTableChange,
      assetOption,
      OnScan,
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