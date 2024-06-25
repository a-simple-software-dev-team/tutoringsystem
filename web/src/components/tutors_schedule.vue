<template>
  <div class="demo-date-picker">

    <div class="block">

      <div class="demonstration">请输入日期：{{ date }}</div>
      <el-date-picker v-model="date" type="date" placeholder="Pick a Date" format="YYYY/MM/DD"
        value-format="YYYY-MM-DD" />
    </div>

  </div>
  <div class="demonstration">请输入时间：{{ time }}</div>

  <el-time-select v-model="time" style="width: 240px" start="08:30" step="00:15" end="18:30"
    placeholder="Select time" />
  <p></p>
  <p></p>
  <div class="demonstration">请输入学生id：{{ student_id }}</div>

  <div>
    <el-input v-model="student_id" style="width: 240px" placeholder="Please input" />
  </div>
  <p></p>
  <el-button type="success">添加日程</el-button>
  <br> <br>
  <el-calendar v-model="value" class="demonstration" />


</template>

<script setup>
import { ref } from 'vue'
const time = ref('')
const date = ref('')
const student_id = ref('')
const value = ref(new Date())





</script>

<style scoped>
.demo-date-picker {
  display: flex;
  width: 100%;
  padding: 0;
  flex-wrap: wrap;
}

.demo-date-picker .block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  flex: 1;
}

.demo-date-picker .block:last-child {
  border-right: none;
}

.demo-date-picker .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}

.canlender {
  width: 40%;
  height: 60%;

  justify-items: center;

}
</style>