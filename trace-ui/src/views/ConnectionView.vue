<template>
  <el-container>
    <el-main>
      <el-button type="primary" @click="listHandler">刷新</el-button>
      <el-divider/>
      <el-table border v-loading="loading" :data="list">
        <el-table-column prop="addr" label="Remote Addr"></el-table-column>
        <el-table-column prop="createTime" label="Connected Time"></el-table-column>
        <el-table-column prop="aliveTime" label="Up Time"></el-table-column>
        <el-table-column prop="netInput" label="Net Input"></el-table-column>
        <el-table-column prop="appStartTime" label="App Start Time"></el-table-column>
        <el-table-column prop="appName" label="App Name"></el-table-column>
        <el-table-column prop="appGroup" label="App Group"></el-table-column>
      </el-table>
    </el-main>
  </el-container>

</template>

<script lang="ts" setup>
import {listAppConn} from '@/api';
import {onMounted, ref} from 'vue';
import type {AppConnRes} from '@/types';

const list = ref<AppConnRes[]>([])

const loading  = ref(false)

const listHandler = async () => {
  loading.value = true
  const res = await listAppConn({})
  list.value = res.data.list
  loading.value = false
}

onMounted(() => {
  listHandler()
})
</script>

<style scoped>

</style>