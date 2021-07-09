<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="10" :md="8">
      <img alt="Vue logo" src="@/assets/logo.png">
      <el-card
          shadow="hover"
          class="login-form"
      >
        <el-input class="login-form-input" v-model="email" placeholder="邮箱"/>

        <el-input class="login-form-input" v-model="pwd" show-password placeholder="密码"/>
        <el-button class="login-form-button" @click="login(email,pwd)">登录</el-button>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      email: '',
      pwd: ''
    }
  },
  methods: {
    checkLogin: function () {
      this.$ajax({url: '/isLogin.php'}).then((response) => {
        if (response.data.code === 20000){
          this.$store.state.isLogin = true
          this.$store.state.right = response.data.right
        }
      })
    },
    login: function (email, pwd) {
      this.$ajax({
        url: '/login.php',
        method: 'post',
        data: {
          email: email,
          pwd: pwd
        }
      }).then((response) => {
        console.log(response.data)
        if(response.data.code === 20000){
          this.checkLogin()
          this.$router.push('/')
        }
      })
    }
  }
}
</script>

<style scoped>
  .login-form-input{
    margin-bottom: 1em;
  }
  .login-form-button{
    width: 100%;
  }
</style>