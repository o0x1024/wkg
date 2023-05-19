<template>
  <a-row type="flex" justify="start" :gutter="16">
    <a-col>
      <a-row :gutter="20">
        <a-col>
          <a-button @click="onAddTask" type="primary">
            <template #icon>
              <read-outlined />
            </template>
            新建任务
          </a-button>
        </a-col>
        <a-col>
          <a-button @click="onClsTaskQueue()" type="primary">
            清除任务缓存队列
          </a-button>
        </a-col>

        <a-col>
          <div style="line-height: 30px;">
            任务队列缓存：
            <a-tag>{{ queueNum }}</a-tag>

          </div>
        </a-col>
      </a-row>
    </a-col>
  </a-row>
  <a-row type="flex" style="margin-top: 10px;">
    <a-col :span="24" :order="4">
      <a-table :row-selection="{ selectedRowKeys: selectedRowKeys, onChange: onSelectChange }" :columns="columns"
        :data-source="taskList.list" :pagination="Mypagination" :align="dbalign" @change="handleTableChange">
        <template #bodyCell="{ column, text, record }">

          <template v-if="column.dataIndex === 'taskType'">
            <div v-if="record.taskType == '0'">
              <span>
                <a-tag color="orange">信息搜集</a-tag>
              </span>
            </div>
            <div v-else>
              <span>
                <a-tag color="blue">漏洞扫描</a-tag>
              </span>
            </div>
          </template>


          <template v-if="column.dataIndex === 'rate'">
            <span>[{{ record.rate }}]分钟</span>
          </template>

          <template v-if="column.dataIndex === 'companyRange'">
            <div v-if="record.companyRange == '1'">
              <span>
                <a-tag color="green">单[多]个公司</a-tag>
              </span>
            </div>
            <div v-else>
              <span>
                <a-tag color="purple">全部公司</a-tag>
              </span>
            </div>
          </template>

          <template v-if="column.dataIndex === 'pocName'">
            <div>
              <span>
                <a-tag color="purple">{{ record.pocName }}</a-tag>
              </span>
            </div>
          </template>

          <template v-if="column.dataIndex === 'companyId'">
            <span>
              <a-tag color="purple">{{ record.companyId }}</a-tag>
            </span>
          </template>

          <template v-if="column.dataIndex === 'assetRange'">
            <div v-if="record.assetRange == '1'">
              <span>
                <a-tag color="red">域名</a-tag>
              </span>
            </div>
            <div v-else-if="record.assetRange == '0'">
              <span>
                <a-tag color="pink">全部方式</a-tag>
              </span>
            </div>
            <div v-else-if="record.assetRange == '2'">
              <span>
                <a-tag color="cyan">IP信息</a-tag>
              </span>
            </div>
            <div v-else-if="record.assetRange == '4'">
              <span>
                <a-tag color="cyan">站点</a-tag>
              </span>
            </div>
            <div v-else>
              <span>
                <a-tag color="cyan">端口[服务]</a-tag>
              </span>
            </div>
          </template>

          <template v-if="column.dataIndex === 'domainCollectionType'">
            <div v-if="record.domainCollectionType == '1'">
              <span>
                <a-tag color="red">接口搜集</a-tag>
              </span>
            </div>
            <div v-else-if="record.domainCollectionType == '0'">
              <span>
                <a-tag color="pink">全部方式</a-tag>
              </span>
            </div>
            <div v-else-if="record.domainCollectionType == '2'">
              <span>
                <a-tag color="cyan">暴力破解</a-tag>
              </span>
            </div>
            <div v-else>
              <span>
                <a-tag color="cyan">——</a-tag>
              </span>
            </div>
          </template>


          <template v-if="column.key === 'operation'">
            <a-row type="flex" justify="center">
              <a-col>
                <a-button size="small" type="primary" @click="onEdit(record.taskId)">编辑</a-button>
              </a-col>
              <a-col>
                <div v-if="record.taskStatus">
                  <a-button size="small" type="primary" @click="onStop(record.taskId)">停止</a-button>
                </div>
                <div v-else>
                  <a-button size="small" type="primary" @click="onRun(record.taskId)">运行</a-button>
                </div>
              </a-col>
              <a-col>
                <a-button size="small" type="primary" @click="OnDelete(record.taskId)" danger>删除</a-button>
              </a-col>
            </a-row>
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
import taskService from '../../service/task.service';
import MdEditor from 'md-editor-v3';
import { useRouter } from "vue-router";
import { message } from 'ant-design-vue';

interface task {
  taskId: string
  taskStatus: boolean
  taskName: string,
  taskType: string,
  rate: number,
  companyRange: string,
  companyId: string,
  pocName: string
  assetRange: string,
  domainCollectionType: string,
  portRange: string
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
    let columns = types.getTaskTableColumns()
    const rolesList: { list: ServiceData[] } = reactive({ list: [] });
    let searchKey = ref('')
    let cmpSelectVal = ref('')
    let curSelectVal = ref('')
    const router = useRouter();
    let drawerVisible = ref(false)
    let queueNum = ref(0)
    const CmpSelectOptions: { list: option[] } = reactive({ list: [] })
    const taskList: { list: task[] } = reactive({ list: [] })


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

    const onSelectChange = () => {
      router.push({ name: "newtask" });

    }


    const onAddTask = () => {
      router.push({ name: "newtask" });

    }

    const selectHandleChange = (value: any) => {
      // curSelectVal.value = value
      // userService.GetUserInfoByCid(curPage.value, pageSize.value, "", parseInt(curSelectVal.value)).then((res: any) => {
      //   if (res.data.code == 400) {
      //     message.error(res.data.msg)
      //   } else if (res.data.code == 200) {
      //     userList.list = (res.data.data)
      //     total.value = res.data.total
      //   }
      // })
    }


    const onStop = (id: string) => {
      taskService.stopTask(id).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
          taskList.list.forEach(element => {
            if (element.taskId == id) {
              element.taskStatus = !element.taskStatus
            }
          });
        }
      })
    }
    const onRun = (id: string) => {
      taskService.runTask(id).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
          taskList.list.forEach(element => {
            if (element.taskId == id) {
              element.taskStatus = !element.taskStatus
            }
          });
        }
      })

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

    const onEdit = (taskid: string) => {
      router.push({ name: "edittask", params: { taskid:taskid } });

    }

    const OnDelete = (taskid: string) => {
      taskService.DelTaskByTaskId(taskid).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          message.success(res.data.msg)
          InitData()
        }
      })
    }

    const onClsTaskQueue = () => {
      taskService.clsTaskQueue().then((res: any) => {
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
      taskService.getTaskList(curPage.value, pageSize.value).then((res: any) => {
        if (res.data.code == 400) {
          message.error(res.data.msg)
        } else if (res.data.code == 200) {
          taskList.list = (res.data.data)
          queueNum.value = res.data.queueNum
          total.value = res.data.total
        }
      })
    }

    return {
      rolesList,
      searchKey,
      columns,
      onAddTask,
      onClose,
      onSelectChange,
      curSelectVal,
      onRun,
      filterOption,
      cmpSelectVal,
      taskList,
      Mypagination,
      drawerVisible,
      onDrawerEdit,
      onEdit,
      CmpSelectOptions,
      onClsTaskQueue,
      onStop,
      dbalign: ref('center'),
      hasSelected,
      selectHandleChange,
      handleTableChange,
      OnDelete,
      queueNum,
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