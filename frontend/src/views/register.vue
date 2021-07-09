<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="10" :md="8">
      <img alt="Vue logo" src="@/assets/logo.png">
      <el-card
          shadow="hover"
          class="login-form"
      >
        <el-input class="register-form-input" v-model="email" placeholder="手机号"/>
        <el-input class="register-form-input" v-model="username" placeholder="用户名"/>
        <el-input class="register-form-input" v-model="pwd" show-password placeholder="密码"/>
        <el-input class="register-form-input" v-model="pwdcheck" show-password placeholder="确认密码"/>
        <el-divider></el-divider>
        <el-input class="register-form-input" v-model="realname" show-password placeholder="姓名"/>
        <el-input class="register-form-input" v-model="cardid" show-password placeholder="校卡卡号（x位）">
          <el-button slot="append">验证</el-button>
        </el-input>
        <el-button class="register-form-button" @click="register(email,username,pwd)">注册</el-button>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      email: '',
      username: '',
      pwd: '',
      realname: '',
      cardid: '',
      pwdcheck: ''
    }
  },
  methods: {
    register: function (email, username, pwd) {
      this.$ajax({
        url: '/register.php',
        method: 'post',
        data: {
          email: email,
          username: username,
          pwd: pwd,

        }
      }).then((response) => {
        if(response.data.code === 20000){
          this.$router.push('/login')
        }else{
          this.$message({
            center: true,
            showClose: true,
            message: response.data.msg,
            type: 'error'
          });
        }
      })
    }
  }
}
</script>

<style scoped>
  .el-divider {
    margin: 0 0 1em;
  }
  .register-form-input {
    margin-bottom: 1em;
  }
  .register-form-button {
    width: 100%;
  }
</style>