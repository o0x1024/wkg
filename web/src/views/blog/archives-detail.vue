<template >
    <div class="smain">
        <a-row justify="start">
            <a-col :xs="{ span: 22, offset: 1 }"  :lg="{ span: 16, offset: 4 }">
                <div class="post-context">
                    <div class="svtitle">
                        <div class="post-title">
                            <strong>{{ title }}</strong>
                        </div>
                        <a-row :gutter="10" class="post-meta">
                            <a-col>
                                <span>Author:<span style="color: forestgreen;">{{ author }}</span></span>
                            </a-col>
                            <a-col>
                                <span>Date:<span style="color: forestgreen;">{{ date }}</span></span>
                            </a-col>
                            <a-col>
                                <span>Category:<span style="color: forestgreen;">{{ category }}</span></span>
                            </a-col>
                        </a-row>
                    </div>

                    <div class="post-detail-content">
                        <md-editor editor-id="my-editor" v-model="content" :previewOnly="true" :showCodeRowNumber="true"
                            @onGetCatalog="onGetCatalogs" preview-theme="default"></md-editor>
                    </div>
                </div>
            </a-col>
            
            <a-col :xs="{ span: 0, offset: 0 }"  :lg="{ span: 3, offset: 0 }">
                <div class="share_catalogTree">
                    <Catelog :heads="catalogList"></Catelog>
                </div>
            </a-col>
        </a-row>


    </div>
</template>


<script lang="ts">

import { defineComponent, ref, onMounted, reactive, watch, computed } from "vue";
import { BugOutlined, ToolOutlined, FileAddOutlined, HomeOutlined, ClockCircleOutlined, RedoOutlined } from "@ant-design/icons-vue";
import { message } from "ant-design-vue";
import MdEditor from 'md-editor-v3';
import ShareKnowledgeService from '../../service/blog.service'
import 'md-editor-v3/lib/style.css';
import { useRouter } from "vue-router";
import Tool from "../../components/tool.vue"
import { useRoute } from 'vue-router';
import Catelog from '../../components/Catalog.vue'


interface HeadList {
    text: string;
    level: 1 | 2 | 3 | 4 | 5 | 6;
}

export default defineComponent({
    components: {
        // MdCatalog,
        BugOutlined,
        Catelog,
        HomeOutlined,
        ToolOutlined,
        FileAddOutlined,
        Tool,
        RedoOutlined,
        MdEditor,
        ClockCircleOutlined,
    },
    setup() {
        const router = useRouter();
        let content = ref("");
        let title = ref("")

        let route = useRoute()
        let key = route.query.key as any

        let updateTime = ref('')
        let author = ref('')
        let date = ref('')
        let category = ref('')


        const state: { catelogs: HeadList[] } = reactive({ catelogs: [] })
        const onGetCatalogs = (list: HeadList[]) => {
            state.catelogs = list

        }


        const toAdmin = () => {
            router.push("/adm");
        }

        const catalogList = computed(() => {
            return state.catelogs
        })

        const backToIndex = () => {
            router.push("/archives");
        }

        onMounted(() => {
            // console.Glog(key)
            ShareKnowledgeService.getKnowledgeByKey(key).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    let data = (res.data.data);
                    title.value = data.title
                    content.value = data.content
                    author.value = data.author
                    date.value = data.createTime
                    category.value = data.category
                }
            });
        });

        return {
            title,
            author,
            date,
            category,

            backToIndex,
            updateTime,
            catalogList,
            toAdmin,
            content,
            onGetCatalogs,
        };
    },
});
</script>


<style>


.post-detail-content{
    margin-bottom: 30px;
}

.post-meta {
    font-size: 1rem;
    color: rgba(85, 85, 85, 0.529) !important;
}


.post-title{
    font-size: 1.5rem;
}

.post-context {
    margin-top: 5%;
    /* max-width: 900px; */
    font-weight: normal;
    line-height: 2rem;
    /* overflow: overlay; */
    background-color: #fff;
    color: #161209;
    transition: color 0.2s ease, border-color 0.2s ease, background 0.2s ease, opacity 0.2s ease;
}


.smain {
    padding-bottom: 30px;
    height: 100%;
    font-size: 1rem;
}

.share_catalogTree {

   margin-top: 20%;

}


</style>
