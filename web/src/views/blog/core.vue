<template>
    <div class="archives-post-wrap">
        <a-layout>
            <a-layout-header class="header">
                <a-row >
                    <a-col :span="24">
                        <a-row justify="space-between">
                            <a-col :xs="{ span: 8, offset: 2 }" :lg="{ span: 4, offset: 2 }">
                                <span class="blog-title" @click="bntBlog()">Gelen's Blog</span>
                            </a-col>
                            <a-col :xs="{ span: 10, offset: 0 }" :lg="{ span: 10, offset: 0 }">
                                <a-row :gutter="20" class="blog-ngv">
                                    <a-col>
                                        <div class=" input-box mb20">
                                            <a-input v-model:value="searchKeyword" @pressEnter="onSearchChange()"
                                                type="text" class="input" />
                                            <span @click="onSearchChange()" class="span">搜</span>
                                        </div>
                                    </a-col>
                                    <!-- <div v-if="!hiddenMenuFlg"> -->
                                    <a-col v-if="!hiddenMenuFlg">
                                        <a-row :gutter="20">
                                            <a-col>
                                                <span class="blog" @click="bntBlog()" style="cursor: pointer;">Posts</span>
                                            </a-col>
                                            <a-col>
                                                <span class="blog" @click="bntCategory()"
                                                    style="cursor: pointer;">Categories</span>
                                            </a-col>
                                            <a-col>
                                                <span class="blog" @click="bntTags()" style="cursor: pointer;">Tags</span>
                                            </a-col>
                                            <a-col>
                                                <a-dropdown>
                                                    <span class="blog" style="cursor: pointer;">Assist</span>
                                                    <template #overlay>
                                                        <a-menu>
                                                            <a-row>
                                                                <a-button @click="btnTool()" type="text">Tool</a-button>
                                                            </a-row>
                                                            <a-row>
                                                                <a-button @click="btnToAdmin()"
                                                                    type="text">manager</a-button>
                                                            </a-row>
                                                        </a-menu>
                                                    </template>
                                                </a-dropdown>
                                            </a-col>
                                            <a-col>
                                                <span class="blog" @click="bntAbout()" style="cursor: pointer;">About</span>
                                            </a-col>
                                        </a-row>
                                    </a-col>
                                    <!-- <a-col v-else>
                                        <div style="cursor: pointer;" @click="disMenu()">☰ Menu</div>

                                    </a-col> -->
                                </a-row>
                            </a-col>
                        </a-row>

                        <!-- <a-row v-if="menuActive" class="drop-menu">
                            <a-col :span="24">
                                <div>
                                    <a class="menu-item" @click="bntBlog()">Posts</a>
                                    <a class="menu-item" @click="bntCategory()">Categories</a>
                                    <a class="menu-item" @click="bntTags()">Tags</a>
                                    <a class="menu-item" @click="bntAbout()">About</a>
                                </div>
                            </a-col>
                        </a-row> -->
                    </a-col>

                </a-row>


            </a-layout-header>
            <a-layout-content class="content">
                <div class="blog-content">
                    <router-view></router-view>
                </div>
            </a-layout-content>
            <a-layout-footer class="footer">
                <div class="powere">
                    <span>@ Gelen | Powered By Gelen</span>
                </div>
            </a-layout-footer>
        </a-layout>

        <a-drawer title="辅助工具窗口" width="1200px" :visible="toolVisible" @close="btnClose">
            <Tool></Tool>
        </a-drawer>

    </div>
</template>



<script lang="ts">
import { defineComponent, ref, watch, onMounted, onUnmounted } from 'vue';
import { useRouter } from 'vue-router';
import Tool from "../../components/tool.vue"
import { debounce } from 'lodash'

export default defineComponent({
    components: {
        Tool,
    },
    setup() {
        const toolVisible = ref<boolean>(false);
        const router = useRouter();
        let searchKeyword = ref('')
        let currentMenu = ref('')
        let hiddenMenuFlg = ref(false)
        let menuActive = ref(false)

        const bntBlog = () => {
            router.push("/archives");
        }

        const bntAbout = () => {
            router.push("/about");
        }

        const bntTags = () => {
            router.push("/tags");
        }


        const bntCategory = () => {
            router.push("/category");
        }

        const btnTool = () => {
            toolVisible.value = true
        }

        const btnToAdmin = () => {
            router.push("/adm");
        }


        const btnClose = () => {
            toolVisible.value = false
        }
        const onSearchChange = () => {
            if (searchKeyword.value) {
                router.push({ name: "archives-search", params: { 'keyword': searchKeyword.value } });
            } else {
                router.push("/");
            }
        }

        const disMenu = () => {
            console.log("active")
            menuActive.value = !menuActive.value
        }


        const screenWidth = ref(window.innerWidth)

        const handleResize = debounce(() => {
            screenWidth.value = window.innerWidth
        }, 250)

        onMounted(() => {
            if (screenWidth.value < 800) {
                hiddenMenuFlg.value = true
            } else {
                menuActive.value = true
                hiddenMenuFlg.value = false
            }
            window.addEventListener('resize', handleResize)
        })

        onUnmounted(() => {
            window.removeEventListener('resize', handleResize)
        })

        watch(
            // 监听的变量
            screenWidth,

            // 变量变化时的回调函数
            (newWidth, oldWidth) => {
                if (newWidth < 576) {
                    menuActive.value = true
                    hiddenMenuFlg.value = true
                } else {
                    menuActive.value = false
                    hiddenMenuFlg.value = false
                }
                console.log(`屏幕宽度从 ${oldWidth}px 变为 ${newWidth}px`)
            }
        )

        watch(searchKeyword, (value) => {
            if (value == "") {
                router.push("/");
            }
        })

        return {
            menuActive,
            toolVisible,
            disMenu,
            bntCategory,
            btnClose,
            searchKeyword,
            currentMenu,
            btnTool,
            onSearchChange,
            bntBlog,
            hiddenMenuFlg,
            bntTags,
            btnToAdmin,
            bntAbout,
        }

    }

})
</script>


<style>
.drop-menu {
    width: 100%;
    text-align: center;
    border-top: 1px solid #000;
    background: #fff;
    padding-top: 1em;
    line-height: 5.5em;
    padding-bottom: 1em;
    z-index: 0;
    box-shadow: 0px 2px 4px rgba(0, 0, 0, 0.1), 0px 4px 8px rgba(0, 0, 0, 0.1);
}

.menu-item {
    margin: 0 1em;
    color: #161209;
    text-decoration: none;
    transition: color 0.2s ease, border-color 0.2s ease, background 0.2s ease, opacity 0.2s ease;
    cursor: pointer;

}


.header {
    width: 100%;
    z-index: 0;
    /* position: fixed; */
    transition: all 0.6s ease 0s;
    background: #fff;
    height: 4rem;
}

.blog-ngv {
    white-space: nowrap;
}

.blog-title {
    cursor: pointer;
    white-space: nowrap;
}

.blog-content {
    z-index: 999;
    background-color: #fff;
}

.footer {
    position: fixed;
    bottom: 0;
    left: 0;
    right: 0;
    width: 100%;
    height: 10px;
    text-align: center;
    line-height: 10px;
    background-color: #fff;
}


.input-box {
    position: relative;
    display: inline-block;
}

.input {
    padding: 0 40px 0 20px;
    width: 50px;
    height: 26px;
    font-size: 14px;
    border: 1px solid #eee;
    border-radius: 40px;
    background: #eee;
    transition: width .5s;
    transition-delay: .1s;
}

.span {
    position: absolute;
    top: 20px;
    right: 0px;
    width: 24px;
    height: 24px;
    line-height: 25px;
    padding: 0;
    color: #969696;
    text-align: center;
    background: #2b2a2a;
    border-radius: 50%;
    font-size: 15px;
    cursor: pointer;
}

.input:focus {
    width: 100px;
    outline: none;
    box-shadow: none;
}

.input:focus+.span {
    background-color: pink;
    color: #fff;
}



.ant-layout-header {
    padding: 0 0px;

}
</style>