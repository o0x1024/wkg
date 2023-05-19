<template>
    <div style="margin-top: 100px; margin-left: 400px;">
        <a-row>
            <a-col style="padding: 5px;">
                <span>任务名称：</span>
            </a-col>
            <a-col>
                <a-input style="width: 300px" v-model:value="taskname" placeholder="任务名称" />
            </a-col>
        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col style="padding: 5px;">
                <span>任务状态：</span>
            </a-col>
            <a-col>
                <a-radio-group v-model:value="taskStatus">
                    <a-radio-button value="true">运行</a-radio-button>
                    <a-radio-button value="false">停止</a-radio-button>
                </a-radio-group>
            </a-col>
        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col style="padding: 5px;">
                <span>任务类型：</span>
            </a-col>
            <a-col>
                <a-radio-group v-model:value="taskType" @change="typeChange()">
                    <a-radio-button value="0">信息搜集</a-radio-button>
                    <a-radio-button value="1">漏洞扫描</a-radio-button>
                </a-radio-group>
            </a-col>
        </a-row>
        <!-- //信息搜集模块 -->
        <div v-if="taskType == '0'">
            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>频&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;率：</span>
                </a-col>
                <a-col>
                    <a-input-number id="inputNumber" v-model:value="taskRate" :min="1" /><span>&nbsp;分钟</span>
                </a-col>
            </a-row>

            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>资产分布：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="CompanyRange">
                        <a-radio-button value="0">所有公司</a-radio-button>
                        <a-radio-button value="1">单[多]个公司</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>

            <div v-if="CompanyRange == '1'" style="margin-top: 10px;">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>公&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;司：</span>
                    </a-col>
                    <a-col>
                        <a-select v-model:value="targetValue" :options="targetOptions.list" mode="multiple" :size="size"
                            placeholder="选择公司" style="width: 400px" @popupScroll="popupScroll"></a-select>
                    </a-col>
                </a-row>
            </div>
            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>范&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;围：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="AssestRange">
                        <a-radio-button value="0">全部</a-radio-button>
                        <a-radio-button value="1">域名</a-radio-button>
                        <a-radio-button value="2">IP信息</a-radio-button>
                        <a-radio-button value="3">端口[服务]</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>
            <div v-if="AssestRange == '1'">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>搜集方式：</span>
                    </a-col>
                    <a-col>
                        <a-radio-group v-model:value="domainCollectionType">
                            <a-radio-button value="0">全部方式</a-radio-button>
                            <a-radio-button value="1">接口搜集</a-radio-button>
                            <a-radio-button value="2">暴力破解</a-radio-button>
                        </a-radio-group>
                    </a-col>
                </a-row>
            </div>
            <div v-if="AssestRange == '3'">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>搜集方式：</span>
                    </a-col>
                    <a-col>
                        <a-radio-group v-model:value="portRange">
                            <a-radio-button value="0">全端口</a-radio-button>
                            <a-radio-button value="1">常用端口</a-radio-button>
                        </a-radio-group>
                    </a-col>
                </a-row>
            </div>
        </div>

        <!-- //漏洞扫描模块,以nuclei进行扫描  -->
        <div v-else>
            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>频&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;率：</span>
                </a-col>
                <a-col>
                    <a-input-number id="inputNumber" v-model:value="taskRate" :min="1" /><span>&nbsp;分钟</span>
                </a-col>
            </a-row>


            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>资产分布：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="CompanyRange">
                        <a-radio-button value="0">所有公司</a-radio-button>
                        <a-radio-button value="1">单[多]个公司</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>

            <div v-if="CompanyRange == '1'" style="margin-top: 10px;">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>公&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;司：</span>
                    </a-col>
                    <a-col>
                        <a-select v-model:value="targetValue" :options="targetOptions.list" mode="multiple" :size="size"
                            placeholder="选择公司" style="width: 400px" @popupScroll="popupScroll"></a-select>
                    </a-col>
                </a-row>
            </div>

            <div v-if="taskType == '1'" style="margin-top: 10px;">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>POC:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>
                    </a-col>
                    <a-col>
                        <a-select v-model:value="PocValue" :options="PocOptions.list" show-search :size="size"
                            placeholder="选择POC" style="width: 400px" @popupScroll="popupScroll"></a-select>
                    </a-col>
                </a-row>
            </div>

            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>范&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;围：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="AssestRange">
                        <a-radio-button value="0">全部</a-radio-button>
                        <a-radio-button value="4">站点</a-radio-button>
                        <a-radio-button value="2">IP信息</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>
        </div>




        <a-row style="margin-top: 30px;">
            <a-col :span="5">
                <a-button @click="onNew()" type="primary">确认</a-button>
            </a-col>
            <a-col>
                <a-button @click="onBack()">返回</a-button>
            </a-col>
        </a-row>
    </div>
</template>


<script lang="ts">
import { defineComponent, reactive, onMounted, computed, ref, Ref, toRaw } from "vue";
import taskService from '../../service/task.service';
import { Rate, SelectProps } from 'ant-design-vue';
import { message } from 'ant-design-vue';
import { useRouter } from "vue-router";

interface Option {
    label: string
    value: string
}


export default defineComponent({
    name: "Home",
    components: {

    },
    setup() {
        let taskType = ref('0')
        let taskRate = ref(12)
        let taskname = ref('')
        const router = useRouter();
        let taskStatus = ref('true')
        let CompanyRange = ref('0')
        let AssestRange = ref('1')
        let portRange = ref('0')
        let domainCollectionType = ref('1')
        let targetValue = ref([]) as Ref<string[]>
        let PocValue = ref() 

        let targetOptions: { list: Option[] } = reactive({ list: [] })
        let PocOptions: { list: Option[] } = reactive({ list: [] })



        const popupScroll = () => {
            console.log('popupScroll');
        };


        const onTaskTargetChange = () => {
            console.log('popupScroll');
        };


        const typeChange = () => {
            if(taskType.value == "1"){
                AssestRange.value = "4"
            }
        };

        
        const onBack = () => {
            router.push({ name: "tasklist" });
        };



        const onNew = () => {
            let ts = true
            if (taskStatus.value == 'false') {
                ts = false
            }
            taskService.NewTask(taskname.value, ts, taskType.value, taskRate.value, CompanyRange.value, targetValue.value,PocValue.value, AssestRange.value, domainCollectionType.value, portRange.value).then((res: any) => {
                if (res.data.code == 200) {
                    message.success(res.data.msg)
                    router.push({ name: "tasklist" });
                } else {
                    message.error(res.data.msg)
                }
            })
        };

        onMounted(() => {
            taskService.getCompIds().then((res: any) => {
                if (res.data.code == 200) {
                    targetOptions.list = res.data.data
                } else {
                    message.error(res.data.msg)
                }
            })
            
            taskService.getPocs().then((res: any) => {
                if (res.data.code == 200) {
                    PocOptions.list = res.data.data
                } else {
                    message.error(res.data.msg)
                }
            })


        });

        return {
            size: ref<SelectProps['size']>('middle'),
            taskType,
            taskRate,
            CompanyRange,
            taskStatus,
            targetValue,
            AssestRange,
            domainCollectionType,
            typeChange,
            taskname,
            onNew,
            popupScroll,
            portRange,
            PocOptions,
            PocValue,
            targetOptions,
            onBack,
            onTaskTargetChange,
        }
    }
});
</script>
    