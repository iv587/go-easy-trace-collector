<template>
  <el-container class="es-fill">
    <el-header style="padding: 0">
      <el-menu background-color="#545c64"
               text-color="#fff"
               active-text-color="#ffd04b"
               router mode="horizontal">
        <el-menu-item index="trace">链路查看</el-menu-item>
        <el-menu-item index="connection">进程查看</el-menu-item>
      </el-menu>
    </el-header>
    <el-main style="padding: 0" class="es-fill">
      <router-view />
    </el-main>
  </el-container>
</template>

<script lang="ts" setup>

</script>

<style scoped>

</style>