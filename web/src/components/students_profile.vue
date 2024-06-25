<template>
  <div>
    <div>
      <span>请选择科目</span>
    </div>
    <div class="demo-button-style">
      <el-checkbox-group v-model="subjects_needed">
        <el-checkbox-button v-for="subject in subjects_choose" :key="subject" :value="subject" size="small">
          {{ subject }}
        </el-checkbox-button>
      </el-checkbox-group>
    </div>
    <div>
      <div class="lay">

        <div class="block">
          <span>请选择日期（可选择多个）</span>
          <el-date-picker v-model="date" type="dates" placeholder="Pick one or more dates" />
        </div>
        <div class="block">
          <el-time-select v-model="startTime" style="width: 240px" :max-time="endTime" class="mr-4"
            placeholder="Start time" start="08:00" step="00:30" end="22:30" />
          <el-time-select v-model="endTime" style="width: 240px" :min-time="startTime" placeholder="End time"
            start="08:00" step="00:30" end="22:30" />
        </div>


      </div>


    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const subjects_needed = ref([''])

const subjects_choose = ['语文', '数学', '英语', '物理', '化学', '生物', '政治', '历史', '地理']

const date = ref('')
const startTime = ref('')
const endTime = ref('')







</script>

<style scoped>
.demo-button-style {
  margin-top: 24px;
}

.block {
  padding: 30px 0;
  text-align: center;
  border-right: solid 1px var(--el-border-color);
  flex: 1;
}

.block:last-child {
  border-right: none;
}

.block .demonstration {
  display: block;
  color: var(--el-text-color-secondary);
  font-size: 14px;
  margin-bottom: 20px;
}

.lay {
  display: flex;
}
</style>