import { createRouter, createWebHistory, RouteRecordRaw } from "vue-router";


const routes: Array<RouteRecordRaw> = [
  { path: "/", redirect: { name: "archives" } },
  {
    path: "/",
    name: "index",
    redirect: "archives",
    meta: { keepAlive: true },
    component: () => import(/* webpackPrefetch: true  */"@/views/blog/core.vue"),
    children: [
      {
        path: "archives",
        name: "archives",
        meta: { keepAlive: true },
        component: () => import(/* webpackPrefetch: true  */"@/views/blog/archives.vue"),
      },
      {
        path: "search/:keyword?",
        name: "archives-search",
        component: () => import(/* webpackPrefetch: true  */"@/views/blog/archives-search.vue"),
      },
      {
        path: "shareview/:key?",
        name: "shareview",
        meta: { keepAlive: true },
        component: () => import(/* webpackPrefetch: true  */"@/views/blog/archives-detail.vue"),
      },
      {
        path: "tags",
        name: "tags",
        meta: { keepAlive: true },
        component: () => import(/* webpackPrefetch: true  */"@/views/blog/tags.vue"),
      },
      {
        path: "detail/:keyword",
        name: "tags-detail",
        component: () => import(/* webpackPrefetch: true  */"@/views/blog/tags-detail.vue"),
      },
      {
        path: "category",
        name: "category",
        meta: { keepAlive: true },
        component: () => import(/* webpackPrefetch: true  */"@/views/blog/category.vue"),
      },
      {
        path: "about",
        name: "about",
        meta: { keepAlive: true },
        component: () => import(/* webpackPrefetch: true  */"@/views/blog/about.vue"),
      }

    ]
  },
  {
    path: "/login",
    name: "login",
    // redirect: "/login",
    component: () => import(/* webpackPrefetch: true  */"@/views/login.vue"),
  },
  {
    path: "/adm",
    name: "adm",
    meta: { title: "后台" },
    redirect: { name: 'login' },
    component: () => import(/* webpackPrefetch: true  */"@/views/admin_index.vue"),
    children: [
      {
        path: "dashboard",
        name: "Dashboard",
        meta: { title: "Dashboard" },
        component: () => import(/* webpackPrefetch: true  */"@/views/dash/dashboard.vue"),
      },

      {
        path: "src",
        name: "SRC管理",
        meta: { title: "SRC管理" },
        redirect: { name: 'gather' },
        component: () => import(/* webpackPrefetch: true  */"@/views/src/src.vue"),
        children: [{
          path: "company",
          name: "公司管理",
          meta: { title: "公司管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/company/company.vue"),
        },
        {
          path: "company/edit/:id",
          name: "editCompany",
          meta: { title: "公司管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/company/edit.vue"),
        },
        {
          path: "company/new",
          name: "newCompany",
          meta: { title: "公司管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/company/new.vue"),
        },
        {
          path: "gather",
          name: "gather",
          meta: { title: "资产总览" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/gather.vue"),
        },
        {
          path: "domain",
          name: "domain",
          meta: { title: "域名管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/domain.vue"),
        },
        {
          path: "website",
          name: "website",
          meta: { title: "站点管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/website.vue"),
        },
        {
          path: "ips",
          name: "ip",
          meta: { title: "IP管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/ips.vue"),
        },
        {
          path: "news",
          name: "news",
          meta: { title: "资讯管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/notice/notice_index.vue"),
        },
        {
          path: "miniprogram",
          name: "miniprogram",
          meta: { title: "小程序管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/mini_program.vue"),
        },
        {
          path: "service",
          name: "service",
          meta: { title: "服务管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/service.vue"),
        },
        {
          path: "webChatAccount",
          name: "webChatAccount",
          meta: { title: "微信公众号管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/webChatAccount.vue"),
        },
        {
          path: "notice",
          name: "notice",
          redirect: { name: 'noticeindex' },
          meta: { keepAlive: true },
          component: () => import(/* webpackPrefetch: true  */"@/views/src/notice/core.vue"),
          children: [
            {
              path: "noticeindex",
              name: "noticeindex",
              meta: { keepAlive: true },
              component: () => import(/* webpackPrefetch: true  */"@/views/src/notice/notice_index.vue"),
            },
            {
              path: "edit/:id",
              name: "editNotice",
              meta: { keepAlive: true },
              component: () => import(/* webpackPrefetch: true  */"@/views/src/notice/notice_edit.vue"),
            },
            {
              path: "view/:id",
              name: "viewNotice",
              meta: { keepAlive: false },
              component: () => import(/* webpackPrefetch: true  */"@/views/src/notice/notice_view.vue"),
            },
            {
              path: "new",
              name: "newNotice",
              meta: { keepAlive: true },
              component: () => import(/* webpackPrefetch: true  */"@/views/src/notice/notice_new.vue"),
            }
          ]
        },
        ]
      },
      {
        path: "vuln",
        name: "vuln",
        meta: { title: "漏洞管理" },
        redirect: { name: 'newVuln' },
        component: () => import(/* webpackPrefetch: true  */"@/views/vuln/vuln.vue"),
        children: [{
          path: "poc",
          name: "poc",
          meta: { title: "POC管理" },
          redirect: { name: 'poclist' },
          component: () => import(/* webpackPrefetch: true  */"@/views/vuln/poc/index.vue"),
          children: [{
            path: "poclist",
            name: "poclist",
            component: () => import(/* webpackPrefetch: true  */"@/views/vuln/poc/list.vue"),
          },
          {
            path: "viewPoc",
            name: "viewPoc",
            component: () => import(/* webpackPrefetch: true  */"@/views/vuln/poc/view.vue"),
          },
          {
            path: "newPoc",
            name: "newPoc",
            component: () => import(/* webpackPrefetch: true  */"@/views/vuln/poc/new.vue"),
          },
          {
            path: "editPoc",
            name: "editPoc",
            component: () => import(/* webpackPrefetch: true  */"@/views/vuln/poc/edit.vue"),
          },]
        }, {
          path: "newVuln",
          name: "newVuln",
          meta: { title: "新增漏洞" },
          component: () => import(/* webpackPrefetch: true  */"@/views/vuln/newVuln.vue"),
        }, {
          path: "historyVuln",
          name: "historyVuln",
          meta: { title: "历史漏洞" },
          component: () => import(/* webpackPrefetch: true  */"@/views/vuln/historyVuln.vue"),
        }, {
          path: "collect",
          name: "collect",
          meta: { title: "收藏漏洞" },
          component: () => import(/* webpackPrefetch: true  */"@/views/vuln/collect.vue"),
        }],
      },
      {
        path: "knowledge",
        name: "knowledge",
        meta: { title: "知识库" },
        redirect: { name: 'treelist' },
        component: () => import(/* webpackPrefetch: true  */"@/views/knowledge/knowledge.vue"),
        children: [{
          path: "treelist",
          name: "treelist",
          meta: { title: "漏洞知识库" },
          component: () => import(/* webpackPrefetch: true  */"@/views/knowledge/index.vue"),
        },
        {
          path: "share",
          name: "share",
          meta: { title: "分享管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/knowledge/shareConfig.vue"),
        },
        {
          path: "tag",
          name: "tag",
          meta: { title: "分享管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/knowledge/tag.vue"),
        },
        ]
      },
      {
        path: "tool",
        name: "tool",
        meta: { title: "工具库" },
        component: () => import(/* webpackPrefetch: true  */"@/views/tool/index.vue"),
        children: [{
          path: "qrcode",
          name: "qrcode",
          meta: { title: "漏洞知识库" },
          component: () => import(/* webpackPrefetch: true  */"@/views/tool/qrcode.vue"),
        },
        {
          path: "endecode",
          name: "endecode",
          meta: { title: "加密解密" },
          component: () => import(/* webpackPrefetch: true  */"@/views/tool/endecode.vue"),
        }]
      },
      {
        path: "task",
        name: "task",
        meta: { title: "任务管理" },
        redirect: { name: 'tasklist' },
        component: () => import(/* webpackPrefetch: true  */"@/views/task/index.vue"),
        children: [{
          path: "tasklist",
          name: "tasklist",
          meta: { title: "任务列表" },
          component: () => import(/* webpackPrefetch: true  */"@/views/task/list.vue")
        }, {
          path: "newtask",
          name: "newtask",
          meta: { title: "新建任务" },
          component: () => import(/* webpackPrefetch: true  */"@/views/task/new.vue")
        }, {
          path: "edittask/:taskid",
          name: "edittask",
          meta: { title: "编辑任务" },
          component: () => import(/* webpackPrefetch: true  */"@/views/task/edit.vue")
        }]
      },
      {
        path: "agent",
        name: "agent",
        meta: { title: "Agent管理" },
        redirect: { name: 'agentlist' },
        component: () => import(/* webpackPrefetch: true  */"@/views/agent/index.vue"),
        children: [{
          path: "list",
          name: "agentlist",
          meta: { title: "Agent列表" },
          component: () => import(/* webpackPrefetch: true  */"@/views/agent/list.vue")
        },
        {
          path: "detail/:agentId",
          name: "AgentDetail",
          meta: { title: "Agent管理" },
          component: () => import(/* webpackPrefetch: true  */"@/views/agent/detail.vue"),
        }]
      },
      {
        path: "onlineScan",
        name: "onlineScan",
        meta: { title: "在线扫描" },
        component: () => import(/* webpackPrefetch: true  */"@/views/scan/scan.vue"),
        children: [{
          path: "domainScan",
          name: "domainScan",
          meta: { title: "任务列表" },
          component: () => import(/* webpackPrefetch: true  */"@/views/scan/domainScan.vue"),
        }, {
          path: "vulnScan",
          name: "vulnScan",
          meta: { title: "漏洞扫描" },
          component: () => import(/* webpackPrefetch: true  */"@/views/scan/vScan.vue"),
        }]
      },
      {
        path: "user",
        name: "user",
        redirect: { name: 'userlist' },
        meta: { keepAlive: true },
        component: () => import(/* webpackPrefetch: true  */"@/views/user/core.vue"),
        children: [
          {
            path: "list",
            name: "userlist",
            meta: { keepAlive: true },
            component: () => import(/* webpackPrefetch: true  */"@/views/user/user_index.vue"),
          },
          {
            path: "edit/:id",
            name: "editUser",
            meta: { keepAlive: true },
            component: () => import(/* webpackPrefetch: true  */"@/views/user/user_edit.vue"),
          },
          {
            path: "view/:id",
            name: "viewUser",
            meta: { keepAlive: false },
            component: () => import(/* webpackPrefetch: true  */"@/views/user/user_view.vue"),
          },
          {
            path: "new",
            name: "newUser",
            meta: { keepAlive: true },
            component: () => import(/* webpackPrefetch: true  */"@/views/user/user_new.vue"),
          }
        ]
      },
    ],
  },

];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});


// router.beforeEach((to, from, next) => {
//   // 判断有没有登录

//   const isLogin = localStorage.token ? true : false;
//   console.Glog(isLogin)
//   // 如果当前访问的是登录页面或者注册页面可以让它进入
//   if (to.path == "/index") {
//     next();
//   } else {
//     // 如果isLogin为true表示已经登录了;就让它正常进入就可以 ,如果没登录就让他进入登录页面
//     isLogin ? next() : next("/");
//   }
// });

export default router;