interface Children {
	key: string
	path: string
	title: string
}

interface Menu {
	key?: string
	path: string
	icon?: string
	title?: string
	children?: Array<Children>
}

export const menuList: Array<Menu> = [
	{
		key: '1',
		path: '/adm/dashboard',
		icon: 'dashboard',
		title: 'Dashboard',
	},
	{
		key: '2',
		title: '任务管理',
		path: '/adm/task',
		children: [{
			key: '2.1',
			path: '/adm/task/tasklist',
			title: '任务列表'
		},
		//  {
		// 	key: '2.2',
		// 	path: '/adm/task/tasklist',
		// 	title: '任务列表'
		// }
	]
	},
	{
		key: '3',
		path: '/adm/src',
		title: 'SRC管理',
		children: [
			{
				key: '3.1',
				path: '/adm/src/company',
				title: '公司管理'
			},
			{
				key: '3.2',
				path: '/adm/src/gather',
				title: '资产总览'
			},
			{
				key: '3.3',
				path: '/adm/src/domain',
				title: '域名管理'
			},
			{
				key: '3.4',
				path: '/adm/src/website',
				title: '站点管理'
			},
			{
				key: '3.5',
				path: '/adm/src/ips',
				title: 'IP管理'
			},
			{
				key: '3.6',
				path: '/adm/src/service',
				title: '服务管理'
			},
			{
				key: '3.7',
				path: '/adm/src/miniprogram',
				title: '小程序管理'
			},
			{
				key: '3.8',
				path: '/adm/src/webChatAccount',
				title: '公众号管理'
			},
			{
				key: '3.9',
				path: '/adm/src/notice',
				title: '挖洞笔记'
			}]
	},
	{
		key: '5',
		path: '/adm/vuln',
		title: '漏洞管理',
		children: [{
			key: '5.1',
			path: '/adm/vuln/poc',
			title: 'POC管理'
		}, {
			key: '5.2',
			path: '/adm/vuln/newVuln',
			title: '新增漏洞'
		}, {
			key: '5.3',
			path: '/adm/vuln/historyVuln',
			title: '历史漏洞'
		}, {
			key: '5.4',
			path: '/adm/vuln/collect',
			title: '收藏漏洞'
		}]
	},
	{
		key: '6',
		path: '/adm/knowledge',
		title: '知识库',
		children: [{
			key: '6.1',
			path: '/adm/knowledge/treelist',
			title: '知识库'
		},
		{
			key: '6.2',
			path: '/adm/knowledge/share',
			title: '分享管理'
		},
		{
			key: '6.3',
			path: '/adm/knowledge/tag',
			title: 'TAG管理'
		}]
	},
	{
		key: '7',
		path: '/adm/tool',
		title: '工具库',
		children: [{
			key: '7.1',
			path: '/adm/tool/qrcode',
			title: '二维码'
		},
		{
			key: '7.2',
			path: '/adm/tool/endecode',
			title: '加密解密'
		}]
	},
	{
		key: '8',
		path: '/adm/agent',
		title: 'Agent管理',
		children: [{
			key: '8.1',
			path: '/adm/agent/list',
			title: 'Agent列表'
		}]
	},
	{
		key: '9',
		path: '/adm/user',
		title: '用户管理',
		children: [{
			key: '9.1',
			path: '/adm/user/list',
			title: '用户列表'
		}]
	},
	{
		key: '10',
		path: '/adm/dashboard',
		title: '系统设置',
		children: [{
			key: '10.1',
			path: '',
			title: '扫描配置'
		},
		{
			key: '10.2',
			path: '',
			title: '通知配置'
		}]
	}
];