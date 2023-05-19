<template>
    <strong>
        <Link :href="`#${tocItem.text}`" :title="tocItem.text">
        <div v-if="tocItem.children" class="catalog-container">
            <CatalogLink v-for="item of tocItem.children" :key="`${item.level}-${item.text}`" :tocItem="item" />
        </div>
        </Link>
    </strong>
</template>

<script setup lang="ts">
import { Anchor } from 'ant-design-vue';
import { defineProps, PropType } from 'vue';

const { Link } = Anchor;

interface TocItem {
    text: string;
    level: number;
    children?: Array<TocItem>;
}
const { tocItem } = defineProps({
    tocItem: {
        type: Object as PropType<TocItem>,
        default: () => ({})
    }
});
</script>
