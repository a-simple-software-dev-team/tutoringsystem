<!--<template>-->

<!--  <div class="certification">-->
<!--    <br>-->
<!--    <br>-->
<!--    <br>-->
<!--    <br>-->
<!--    <span>家教资格认证: </span> &nbsp;&nbsp;-->
<!--    <el-input v-model="tutor_id" style="width: 240px" placeholder="Please input" />&nbsp;&nbsp;-->
<!--    <br>-->
<!--    <br>-->
<!--    <br>-->
<!--    <el-button plain @click="open">查询</el-button>-->
<!--    <br>-->
<!--    <br>-->
<!--    <br>-->




<!--  </div>-->

<!--</template>-->

<!--<script setup>-->
<!--import { ref } from 'vue'-->
<!--const tutor_id = ref('')-->
<!--import { ElMessage, ElMessageBox } from 'element-plus'-->
<!--import { Action } from 'element-plus'-->

<!--const open = () => {-->
<!--  ElMessageBox.alert('This is a message', '查询结果', {-->

<!--    confirmButtonText: 'OK',-->
<!--    callback: (action) => {-->
<!--      ElMessage({-->
<!--        type: 'info',-->
<!--        message: 'action: ' + action,-->
<!--      })-->
<!--    },-->
<!--  })-->
<!--}-->

<!--</script>-->
<!--<style>-->
<!--.certification {-->
<!--  position: relative;-->
<!--  height: 400px;-->
<!--  border: 2px solid #071a52;-->

<!--}-->
<!--</style>-->
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
    <el-button plain @click="queryCertification">查询</el-button>
    <br>
    <br>
    <br>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { ElMessage, ElMessageBox } from 'element-plus'

const tutor_id = ref('')

const queryCertification = async () => {
  // 检查tutor_id是否为空
  if (!tutor_id.value) {
    ElMessage({
      type: 'warning',
      message: '请填写tutor_id'
    })
    return
  }

  try {
    const response = await axios.get('http://127.0.0.1:8080/api/tutors/certification', {
      params: {
        id: tutor_id.value
      }
    })

    const data = response.data
    if (data.success) {
      ElMessageBox.alert(JSON.stringify(data.certifications, null, 2), '查询结果', {
        confirmButtonText: 'OK',
        callback: (action) => {
          ElMessage({
            type: 'info',
            message: '你已确认'
          })
        }
      })
    } else {
      ElMessage({
        type: 'error',
        message: '查询失败'
      })
    }
  } catch (error) {
    ElMessage({
      type: 'error',
      message: `请求失败: ${error}`
    })
  }
}
</script>

<style>
.certification {
  position: relative;
  height: 400px;
  border: 2px solid #071a52;
}
</style>