import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import './assets/css/global.css'
import { useRouter } from 'vue-router'
import { ElButton } from 'element-plus'
import { ElForm, ElFormItem } from 'element-plus'
import { ElInput } from 'element-plus'
import { h } from 'vue'
import { ElMessage } from 'element-plus'
import * as ElementPlusIconsVue from '@element-plus/icons-vue'
//import axios from 'axios'

const app = createApp(App);
//axios.defaults.baseURL='http://127.0.0.1:8888/register'
//app.config.globalProperties.$http = axios

app.use(router).use(ElMessage)
    .use(ElementPlus).use(ElButton).use(ElForm).use(ElFormItem).use(ElInput);
for (const [key, component] of Object.entries(ElementPlusIconsVue)) {
    app.component(key, component)
}

app.mount('#app');
