<template>
    <!-- adsfsdfsf -->
    <div>
        <span><strong>&nbsp;&nbsp;&nbsp;&nbsp;热度排行</strong></span>
    </div>
    <div>
        <div v-for="line in hotlist">
            <a-row>
                <a-col>
                    <a-button :href="line.href" type="text">#{{ line.title }}# </a-button>
                </a-col>

                <a-col style="align-self:center">
                    <eye-outlined />&nbsp;&nbsp;{{ line.hit }}
                </a-col>
            </a-row>

            <!-- <a :href="line.href">{{ line.title }}</a>&nbsp;&nbsp;[{{ line.hit }}] -->
        </div>
    </div>
</template>



<script setup lang="ts">
import {
    EyeOutlined
} from "@ant-design/icons-vue";
import { computed, PropType, onMounted, defineProps, onActivated } from 'vue';

interface TocItem {
    title: string;
    href: string;
    hit: number
}

const props = defineProps({
    list: {
        type: Array as PropType<Array<any>>
    }
});



const Items: TocItem[] = [];

const hotlist = computed(() => {
    props.list?.forEach(({ title, href, hit }) => {
        const item = { title, href, hit };
        Items.push(item);
    })
    return Items
});

</script>

