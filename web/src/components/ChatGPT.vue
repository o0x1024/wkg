<template>
    <div>
        <a-row :gutter="10">
            <a-col style="padding-top: 5px;">
                <span>API KEY:</span>
            </a-col>
            <a-col>
                <a-input v-model:value="apiKey" style="width: 300px;" placeholder="api key" />
            </a-col>
            <a-col></a-col>
            <a-col style="padding-top: 5px;">
                <span>SECRECT KEY:</span>
            </a-col>
            <a-col>
                <a-input v-model:value="secrectKey" style="width: 200px;" placeholder="secrect key" />
            </a-col>
        </a-row>
    </div>

    <div class="chat-content">
        <div v-html="ChatMessage"></div>
        <p></p>
        <a-spin :spinning="spinning" :delay="1000">
        </a-spin>
    </div>

    <div class="question-warp">
        <a-input v-model:value="input_message" @pressEnter="bntSendQuestion()" class="question-input"
            placeholder="input you question">
            <!-- <template #addonAfter>
                <caret-right-outlined />
            </template> -->
        </a-input>

    </div>
</template>


<script lang="ts">
import { defineComponent, onMounted, reactive, computed, ref } from 'vue';
import { CaretRightOutlined } from '@ant-design/icons-vue';
import toolService from "@/service/tool.service"
import { message } from 'ant-design-vue';


export default defineComponent({
    name: "ChatGPT",
    components: {
        CaretRightOutlined,
    },
    setup() {
        let ChatMessage = ref('')
        let apiKey = ref('')
        let secrectKey = ref('')
        let input_message = ref('')
        const spinning = ref<boolean>(false);


        const bntSendQuestion = () => {
            spinning.value = true
            ChatMessage.value += "<p><span class=\"question chatgpt\">问题:</span><span class=\"question-content chatgpt\">"+ input_message.value+"</span></p>"
            toolService.SendQuestion(apiKey.value, secrectKey.value, input_message.value).then((res: any) => {
                if (res.data.code==200) {
                    ChatMessage.value +="<p><span class=\"gpt-name chatgpt\">ChatGTP:</span> <span class=\"gpt-content chatgpt\">" +res.data.data +"</span></p>"
                    spinning.value = false
                }else if (res.data.code == 400){
                    spinning.value = false
                    message.warning(res.data.msg)
                    
                }
            })
        }
        return {
            spinning,
            apiKey,
            input_message,
            secrectKey,
            ChatMessage,
            bntSendQuestion,
        }
    }
})


</script>

<style>


span.chatgpt {
    font-size: 13pt;
    font-family:"Microsoft YaHei",微软雅黑
}
span.question {
    color: #e21414;
}

span.question-content{
    color: #000000;
}

span.gpt-name{
    color: #168124;
}

span.gpt-content{
    color: #000000;
}


.question-warp {
    position: absolute;
    bottom: 30px;
    left: 30%;
}

.chat-content {
    overflow-y: scroll;
    margin-top: 10px;
    width: 100%;
    max-height: 580px;
    background-color: rgb(255, 255, 255);
    font-family: "Arial", "Microsoft YaHei", "黑体", "宋体", sans-serif;
    /* box-shadow: 2px 2px 5px rgb(161, 154, 154), 2px 2px 5px rgb(160, 170, 160); */

}


.question-input {
    border-radius: 20px;
    padding: 20px;
    height: 20px;
    width: 500px;
    box-shadow: 2px 2px 5px rgb(161, 154, 154), 2px 2px 5px rgb(160, 170, 160);
}
</style>