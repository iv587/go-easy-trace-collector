<template>
  <el-table highlight-current-row @row-click="onColumnClick" v-loading="loading" height="100%" :data="data.list">
    <el-table-column>
      <template #header>
        <el-pagination layout=" prev,slot,next,total" :total="data.total" v-model:current-page="queryPage.pageNo" @current-change="pageChangeHandler">
          <template #default>
            <el-space>
              <el-input @change="pageChangeHandler" size="small" style="width: 30px" v-model="queryPage.pageNo" />
              <span>/</span>
              <span>{{Math.ceil(data.total/queryPage.size)}}</span>
            </el-space>
          </template>
        </el-pagination>
      </template>
      <template #default="{row}">
        <div class="span-item">
          <el-link :type="row.error == 1 ? 'error' : 'success'" :underline="false">{{row.operationName}}</el-link>
          <el-space style="margin: 10px 0px">
            <el-tag effect="plain">
              {{row.startTimeText}}
            </el-tag>
            <el-tag effect="plain" :type="row.tagType">
              {{row.elapsedTime}}ms
            </el-tag>
          </el-space>
          <el-space style="margin: 0 0 10px 0">
            <el-tag effect="plain">
              {{row.application}}
            </el-tag>
            <el-tag effect="plain" type="success">
              {{row.applicationGroup}}
            </el-tag>
          </el-space>
          <el-tag>{{row.traceId}}</el-tag>
        </div>
      </template>
    </el-table-column>
  </el-table>
</template>

<script lang="ts" setup>

import {ref} from 'vue';
import type {TraceListRes, TraceSearchForm} from '@/types';
import {listTraceApi } from '@/api'
import {useDateFormat} from '@vueuse/core';
import {dayjs} from 'element-plus';

const queryPage = ref<{
  pageNo: number,
  size: number,
}>({
  pageNo: 1,
  size: 10,
})

const loading = ref<boolean>(false)

const data = ref<TraceListRes>({
  list: [],
  total: 0
})

const emit = defineEmits<{
  (e: 'queryChange'):void,
  (e: 'click', id:number, startTime: number):void
}>()

const pageChangeHandler = () => {
  emit('queryChange')
}

const listTrace = async (searchData: TraceSearchForm, pageNo?: number) => {
  if (pageNo) {
    queryPage.value.pageNo = pageNo
  }
  loading.value = true
  const res = await listTraceApi({
    ...searchData,
    ...queryPage.value,
  })
  const {total, list} = res.data
  data.value.total = total
  data.value.list = list.map(it => {
    it.startTimeText = dayjs(it.startTime).format('YYYY-MM-DD HH:mm:ss.SSS')
    it.elapsedTime = it.finishTime - it.startTime
    it.tagType = 'success'
    if (it.elapsedTime >= 2500) {
      it.tagType = 'warning'
    }
    return it
  })
  loading.value = false
}

const onColumnClick = (row: any) => {
  emit('click', row.id, row.startTime)
}

defineExpose({
  listTrace
})

</script>

<style scoped>
.span-item {
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 5px 0;
}
</style>