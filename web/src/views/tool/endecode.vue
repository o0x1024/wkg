<template>
    <div>
        <iframe src="/CyberChef/index.html" :height="Iheigth" width="100%" frameborder="0"></iframe>
    </div>
</template>


<script lang="ts">
import { computed, defineComponent, ref } from 'vue';


export default defineComponent({

    setup() {
        let Iheigth = computed(() => {
            const deviceHeight = document.documentElement.clientHeight;
            return (Number(deviceHeight)-150); //数字是页面布局高度差
        })
        // computed
        return {
            Iheigth,
        };
    },
});

</script>