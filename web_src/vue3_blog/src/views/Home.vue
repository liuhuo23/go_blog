<template>
  <div class="common-layout">
    <el-container>
      <el-header>Header</el-header>
      <el-main>
        {{name}}
        <el-button type="primary" @click="show()">点击</el-button>
      </el-main>
      <el-footer>Footer</el-footer>
    </el-container>
  </div>
</template>

<script setup>
// @ is an alias to /src
import { ElMessage } from 'element-plus'
import { useStore } from 'vuex'
import {ref} from 'vue'
const store = useStore()
const name = ref('hello')
const show = ()=>{
  ElMessage(store.state.username)
}
</script>
