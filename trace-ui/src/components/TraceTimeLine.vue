<template>
  <el-tooltip>
    <template #content>
      <p>时间: {{dayjs(props.startTime).format('HH:mm:ss.SSS')}} ~ {{dayjs(props.startTime + props.elapsedTime).format('HH:mm:ss.SSS')}} </p>
      <p>耗时: {{props.elapsedTime}}2ms</p>
    </template>
    <div class="trace-time-line" :style="{width: width, marginLeft: marginLeft}"></div>
  </el-tooltip>
</template>

<script lang="ts" setup>
import {computed} from 'vue';
import {dayjs} from 'element-plus';

const props = defineProps<{
  rootStartTime: number,
  rootElapsedTime: number,
  startTime: number,
  elapsedTime: number
}>();

const width = computed(() => {
  return (props.elapsedTime / props.rootElapsedTime * 100) + '%';
});

const marginLeft = computed(() => {
  return ((props.startTime - props.rootStartTime) / props.rootElapsedTime * 100) + '%';
});
</script>

<style scoped>
.trace-time-line {
  height: 8px;
  background-color: rgba(110, 64, 170, .4117647058823529);
  border-radius: var(--el-border-radius-base);
}
</style>