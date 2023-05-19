<template>
  <div>
    <a-layout>
      <a-layout-sider v-model:collapsed="collapsed" :trigger="null" collapsible :style="{ height: '100vh', left: 0 }">
        <div>
          <img class="logo" src="../assets/logo.png" />
        </div>
        <a-menu v-model:selectedKeys="selectedKeys" @click="menuClcik" theme="dark" mode="inline">
          <template v-for="(item, index) in menuList" :key="index">
            <template v-if="item.children">
              <a-sub-menu :key="item.key">
                <template #title>{{ item.title }}</template>
                <a-menu-item v-for="subitem in item.children" :key="subitem.key">
                  <router-link :to="subitem.path">
                    <span>{{ subitem.title }}</span>
                  </router-link>
                </a-menu-item>
              </a-sub-menu>
            </template>
            <template v-else>
              <a-menu-item :key="item.key">
                <router-link :to="item.path">
                  <span>{{ item.title }}</span>
                </router-link>
              </a-menu-item>
            </template>
          </template>
        </a-menu>
      </a-layout-sider>
      <a-layout>
        <a-layout-header style="background: #fff; padding: 0">

          <menu-unfold-outlined v-if="collapsed" class="trigger" @click="() => (collapsed = !collapsed)" />
          <menu-fold-outlined v-else class="trigger" @click="() => (collapsed = !collapsed)" />
          <a-dropdown placement="bottom">
            <a-avatar class="user">
              <template #icon>
                <user-outlined @click="BntUser"></user-outlined>
              </template>
            </a-avatar>

            <template #overlay>
              <a-menu>
                <a-menu-item>
                  <a target="_blank" @click="OnModifyPasswd" rel="noopener noreferrer">修改密码</a>
                </a-menu-item>
                <a-menu-item style="text-align: center;">
                  <a target="_blank" @click="OnLogout" rel="noopener noreferrer">登出</a>
                </a-menu-item>
              </a-menu>
            </template>

          </a-dropdown>
          <!-- <a-button @click="toBlog()" type="primary" shape="circle">
            <template #icon>
              <home-outlined />
            </template>
          </a-button> -->
        </a-layout-header>

        <a-modal v-model:visible="visible" title="修改密码" @ok="handleOk">
          <a-row>
            <a-col>
              <a-row>
                <a-col style="margin-top:5px;">
                  <label>旧密码：</label>
                </a-col>
                <a-col>
                  <a-input v-model:value="oldPasswd" placeholder />
                </a-col>
              </a-row>
              <a-row style="margin-top: 10px;">
                <a-col style="margin-top:5px;">
                  <label>新密码：</label>
                </a-col>
                <a-col>
                  <a-input v-model:value="newPasswd" placeholder />
                </a-col>
              </a-row>
              <a-row style="margin-top: 10px;">
                <a-col style="margin-top:5px;">
                  <label>确认新密码：</label>
                </a-col>
                <a-col>
                  <a-input v-model:value="confimNewPasswd" placeholder />
                </a-col>
              </a-row>
            </a-col>
          </a-row>
        </a-modal>

        <a-layout-content>
          <a-breadcrumb style="margin-left:1% ; padding: 5px; font-size: 14px;">
            <a-breadcrumb-item v-for="(item, index) in routes" :key="item.name">
              <router-link v-if="index !== routes.length - 1" :to="{ path: item.path === '' ? '/' : item.path }">{{
                  item.meta.title
              }}</router-link>
              <span v-else>{{ item.meta.title }}</span>
            </a-breadcrumb-item>
          </a-breadcrumb>
          <div :style="{
            margin: '0px 16px',
            padding: '10px',
            background: '#fff',
            minHeight: '700px',
          }">
            <router-view />
          </div>
        </a-layout-content>
      </a-layout>
    </a-layout>
  </div>

</template>

<script lang="ts">
import {
  UserOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  HomeOutlined,
  // DashboardOutlined,
} from "@ant-design/icons-vue";
import { defineComponent, ref, computed } from "vue";
import { useRouter } from "vue-router";
import { menuList } from "../common/menu"
import axiosInstance from '../service/api';
import { message } from 'ant-design-vue';



export default defineComponent({
  components: {
    UserOutlined,
    HomeOutlined,
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    // DashboardOutlined,
  },
  setup() {
    const router = useRouter();
    const routes = computed(() => { return router.currentRoute.value.matched.filter(item => item.meta.title) })
    const visible = ref(false)
    const oldPasswd = ref('')
    const newPasswd = ref('')
    const confimNewPasswd = ref()

    const KeepAliveList = computed(() => {
      // return router.getRoutes.
      return "notice"
    })

    const menuClcik = (item: any) => {
      if (item.key == 1) {
        // console.Glog(router.getRoutes())
      }
    };
    const BntUser = () => {
      console.log(123)
    }

    const toBlog = () => {
      router.push("/archives");
    }

    const handleOk = () => {
      const data = { "oldPasswd": oldPasswd.value, "newPasswd": newPasswd.value, "confimNewPasswd": confimNewPasswd.value }
      axiosInstance.post('/api/v1/modifyPasswd', data,).then((res: any) => {
        if (res.data.code == 200) {
          message.success(res.data.msg)
        } else if (res.data.code == 400) {
          message.error(res.data.msg)
        }
      })
      visible.value = false
    }

    const OnLogout = () => {
      localStorage.removeItem('token')
      router.push("/login");
    }

    const OnModifyPasswd = () => {
      visible.value = true
    }



    return {
      selectedKeys: ref<string[]>(["1"]),
      collapsed: ref<boolean>(false),
      routes,
      oldPasswd,
      confimNewPasswd,
      newPasswd,
      visible,
      toBlog,
      menuList,
      KeepAliveList,
      basePath: '',
      OnLogout,
      handleOk,
      menuClcik,
      OnModifyPasswd,
      BntUser,
    };
  },
});



</script>

<style>
.opanchor {
  min-width: 200px;
  border-radius: 25px;
  max-width: 220px;
  font-weight: bold;
  top: 8%;
  position: fixed;
  right: 1px;
  z-index: 999;
  background: rgb(211, 210, 210);
}

.user {
  font-size: 18px;
  background-color: #1890ff;
  /* line-height: 64px; */
  /* padding: 0 24px; */
  /* cursor: pointer; */
  transition: color 0.3s;
  margin-left: 92%;
}

.trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 0px;
  cursor: pointer;
  transition: color 0.3s;
  /* margin-top: 1px; */
  margin-left: 10px;
}

.trigger:hover {
  color: #1890ff;
}

.logo {
  height: 44px;
  /* margin: auto; */
  margin: 16px;
  /* margin-left: 75px; */
  /* width: 0px; */
}

.site-layout .site-layout-background {
  background: #fff;
}
</style>
