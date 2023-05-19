<template>
    <div class="treeGlobal">
        <a-row :gutter="20">
            <a-col class="oneCol"  :xs="{ span: 15, offset: 0 }"  :lg="{ span: 4, offset: 0 }">
                <a-row justify="start">
                    <a-col :span="24">
                        <a-input-search v-model:value="searchValue" placeholder="Search">
                            <template #enterButton>
                                <redo-outlined @click="OnFlush" />
                            </template>
                        </a-input-search>
                    </a-col>
                </a-row>
                <a-row>
                    <a-col>
                        <a-tree v-model:expandedKeys="expandedKeys" v-model:selectedKeys="selectedKeys"
                            :auto-expand-parent="autoExpandParent" :show-icon="true" blockNode :tree-data="tdList.list"
                            style="padding-top: 10px;" @expand="Onexpand" @select="onSelect">

                            <template #title="{ isLeaf, id, key: key, title }">
                                <a-dropdown :trigger="['contextmenu']">
                                    <span v-if="title.indexOf(searchValue) > -1">
                                        {{ title.substr(0, title.indexOf(searchValue)) }}
                                        <span style="color: #f50">{{ searchValue }}</span>
                                        {{
                                            title.substr(
                                                title.indexOf(searchValue) + searchValue.length
                                            )
                                        }}
                                    </span>
                                    <span v-else>{{ title }}</span>
                                    <template #overlay>
                                        <a-menu>
                                            <a-menu-item @click="() => ShowAddRootNodeModel()" key="1">添加根节点
                                            </a-menu-item>
                                            <a-menu-item @click="() => ShowAddNodeModel(isLeaf, id)" key="1">添加子节点
                                            </a-menu-item>
                                            <a-menu-item @click="() => OnNewKnowledge(isLeaf, id)" key="1">新建文档
                                            </a-menu-item>
                                            <a-menu-item @click="() => OnModifyClick(isLeaf, id)" key="1">修改节点
                                            </a-menu-item>
                                            <a-menu-item @click="() => OnDeleteClick(isLeaf, id)" key="1">删除节点
                                            </a-menu-item>
                                        </a-menu>
                                    </template>
                                </a-dropdown>
                            </template>
                            <!-- 
                            <template #switcherIcon="{ dataRef, defaultIcon }">
                                <div v-if="!dataRef.isLeaf"><folder-filled /></div>
                                <file-text-filled v-else />
                                <component :is="defaultIcon" v-else />
                            </template> -->

                            <template #icon="{ isLeaf }">
                                <template v-if="!isLeaf">
                                    <folder-filled />
                                </template>
                                <template v-else-if="isLeaf">
                                    <file-text-filled />
                                </template>
                            </template>
                        </a-tree>
                    </a-col>
                </a-row>
            </a-col>
            <!-- <a-col :span="1"></a-col> -->
            <a-col class="twoCol"  :xs="{ span: 0, offset: 0 }"  :lg="{ span: 18, offset: 0 }">
                <a-row>
                    <a-col v-if="showInit" :span="24">
                        <!-- <a-row>
                            <a-button @click="OnAdd" type="primary">
                                <template #icon>
                                    <FileAddOutlined />
                                </template>
                                新增知识页
                            </a-button>
                        </a-row> -->
                        <a-row v-if="showInit">
                            <a-col :span="24" style="margin-top: 10px">
                                <a-timeline mode="alternate" style="margin: 50px 300px 0 0">
                                    <div v-for="(item, index) in tlList" :key="index">
                                        <a-timeline-item>
                                            <div class="information-title" @click="onClickTimeLine(item.key)">{{
                                                item.title
                                            }} | {{ item.updateTime }}</div>
                                        </a-timeline-item>
                                    </div>
                                    <a-timeline-item>
                                        <template #dot>
                                            <ClockCircleOutlined style="font-size: 16px" />
                                        </template>
                                        信息展示
                                    </a-timeline-item>
                                </a-timeline>
                            </a-col>
                        </a-row>
                    </a-col>
                </a-row>
                <!-- 第一行做为工具栏显示 -->
                <a-row :gutter="16">
                    <a-col v-if="showEdit" :xs="{ span: 24, offset: 0 }"  :lg="{ span: 24, offset: 0 }">
                        <a-row :gutter="16">
                            <a-col>
                                <a-button @click="OnSave" type="primary">保存</a-button>
                            </a-col>
                            <a-col>
                                <a-button @click="OnCancel" type="primary">取消</a-button>
                            </a-col>
                        </a-row>
                        <a-row :gutter="16" style="margin-top: 10px; margin-bottom: 10px">
                            <a-col style="padding: 5px">标题：</a-col>
                            <a-col>
                                <a-input style="width: 300px" v-model:value="MdEdit.title" placeholder="输入标题." />
                            </a-col>
                            <a-col style="padding: 5px">TAG：</a-col>
                            <a-col>
                                <a-select v-model:value="seleteTags" mode="tags" style="width: 200px"
                                    :token-separators="[',']" placeholder="tags"></a-select>
                            </a-col>
                        </a-row>

                        <a-row>
                            <a-col :xs="{ span: 24, offset: 0 }"  :lg="{ span: 24, offset: 0 }">
                                <div>
                                    <md-editor style="padding: 10px;height:600px" v-model="MdEdit.content"
                                        :showCodeRowNumber="true" @onUploadImg="onUploadImg" preview-theme="default" />
                                </div>
                            </a-col>
                        </a-row>
                    </a-col>

                    <a-col v-if="showPriview" :span="24">
                        <a-row>
                            <a-col :span="16">
                                <span style="margin-left: 8px;font-weight:bold;font-size:20px">{{ kwge.title }}</span>
                            </a-col>

                            <a-col :span="2">
                                <a-button @click="OnEdit" type="primary">编辑</a-button>
                            </a-col>
                            <a-col :span="2">
                                <a-button @click="OnEditReturn" type="primary">返回</a-button>
                            </a-col>
                            <a-col :span="2">
                                <a-button @click="OnShare" type="primary">{{ share }}</a-button>
                            </a-col>
                        </a-row>
                        <a-row :gutter="20">
                            <a-col>
                                <span style="margin-left: 8px;font-size:5px">作者:{{ kwge.author
                                }}</span>

                            </a-col>
                            <a-col>
                                <span style="margin-left: 8px;font-size:5px">更新时间:{{ kwge.updateTime
                                }}</span>
                            </a-col>
                            <a-col>
                                tag:
                                <span style="font-size:5px">{{ kwge.tags }}</span>
                            </a-col>
                        </a-row>
                        <a-row>
                            <a-col :span="24"
                                style=" margin-top: 10px; border-style:solid; border-width:2px;border-color:#e2e2e2;">
                                <div>
                                    <md-editor editor-id="my-editor" style="padding: 10px;" v-model="kwge.content"
                                        :previewOnly="viewOnly" :showCodeRowNumber="true"
                                        preview-theme="default"></md-editor>
                                </div>
                                <!-- @onGetCatalog="onGetCatalogs" -->
                            </a-col>
                        </a-row>
                    </a-col>
                </a-row>
            </a-col>
        </a-row>
    </div>
    <a-modal v-model:visible="showModifyTreeNode" title="修改节点信息" ok-text="确认" cancel-text="取消" @ok="ModifyTreeNode">
        <a-row>
            <a-col style="margin-top:5px;">
                <label>名称：</label>
            </a-col>
            <a-col>
                <a-input v-model:value="modifyTreeValue" placeholder />
            </a-col>
        </a-row>
    </a-modal>

    <a-modal v-model:visible="modelOption.showAddNodeModel" title="添加节点" ok-text="确认" cancel-text="取消" @ok="AddNode">
        <a-row>
            <a-col style="margin-top:5px;">
                <label>名称：</label>
            </a-col>
            <a-col>
                <a-input v-model:value="newNodename" placeholder />
            </a-col>
        </a-row>
    </a-modal>

    <a-modal v-model:visible="modelOption.ShowRootNodeModel" title="添加根节点" ok-text="确认" cancel-text="取消" @ok="AddRootNode">
        <a-row>
            <a-col style="margin-top:5px;">
                <label>名称：</label>
            </a-col>
            <a-col>
                <a-input v-model:value="newNodename" placeholder />
            </a-col>
        </a-row>
    </a-modal>
</template>


<script lang="ts">
interface TreeProps {
    id?: number;
    parentId?: number;
    title?: string;
    key?: string;
    level?: number;
    isLeaf?: boolean;
    children?: TreeProps[];
}

interface Konwledge {
    author?: string
    key?: string;
    title?: string;
    content?: string;
    updateTime?: string;
    shared?: boolean
    tags?: Array<string>
    parentId?: number
}


interface NodeOption {
    showAddNodeModel: boolean;
    ShowRootNodeModel: boolean;
}

interface HeadList {
    text: string;
    level: 1 | 2 | 3 | 4 | 5 | 6;
}

import { defineComponent, ref, onMounted, reactive, watch, computed, h, } from "vue";
import knowledgeService from "../../service/knowledge.service";
import { FileAddOutlined, ClockCircleOutlined, RedoOutlined, FolderFilled, FileTextFilled, FolderOpenFilled } from "@ant-design/icons-vue";

import { message, notification, Button } from "ant-design-vue";
import MdEditor from 'md-editor-v3';
import Catelog from '@/components/Catalog.vue'
import 'md-editor-v3/lib/style.css';
import useClipboard from 'vue-clipboard3'


const { toClipboard } = useClipboard()

export default defineComponent({
    components: {
        FolderFilled,
        FolderOpenFilled,
        FileTextFilled,
        Catelog,
        // MdCatalog,
        FileAddOutlined,
        RedoOutlined,
        MdEditor,
        ClockCircleOutlined,
    },
    setup() {
        let searchValue = ref("");
        const expandedKeys = ref<(string | number)[]>([]);
        const selectedKeys = ref<string[]>([]);
        let tdList: { list: TreeProps[] } = reactive({ list: [] });

        let newNodename = ref('')

        let seleteTags = ref<string[]>([]);
        let showCatlog = ref(false)
        let share = ref("分享")
        const kwge = ref<Konwledge>({});
        let MdEdit: Konwledge = reactive({});
        let tlList: Konwledge[] = reactive([]);
        let category = ref("")

        let autoExpandParent = ref<boolean>(true);
        let ShowNodeKey: (string | number)[] = [];
        let CurSelectKey: string;
        let showPriview = ref(false);
        let modifyTreeValue = ref('')
        let showInit = ref(true);
        let showSelect = ref(true)
        let showEdit = ref(false);
        let parentId = 0

        let curSelectNodeId = 0
        let curlSelectNodeIsLeaf = true

        let Editing = ref(true);
        let modifyCurTreeKey = ref('')
        let viewOnly = ref(true)
        let IsAdd = false;
        let showModifyTreeNode = ref(false);
        let modelOption: NodeOption = reactive({
            showAddNodeModel: false,
            ShowRootNodeModel: false,
        })


        const OnNewKnowledge = (isLeaf: boolean, _parentid: number) => {

            MdEdit.title = ''
            MdEdit.content = ''
            seleteTags.value = []

            parentId = _parentid
            showEdit.value = true;
            showPriview.value = false;
            showInit.value = false;
            IsAdd = true;
            showSelect.value = true;

        };

        const Onexpand = (exkey: any, event: any) => {
            expandedKeys.value = exkey;
            autoExpandParent.value = false;

        };

        const copyRuleCode = async (item: any) => {
            try {
                //item为要复制内容
                await toClipboard(item)
                //复制成功提示
                message.success('复制成功')
            } catch (e) {
                //复制失败回调
                console.error(e)
                message.warning('复制失败')
            }
        }


        const filterOption = (input: string, option: any) => {
            return option.label.toLowerCase().indexOf(input.toLowerCase()) >= 0;

        }


        async function onUploadImg(files: FileList, callback: (urls: string[]) => void) {
            let urls: string[] = []

            Array.from(files).map(async (file) => {
                const form = new FormData();
                form.append('mkimg', file);
                await knowledgeService.UploadImg(form).then((res: any) => {
                    let url = process.env.VUE_APP_API_URL + res.data.data
                    urls.push(url)
                    callback(urls)
                }).catch((res: any) => {
                    console.log()
                })

            })
        }


        const Init = () => {
            knowledgeService.getTree().then((res: any) => {
                if (res.data.data) {
                    let jdata = res.data.data;
                    if (jdata.length < 0) {
                        return;
                    }
                    tdList.list = jdata;
                }

            });
            knowledgeService.getSummary().then((res: any) => {
                if (res.data.data) {
                    let _kList: Konwledge = {};
                    let jdata = res.data.data;
                    if (jdata.length < 0) {
                        return;
                    }
                    for (let i = 0; i < jdata.length; i++) {
                        _kList.title = jdata[i].title;
                        _kList.key = jdata[i].key;
                        _kList.updateTime = jdata[i].updateTime;
                        tlList.push(_kList);
                        _kList = {};
                    }
                }

            });
        };




        onMounted(() => {
            Init();
        });

        const getParentKey = (
            key: string | number,
            tree: TreeProps[]
        ): string | number | undefined => {
            let parentKey;
            let xtree = tree || [];

            for (let i = 0; i < xtree.length; i++) {
                const node = xtree[i];
                if (node.children) {
                    if (node.children.some((item) => item.key === key)) {
                        parentKey = node.key;
                    } else if (getParentKey(key, node.children)) {
                        parentKey = getParentKey(key, node.children);
                    }
                }
            }
            return parentKey;
        };

        const getMatchKey = (node: TreeProps[]) => {
            if (!node) {
                return;
            }
            for (let i = 0; i < node.length; i++) {
                if ((node[i].title || "").indexOf(searchValue.value) > -1) {
                    let xkey = node[i].key || "";
                    // ShowNodeKey.push(xkey)
                    let pkey = getParentKey(xkey, tdList.list) || "";
                    ShowNodeKey.push(pkey);
                }
                let _child = node[i].children || [];
                if (_child) {
                    getMatchKey(_child);
                }
            }
        };



        watch(searchValue, (value) => {
            if (value.length > 0) {
                getMatchKey(tdList.list);
                searchValue.value = value;
                expandedKeys.value = ShowNodeKey;
                ShowNodeKey = [];
                autoExpandParent.value = true;
            } else {
                searchValue.value = value;
                expandedKeys.value = [];
                ShowNodeKey = [];
                autoExpandParent.value = true;
            }
        });


        const ShowAddRootNodeModel = () => {
            modelOption.ShowRootNodeModel = true;
        };

        const ShowAddNodeModel = (isLeaf: boolean, _parentId: number) => {
            parentId = _parentId
            if (!isLeaf) {
                modelOption.showAddNodeModel = true;
            } else {
                message.warning("叶子节点无法添加节点")
            }
        };

        const AddNode = () => {
            knowledgeService.AddNode(parentId, newNodename.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    modelOption.showAddNodeModel = false;
                    const _tree = res.data.data
                    addNode(tdList.list, parentId, _tree);
                }
            })

        };


        const AddRootNode = () => {
            modelOption.ShowRootNodeModel = false;
            knowledgeService.AddRootNode(newNodename.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    modelOption.showAddNodeModel = false;

                    let tp: TreeProps[] = [res.data.data]
                    tdList.list = tdList.list?.concat(tp)


                }
            })

        };



        const OnEditReturn = () => {
            showCatlog.value = false
            showInit.value = true;
            Editing.value = true
            showEdit.value = false;
            showPriview.value = false;
        };

        const OnEdit = () => {

            Editing.value = false;
            showEdit.value = true;
            showPriview.value = false;
            showInit.value = false;
            showSelect.value = false;

            // console.log(kwge)
            MdEdit.title = kwge.value.title
            MdEdit.content = kwge.value.content
            if (kwge.value.tags ){
                seleteTags.value = kwge.value.tags
            }
           
        };


        const OnAdd = () => {
            showEdit.value = true;
            showPriview.value = false;
            showInit.value = false;
            IsAdd = true;
            showSelect.value = true;

        };

        const onSelect = (exkey: any, expand: any) => {
            share.value = "分享"
            if (expand.node.isLeaf) {
                showCatlog.value = true
                CurSelectKey = expand.node.key;
                showEdit.value = false;
                showPriview.value = true;
                showInit.value = false;
                knowledgeService.GetKnowledgeByKey(CurSelectKey).then((res: any) => {
                    let jdata = res.data.data;
                    kwge.value = jdata
                });
            }
        };

        const OnCancel = () => {
            showCatlog.value = false
            showEdit.value = false;
            showPriview.value = false;
            showInit.value = true;
            Editing.value = true

            kwge.value = {};
            seleteTags.value = []

        };

        const OnSave = () => {

            let title = MdEdit.title || "";
            let content = MdEdit.content || "";
            if (title == "" || content == "") {
                message.info("标题或内容未填写.");
                return;
            }
            if (IsAdd) {
                kwge.value.shared = false
                knowledgeService
                    .SaveNewKnowledge(seleteTags.value, parentId, title, content)
                    .then((res: any) => {
                        if (res.data.code == 400) {
                            message.error(res.data.msg);
                        } else if (res.data.code == 200) {
                            message.success(res.data.msg);
                            let kdata = res.data.data;
                            showPriview.value = true;
                            Editing.value = true
                            showEdit.value = false;
                            showInit.value = false;
                            expandedKeys.value.push(kdata.key)
                            selectedKeys.value.push(kdata.key)
                            CurSelectKey = kdata.knowledge.key
                            kwge.value = kdata.knowledge

                            MdEdit.title = ''
                            MdEdit.content = ''
                            seleteTags.value = []
                            //更新目录树
                            addNode(tdList.list, parentId, kdata.category);

                        }
                    });
            } else {
                knowledgeService
                    .SaveEditKnowledge(
                        category.value,
                        seleteTags.value,
                        parentId,
                        title,
                        content,
                        CurSelectKey
                    )
                    .then((res: any) => {
                        if (res.data.code == 400) {
                            message.error(res.data.msg);
                        } else if (res.data.code == 200) {
                            message.success(res.data.msg);
                            Editing.value = true;
                            showPriview.value = true;
                            showEdit.value = false;
                            showInit.value = false;
                            const data = res.data.data.category

                            modifyNode(tdList.list, data.id, data);

                            // UpdateTreeMenu()
                            kwge.value.title = title
                            kwge.value.content = content
                            kwge.value.tags = seleteTags.value
                            MdEdit.title = ''
                            MdEdit.content = ''
                            seleteTags.value = []

                        }
                    });
            }
            IsAdd = false;

        };

        const OnFlush = () => {
            // UpdateTreeMenu()
            tlList.length = 0
            Init()
        }


        // 添加节点
        function addNode(tree: TreeProps[], parentId: number, node: TreeProps): void {
            for (const item of tree) {
                if (item.id === parentId) {
                    if (!item.children) {
                        item.children = [node];
                    } else {
                        item.children.push(node);
                    }
                    return;
                } else if (item.children) {
                    addNode(item.children, parentId, node);
                }
            }
        }

        // 修改节点
        function modifyNode(tree: TreeProps[], id: number, modifiedNode: TreeProps): void {
            for (const item of tree) {
                if (item.id === id) {
                    Object.assign(item, modifiedNode);
                    return;
                } else if (item.children) {
                    modifyNode(item.children, id, modifiedNode);
                }
            }
        }

        // 删除节点
        function deleteNode(tree: TreeProps[], id: number): void {
            for (const item of tree) {
                if (item.children) {
                    const index = item.children.findIndex(child => child.id === id);
                    if (index !== -1) {
                        item.children.splice(index, 1);
                        return;
                    } else {
                        deleteNode(item.children, id);
                    }
                }
            }
        }

        const OnDeleteClick = (isLeaf: boolean, id: number) => {
            knowledgeService.DelTreeNode(isLeaf, id).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    message.success(res.data.msg);
                    deleteNode(tdList.list, id)

                }
            })

        }
        const ModifyTreeNode = () => {
            knowledgeService.ModifyTreeNode(curlSelectNodeIsLeaf, curSelectNodeId, modifyTreeValue.value).then((res: any) => {
                if (res.data.code == 400) {
                    message.error(res.data.msg);
                } else if (res.data.code == 200) {
                    message.success(res.data.msg);
                    //更新目录树的信息，前端直接更新，不使用接口
                    console.log(tdList.list)
                    modifyNode(tdList.list, curSelectNodeId, { title: modifyTreeValue.value });

                }
            })
            showModifyTreeNode.value = false
        };

        const OnModifyClick = (isLeaf: boolean, id: number) => {
            curSelectNodeId = id
            curlSelectNodeIsLeaf = isLeaf
            showModifyTreeNode.value = true
        }


        watch(kwge, (newkwge) => {
            if (newkwge.shared) {
                share.value = "取消分享"
            } else {
                share.value = "分享"
            }
        })

        const onClickTimeLine = (key: string | undefined) => {
            showCatlog.value = true
            let ckey = key || ''
            CurSelectKey = ckey;
            showEdit.value = false;
            // showCatlog.value = false
            showPriview.value = true;
            showInit.value = false;
            knowledgeService.GetKnowledgeByKey(ckey).then((res: any) => {
                let jdata = res.data.data;
                kwge.value = jdata
                seleteTags.value = jdata.tags
            });
        }

        // const state: { catelogs: HeadList[] } = reactive({ catelogs: [] })
        // const onGetCatalogs = (list: HeadList[]) => {
        //     state.catelogs = list
        // }

        // const catalogList = computed(() => {
        //     return state.catelogs
        // })


        const OnShare = () => {
            if (share.value == "分享") {
                knowledgeService.ShareKnowledge(CurSelectKey).then((res: any) => {
                    if (res.data.code == 200) {
                        message.success(res.data.msg)
                        share.value = "取消分享"
                        const key = `open${Date.now()}`;
                        let url = res.data.data
                        notification.open({
                            message: '分享知识',
                            description:
                                url,
                            btn: () =>
                                h(
                                    Button,
                                    {
                                        type: 'primary',
                                        size: 'small',
                                        onClick: () => {
                                            copyRuleCode(url)
                                            notification.close(key)
                                        },
                                    },
                                    { default: () => '点击复制' },
                                ),
                            key,
                            onClose: close,
                        });
                    } else if (res.data.code == 401) {
                        message.warning(res.data.msg)
                    } else {
                        message.error(res.data.msg)
                    }
                });
            } else {
                knowledgeService.cancelShareByKey(CurSelectKey).then((res: any) => {
                    if (res.data.code == 400) {
                        message.error(res.data.msg)
                    } else if (res.data.code == 200) {
                        message.success(res.data.msg)
                        share.value = "分享"
                    }
                })
            }
        }

        return {
            // onGetCatalogs,
            OnShare,
            // catalogList,
            seleteTags,
            expandedKeys,
            selectedKeys,
            searchValue,
            tlList,
            kwge,
            MdEdit,
            Editing,
            onClickTimeLine,
            onUploadImg,
            OnModifyClick,
            viewOnly,
            ModifyTreeNode,
            OnDeleteClick,
            AddRootNode,
            tdList,
            autoExpandParent,
            showPriview,
            newNodename,
            showInit,
            showEdit,
            modelOption,
            share,
            showModifyTreeNode,
            showSelect,
            modifyCurTreeKey,
            modifyTreeValue,
            OnEdit,
            ShowAddRootNodeModel,
            OnFlush,
            filterOption,
            OnSave,
            OnEditReturn,
            AddNode,
            showCatlog,
            OnAdd,
            ShowAddNodeModel,
            OnCancel,
            onSelect,
            OnNewKnowledge,
            Onexpand,
        };
    },
});
</script>


<style>
.ant-tree-node-content-wrapper {

    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    display: -webkit-box;
    overflow: hidden;
    text-overflow: ellipsis;
}


.editor {
    width: 100%;
    margin-top: 10px;
    border: 2px solid rgb(240, 242, 245);
    /* box-shadow: 2px 2px 5px #000; */
}

.information-title {
    cursor: pointer;
    /*鼠标悬停变小手*/
}

.oneCol {
    /* margin-right: 5px; */
    /* padding: 15px; */
    background: #fff;
    border: 2px solid rgb(240, 242, 245);

}

.twoCol {
    /* padding: 10px; */
    background: #fff;
}

.treeGlobal {
    /* background: red; */
    padding-left: 10px;
}

.catalogTree {
    font-weight: bold;
    position: fixed;
    top: 20%;
    right: 0%;
    min-width: 200px;
    max-width: 220px;
    background: #fff;
    word-wrap: break-word;
}

/* .ant-tree-title {
    width: 250px;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    display: -webkit-box;
    overflow: hidden;
    text-overflow: ellipsis;
} */
</style>
