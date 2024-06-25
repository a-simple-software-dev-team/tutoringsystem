<template>
  <div class="login_container">
    <div class="login_box">

      <div class="avatar_box">
        <img src="../assets/logo.png" alt="">
      </div>
      <el-form ref="loginformref" :model="loginForm" label-width="auto" style="max-width: 600px" :rules="loginFormRules"
        class="fp">
        <el-form-item label="用户名" prop="username">
          <el-input prefix-icon="User" v-model="loginForm.username" />
        </el-form-item>



        <el-form-item label="密码" prop="password">
          <el-input prefix-icon="Lock" v-model="loginForm.password" type="password" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="loginForm.email" />
        </el-form-item>

        <el-form-item label="角色">
          <el-select v-model="loginForm.role" placeholder="tutor" clearable>
            <el-option label="tutor" value="tutor" />
            <el-option label="student" value="student" />
          </el-select>
        </el-form-item>

        <el-form-item class="bp">
          <div class="bp">
            <el-button type="primary" v-on:click="login_a">注册</el-button>
          </div>
        </el-form-item>


      </el-form>
    </div>
  </div>
</template>

<script>
import { ref } from 'vue';
import axios from 'axios'
export default {

  data() {
    return {
      loginForm: {
        username: 'admin',
        password: '123456',
        email: '',
        role: '',
      },
      loginFormRules: {
        username: [
          { required: true, message: '用户名不能为空', trigger: 'blur' },
          { min: 3, max: 10, message: '用户名长度为3-10个字符', trigger: 'blur' }
        ],
        password: [
          { required: true, message: '密码不能为空', trigger: 'blur' },
          { min: 6, max: 15, message: '密码长度为6-15个字符', trigger: 'blur' }
        ],

      },
    }
  },
  methods: {
    login_a() {
      this.$refs.loginformref.validate(a => {
        console.log(a);
        if (!a) return;
        const res1 = this.loginForm;
        console.log(res1);
        //    console.log(res.data);
        axios.post('http://127.0.0.1:8080/api/register', res1)
          .then(response => {
            console.log('Response:', response);

            // 根据需求处理响应数据
            this.$router.push('/Login');
          })
          .catch(error => {
            console.error('Error:', error);

            // 处理错误
          });
        // this.$router.push('/Login');



      });
    }

  }
};


</script>

<style lang="less" scoped>
.login_container {
  background-color: #2b4b6b;
  height: 100%;

}

.login_box {
  width: 450px;
  height: 300px;
  background-color: #fff;
  border-radius: 3px;
  position: absolute;
  left: 50%;
  top: 50%;
  transform: translate(-50%, -50%);

  .avatar_box {
    height: 130px;
    width: 130px;
    border: 1px solid #eee;
    border-radius: 50%;
    padding: 10px;
    box-shadow: 0 0 10px #eee;
    position: absolute;
    left: 50%;
    transform: translate(-50%, -50%);

    img {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-color: #eee;
    }

  }
}

.bp {
  display: flex;
  width: 100%;
  align-items: center;
  justify-content: center !important;
}

.fp {
  position: absolute;
  bottom: 0;
  width: 100%;
  padding: 0 20px;
  box-sizing: border-box;
}
</style>
