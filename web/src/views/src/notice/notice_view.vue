<template>
    <a-row>
        <a-col :span="3">
            <div>
                <a-button @click="onReturn" type="primary">
                    返回
                </a-button>
            </div>
        </a-col>

        <a-col style="margin-top: 50px;">
            <a-row>
                <a-col>
                    <span style="font-weight:bold;font-size:35px">{{ viewNotice.title }}</span>

                    <md-editor editor-id="my-editor" style="margin-top: 10px;" v-model="viewNotice.content" :previewOnly="true"
                        :showCodeRowNumber="true" preview-theme="default"></md-editor>
                </a-col>
            </a-row>
        </a-col>
    </a-row>
</template>


<script lang="ts">

import { SearchOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { defineComponent, reactive, onMounted, ref } from "vue";
import noticeService from '../../../service/notice.service';
import { useRouter } from "vue-router";
import MdEditor from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { message } from 'ant-design-vue';


interface Notice {
    id?: number
    cid?: number
    title?: string
    content?: string
    updateTime?: string
}


export default defineComponent({
    components: { SearchOutlined, ReadOutlined, MdEditor },
    setup() {
        const router = useRouter();
        let viewNotice: Notice = reactive({})
        let content = ref('')
        let title = ref('')
        let id = ref(router.currentRoute.value.params.id);

        const InitData = () => {
            noticeService.GetNoticeDetailById(id.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    let jdata = (res.data.data)
                    viewNotice.title = jdata.title
                    viewNotice.content = jdata.content
                }
            })
        }

        const onReturn = () => {

            router.push({ name: "notice" });

        }

        onMounted(() => {
            InitData()
        });

        return {
            content,
            title,
            onReturn,
            viewNotice
        }
    }
})
</script>


<style>
.title {
    margin-left: 35%;
}

.notice_content {
    margin-left: 15%;
    width: 900px;
    margin-right: 20%;
    margin-top: 10px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    /* height: 100px; */
    /* border: 1px solid #f00; */
}
</style>