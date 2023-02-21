<template>
  <el-container class="es-fill">
    <el-header style="display: flex ;align-items: center ;padding: 0; border-bottom: 1px solid var(--el-border-color)">
      <el-form inline label-width="120px" label-position="right">
        <el-form-item style="margin: 0" label="日期:">
          <el-date-picker @change="queryHandler(1)" value-format="YYYY-MM-DD" clearable v-model="searchForm.day"/>
        </el-form-item>
        <el-form-item style="margin: 0" label="状态:">
          <el-select @change="queryHandler(1)" v-model="searchForm.error">
            <el-option :value="-1" label="全部">全部</el-option>
            <el-option :value="0" label="成功">成功</el-option>
            <el-option :value="1" label="失败">失败</el-option>
          </el-select>
        </el-form-item>
        <el-form-item style="margin: 0" label="端点搜索">
          <el-input @change="queryHandler(1)" v-model="searchForm.operationKey" placeholder="请输入端点"/>
        </el-form-item>
        <el-form-item style="margin: 0" label="选择应用">
          <el-cascader @change="queryHandler(1)" v-model="appValues" :options="appList" @focus="showAppHandler">
          </el-cascader>
        </el-form-item>
        <el-form-item style="margin: 0; padding: 0 20px">
          <el-button @click="queryHandler(1)" type="primary">刷新</el-button>
        </el-form-item>
      </el-form>
    </el-header>
    <el-main  class="es-fill" style="padding: 0;">
      <el-container class="es-fill">
        <el-aside style="border-right: var(--el-border)" width="350px">
          <trace-list @click="showTraceTreeHandler" ref="traceListRef" @query-change="queryHandler()"/>
        </el-aside>
        <el-main>
          <trace-tree ref="traceTreeRef"/>
        </el-main>
      </el-container>
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>
import {onMounted, ref} from 'vue';
import TraceList from '@/components/TraceList.vue';
import TraceTree from '@/components/TraceTree.vue';
import {useDateFormat, useNow} from '@vueuse/core';
import type {AppRes, TraceSearchForm} from '@/types';
import {listAppApi} from '@/api'

const currentDay = useDateFormat(useNow(), 'YYYY-MM-DD')

const appList = ref<AppRes[]>([])

const searchForm = ref<TraceSearchForm>({
  day: currentDay.value,
  error: -1,
});

const appValues = ref<string[]>()

const traceListRef = ref<InstanceType<typeof TraceList> | null>(null)
const traceTreeRef = ref<InstanceType<typeof TraceTree> | null>(null)

const queryHandler = (pageNo?: number) => {
  if(appValues.value && appValues.value?.length > 1) {
    searchForm.value.application = appValues.value[1]
    searchForm.value.applicationGroup = appValues.value[0]
  } else {
    searchForm.value.application = undefined
    searchForm.value.applicationGroup = undefined
  }
  traceListRef?.value?.listTrace(searchForm.value, pageNo)
}

const showAppHandler = async () => {
  const res = await listAppApi({})
  appList.value = res.data
}

const showTraceTreeHandler = async (id: number, startTime: number) => {
  traceTreeRef?.value?.treeTraceHandler(id, startTime)
}

onMounted(() => {
  queryHandler()
})
</script>

<style scoped>

</style>