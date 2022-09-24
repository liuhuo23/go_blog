import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from "./router/index.js";
import axios from "axios";
import ElementPlus from 'element-plus'
import zhCn from 'element-plus/es/locale/lang/zh-cn'
import 'element-plus/dist/index.css'
import store from './store'
import 'bootstrap/dist/css/bootstrap.css'
createApp(App)
    .provide('$axios', axios)
    .use(router)
    .use(store)
    .use(ElementPlus, {zIndex: 3000, local:zhCn})
    .mount('#app')
