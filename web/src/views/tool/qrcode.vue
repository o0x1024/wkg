<template>
    <a-row>
        <a-col :span="24">
            <a-upload-dragger
                v-model:fileList="fileList"
                name="file"
                :multiple="true"
                action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                @change="handleChange"
            >
                <p class="ant-upload-drag-icon">
                    <inbox-outlined></inbox-outlined>
                </p>
                <p class="ant-upload-text">Click or drag file to this area to upload</p>
                <p class="ant-upload-hint">
                    Support for a single or bulk upload. Strictly prohibit from uploading company data or other
                    band files
                </p>
            </a-upload-dragger>
        </a-col>
    </a-row>
    <a-row style="margin-top: 10px;">
        <a-col :span="24">
            <a-textarea v-model:value="result" placeholder="result" :rows="4" />
        </a-col>
    </a-row>
    <a-divider>qrcode decode</a-divider>  
</template>

<script lang="ts">
import { InboxOutlined } from '@ant-design/icons-vue';
import { message } from 'ant-design-vue';
import { defineComponent, ref } from 'vue';
    

export default defineComponent({
    components: {
        InboxOutlined,
    },
    setup() {
        let result = ref('')

        const handleChange = (info: any) => {
            const status = info.file.status;
            if (status !== 'uploading') {
            }
            if (status === 'done') {
                message.success(`${info.file.name} file uploaded successfully.`);
            } else if (status === 'error') {
                message.error(`${info.file.name} file upload failed.`);
            }
        };
        return {
            handleChange,
            result,
            fileList: ref([]),
        };
    },
});

</script>