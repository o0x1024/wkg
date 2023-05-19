<template>
  <a-row type="flex" justify="start" :gutter="16">
    <a-col :span="5">
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
    <a-col>
      <a-row>
        <a-col :span="12">
          <a-button @click="onAddUser" type="primary">
            <template #icon>
              <read-outlined />
            </template>
            新增用户
          </a-button>
        </a-col>
      </a-row>
    </a-col>
  </a-row>
  <a-row type="flex" style="margin-top: 10px;">
    <a-col :span="24" :order="4">
      <a-table :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }" :columns="columns"
        :data-source="userList.list" :pagination="Mypagination" :align="dbalign" @change="handleTableChange">
        <template #bodyCell="{ column, text, record }">
          <template v-if="column.dataIndex === 'website'">
            <a :href="text" target="_blank">{{ text }}</a>
          </template>
          <template v-if="column.key === 'operation'">
            <a-row type="flex" justify="center">
              <a-col :span="8">
                <a-button size="small" type="primary" @click="onResetPwd(record.id)">重置密码</a-button>
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

import { SearchOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { defineComponent, computed, toRefs, reactive, onMounted, ref, onActivated } from "vue";
import types from "../../common/types"
import userService from '../../service/user.service';
import MdEditor from 'md-editor-v3';
import { useRouter } from "vue-router";
import { message } from 'ant-design-vue';

interface user {
  id?: number
  username?: number
  password?: string
  createTime?: string
}


interface ServiceData { }

interface option {
  value: string,
  label: string
}

type Key = string | number;

export default defineComponent({
  components: { SearchOutlined, ReadOutlined, MdEditor },

  beforeRouteEnter(to, from, next) {
    // to.meta.keepAlive = false
    from.meta.keeplive = false;
    next()
  },
  beforeRouteLeave(to, from, next) {
    to.meta.keepAlive = false
    next()
  },

  activated() {
    if (!this.$route.meta.isBack) {

    }
    // 从新设置页面得路由元信息
    this.$route.meta.isBack = false
  },
  setup() {
    let total = ref(1)
    let curPage = ref(1)
    let pageSize = ref(10)
    let columns = types.getUserTableColumns()
    let delIdList = new Map();
    const rolesList: { list: ServiceData[] } = reactive({ list: [] });
    let searchKey = ref('')
    let cmpSelectVal = ref('')
    let curSelectVal = ref('')
    const router = useRouter();
    let viewNotice: user = reactive({})
    let drawerVisible = ref(false)
    const CmpSelectOptions: { list: option[] } = reactive({ list: [] })
    const userList: { list: user[] } = reactive({ list: [] })


    // router.beforeEach((to, from ,next) => {
    //   if (from.name === 'newNotice' || from.name === 'editNotice') {
    //       router.meta
    //   }
    //   next()
    // })
    const filterOption = (input: string, option: any) => {
      return option.label.indexOf(input.toLowerCase()) >= 0;
    };

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
      pageSizeOptions: ['10', '20', '50', '100'], // 可不设置使用默认
      showSizeChanger: true, // 是否显示pagesize选择
      showQuickJumper: true, // 是否显示跳转窗

    }));



    const handleTableChange = (pag: any, filters: any, sorter: any) => {
      state.selectedRowKeys = []
      curPage.value = pag.current;
      pageSize.value = pag.pageSize;
      // svcService.GetServiceInfo(curPage.value, pageSize.value, selectVal.value, searchKey.value).then((res: any) => {
      //   if (res.data.code == 400) {
      //     alert(res.data.data)
      //   } else if (res.data.code == 200) {
      //     rolesList.list = (res.data.data)
      //     total.value = res.data.total
      //   }
      // })
    };


    const onSelectChange = (selectedRowKeys: Key[], selectedRows: Array<ServiceData>) => {

      // delIdList.clear()
      // for (let row of selectedRows) {
      //   delIdList.set(row.id, row.id)
      // }
      // state.selectedRowKeys = selectedRowKeys;
    };

    const onAddUser = () => {
      router.push({ name: "newUser", params: { cid: curSelectVal.value } });

    }

    const selectHandleChange = (value: any) => {
      curSelectVal.value = value
      userService.GetUserInfoByCid(curPage.value, pageSize.value, "", parseInt(curSelectVal.value)).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          userList.list = (res.data.data)
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
      router.push({ name: "editUser", params: { id: id, cid: curSelectVal.value } });
    }


    const OnSearch = () => {
      curPage.value = 1
      pageSize.value = 10
      // svcService.GetServiceInfo(curPage.value, pageSize.value, selectVal.value, searchKey.value).then((res: any) => {
      //   if (res.data.code == 400) {
      //     message.error(res.data.msg)
      //   } else if (res.data.code == 200) {
      //     rolesList.list = (res.data.data)
      //     total.value = res.data.total

      //   }
      // })
    }


    const OnDelete = (id: number) => {
      userService.DelUserInfo(id).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
          InitData()
        }
      })
    }

    const onResetPwd = (uid: string) => {
      userService.ResetPwd(uid).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
        }
      })
    }

    onMounted(() => {
      InitData()
    });

    const InitData = () => {
      userService.getUserList(curPage.value, pageSize.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          userList.list = (res.data.data)
          total.value = res.data.total
        }
      })
    }

    return {
      rolesList,
      searchKey,
      columns,
      onAddUser,
      onClose,
      curSelectVal,
      onResetPwd,
      onView,
      filterOption,
      cmpSelectVal,
      Mypagination,
      onDrawerVisible,
      drawerVisible,
      userList,
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