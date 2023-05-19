<template>
    <a-row :gutter="10">
        <a-col>
            <a-button @click="onReturn" type="primary">
                返回
            </a-button>
        </a-col>
        <a-col>
            <a-button @click="onSave" type="primary">
                保存
            </a-button>
        </a-col>
        <a-col>
            <a-button @click="onReturn" type="primary">
                返回
            </a-button>
        </a-col>
    </a-row>


    <a-row style="margin-top: 10px; line-height: 30px;" :gutter="10">
        <a-col :span="1">
            <span><b>标题:</b></span>
        </a-col>

        <a-col :span="8">
            <a-input v-model:value="viewNotice.title" placeholder />
        </a-col>
    </a-row>


    <a-row style="margin-top: 10px;">
        <a-col :span="24">
            <md-editor editor-id="my-editor" style="padding: 10px; height:550px" v-model="viewNotice.content"
                :showCodeRowNumber="true" preview-theme="default"></md-editor>
        </a-col>
    </a-row>

    <a-row class="button_group">
        <a-col :span="8">
            <a-button @click="onSave" type="primary">
                保存
            </a-button>
        </a-col>
        <a-col :span="3">
            <a-button @click="onReturn" type="primary">
                返回
            </a-button>
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
    beforeRouteLeave(to, from, next) {
        if (this.hasEdit) {
            to.meta.keepAlive = false
        } else {
            to.meta.keepAlive = true
        }
        next()
    },
    setup() {
        const router = useRouter();
        let viewNotice: Notice = reactive({})
        let id = router.currentRoute.value.params.id as string;
        let cid = router.currentRoute.value.params.cid as string;
        let hasEdit = ref(false)
        const onReturn = () => {
            router.push({ name: "notice" });
        }

        onMounted(() => {
            noticeService.GetNoticeDetailById(parseInt(id)).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    let jdata = (res.data.data)
                    viewNotice.title = jdata.title
                    viewNotice.content = jdata.content
                }
            })
        });

        const onSave = () => {
            hasEdit.value = true
            noticeService.SaveEditNotice(parseInt(id), parseInt(cid), viewNotice.title, viewNotice.content).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    message.success(res.data.msg)
                }
            })
            router.push({ name: "notice" });
        }

        return {
            hasEdit,
            onReturn,
            onSave,
            viewNotice
        }
    }
})
</script>


<style>
.title {
    /* padding: 5px; */
    line-height: 30px;
    margin-left: 10%;
}


.mdeditor {
    /* min-height: 500px; */
    margin-top: 10px;
    margin-left: 10%;
}

.button_group {
    margin-left: 30%;
    margin-top: 1%;
}
</style>