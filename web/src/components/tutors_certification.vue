<template>

  <div class="certification">
    <br>
    <br>
    <br>
    <br>
    <span>家教资格认证: </span> &nbsp;&nbsp;
    <el-input v-model="tutor_id" style="width: 240px" placeholder="Please input" />&nbsp;&nbsp;
    <br>
    <br>
    <br>
    <el-button plain @click="open">查询</el-button>
    <br>
    <br>
    <br>




  </div>

</template>

<script setup>
import { ref } from 'vue'
const tutor_id = ref('')
import { ElMessage, ElMessageBox } from 'element-plus'
import { Action } from 'element-plus'

const open = () => {
  ElMessageBox.alert('This is a message', '查询结果', {

    confirmButtonText: 'OK',
    callback: (action) => {
      ElMessage({
        type: 'info',
        message: 'action: ' + action,
      })
    },
  })
}

</script>
<style>
.certification {
  position: relative;
  height: 400px;
  border: 2px solid #071a52;

}
</style>