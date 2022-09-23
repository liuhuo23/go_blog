import { createApp } from 'vue'
import './style.css'
import App from './App.vue'
import router from "./router/index.js";
import axios from "axios";
createApp(App)
    .provide('$axios', axios)
    .use(router)
    .mount('#app')
