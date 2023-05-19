<template>
    <div>
        <a-tabs v-model:activeKey="activeKey">
            <a-tab-pane key="1">
                <template #tab>
                    <span>
                        <apple-outlined />
                        加密解密
                    </span>
                </template>
                <div>
                    <iframe src="/CyberChef/index.html" :height="Iheigth" width="100%" frameborder="0"></iframe>
                </div>
            </a-tab-pane>
            <a-tab-pane key="2">
                <template #tab>
                    <span>
                        <android-outlined />
                        ChatGPT
                    </span>
                </template>
               <ChatGPT />
            </a-tab-pane>
        </a-tabs>
    </div>
</template>


<script lang="ts">

import { StarOutlined, BugOutlined, ToolOutlined, HomeOutlined, LikeOutlined, MessageOutlined, FireOutlined } from '@ant-design/icons-vue';
import { defineComponent, onMounted, reactive, computed, ref } from 'vue';
import ChatGPT  from '@/components/ChatGPT.vue'


export default defineComponent({
    name: "Tool",
    components: {
        ChatGPT,
        StarOutlined,
        BugOutlined,
        HomeOutlined,
        ToolOutlined,
        LikeOutlined,
        FireOutlined,
        MessageOutlined,
    },
    setup() {
        let Iheigth = computed(() => {
            const deviceHeight = document.documentElement.clientHeight;
            return (Number(deviceHeight) - 150); //数字是页面布局高度差
        })


        return {
            activeKey: ref('1'),
            Iheigth,
        }
    }

})

</script>