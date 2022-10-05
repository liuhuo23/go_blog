<template>
  <div>
    <el-row>
      <el-col :span="1" align="middle">
        <el-image src="src/assets/vue.svg" fit="contain"/>
      </el-col>
      <el-col :span="7" align="middle">
        <el-menu
            default-active="/"
            class="el-menu-demo menu"
            mode="horizontal"
            :router="true"
        >
          <el-menu-item index="/">博客</el-menu-item>
          <el-menu-item index="/quession">问答</el-menu-item>
          <el-menu-item index="/square">广场</el-menu-item>
          <el-menu-item index="/bbs">论坛</el-menu-item>
        </el-menu>
      </el-col>
      <el-col :span="9" align="middle" class="w-full">
        <el-select
            v-model="value"
            placeholder="选择"
            multiple
            :reserve-keyword="true"
            :remote-show-suffix="false"
            :remote-method="search"
            class="w-10/12"
        >
          <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              class="rounded-bl-full rounded-tl-full"
              popper-class="rounded-bl-full rounded-tl-full"
              :value="item.value">

          </el-option>
        </el-select>
        <el-button class="w-2/12 rounded-br-full rounded-tr-full bg-red-500 ml-0">搜索</el-button>
      </el-col>
      <el-col :span="7"  class="text-right align-middle pb-2" >

        <el-avatar
            src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"
        />
      </el-col>
    </el-row>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const value =ref("")
const options = ref([{
  label: 'python',
  value: 'python'
}])
const search = () => {

}
</script>

<style scoped>

.menu{
  height: 39px;
  border-bottom: 0;
}
</style>