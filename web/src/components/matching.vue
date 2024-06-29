<template>
  <br /><br />

  <div class="container">
    <el-form-item label="角色" class="form-item">
      <el-select class="select" v-model="loginForm.role" placeholder="选择角色" clearable>
        <el-option label="tutor" value="tutor" />
        <el-option label="student" value="student" />
      </el-select>
    </el-form-item>

    <el-button type="primary" @click="matchRole">匹配</el-button>

    <br /><br />

    &emsp;&emsp;<el-text class="result" type="primary">匹配结果：{{ matchResult }}</el-text>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      loginForm: {
        role: '',
      },
      matchResult: '',
    };
  },
  methods: {
    async matchRole() {
      if (!this.loginForm.role) {
        this.$message.warning('请选择一个角色');
        return;
      }

      try {
        const response = await axios.get('http://localhost:8080/api/matches', {
          user_id: this.loginForm.role,
        });

        if (response.data.success) {
          this.matchResult = '匹配成功';
        } else {
          this.matchResult = '匹配失败';
        }
      } catch (error) {
        this.matchResult = `请求错误: ${error.message}`;
      }
    },
  },
};
</script>

<style>
.container {
  width: 320px;
  margin: auto;
  padding: 20px;
  border: 2px solid #071a52;
  border-radius: 10px;
  background-color: #f5f5f5;
  text-align: center;
}

.form-item {
  margin-bottom: 20px;
}

.select {
  width: 240px;
}

.result {
  font-size: 18px;
  margin-top: 20px;
}
</style>