<template>
    <div class="tag-main">
        <div class="tag-container">
            <div class="tag-post-wrap tags">
                <h2 class="tag-post-title">-&nbsp;Tag Cloud&nbsp;-</h2>
                <div class="tag-cloud-tags">
                    <template v-for="item in state.list">
                        <a @click="bntTag(item.tagName)">{{ item.tagName }} <small>({{ item.count }})</small></a>
                    </template>
                </div>
            </div>
        </div>
    </div>
</template>



<script lang="ts">

interface tag {
    tagName:string
    count:number
}

import { defineComponent, ref, onMounted, reactive, watch, computed } from "vue";
import { message } from "ant-design-vue";
import ShareKnowledgeService from '../../service/blog.service'
import 'md-editor-v3/lib/style.css';
import { useRouter } from 'vue-router';


interface HeadList {
    text: string;
    level: 1 | 2 | 3 | 4 | 5 | 6;
}

export default defineComponent({
    components: {},
    setup(props, ctx) {
        const router = useRouter();
        const state: { list: tag[] } = reactive({ list: [] })

        
        const getTags = () =>{
            ShareKnowledgeService.getTags().then((res:any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    let data = (res.data.data);
                    state.list = data
                }
            })
        }


        const bntTag = (tagName:string) =>{
            router.push({ name: "tags-detail", params: { keyword: tagName } });
        }
        

        onMounted(() => {
            getTags()
        });

        return{
            bntTag,
            getTags,
            state,
        }
    },
})
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