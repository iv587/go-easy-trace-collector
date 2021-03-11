<template>
  <div id="app">
    <Main></Main>
  </div>
</template>

<script>
  import Main from "./components/main";
  export default {
  name: 'App',
  components: {
    Main

  }
}
</script>

<style>
  * {
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  }
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  color: #2c3e50;
}
</style>
