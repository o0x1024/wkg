<template>
    <div class="info_panel">
        <a-divider>基础信息</a-divider>

        <a-row style="margin-top: 10px;">
            <a-col :span="6">
                <a-row>
                    <a-col><strong>主机名:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.host_name }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>公网Ipv4:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.IntranetIpv4 }}</span>
                    </a-col>

                </a-row>
            </a-col>


            <a-col :span="5">
                <a-row>
                    <a-col><strong>地域:</strong></a-col>
                    <a-col>
                        <span>default</span>
                    </a-col>
                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>系统启动时间:</strong> </a-col>
                    <a-col>
                        <span>{{ agentData.agent.boot_time }}</span>
                    </a-col>
                </a-row>
            </a-col>

        </a-row>
        <a-row style="margin-top: 10px;">
            <a-col :span="6">
                <a-row>
                    <a-col><strong>主机ID:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.agent_id }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>公网Ipv6:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.ExtranetIpv6 }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="5">
                <a-row>

                    <a-col><strong>内核版本:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.kernel_version }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>最近活跃时间:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.update_time }}</span>
                    </a-col>

                </a-row>
            </a-col>
        </a-row>

        <a-row style="margin-top: 10px;">
            <a-col :span="6">
                <a-row>
                    <a-col><strong>操作系统:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.platform }} {{ agentData.agent.platform_version }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>私网Ipv4:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.IntranetIpv4 }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="5">
                <a-row>

                    <a-col><strong>网络模式:</strong></a-col>
                    <a-col>
                        <span>SD</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>客户端安装时间:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.start_time }}</span>
                    </a-col>

                </a-row>
            </a-col>
        </a-row>

        <a-row style="margin-top: 10px;">
            <a-col :span="6">
                <a-row>
                    <a-col><strong>设置型号:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.host_model }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>CPU信息:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.cpu_name }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="5">
                <a-row>

                    <a-col><strong>CPU使用率:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.cpu }} </span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>设备序列号:</strong></a-col>
                    <a-col>
                        <div  style="width:200px; overflow: hidden; text-overflow : ellipsis; white-space: nowrap;">
                            {{ agentData.agent.host_serial }}
                        </div>
                    </a-col>

                </a-row>
            </a-col>
        </a-row>

        <a-row style="margin-top: 10px;"> 
            <a-col :span="6">
                <a-row>
                    <a-col><strong>生产商:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.host_vendor }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>设备ID:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.host_id }}</span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="5">
                <a-row>

                    <a-col><strong>内存大小:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.total_mem }} </span>
                    </a-col>

                </a-row>
            </a-col>

            <a-col :span="6">
                <a-row>
                    <a-col><strong>内存使用率:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.mem_usage }}</span>
                    </a-col>

                </a-row>
            </a-col>
        </a-row>

        <a-row style="margin-top: 10px;"> 
            <a-col :span="6">
                <a-row>
                    <a-col><strong>CPU使用率:</strong></a-col>
                    <a-col>
                        <span>{{ agentData.agent.cpu_usage }}</span>
                    </a-col>

                </a-row>
            </a-col>



        </a-row>
        <a-divider></a-divider>
    </div>
</template>


<script lang="ts">
import agentService from "@/service/agent.service";
import { ReconciliationFilled } from "@ant-design/icons-vue";
import {
    defineComponent,
    reactive,
    computed,
    onMounted,
    ref,
    toRefs,
} from "vue";

import { useRouter } from "vue-router";

interface Agent {
    agent_id?: string
    IntranetIpv4?: string
    IntranetIpv6?: string
    ExtranetIpv4?: string
    ExtranetIpv6?: string
    host_name?: string
    version?: string
    product?: string
    kernel_version?: string
    arch?: string
    platform?: string
    platform_family?: string
    platform_version?: string
    cpu?: string
    rss?: string
    total_mem?: string
    read_speed?: string
    write_speed?: string
    pid?: string
    nfd?: string
    start_time?: string
    host_serial?: string
    host_id?: string
    host_model?: string
    host_vendor?: string
    cpu_name?: string
    boot_time?: string
    cpu_usage?: string
    mem_usage?: string
    update_time?: string
}

export default defineComponent({
    setup() {
        const router = useRouter();
        let agentId = router.currentRoute.value.params.agentId;
        const agentData: { agent: Agent } = reactive({ agent: {} });
        let env: string[] = reactive([])
        let ellipsis = ref(true)

        onMounted(() => {
            InitData();
        });

        const InitData = () => {
            agentService
                .AgentDetail(agentId)
                .then((res: any) => {
                    if (res.status == 400) {
                        alert(res.data.msg);
                    } else if (res.status == 200) {
                        agentData.agent = res.data.data;
                        let size = Math.round(Number(agentData.agent.total_mem) / 1000 / 1000 / 1000)
                        agentData.agent.total_mem = String(size) + 'GB'

                        let mem_usage = Math.round(Number(agentData.agent.mem_usage)* 100)
                        agentData.agent.mem_usage = String(mem_usage) + '%'

                        let cpu_usage = Math.round(Number(agentData.agent.cpu_usage)* 100)
                        agentData.agent.cpu_usage = String(cpu_usage) + '%'
                    }
                });
        };


        return {
            agentData,
            ellipsis,
        }
    }
});


</script> 

<style>
.info_panel {
    width: 600;
}
</style>