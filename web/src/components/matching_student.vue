<template>
  <br><br>
  <div>
    <el-form-item label="角色" class="mx-2">
      <el-select class="mx-2" v-model="loginForm.role" placeholder="tutor" clearable>
        <el-option label="tutor" value="tutor" />
        <el-option label="student" value="student" />
      </el-select>
    </el-form-item>
    <br><br>
  </div>


  &emsp;&emsp;<el-text class="mx-1" type="primary">匹配结果：</el-text>
</template>

<script>
// @ is an alias to /src

export default {


  data() {
    return {
      loginForm: {

        role: '',
      },

    }
  },

};

</script>
<style>
.mx-1 {
  width: 300px;
  height: 100px;
  display: flex;
  justify-items: left;
  border: 2px solid #071a5229;
}

.mx-2 {
  width: 300px;
  height: 100px;
  display: flex;
  justify-items: left;
}
</style>