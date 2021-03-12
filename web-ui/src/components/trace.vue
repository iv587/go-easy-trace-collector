<template>
<div>
    <!-- 抽屉 -->
    <Drawer placement="left" v-model="drawerOpen" :transfer="false" :inner="true" :width="650">
        <Divider orientation="center">基本信息</Divider>
        <div >
            <Row>
                <i-col span="4"><h4>应用名称:</h4></i-col><i-col span="6">{{spanInfo.application}}</i-col>
            </Row>
            <Row>
                <i-col span="4"><h4>应用租:</h4></i-col><i-col span="6">{{spanInfo.applicationGroup}}</i-col>
            </Row>
            <Row>
                <i-col span="4"><h4>实例:</h4></i-col><i-col span="18">{{spanInfo.appInstance}}</i-col>
            </Row>
            <Row>
                <i-col span="4"><h4>Method:</h4></i-col><i-col span="18">{{spanInfo.operationName}}</i-col>
            </Row>
        </div >
        <Divider orientation="center">Span标记</Divider>
        <div>
            <Row v-for="(v,i) in spanInfo.tagObj" :key="i">
                <i-col span="10"><h4>{{i}}:</h4></i-col><i-col span="14"><p style="word-break: break-all;overflow-wrap: break-word;">{{v}}</p></i-col>
            </Row>
        </div>
        <Divider orientation="center">Log</Divider>
        <Card v-for="(v,i) in spanInfo.logArray" :key="i">
            <p slot="title"><Time type="datetime" :time="v.timestampMicros"></Time></p>
            <div v-for=" (vv,ii) in v.fields" :key="ii">
                <Divider orientation="center">{{ii}}</Divider>
                <pre style="white-space: pre-wrap;word-wrap: break-word;">{{vv}}</pre>
            </div>
        </Card>
    </Drawer>
    <Row>
        <i-col span="24">
            <Card :dis-hover="true">
                <Row>
                    <Form :model="search">
                        <i-col span="6">
                            <FormItem label="选择日期">
                                <DatePicker @on-change="searchDateHandler" :value="new Date()"
                                            format="yyyy-MM-dd" type="date" :clearable="false"
                                            placeholder="选择日期" style="width: 200px"></DatePicker>
                            </FormItem>
                        </i-col>
                        <i-col span="4">
                            <FormItem label="状态查询:">
                                <Select v-model="search.error" style="width: 100px">
                                    <Option :value="-1">全部</Option>
                                    <Option :value="0">成功</Option>
                                    <Option :value="1">失败</Option>
                                </Select>
                            </FormItem>
                        </i-col>
                        <i-col span="5">
                            <FormItem label="端点搜索:">
                                <Input v-model="search.operationKey" style="width: 150px"/>
                            </FormItem>
                        </i-col>
                        <i-col span="3">
                            <Cascader v-model="appValue" placeholder="选择应用" :data="appData" @on-change="appChangeHandler" @on-visible-change="loadAppData"></Cascader>
                        </i-col>
                        <i-col span="6">
                            <Button @click="searchData" style="margin-left: 20px" type="info">刷新</Button>
                        </i-col>
                    </Form>
                </Row>
            </Card>
        </i-col>
    </Row>
    <Row>
        <i-col span="6">
            <Card :dis-hover="true">
                <div slot="title">
                    <Page slot="title" :current="search.pageNo" :page-size="search.size"
                          :total="search.total" simple :show-total="true" @on-change="pageChangeHandler"/>
                </div>
                <div style="height: calc(100vh - 240px) ; overflow-y: auto ;">
                    <div :ref="item.id+'div'" class="trace-div-cls" v-for="item in traceList" :key="item.id"
                         @click="getTraceTree(item.traceId,item.id)">
                        <Card :ref="item.id+'-card'">
                            <p class="span-operation-p-cls">
                                <a href="javascript:void(0)" @click="getTraceTree(item.traceId)"
                                   :style="{color: item.error === 0 ? 'green' : 'red' }">
                                    <Tag color="blue">{{item.component}}</Tag>
                                    {{item.operationName}}
                                </a>
                            </p>
                            <p>
                                <Tag color="cyan">
                                    <Time :time="new Date(item.startTime)" type="datetime"/>
                                </Tag>&nbsp;
                                <Tag v-if="item.elapsedTime < 2000 " color="green">{{item.elapsedTime}}ms
                                </Tag>
                                <Tag v-if="item.elapsedTime > 2000 " color="orange">{{item.elapsedTime}}ms
                                </Tag>
                            </p>
                            <div>
                                <Tag color="geekblue">{{item.applicationGroup}}</Tag>
                                <Tag color="geekblue">{{item.application}}</Tag>
                            </div>
                        </Card>
                    </div>
                </div>
            </Card>
        </i-col>
        <i-col span="18">
            <Card :dis-hover="true" style="height: calc(100vh - 153px) ; overflow-y: auto ;">
                <Divider/>
                <Table row-key="id" :columns="traceColumnNames" :data="traceData" border
                       :highlight-row="true" @on-row-click="rowClick"></Table>
            </Card>
            <Spin size="large" fix v-if="loading"></Spin>
        </i-col>
    </Row>
</div>
</template>

<script>
import {getAppApi, getSpanByIdApi, getTraceListApi, getTraceTreeApi} from "../lib/api";

export default {
    name: 'trace',
    data() {
        return {
            search: {
                pageNo: 1,
                error: -1,
                size: 10,
                day: '',
                total: 0,
                operationKey: ''
            },
            spanInfo: {},
            drawerOpen: false,
            appData: [],
            appValue: [],
            loading: false,
            rootStartTime: 0,
            rootElapsedTime: 0,
            traceColumnNames: [
                {
                    title: 'Method',
                    key: 'operationName',
                    tree: true,
                    width: 350,
                    tooltip: true,
                    render: (h, params) => {
                        let a = h('span', {
                            style: {
                                color: params.row.error === 0 ? 'green' : 'red'
                            }
                        }, params.row.operationName);
                        let n = h('Tag', {
                            props: {
                                color: 'blue',
                            },
                            style: {
                                marginLeft: "5px"
                            }
                        }, params.row.component);
                        return h('div', {
                            style: {
                                marginLeft: params.row.deepth * 5 + "%",
                            }
                        }, [a, n]);
                    }
                },
                {
                    title: '时间',
                    align: 'center',
                    key: 'startTimeText',
                    width: 120,
                },
                {
                    title: '耗时(ms)',
                    key: 'elapsedTime',
                    align: 'center',
                    width: 120,
                    render: (h, params) => {
                        return h('Tag', {
                            props: {
                                color: params.row.elapsedTime >= 2000 ? 'orange' : 'green',
                            },
                        }, params.row.elapsedTime + 'ms')
                    }
                },
                {
                    title: '时间轴',
                    key: 'timeLine',
                    render: (h, params) => {
                        let elapsedTime = params.row.elapsedTime;
                        let startTime = params.row.startTime;
                        let parentStartTime = params.row.parentStartTime;
                        let parentElapsedTime = params.row.parentElapsedTime;
                        let width = '100%';
                        let marginLeft = '0';
                        if (parentStartTime > 0) {
                            parentStartTime = this.rootStartTime;
                            parentElapsedTime = this.rootElapsedTime;
                            marginLeft = ((startTime - parentStartTime) / parentElapsedTime) * 100 + '%';
                            width = (elapsedTime / parentElapsedTime) * 100 + '%';
                        }

                        let showContent = h('div', {
                            style: {
                                width: width,
                                marginLeft: marginLeft
                            },
                            class: {
                                'time-line-cls-server': params.row.spanKind === 'server',
                                'time-line-cls-client': params.row.spanKind === 'client'
                            },
                        })
                        let content = h('div', {
                            slot: 'content'
                        }, [
                            h('p', '时间:' + params.row.startTimeText + '~' + params.row.finishTimeText),
                            h('p', '耗时' + elapsedTime + 'ms')
                        ])


                        let toolTip = h('Tooltip', {
                            style: {
                                width: '100%'
                            },
                            props: {
                                transfer: true,
                            },
                        }, [content, showContent])
                        return toolTip;
                    }
                },
                {
                    title: '服务名',
                    key: 'application',
                    align: 'center',
                    width: 120,
                    render: (h, params) => {
                        let t1 = h('Tag', {
                            props: {
                                color: 'geekblue',
                            },
                        }, params.row.application);
                        let t2 = h('Tag', {
                            props: {
                                color: 'geekblue',
                            },
                        }, params.row.applicationGroup);
                        return h('div', {}, [t1, t2])
                    }
                }
            ],
            traceData: [],
            traceList: []
        }
    },
    methods: {
        rowClick(row) {
            let _this = this;
            getSpanByIdApi({id: row.id}).then(res => {
                _this.spanInfo = res.data
                _this.spanInfo.tagObj = JSON.parse(res.data.tags)
                _this.spanInfo.logArray = JSON.parse(res.data.logDatas)
                _this.drawerOpen = true;
            })
        },
        searchDateHandler(it) {
            this.search.day = it;
        },
        pageChangeHandler(num) {
            this.search.pageNo = num;
            this.getTraceList();
        },
        appChangeHandler(value) {
            this.appValue = value;
            this.searchData()
        },
        searchData() {
            this.search.pageNo = 1;
            this.getTraceList();
        },
        getTraceList() {
            let _this = this;
            if (this.appValue.length > 0) {
                this.search.applicationGroup = this.appValue[0]
            }
            if (this.appValue.length > 1) {
                this.search.application = this.appValue[1]
            }
            if (this.appValue <= 0) {
                this.search.applicationGroup = ''
                this.search.application = ''
            }
            getTraceListApi(this.search).then(res => {
                let list = res.data.list;
                _this.search.total = res.data.total
                _this.traceList = list.map(function (it) {
                    it.elapsedTime = (it.finishTime - it.startTime);
                    return it;
                });
                if (_this.traceList.length > 0) {
                    _this.$nextTick(function () {
                        _this.getTraceTree(_this.traceList[0].traceId, _this.traceList[0].id)
                    });

                }
            });
        },
        getTraceTree(traceId) {
            let _this = this;
            this.loading = true;
            getTraceTreeApi({traceId}).then(res => {
                _this.rootStartTime = res.data.startTime;
                _this.rootElapsedTime = res.data.elapsedTime;
                _this.traceData = [res.data]
                _this.$nextTick(function () {
                    _this.loading = false;
                });

            })
        },
        loadAppData(visible) {
            if (visible) {
                let _this = this
                getAppApi({}).then(res => {
                    _this.appData = res.data;
                })
            }

        }
    },
    mounted() {
        this.getTraceList();
    }
}
</script>

<style scoped>
    .time-line-row-cls {
        padding-left: 0;
        padding-right: 0;
        margin: 0;

    }

    .span-operation-p-cls {
        width: 100%;
        word-break: break-all;
        overflow-wrap: break-word;
        font-size: 15px;
        margin-bottom: 10px;
    }
</style>

<style>
    .time-line-cls-server {
        height: 10px;
        border-radius: 5px;
        background-color: #6e40aa69;
    }

    .time-line-cls-client {
        height: 10px;
        border-radius: 5px;
        background-color: #16bf98a1;
    }

    .trace-div-cls {
        margin-top: 10px;
    }

    .trace-div-cls:hover {

    }
</style>