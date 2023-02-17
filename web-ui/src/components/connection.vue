<template>
    <div>
        <card :bordered="false" dis-hover>
            <div slot="title">
                <Button icon="md-refresh" shape="circle" @click="getData"></Button>
            </div>
            <i-table border :columns="columns" :data="tableData"></i-table>
        </card>
    </div>
</template>

<script>

    import { getConnectList } from  '../lib/api'

    export default {
        name: 'connection',
        data() {
          return {
              columns: [
                  {
                      title: 'Remote Addr',
                      key: 'addr'
                  },
                  {
                      title: 'Connected Time ',
                      key: 'createTime'
                  },
                  {
                      title: 'Up Time',
                      key: 'aliveTime'
                  },
                  {
                      title: 'Net Input',
                      key: 'netInput'
                  },
                  {
                      title: 'App Start Time',
                      key: 'appStartTime'
                  },
                  {
                      title: 'App Name',
                      key: 'appName'
                  },
                  {
                      title: 'App Group',
                      key: 'appGroup'
                  }
              ],
              tableData: [],
          }
        },
        methods: {
            getData() {
                getConnectList({}).then(res => {
                    this.tableData = res.data.list
                })
            }
        },
        mounted() {
            this.getData()
        }
    }
</script>

<style scoped>

</style>