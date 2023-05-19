<template>
    <div style="background: #ececec; padding: 30px">
        <a-row style="margin-left: 50px;" :gutter="30">
            <a-col :span="3">
                <a-card>
                    <a @click="toDomain()">
                        <a-statistic title="新增域名" :value="Data.sinfo.newDomain" :value-style="{ color: '#3f8600' }">
                            <!-- <template #prefix>
                            <arrow-up-outlined />
                        </template> -->
                        </a-statistic>
                    </a>
                </a-card>
            </a-col>
            <a-col :span="3">
                <a-card>
                    <a @click="toWebsite()">
                        <a-statistic title="新增站点" :value="Data.sinfo.newWebSite" :value-style="{ color: '#3f8600' }">
                            <!-- <template #prefix>
                            <arrow-down-outlined />
                        </template> -->
                        </a-statistic>
                    </a>
                </a-card>
            </a-col>
            <a-col :span="3">
                <a-card>
                    <a @click="toIP()">
                        <a-statistic title="新增IP" :value="Data.sinfo.newIP" class="demo-class"
                            :value-style="{ color: '#3f8600' }">
                            <!-- <template #prefix>
                            <arrow-down-outlined />
                        </template> -->
                        </a-statistic>
                    </a>
                </a-card>
            </a-col>
            <a-col :span="3">
                <a-card>
                    <a @click="toService()">
                        <a-statistic title="新增服务" :value="Data.sinfo.newService" class="demo-class"
                            :value-style="{ color: '#3f8600' }">
                            <!-- <template #prefix>
                            <arrow-down-outlined />
                        </template> -->
                        </a-statistic>

                    </a>
                </a-card>
            </a-col>
            <a-col :span="3">
                <a-card>
                    <a @click="toMiniprogram()">
                        <a-statistic title="新增小程序" :value="Data.sinfo.newMini" class="demo-class"
                            :value-style="{ color: '#3f8600' }">
                            <!-- <template #prefix>
                            <arrow-down-outlined />
                        </template> -->
                        </a-statistic>
                    </a>
                </a-card>
            </a-col>
            <a-col :span="3">
                <a-card>
                    <a @click="toWebChatAccount()">
                        <a-statistic title="新增公众号" :value="Data.sinfo.newWOA" class="demo-class"
                            :value-style="{ color: '#3f8600' }">
                            <!-- <template #prefix>
                            <arrow-down-outlined />
                        </template> -->
                        </a-statistic>
                    </a>
                </a-card>
            </a-col>
            <a-col :span="3">
                <a-card>
                    <a @click="toNews()">
                        <a-statistic title="新增资讯" :value="Data.sinfo.newMini" class="demo-class"
                            :value-style="{ color: '#3f8600' }">
                            <!-- <template #prefix>
                            <arrow-down-outlined />
                        </template> -->
                        </a-statistic>
                    </a>
                </a-card>
            </a-col>


        </a-row>
    </div>
</template>

<script lang="ts">
// import { ArrowUpOutlined, ArrowDownOutlined } from '@ant-design/icons-vue';

import dashService from "@/service/dash.service";
import { defineComponent, onMounted, reactive, ref } from "vue";
import { useRouter } from "vue-router";

interface StatisticInformation {
    totalWebSite?: number
    newWebSite?: number
    totalDomain?: number
    newDomain?: number
    totalIP?: number
    newIP?: number
    totalService?: number
    newService?: number
    totalMini?: number
    newMini?: number
    totalWOA?: number
    newWOA?: number
}

export default defineComponent({
    // components: {
    //     ArrowUpOutlined,
    //     ArrowDownOutlined,
    // },
    setup() {
        let Data: { sinfo: StatisticInformation } = reactive({ sinfo: {} })
        const router = useRouter();

        onMounted(() => {
            dashService.GetDashboard().then((res: any) => {
                if (res.status == 200) {
                    Data.sinfo = res.data.data
                }
            })
        });

        const toDomain = () => {
            router.push({ name: "domain" });
        }
        const toWebsite = () => {
            router.push({ name: "website" });
        }
        const toIP = () => {
            router.push({ name: "ip" });
        }
        const toNews = () => {
            router.push({ name: "news" });
        }
        const toMiniprogram = () => {
            router.push({ name: "miniprogram" });
        }
        const toService = () => {
            router.push({ name: "service" });
        }
        const toWebChatAccount = () => {
            router.push({ name: "webChatAccount" });
        }

        return {
            Data,
            toDomain,
            toWebsite,
            toIP,
            toWebChatAccount,
            toNews,
            toMiniprogram,
            toService,
        }
    }
})

</script>