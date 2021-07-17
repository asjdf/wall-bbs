<template>
  <el-row type="flex" justify="center">
    <el-col :xs="24" :sm="10" :md="8">
<!--      <img alt="logo" src="@/assets/logo.jpg">-->
      <el-card
          shadow="hover"
          class="login-form"
      >
        <el-input class="register-form-input" v-model="tel" placeholder="手机号"/>
        <el-input class="register-form-input" v-model="username" placeholder="用户名"/>
        <el-input class="register-form-input" v-model="pwd" show-password placeholder="密码"/>
        <el-input class="register-form-input" v-model="pwd_check" show-password placeholder="确认密码"/>
        <el-divider></el-divider>
        <el-input class="register-form-input" v-model="real_name" @input="info_verified=false" placeholder="姓名"/>
        <el-input class="register-form-input" v-model="card_id" @input="info_verified=false" placeholder="校卡卡号（10位）">
          <el-button slot="append" @click="verifyCardInfo">验证</el-button>
        </el-input>
        <el-button class="register-form-button" :disabled="!info_verified" @click="register">注册</el-button>
      </el-card>
    </el-col>
  </el-row>
</template>

<script>
export default {
  name: "register",
  data() {
    return {
      tel: '',
      username: '',
      pwd: '',
      real_name: '',
      card_id: '',
      pwd_check: '',
      info_verified: false,
    }
  },
  methods: {
    register() {
      if(this.tel.length !== 11){
        this.$message({
          message: "手机号错误",
          type: 'error'
        });
        return;
      }
      if(this.pwd !== this.pwd_check){
        this.$message({
          message: "密码与确认密码不同",
          type: 'error'
        });
        return;
      }
      this.$ajax({
        url: '/user/register',
        method: 'post',
        data: {
          tel: this.tel,
          name: this.username,
          password: this.pwd,
          real_name: this.real_name,
          card_code: this.card_id
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
    },
    verifyCardInfo() {
      this.$ajax({
        url: '/user/verify-card-info',
        method: 'post',
        data: {
          real_name: this.real_name,
          card_id: this.card_id
        }
      }).then((response) => {
        if(response.data.code === 20000){
          if(response.data.data.result === true){
            this.$message({
              center: true,
              showClose: true,
              message: "校验成功",
              type: 'success'
            });
            this.info_verified = true;
          }else{
            this.$message({
              center: true,
              showClose: true,
              message: "信息不正确",
              type: 'warning'
            });
          }
        }else{
          this.$message({
            center: true,
            showClose: true,
            message: "校验失败",
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