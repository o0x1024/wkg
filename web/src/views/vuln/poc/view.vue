<template>
    <div>
        <a-button @click="onReturn" type="primary">
            退回
        </a-button>
    </div>
    <div class="poc_title">
        <a-row>
            <a-col :span="24">
                <span style="font-weight:bold;font-size:20px">{{ Poc.pocName }}</span>
            </a-col>
        </a-row>
    </div>
    <div class="poc_content">

        <a-row>
            <a-col :span="24">
                <Codemirror v-model:value="Poc.pocContent" :options="cmOptions" border placeholder="" :height="500"
                     />
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
    setup() {
        const router = useRouter();
        let Poc: Poc = reactive({})
        let content = ref('')
        let title = ref('')
        let id = ref(router.currentRoute.value.params.id);

           
        const cmOptions = { 
        mode: "text/yaml", // 语言模式
        theme: "dracula", // 主题
        lineNumbers: true, // 显示行号
        smartIndent: true, // 智能缩进
        indentUnit: 2, // 智能缩进单位为4个空格长度
        foldGutter: true, // 启用行槽中的代码折叠
        styleActiveLine: true, // 显示选中行的样式}
        readOnly:true
        }


        const InitData = () => {
            pocService.GetPocDetailById(id.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg)
                } else if (res.data.code == 200) {
                    let jdata = (res.data.data)
                    Poc.pocName = jdata.pocName
                    Poc.pocContent = jdata.pocContent
                }
            })
        }

        const onReturn = () => {
            router.push({ name: "poclist" });

        }

        onMounted(() => {
            InitData()
        });

        return {
            content,
            title,
            onReturn,
            cmOptions,
            Poc
        }
    }
})
</script>


<style>
.poc_title {
    margin-left: 15%;
}

.poc_content {
    margin-left: 15%;
    width: 900px;
    margin-right: 20%;
    margin-top: 10px;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2), 0 6px 20px 0 rgba(0, 0, 0, 0.19);
    /* height: 100px; */
    /* border: 1px solid #f00; */
}
</style>