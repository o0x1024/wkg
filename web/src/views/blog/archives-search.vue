<template>
	<div class="archives-post-wrap">
		<a-row justify="center" :gutter="16">
			<a-col :xs="{ span: 22, offset: 2 }" :md="{ span: 24, offset: 2 }" :lg="{ span: 16, offset: 2 }">
				<div>
					<a-row>
						<a-col :span="24">
							<div class="clist">
								<div v-for="item in listData.list">
									<div>
										<h3> {{ item.year }} </h3>
										<div v-for="itm in item.knows">
											<a-row justify="space-between">
												<a-col :span="24" :offset="1">
													<span class="archive-item" @click="btnDetail(itm.href)"> {{ itm.title
													}}</span>
												</a-col>
												<!-- <a-col :span="8" :offset="1" style="color: #a9a9b3;">
													<span> {{ itm.createTime }}</span>

												</a-col> -->
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
				<a-pagination :total="totals" size="small" @change="onChange" />
			</a-col>
		</a-row>
		<div style="height: 50px;">
		</div>
	</div>
</template>

<script lang="ts">

import { StarOutlined, BugOutlined, EyeOutlined, ToolOutlined, HomeOutlined, LikeOutlined, MessageOutlined, FireOutlined } from '@ant-design/icons-vue';
import { defineComponent, onMounted, reactive, computed, ref } from 'vue';
import { message } from "ant-design-vue";
import { useRouter } from "vue-router";
import HostList from '../../components/hotList.vue';
import Tool from "../../components/tool.vue"
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
		let pagesize = ref(14);
		let size = ref(900)

		let searchKey = router.currentRoute.value.params.keyword as string;

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

			shareKnowledgeService.searchsharedKnowledgebyWord(pageNum.value, pagesize.value, searchKey).then((res: any) => {
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
			shareKnowledgeService.searchsharedKnowledgebyWord(pageNum.value, pagesize.value, searchKey).then((res: any) => {
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
