<template>
  <el-row>
    <ChatDotRound style="width: 42px"/>
    <el-text :style="`font-size: 24px; font-family: '等线';font-weight: bold;`">家教消息通知</el-text>
  </el-row>

  <el-row style="margin-top: 10px; height: calc(90% - 100px);">
    <div class="infinite-list-wrapper" style="overflow: auto">
      <ul
        v-infinite-scroll="load"
        class="list"
        :infinite-scroll-disabled="true"
      >
        <li v-for="i in msgNum" :key="i" class="list-item">
          <div style="margin-top: 10px; margin-bottom: 10px; margin-left: 15px">
            <span style="font-size: 20px; font-family: '等线';font-weight: bold;">
              {{msgList[i-1].sender_id}}
            </span>
            <div style="height: 5px" />
            <span style="font-size: 18px; font-family: '黑体';font-weight: lighter;white-space: pre-line;">
              {{msgList[i-1].content}}
            </span>
          </div>
        </li>
      </ul>
    </div>
  </el-row>

  <el-row>
    <div class="con">
      <el-input v-model="currentUser" style="width: 240px" placeholder="Current User ID" />
      <el-input v-model="targetUser" style="width: 240px" placeholder="Target User ID" />
      <el-input v-model="textarea" style="width: 240px" :rows="2" type="textarea" placeholder="Message Content" />
      <el-button type="primary" @click="sendMessage">Send</el-button>
    </div>
  </el-row>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'

export default {
  setup() {
    const currentUser = ref('') // 当前用户ID
    const targetUser = ref('') // 目标用户ID
    const textarea = ref('')
    const msgList = ref([])
    const msgNum = ref(0)

    onMounted(async () => {
      if (currentUser.value) {
        await loadMessages()
      }
    })

    const loadMessages = async () => {
      try {
        const response = await axios.get(`http://localhost:8080/api/messages/${currentUser.value}`)
        msgList.value = response.data
        msgNum.value = response.data.length
      } catch (error) {
        console.error('Failed to load messages:', error)
      }
    }

    const sendMessage = async () => {
      if (!currentUser.value || !targetUser.value || !textarea.value) {
        console.warn('请确保所有字段都已填写完毕')
        return
      }

      try {
        // axios.defaults.headers.common['Cookie'] = document.cookie;
        // axios = axios.create({
        //   headers: {
        //     Cookie: document.cookie
        //   }
        // })
        await axios.post('http://localhost:8080/api/send', {
          sender_id: currentUser.value, // 替换为实际的发送者ID
          receiver_id: targetUser.value, // 替换为实际的接收者ID
          content: textarea.value
        })

        textarea.value = ''
        await loadMessages()
      } catch (error) {
        console.error('Failed to send message:', error)
      }
    }

    return {
      currentUser,
      targetUser,
      textarea,
      msgList,
      msgNum,
      loadMessages,
      sendMessage
    }
  }
}
</script>

<style scoped>

.infinite-list-wrapper {
  width: 100%;
  border: 2px solid #8f8f8f;
  text-align: left;
}
.infinite-list-wrapper .list {
  padding: 0;
  margin: 0;
  list-style: none;
}

.infinite-list-wrapper .list-item {
  display: flex;
  align-items: center;
  justify-content: left;
  border-bottom: 1px solid #8f8f8f;
  border-image: -webkit-linear-gradient(right, transparent 60%, #8f8f8f 60%, #8f8f8f 100%) 1;
  color: #696969;
}
.infinite-list-wrapper .list-item + .list-item {
  margin-top: 10px;
}

.con {
  display: flex;
  align-items: center;
  justify-content: space-between;
  margin-top: 10px;
}

.con .el-input {
  margin-right: 10px;
}

</style>
<!--<template>-->
<!--  <el-row >-->
<!--    <ChatDotRound style="width: 42px"/>-->
<!--    <el-text :style="`font-size: 24px; font-family: '等线';font-weight: bold;`">家教消息通知</el-text>-->
<!--  </el-row>-->
<!--  <el-row style="margin-top: 10px; height: calc(90% - 100px);">-->
<!--    <div class="infinite-list-wrapper" style="overflow: auto">-->
<!--      <ul-->
<!--        v-infinite-scroll="load"-->
<!--        class="list"-->
<!--        :infinite-scroll-disabled="true"-->
<!--      >-->
<!--        <li v-for="i in msgNum" :key="i" class="list-item">-->
<!--          <div style="margin-top: 10px; margin-bottom: 10px; margin-left: 15px">-->
<!--            <span style="font-size: 20px; font-family: '等线';font-weight: bold;">-->
<!--              {{msgList[i-1][0]}}-->
<!--            </span>-->
<!--            <div style="height: 5px" />-->
<!--            <span style="font-size: 18px; font-family: '黑体';font-weight: lighter;white-space: pre-line;">-->
<!--              {{msgList[i-1][1]}}-->
<!--            </span>-->
<!--          </div>-->
<!--        </li>-->
<!--      </ul>-->
<!--    </div>-->
<!--  </el-row>-->
<!--  <ElButton v-on:click="test"></ElButton>-->
<!--  -->
<!--  <el-row >-->
<!--    <div class="con">-->

<!--      <el-icon class="place1">-->
<!--        <User />-->
<!--      </el-icon>-->
<!--      <el-input v-model="textarea" style="width: 240px" :rows="2" type="textarea" placeholder="Please input"-->
<!--        class="place3" />-->
<!--      <el-button type="primary" class="place2" :icon="Position" color="#f9f7f7" />-->

<!--    </div>-->
<!--  </el-row>-->
<!--</template>-->
<!--<script>-->
<!--  import { ref } from 'vue'-->
<!--  import { Position } from '@element-plus/icons-vue'-->
<!--import { ElButton } from 'element-plus';-->

<!--  export default {-->
<!--      setup(){-->
<!--          const textarea = ref('');-->
<!--          const msgList = ref([["ABC", "HELLO\nThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a startThis is a start"]]);-->
<!--          const msgNum = ref(1);-->
<!--          const load = () => {-->
<!--            console.log("add");-->
<!--          }-->
<!--          const test = () => {-->
<!--            if(msgNum.value < 20){-->
<!--              msgNum.value += 1;-->
<!--              msgList.value = [-->
<!--                ...msgList.value,-->
<!--                ["ABC", "HELLO"]-->
<!--              ]-->
<!--            }-->
<!--          }-->
<!--          return {-->
<!--            textarea,-->
<!--            msgList,-->
<!--            msgNum,-->
<!--            load,-->
<!--            test-->
<!--          }-->
<!--      }-->
<!--  }-->
<!--</script>-->

<!--<style scoped>-->

<!--.infinite-list-wrapper {-->
<!--  width: 100%;-->
<!--  border: 2px solid #8f8f8f;-->
<!--  text-align: left;-->
<!--}-->
<!--.infinite-list-wrapper .list {-->
<!--  padding: 0;-->
<!--  margin: 0;-->
<!--  list-style: none;-->
<!--}-->

<!--.infinite-list-wrapper .list-item {-->
<!--  display: flex;-->
<!--  align-items: center;-->
<!--  justify-content: left;-->
<!--  border-bottom: 1px solid #8f8f8f;-->
<!--  border-image: -webkit-linear-gradient(right, transparent 60%, #8f8f8f 60%, #8f8f8f 100%) 1;-->
<!--  color: #696969;-->
<!--}-->
<!--.infinite-list-wrapper .list-item + .list-item {-->
<!--  margin-top: 10px;-->
<!--}-->

<!--</style>-->