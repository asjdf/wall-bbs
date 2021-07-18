<template>
  <el-menu
      mode="horizontal"
      background-color="#545c64"
      text-color="#fff"
      active-text-color="#ffd04b"
      v-bind:router = true>
    <el-menu-item index="/">PLBBQ</el-menu-item>
<!--    <el-menu-item v-if="this.$store.state.isLogin" index="/post">发表</el-menu-item>-->
    <el-menu-item v-if="!this.$store.state.hasToken" index="/login">登录</el-menu-item>
    <el-menu-item v-if="!this.$store.state.hasToken" index="/register">注册</el-menu-item>
    <el-submenu v-if="this.$store.state.hasToken" index="4">
      <template slot="title">个人中心</template>
      <el-menu-item @click="logout">登出</el-menu-item>
    </el-submenu>
  </el-menu>
</template>

<script>
export default {
  name: "navbar",
  methods: {
    logout() {
      this.$ajax({
        url: '/user/logout',
        method: 'get',
      })
      localStorage.removeItem('token');
      localStorage.removeItem('hasToken');
      localStorage.removeItem('right');
      localStorage.removeItem('uid');
      this.$store.state.hasToken = false;
      this.$store.state.right = 0;
      this.$store.state.uid = -1;
      this.$router.push('/')
    },
  }
}
</script>

<style scoped>

</style>