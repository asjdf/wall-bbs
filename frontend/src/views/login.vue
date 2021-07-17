<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="10" :md="8">

      <el-card
          shadow="hover"
          class="login-form"
      >
<!--        <img alt="logo" src="@/assets/logo.jpg">-->
        <el-input class="login-form-input" v-model="tel" placeholder="手机号"/>

        <el-input class="login-form-input" v-model="pwd" show-password placeholder="密码"/>
        <el-button class="login-form-button" @click="login()">登录</el-button>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      tel: '',
      pwd: '',
    }
  },
  methods: {
    checkLogin() {
      this.$ajax({url: '/isLogin.php'}).then((response) => {
        if (response.data.code === 20000){
          this.$store.state.isLogin = true
          this.$store.state.right = response.data.right
        }
      })
    },
    goBack() {
      window.history.length > 1 ? this.$router.go(-1) : this.$router.push('/')
    },
    login() {
      this.$ajax({
        url: '/user/login',
        method: 'post',
        data: {
          tel: this.tel,
          password: this.pwd,
        }
      }).then((response) => {
        console.log(response.data)
        if(response.data.code === 20000){
          this.$message({
            message: '登录成功',
            type: 'success'
          });
          // this.checkLogin()
          this.goBack()
          //全局存储token
          window.localStorage["token"] = response.data.data.token;
          this.$store.state.hasToken = true;
          this.$store.state.right = response.data.data.right;
          this.$store.state.uid = response.data.data.uid;
        }else{
          this.$message({
            message: response.data.msg,
            type: 'error'
          });
        }
        // eslint-disable-next-line no-unused-vars
      }).catch(err => {
        this.$message({
          message: "登录错误",
          type: 'error'
        });
      })
    },
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