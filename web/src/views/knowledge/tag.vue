<template>
    <div>
      <div style="margin-bottom: 16px">
        showLine:
        <a-switch v-model:checked="showLine" />
        <br />
        <br />
        showIcon:
        <a-switch v-model:checked="showIcon" />
      </div>
      <a-tree
        :show-icon="showIcon"
        :default-expanded-keys="['0-0-0']"
        :tree-data="treeData"
        @select="onSelect"
      >
        <template #icon><carry-out-outlined /></template>
        <template #title="{ dataRef }">
          <template v-if="dataRef.key === '0-0-0-1'">
            <div>multiple line title</div>
            <div>multiple line title</div>
          </template>
          <template v-else>{{ dataRef.title }}</template>
        </template>
        <template #switcherIcon="{ dataRef }">
          <SmileTwoTone v-if="dataRef.key === '0-0-2'" />
          <SmileTwoTone v-else />
        </template>
      </a-tree>
    </div>
  </template>
  <script lang="ts">
  import { CarryOutOutlined, SmileTwoTone } from '@ant-design/icons-vue';
  import type { TreeProps } from 'ant-design-vue';
  import { defineComponent, ref } from 'vue';
  export default defineComponent({
    components: {
      CarryOutOutlined,
      SmileTwoTone,
    },
    setup() {
      const showLine = ref<boolean>(true);
      const showIcon = ref<boolean>(false);
      const treeData = ref<TreeProps['treeData']>([
        {
          title: 'parent 1',
          key: '0-0',
          children: [
            {
              title: 'parent 1-0',
              key: '0-0-0',
              children: [
                { title: 'leaf', key: '0-0-0-0'},
                {
                  key: '0-0-0-1',
                },
                { title: 'leaf', key: '0-0-0-2' },
              ],
            },
            {
              title: 'parent 1-1',
              key: '0-0-1',
              children: [{ title: 'leaf', key: '0-0-1-0' }],
            },
            {
              title: 'parent 1-2',
              key: '0-0-2',
              children: [
                { title: 'leaf 1', key: '0-0-2-0' },
                {
                  title: 'leaf 2',
                  key: '0-0-2-1',
                },
              ],
            },
          ],
        },
        {
          title: 'parent 2',
          key: '0-1',
          children: [
            {
              title: 'parent 2-0',
              key: '0-1-0',
              children: [
                { title: 'leaf', key: '0-1-0-0' },
                { title: 'leaf', key: '0-1-0-1' },
              ],
            },
          ],
        },
      ]);
      const onSelect: TreeProps['onSelect'] = (selectedKeys, info) => {
        console.log();
      };
      return {
        showLine,
        showIcon,
        onSelect,
        treeData,
      };
    },
  });
  </script>
  
  