<template>
  <el-table highlight-current-row border v-loading="loading" row-key="id" :data="list" height="100%" default-expand-all>
    <el-table-column width="300" label="Method">
      <template #default="{row}">
        <el-link @click="showSpanInfoHandler(row.id, row.startTime)" :underline="false" :style="{marginLeft: (row.deepth * 10) + '%' }" :type="row.error == 1 ? 'error' : 'success'">{{row.operationName}}</el-link>
      </template>
    </el-table-column>
    <el-table-column align="center" width="120" label="时间">
      <template #default="{row}">
        {{dayjs(row.startTime).format('HH:mm:ss.SSS')}}
      </template>
    </el-table-column>
    <el-table-column align="center" width="100" label="耗时">
      <template #default="{row}">
        <el-tag :type="row.elapsedTime < 2500 ? 'success' : 'warning'">
          {{row.elapsedTime}}ms
        </el-tag>
      </template>
    </el-table-column>
    <el-table-column align="center" label="时间轴">
      <template #default="{row}">
        <trace-time-line
          :elapsed-time="row.elapsedTime"
          :start-time="row.startTime"
          :root-elapsed-time="rootElapsedTime"
          :root-start-time="rootStartTime"
        />
      </template>
    </el-table-column>
    <el-table-column align="center" label="服务名" width="200">
      <template #default="{row}">
        <el-space>
          <el-tag>{{row.application}}</el-tag>
          <el-tag type="success">{{row.applicationGroup}}</el-tag>
        </el-space>
      </template>
    </el-table-column>
  </el-table>
  <span-info ref="spanInfoRef" />
</template>

<script lang="ts" setup>
import {ref} from 'vue';
import type {TraceSpan} from '@/types';
import {treeTraceApi} from '@/api';
import {dayjs} from 'element-plus';
import TraceTimeLine from '@/components/TraceTimeLine.vue';
import SpanInfo from '@/components/SpanInfo.vue';


const spanInfoRef = ref<InstanceType<typeof SpanInfo> | null>(null)

const list = ref<TraceSpan[]>([])
const loading = ref<boolean>(false)
const rootStartTime = ref(0)
const rootElapsedTime = ref(0)

const treeTraceHandler = async (id: number, startTime: number) => {
  loading.value = true
  const res = await treeTraceApi({id, startTime})
  const rootSpan = res.data
  rootStartTime.value = rootSpan.startTime
  rootElapsedTime.value = rootSpan.elapsedTime
  list.value = Array.of(rootSpan)
  loading.value = false
}

defineExpose({
  treeTraceHandler
})

const showSpanInfoHandler = (id: number, startTime: number) => {
  spanInfoRef?.value?.openHandler(id, startTime)
}

</script>

<style scoped>

</style>