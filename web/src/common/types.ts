interface Columns {
    title: string
    dataIndex?: string
    key: string
    align?: string
    fixed?: string
    width?: number
    minWidth?: number
    maxWidth?: number
    resizable?: boolean
    ellipsis?: boolean
}





class Types {

    getAgentColumns() {
        let column: Columns[]
        return column = [
            {
                title: "主机名",
                dataIndex: "host_name",
                key: "2",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "操作系统",
                dataIndex: "platform",
                key: "2",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "状态",
                dataIndex: "status",
                key: "3",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "资源使用",
                dataIndex: "source_usage",
                key: "4",
                align: 'center',
            },
            {
                title: "更新时间",
                dataIndex: "update_time",
                key: "11",
                align: 'center',
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
            }
        ];
    }

    getComapnyTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                resizable: true,
                width: 80,
            },
            {
                title: "类型",
                dataIndex: "projectType",
                key: "2",
                align: 'center',
                minWidth: 100,
                maxWidth: 200,
                resizable: true,
                width: 100,
            },
            {
                title: "公司名称",
                dataIndex: "companyName",
                key: "3",
                align: 'center',
                resizable: true,

                width: 150,
                ellipsis: true,
            },
            {
                title: "域名",
                dataIndex: "domain",
                key: "4",
                align: 'center',
                resizable: true,

                width: 200,
                ellipsis: true,
            },
            {
                title: "SRC地址",
                dataIndex: "srcUrl",
                resizable: true,

                key: "5",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "关键字",
                dataIndex: "keyWord",
                key: "6",
                resizable: true,

                align: 'center',
                width: 100,
                ellipsis: true,
            },
            // {
            //     title: "监控状态",
            //     resizable: true,

            //     dataIndex: "monitorStatus",
            //     key: "7",
            //     align: 'center',
            //     width: 90,
            // },
            // {
            //     title: "监控频率",
            //     resizable: true,

            //     dataIndex: "monitorRate",
            //     key: "8",
            //     align: 'center',
            //     width: 90,
            // },
            {
                title: "创建时间",
                resizable: true,

                dataIndex: "lastUpdateTime",
                key: "9",
                align: 'center',
                width: 150,
            },
            {
                title: '操作',
                resizable: true,

                key: 'operation',
                fixed: 'right',
                // align: 'center',
                width: 200,
            }
        ];
    }

    getDomainTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "Cid",
                dataIndex: "cid",
                key: "2",
                align: 'center',
                width: 100,
            },
            {
                title: "域名",
                dataIndex: "domain",
                key: "3",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: "来源",
                dataIndex: "source",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "5",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "备注",
                dataIndex: "remarks",
                key: "6",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                //  align: 'center',
                width: 200,
            }
        ];
    }


    getShareTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "标题",
                dataIndex: "title",
                key: "2",
                align: 'center',
                width: 100,
            },
            {
                title: "key",
                dataIndex: "key",
                key: "3",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "5",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                //  align: 'center',
                width: 200,
            }
        ];
    }

    getGatherDomainTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "Cid",
                dataIndex: "cid",
                key: "2",
                align: 'center',
                width: 100,
            },
            {
                title: "域名",
                dataIndex: "domain",
                key: "3",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: "来源",
                dataIndex: "source",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "5",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            }
        ];
    }

    getWebsiteTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "网站",
                dataIndex: "website",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "favicon",
                dataIndex: "favicon",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "标题",
                dataIndex: "title",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "ips",
                dataIndex: "ips",
                key: "5",
                align: 'center',
                width: 120,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 60,
                ellipsis: true,
            },
            // {
            //     title: "指纹",
            //     dataIndex: "finger",
            //     key: "7",
            //     align: 'center',
            //     width: 100,
            //     ellipsis: true,
            // },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            },

            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 200,
            }
        ];
    }

    getGatherServiceTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "#",
                dataIndex: "cid",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "服务名",
                dataIndex: "service",
                key: "2",
                align: 'center',
                width: 150,
            },
            {
                title: "主机/IP",
                dataIndex: "host",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "端口",
                dataIndex: "port",
                key: "4",
                align: 'center',
                width: 60,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "5",
                align: 'center',
                width: 120,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            }
        ];
    }

    getServiceTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "#",
                dataIndex: "cid",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "服务名",
                dataIndex: "service",
                key: "2",
                align: 'center',
                width: 150,
            },
            {
                title: "主机/IP",
                dataIndex: "host",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "端口",
                dataIndex: "port",
                key: "4",
                align: 'center',
                width: 60,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "5",
                align: 'center',
                width: 120,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            },
            {
                title: "备注",
                dataIndex: "remarks",
                key: "6",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 200,
            }
        ];
    }

    getNoticeTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "标题",
                dataIndex: "title",
                key: "2",
                align: 'center',
                width: 150,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 100,
            }
        ];
    }

    getPocTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "标题",
                dataIndex: "pocName",
                key: "2",
                align: 'center',
                width: 150,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 100,
            }
        ];
    }

    getTagColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "TAG名称",
                dataIndex: "key",
                key: "2",
                align: 'center',
                width: 150,
            },

        ];
    }

    getUserTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 60,
            },
            {
                title: "用户名",
                dataIndex: "username",
                key: "2",
                align: 'center',
                width: 150,
            },
            {
                title: "创建时间",
                dataIndex: "createTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 100,
            }
        ];
    }

    getTaskTableColumns() {
        let columns: Columns[]
        return columns = [
            // {
            //     title: "#",
            //     dataIndex: "taskId",
            //     key: "1",
            //     align: 'center',
            //     width:300,
            // },
            {
                title: "任务名称",
                dataIndex: "taskName",
                key: "2",
                align: 'center',
            },
            {
                title: "任务类型",
                dataIndex: "taskType",
                key: "3",
                align: 'center',
            },
            {
                title: "频率",
                dataIndex: "rate",
                key: "4",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "公司范围",
                dataIndex: "companyRange",
                key: "5",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "范围",
                dataIndex: "assetRange",
                key: "6",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "域名搜集方式",
                dataIndex: "domainCollectionType",
                key: "7",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "公司",
                dataIndex: "companyId",
                key: "8",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "POC",
                dataIndex: "pocName",
                key: "8",
                align: 'center',
                ellipsis: true,
            },
            {
                title: "最后执行时间",
                dataIndex: "updateTime",
                key: "9",
                align: 'center',
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                width: 200,
                align: 'center',
            }
        ];
    }

    getGatherWebsiteTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "网站",
                dataIndex: "website",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "favicon",
                dataIndex: "favicon",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "标题",
                dataIndex: "title",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "ips",
                dataIndex: "ips",
                key: "5",
                align: 'center',
                width: 120,
                ellipsis: true,
            },
            {
                title: "new?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 60,
                ellipsis: true,
            },
            {
                title: "指纹",
                dataIndex: "finger",
                key: "7",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "8",
                align: 'center',
                width: 150,
                ellipsis: true,
            }

        ];
    }

    getIPsTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "ip",
                dataIndex: "ip",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "os",
                dataIndex: "os",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "NEW?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "备注",
                dataIndex: "remarks",
                key: "6",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 220,
            }
        ];
    }

    getMiniTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "#",
                dataIndex: "cid",
                key: "2",
                align: 'center',
                width: 100,
            },
            {
                title: "小程序名",
                dataIndex: "name",
                key: "3",
                align: 'center',
                width: 250,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "NEW?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "备注",
                dataIndex: "remarks",
                key: "6",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 220,
            }
        ];
    }

    getWechatOfficeAccountTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "#",
                dataIndex: "cid",
                key: "2",
                align: 'center',
                width: 100,
            },
            {
                title: "公众号名称",
                dataIndex: "name",
                key: "3",
                align: 'center',
                width: 250,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "NEW?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "备注",
                dataIndex: "remarks",
                key: "6",
                align: 'center',
                width: 300,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 220,
            }
        ];
    }

    getGatherIPsTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "ip",
                dataIndex: "ip",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "os",
                dataIndex: "os",
                key: "3",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "4",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "NEW?",
                dataIndex: "isNew",
                key: "6",
                align: 'center',
                width: 100,
                ellipsis: true,
            }
        ];
    }

    getBugReportTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "Detail",
                dataIndex: "detail",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "3",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: "NEW?",
                dataIndex: "isNew",
                key: "4",
                align: 'center',
                width: 100,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 200,
            }
        ];
    }


    getBugCollectTableColumns() {
        let columns: Columns[]
        return columns = [
            {
                title: "#",
                dataIndex: "id",
                key: "1",
                align: 'center',
                width: 100,
            },
            {
                title: "Detail",
                dataIndex: "detail",
                key: "2",
                align: 'center',
                width: 300,
            },
            {
                title: "更新时间",
                dataIndex: "updateTime",
                key: "3",
                align: 'center',
                width: 200,
                ellipsis: true,
            },
            {
                title: '操作',
                key: 'operation',
                fixed: 'right',
                align: 'center',
                width: 100,
            }
        ];
    }
}




const types = new Types()

export default types;