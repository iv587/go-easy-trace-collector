<template>
  <div class="es-fill" style="background-color: var(--el-bg-color-page) ;display: flex; justify-content: center;">
    <section style="width: 25vw; margin-top: 25vh">
      <h1 style="margin-bottom: 10px;">欢迎登陆</h1>
      <el-form ref="loginRef">
        <el-form-item >
          <el-input v-model="loginForm.userName"  placeholder="请输入用户名" :prefix-icon="User"></el-input>
        </el-form-item>
        <el-form-item>
          <el-input
            v-model="loginForm.password"
            type="password"
            placeholder="请输入密码"
            show-password
            :prefix-icon="Lock"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <el-button @click="loginHandler"  style="width: 100%"  type="primary">登陆</el-button>
        </el-form-item>
      </el-form>
    </section>
  </div>

</template>

<script lang="ts" setup>
import {User, Lock} from '@element-plus/icons-vue';
import {ref} from 'vue';
import {ElForm} from 'element-plus';
import {loginApi} from '@/api';
import auth from '@/utils/auth';
import {useRouter} from 'vue-router';

const router = useRouter()

const loginForm = ref<{
  userName: string,
  password: string
}>({
  userName: '',
  password: ''
})

const loginRef = ref<InstanceType<typeof ElForm> | null>(null)

const loginHandler = async () => {
  await loginRef.value?.validate()
  const res = await loginApi(loginForm.value)
  const {token }= res.data
  auth.setToken(token)
  await router.push({path: '/'})
}


</script>

<style scoped>

</style>