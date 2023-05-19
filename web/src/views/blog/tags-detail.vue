<template>
    <div class="tag-main">
        <div class="tag-container">
            <div class="tag-post-wrap tags">
                <h2 class="tag-post-title">-&nbsp;Tag {{ tagsWord }}&nbsp;-</h2>

                <div style="padding-top: 2rem;">
                    <div v-for="item in listData.list">
                        <div>
                            <div class="post-title"> {{ item.year }} </div>
                            <div v-for="itm in item.knows">
                                <a-row justify="space-between">
                                    <a-col :span="4" :offset="1">
                                        <span class="archive-item" @click="btnDetail(itm.href)"> {{
                                            itm.title
                                        }}</span>
                                    </a-col>
                                    <a-col :span="8" :offset="1" style="color: #a9a9b3;">
                                        <span> {{ itm.createTime }}</span>

                                    </a-col>
                                </a-row>
                            </div>
                        </div>
                    </div>


                    <a-row justify="center" style="margin-top: 20px; margin-bottom: 50px;">
                        <a-col>
                            <a-pagination :total="totals" size="small" @change="onChange" />
                        </a-col>
                    </a-row>
                    <div style="height: 50px;">
                    </div>
                </div>

            </div>
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
        let tagsWord = router.currentRoute.value.params.keyword as string;

        let totals = ref(0)
        let loading = ref(true)
        const listData: { list: ListData[] } = reactive({ list: [] });

        const onSearch = () => {
            GetsharedknowledgeListByTag()
        }

        const onChange = (page: number) => {
            pageNum.value = page
            GetsharedknowledgeListByTag()
        }

        const onSearchChange = () => {
            pageNum.value = 1
            pagesize.value = 10

            GetsharedknowledgeListByTag()
        }


        const GetsharedknowledgeListByTag = () => {
            shareKnowledgeService.GetsharedKnowledgeListbyTag(pageNum.value, pagesize.value, tagsWord).then((res: any) => {
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
            GetsharedknowledgeListByTag()
            loading.value = false
        });

        const btnDetail = (href: string) => {
            router.push(href);
        }
        return {

            tagsWord,
            btnDetail,
            onChange,
            loading,
            size,
            onSearch,
            listData,
            activeKey: ref('1'),
            onSearchChange,
            pagesize,
            totals,
        };
    },
});



</script>



<style>
.tag-main {
    flex-grow: 1;
    flex-shrink: 0;
    flex-basis: auto;
    display: flex;
    flex-direction: column;
}


.tag-main .tag-container {
    padding-left: 1em;
    padding-right: 1em;
    height: 100%;
    display: flex;
    flex-direction: column;
    flex: 1;
}

.tag-post-wrap {
    position: relative;
    width: 100%;
    max-width: 780px;
    margin: 0 auto;
    padding-top: 2rem;
}

.tag-post-wrap .tag-post-title {
    font-size: 2em;
    line-height: 1.5em;
}

.tag-cloud-tags {
    margin: 10px 0;
    padding-top: 2em;
}

.tag-cloud-tags a {
    display: inline-block;
    position: relative;
    margin: 5px 10px;
    word-wrap: break-word;
    transition-duration: 0.3s;
    transition-property: transform;
    transition-timing-function: ease-out;
}

a {
    color: #161209;
    text-decoration: none;
    transition: color 0.2s ease, border-color 0.2s ease, background 0.2s ease, opacity 0.2s ease;
    cursor: pointer;
}
</style>