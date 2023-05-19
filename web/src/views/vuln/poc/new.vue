<template>
    <div>
        <a-button @click="onReturn" type="primary">
            返回
        </a-button>
        <a-row class="title">
            <a-col :span="1">
                <span><b>标题:</b></span>
            </a-col>

            <a-col :span="8">
                <a-input v-model:value="Poc.pocName" placeholder />
            </a-col>
        </a-row>
        <a-row class="mdeditor">
            <a-col :span="22">
                <Codemirror v-model:value="Poc.pocContent" :options="cmOptions" border placeholder="" :height="500" />
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
    </div>
</template>


<script lang="ts">

import { SearchOutlined, ReadOutlined } from '@ant-design/icons-vue';
import { defineComponent, reactive, onMounted, ref } from "vue";
import pocService from '../../../service/poc.service';
import { useRouter } from "vue-router";
import 'md-editor-v3/lib/style.css';
import { message } from 'ant-design-vue';
import "codemirror/theme/dracula.css";
import "codemirror/mode/yaml/yaml.js";
import Codemirror from "codemirror-editor-vue3";

interface Poc {
    id?: number
    pocName?: string
    pocContent?: string
    updateTime?: string
}


export default defineComponent({
    components: { SearchOutlined, ReadOutlined, Codemirror },
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
        let Poc: Poc = reactive( {})
        // let cid = router.currentRoute.value.params.cid as string;
        let hasEdit = ref(false)

        const cmOptions = { 
        mode: "text/yaml", // 语言模式
        theme: "dracula", // 主题
        lineNumbers: true, // 显示行号
        smartIndent: true, // 智能缩进
        indentUnit: 2, // 智能缩进单位为4个空格长度
        foldGutter: true, // 启用行槽中的代码折叠
        styleActiveLine: true, // 显示选中行的样式}
        }
        const onReturn = () => {
            router.push({ name: "poclist" });
        }

        onMounted(() => {
            console.log("cid:")
            // InitData()
        });

        const onSave = () => {
            hasEdit.value = true
            pocService.SavePoc(Poc.pocName, Poc.pocContent).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    message.success(res.data.msg)
                }
            })
            router.push({ name: "poclist" });
        }

        return {
            onReturn,
            hasEdit,
            onSave,
            cmOptions,
            Poc
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