<template>
<el-drawer size="50%" :with-header="false"  v-model="show" direction="ltr">
  <el-divider>Basic Info</el-divider>
  <div>
    <el-row class="info-item">
      <el-col :span="5"><h1>应用名称:</h1></el-col>
      <el-col :span="19">
        <span class="label-value">{{spanInfo?.application}}</span>
      </el-col>
    </el-row>
    <el-row class="info-item">
      <el-col :span="5"><h1>应用组:</h1></el-col>
      <el-col :span="19">
        <span class="label-value">{{spanInfo?.applicationGroup}}</span>
      </el-col>
    </el-row>
    <el-row class="info-item">
      <el-col :span="5"><h1>实例:</h1></el-col>
      <el-col :span="19">
        <span class="label-value">{{spanInfo?.appInstance}}</span>
      </el-col>
    </el-row>
    <el-row class="info-item">
      <el-col :span="5"><h1>Method:</h1></el-col>
      <el-col :span="19">
        <span class="label-value">{{spanInfo?.operationName}}</span>
      </el-col>
    </el-row>
  </div>
  <el-divider>Span Tags</el-divider>
  <el-row class="info-item" v-for="(item, index) in tagInfo" :key="index">
    <el-col :span="8"><h1>{{ index }}:</h1></el-col>
    <el-col :span="16">
      <p class="label-value">{{item}}</p>
    </el-col>
  </el-row>
  <el-divider>Logs</el-divider>
  <el-timeline>
    <el-timeline-item v-for="(item, index) in logData" :key="index" :timestamp="`${dayjs(item.timestampMicros).format('YYYY-MM-DD HH:mm:ss')}`" placement="top">
      <el-card :header="index1" v-for="(item1, index1) in item.fields" :key="index1">
        <pre class="label-value">{{item1}}</pre>
      </el-card>
    </el-timeline-item>
  </el-timeline>

</el-drawer>
</template>

<script lang="ts" setup>
import {ref} from 'vue';
import type {TraceSpan} from '@/types';
import {getSpanInfo} from '@/api';
import {dayjs} from 'element-plus';

interface Log {
  timestampMicros: number,
  fields: string
}

const show = ref(false)

const spanInfo = ref<TraceSpan>()

const tagInfo = ref<any>({})
const logData = ref<any[]>([])


const openHandler = async (id: number, startTime: number) => {
  const res = await getSpanInfo({id, startTime})
  const span = res.data
  tagInfo.value = JSON.parse(span.tags)
  logData.value = JSON.parse(span.logDatas)
  spanInfo.value = res.data
  show.value = true
}

defineExpose({
  openHandler,
})

</script>

<style scoped lang="scss">
.label-value {
  color: var(--el-color-info);
  white-space: pre-wrap;
  overflow-wrap: break-word;
}
.info-item {
  margin-bottom: 10px;
}
</style>