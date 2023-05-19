<template>
	<div class="archives-post-wrap">
		<a-row justify="center" :gutter="16">
			<a-col :xs="{ span: 24, offset: 1 }" :md="{ span: 24, offset: 2 }" :lg="{ span: 16, offset: 2 }">
				<div>
					<a-row>
						<a-col :xs="{ span: 20, offset: 0 }" :md="{ span: 24, offset: 0 }" :lg="{ span: 24, offset: 0 }">
							<div>
								<div v-for="item in listData.list">
									<div>
										<h3> {{ item.year }} </h3>
										<div v-for="itm in item.knows">
											<a-row justify="space-between">
												<a-col :xs="{ span: 15, offset: 1 }" :lg="{ span: 10, offset: 1 }">
													<span class="archive-item" @click="btnDetail(itm.href)"> {{ itm.title
													}}</span>
												</a-col>
												<a-col v-if="!screenRotateFlag" :xs="{ span: 3, offset: 1 }" :lg="{ span: 8, offset: 2 }"
													style="color: #a9a9b3;">
													<span> {{ itm.createTime }}</span>

												</a-col>
											</a-row>
										</div>
									</div>
								</div>
							</div>
						</a-col>
					</a-row>
				</div>
			</a-col>
		</a-row>

		<a-row justify="center" style="margin-top: 20px; margin-bottom: 50px;">
			<a-col>
				<a-pagination :defaultPageSize="13" :hideOnSinglePage="true" :total="totals" size="small" @change="onChange" />
			</a-col>
		</a-row>
		<!-- <div style="height: 50px;"> -->
		<!-- </div> -->
	</div>
</template>

<script lang="ts">

import { StarOutlined, BugOutlined, EyeOutlined, ToolOutlined, HomeOutlined, LikeOutlined, MessageOutlined, FireOutlined } from '@ant-design/icons-vue';
import { defineComponent, onMounted, reactive, computed,onUnmounted,watch, ref } from 'vue';
import { message } from "ant-design-vue";
import { useRouter } from "vue-router";
import HostList from '../../components/hotList.vue';
import Tool from "../../components/tool.vue"
import { debounce } from 'lodash'

import shareKnowledgeService from '../../service/blog.service'

interface TocItem {
	title: string;
	href: string;
	hit: number
}

interface Knows {
	href: string
	title: string
	createTime: string
}

interface ListData {
	year: string
	knows: Array<Knows>
}


export default defineComponent({
	components: {
		StarOutlined,
		HostList,
		EyeOutlined,
		BugOutlined,
		HomeOutlined,
		Tool,
		ToolOutlined,
		LikeOutlined,
		FireOutlined,
		MessageOutlined,
	},
	setup() {
		const router = useRouter();
		let pageNum = ref(1);
		let pagesize = ref(13);
		let size = ref(900)
		let screenRotateFlag = ref(false)

		let totals = ref(0)
		let search_word = ref('')
		let loading = ref(true)
		const listData: { list: ListData[] } = reactive({ list: [] });
		const actions: Record<string, string>[] = reactive([]);
		let Iheigth = computed(() => {
			const deviceHeight = document.documentElement.clientHeight;
			return (Number(deviceHeight) - 150); //数字是页面布局高度差
		})

		const state: { hotList: TocItem[] } = reactive({ hotList: [] })

		const proHotList = computed(() => {
			return state.hotList
		})




		const screenWidth = ref(window.innerWidth)

		const handleResize = debounce(() => {
			screenWidth.value = window.innerWidth
		}, 250)

		onMounted(() => {
			if (screenWidth.value < 576) {
				screenRotateFlag.value = true
			} else {
				screenRotateFlag.value = false
			}
			window.addEventListener('resize', handleResize)
		})

		onUnmounted(() => {
			window.removeEventListener('resize', handleResize)
		})

		watch(
			// 监听的变量
			screenWidth,

			// 变量变化时的回调函数
			(newWidth, oldWidth) => {
				if (newWidth < 576) {
					screenRotateFlag.value = true
				} else {
					screenRotateFlag.value = false
				}
				console.log(`屏幕宽度从 ${oldWidth}px 变为 ${newWidth}px`)
			}
		)

		const onChange = (page: number) => {
			pageNum.value = page
			shareKnowledgeService.getKnowledgeList(pageNum.value, pagesize.value).then((res: any) => {
				if (res.data.code == 200) {
					listData.list.length = 0
					let data = res.data.data
					data.forEach((val: any) => {
						listData.list.push(val)
					});
				} else {
					message.error(res.data.msg)
				}
			})
		}

		// const onShowSizeChange = (current: number, pageSize: number) => {
		// 	pageNum.value = current;
		// 	pagesize.value = pageSize
		// 	shareKnowledgeService.getKnowledgeList(pageNum.value, pagesize.value).then((res: any) => {
		// 		if (res.data.code == 200) {
		// 			listData.list.length = 0
		// 			let data = (res.data.data)
		// 			data.forEach((val: any) => {
		// 				listData.list.push(val)
		// 			});
		// 		} else {
		// 			message.error(res.data.msg)
		// 		}
		// 	})
		// }

		const onSearch = () => {
			shareKnowledgeService.searchsharedKnowledgebyWord(pageNum.value, pagesize.value, search_word.value).then((res: any) => {
				if (res.data.code == 200) {
					listData.list.length = 0
					let data = (res.data.data)
					data.forEach((val: any) => {
						listData.list.push(val)
					});
				} else {
					message.error(res.data.msg)
				}
			})
		}


		const toAdmin = () => {
			router.push("/adm");
		}

		const onSearchChange = () => {
			pageNum.value = 1
			pagesize.value = 10

			shareKnowledgeService.searchsharedKnowledgebyWord(pageNum.value, pagesize.value, search_word.value).then((res: any) => {
				if (res.data.code == 200) {
					listData.list.length = 0
					let data = (res.data.data)
					data.forEach((val: any) => {
						listData.list.push(val)
					});
				} else {
					message.error(res.data.msg)
				}
			})
		}
		onMounted(() => {
			shareKnowledgeService.getKnowledgeList(pageNum.value, pagesize.value).then((res: any) => {
				if (res.data.code == 200) {
					let data = (res.data.data)
					data.forEach((val: any) => {
						listData.list.push(val)
					});
					// 
					totals.value = res.data.total
				} else {
					message.error(res.data.msg)
				}
			})
			loading.value = false
		});

		const btnDetail = (href: string) => {
			router.push(href);
		}
		return {
			toAdmin,
			// onShowSizeChange,
			actions,
			screenRotateFlag,
			btnDetail,
			loading,
			search_word,
			size,
			proHotList,
			state,
			onSearch,
			listData,
			activeKey: ref('1'),
			onSearchChange,
			pagesize,
			onChange,
			totals,
			Iheigth,
		};
	},
});



</script>

<style>
.year {
	font-size: 1rem;
	font-weight: normal;
	font-family: 'Microsoft Jhenghei', Lantinghei SC, 'lanting', PingFang SC, Seguo UI, Microsoft Yahei, Arial;
	line-height: 2em;
}

.archive-item {
	font-size: 1rem;
	color: #161209;
	cursor: pointer;
	transition: color 0.2s ease, border-color 0.2s ease, background 0.2s ease, opacity 0.2s ease;
	margin-left: 2rem;

	display: inline-block;
	text-decoration: none;
	overflow: hidden;
	white-space: nowrap;
	font-family: 'Microsoft Jhenghei', Lantinghei SC, 'lanting', PingFang SC, Seguo UI, Microsoft Yahei, Arial;

}

h3 {
	display: block;
	font-size: 1.17rem;
	margin-block-start: 1em;
	margin-block-end: 1em;
	margin-inline-start: 0px;
	margin-inline-end: 0px;
	font-weight: bold;
}

/* .archives-post-wrap { */
/* position: relative; */
/* width: 100%; */
/* max-width: 1500px; */
/* margin: 0 auto; */
/* padding-top: 2rem; */
/* } */


.archives-post-wrap {
	font-size: 1rem;
	font-weight: normal;
	line-height: 2em;
	overflow: hidden;
	background-color: #fff;
	color: #161209;
	transition: color 0.2s ease, border-color 0.2s ease, background 0.2s ease, opacity 0.2s ease;
}
</style>
