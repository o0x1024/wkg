<template>
    <div style="margin-top: 100px; margin-left: 400px;">
        <a-row>
            <a-col style="padding: 5px;">
                <span>任务名称：</span>
            </a-col>
            <a-col>
                <a-input style="width: 300px" v-model:value="task.taskName" placeholder="任务名称" />
            </a-col>
        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col style="padding: 5px;">
                <span>任务状态：</span>
            </a-col>
            <a-col>
                <a-radio-group v-model:value="task.taskStatus">
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
                <a-radio-group v-model:value="task.taskType" @change="typeChange()">
                    <a-radio-button value="0">信息搜集</a-radio-button>
                    <a-radio-button value="1">漏洞扫描</a-radio-button>
                </a-radio-group>
            </a-col>
        </a-row>
        <!-- //信息搜集模块 -->
        <div v-if="task.taskType == '0'">
            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>频&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;率：</span>
                </a-col>
                <a-col>
                    <a-input-number id="inputNumber" v-model:value="task.rate" :min="1" /><span>&nbsp;分钟</span>
                </a-col>
            </a-row>

            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>资产分布：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="task.companyRange">
                        <a-radio-button value="0">所有公司</a-radio-button>
                        <a-radio-button value="1">单[多]个公司</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>

            <div v-if="task.companyRange == '1'" style="margin-top: 10px;">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>公&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;司：</span>
                    </a-col>
                    <a-col>
                        <a-select v-model:value="task.targetComp" :options="targetOptions.list" mode="multiple" :size="size"
                            placeholder="选择公司" style="width: 400px" @popupScroll="popupScroll"></a-select>
                    </a-col>
                </a-row>
            </div>
            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>范&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;围：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="task.assetRange">
                        <a-radio-button value="0">全部</a-radio-button>
                        <a-radio-button value="1">域名</a-radio-button>
                        <a-radio-button value="2">IP信息</a-radio-button>
                        <a-radio-button value="3">端口[服务]</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>
            <div v-if="task.assetRange == '1'">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>搜集方式：</span>
                    </a-col>
                    <a-col>
                        <a-radio-group v-model:value="task.domainCollectionType">
                            <a-radio-button value="0">全部方式</a-radio-button>
                            <a-radio-button value="1">接口搜集</a-radio-button>
                            <a-radio-button value="2">暴力破解</a-radio-button>
                        </a-radio-group>
                    </a-col>
                </a-row>
            </div>
            <div v-if="task.assetRange == '3'">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>搜集方式：</span>
                    </a-col>
                    <a-col>
                        <a-radio-group v-model:value="task.portRange">
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
                    <a-input-number id="inputNumber" v-model:value="task.rate" :min="1" /><span>&nbsp;分钟</span>
                </a-col>
            </a-row>


            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>资产分布：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="task.companyRange">
                        <a-radio-button value="0">所有公司</a-radio-button>
                        <a-radio-button value="1">单[多]个公司</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>

            <div v-if="task.companyRange == '1'" style="margin-top: 10px;">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>公&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;司：</span>
                    </a-col>
                    <a-col>
                        <a-select v-model:value="task.targetComp" :options="targetOptions.list" mode="multiple" :size="size"
                            placeholder="选择公司" style="width: 400px" @popupScroll="popupScroll"></a-select>
                    </a-col>
                </a-row>
            </div>

            <div v-if="task.taskType == '1'" style="margin-top: 10px;">
                <a-row style="margin-top: 10px;">
                    <a-col style="padding: 5px;">
                        <span>POC:&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>
                    </a-col>
                    <a-col>
                        <a-select v-model:value="task.poc" :options="PocOptions.list" show-search :size="size"
                            placeholder="选择POC" style="width: 400px" @popupScroll="popupScroll"></a-select>
                    </a-col>
                </a-row>
            </div>

            <a-row style="margin-top: 10px;">
                <a-col style="padding: 5px;">
                    <span>范&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;围：</span>
                </a-col>
                <a-col>
                    <a-radio-group v-model:value="task.assetRange">
                        <a-radio-button value="0">全部</a-radio-button>
                        <a-radio-button value="4">站点</a-radio-button>
                        <a-radio-button value="2">IP信息</a-radio-button>
                    </a-radio-group>
                </a-col>
            </a-row>
        </div>




        <a-row style="margin-top: 30px;">
            <a-col :span="5">
                <a-button @click="onUpdate()" type="primary">更新</a-button>
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


interface Task {
    taskId:string
    taskType:string
    rate:number
    taskName:string
    taskStatus:string
    companyRange:string
    assetRange:string
    portRange:string
    domainCollectionType:string
    targetComp:string[]
    poc:string
}

export default defineComponent({
    name: "Home",
    components: {

    },
    setup() {
        const router = useRouter();

        // let taskType = ref('0')
        // let taskRate = ref(12)
        // let taskname = ref('')
        // let taskStatus = ref('true')
        // let CompanyRange = ref('0')
        // let AssestRange = ref('1')
        // let portRange = ref('0')
        // let domainCollectionType = ref('1')
        // let targetComp = ref([]) as Ref<string[]>
        // let PocValue = ref()
        let taskid = router.currentRoute.value.params.taskid as string;

        let task:Task = reactive({taskId:taskid, taskType:"1",rate:12,taskName:'',taskStatus:'true',companyRange:'0',assetRange:'1',portRange:'0',domainCollectionType:'1',targetComp:[],poc:''})

        let targetOptions: { list: Option[] } = reactive({ list: [] })
        let PocOptions: { list: Option[] } = reactive({ list: [] })



        const popupScroll = () => {
            console.log('popupScroll');
        };


        const onTaskTargetChange = () => {
            console.log('popupScroll');
        };


        const typeChange = () => {
            if(task.taskType == "1"){
                task.assetRange = "4"
            }
        };

        
        const onBack = () => {
            router.push({ name: "tasklist" });
        };



        const onUpdate = () => {
            let ts = true
            if (task.taskStatus == 'false') {
                ts = false
            }
            taskService.UpdateTask(task.taskId, task.taskName, ts,  task.taskType, task.rate , task.companyRange , task.targetComp,task.poc, task.assetRange,task.domainCollectionType,task.portRange).then((res: any) => {
                if (res.data.code == 200) {
                    message.success(res.data.msg)
                    router.push({ name: "tasklist" });
                } else {
                    message.error(res.data.msg)
                }
            })
        };

        onMounted(() => {
            taskService.TaksInfo(taskid).then((res: any) => {
                if (res.data.code == 200) {
                    task.taskName  =res.data.data.taskName
                    task.assetRange  =res.data.data.assetRange
                    task.companyRange  =res.data.data.companyRange
                    task.domainCollectionType  =res.data.data.domainCollectionType
                    task.portRange  =res.data.data.portRange
                    task.rate  =res.data.data.rate
                    task.taskStatus  =res.data.data.taskStatus
                    task.taskType  =res.data.data.taskType

                    let arr = (res.data.data.companyId as string).split(',')
                    arr.forEach(element => {
                        if (element != ""){
                            task.targetComp.push(element)
                        }
                    });

                } else {
                    message.error(res.data.msg)
                }
            })

            
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
            task,
            typeChange,
            onUpdate,
            popupScroll,
            PocOptions,
            targetOptions,
            onBack,
            onTaskTargetChange,
        }
    }
});
</script>
    