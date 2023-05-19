<template>
    <div class="user_from">
        <a-row>
            <a-col :span="5">
                <span>username:</span>
            </a-col>
            <a-col :span="12">
                <a-input v-model:value="formState.username" />
            </a-col>
        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col :span="5">
                <span>password:</span>
            </a-col>
            <a-col :span="12">
                <a-input-password v-model:value="formState.password" />
            </a-col>
        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col :span="8">
                <a-button type="primary" @click="onFinish">确定</a-button>
            </a-col>
            <a-col :span="12">
                <a-button type="primary" @click="onReturn">返回</a-button>
            </a-col>
        </a-row>

    </div>
</template>


<script lang="ts">

import { SearchOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { defineComponent, reactive, onMounted, ref } from "vue";
import userService from '../../service/user.service';
import { useRouter } from "vue-router";
import MdEditor from 'md-editor-v3';
import 'md-editor-v3/lib/style.css';
import { message } from 'ant-design-vue';


interface user {
    id?: number
    username?: number
    password?: string
    createTime?: string
}

interface FormState {
    username: string;
    password: string;
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
        const formState = reactive<FormState>({
            username: '',
            password: '',
        });
        let cid = router.currentRoute.value.params.cid as string;
        let hasEdit = ref(false)

        const onReturn = () => {
            router.push({ name: "userlist" });
        }

        const onFinish = () => {

            userService.SaveUser(formState.username, formState.password).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    message.success(res.data.msg)
                }
            })
            router.push({ name: "userlist" });
        }

        const onFinishFailed = (errorInfo: any) => {
            message.error(errorInfo)
        }


        return {
            onReturn,
            hasEdit,
            onFinish,
            onFinishFailed,
            formState
        }
    }
})
</script>


<style>
.user_from {
    margin-top: 10%;
    width: 400px;
    margin-left: 30%;
}
</style>