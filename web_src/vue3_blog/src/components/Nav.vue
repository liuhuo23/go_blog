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
      <el-col :span="9" align="middle" class="search">
        <el-select
            v-model="value"
            placeholder="选择"
            multiple
            :reserve-keyword="true"
            :remote-show-suffix="false"
            :remote-method="search"
            class="my-select"
        >
          <el-option
              v-for="item in options"
              :key="item.value"
              :label="item.label"
              :value="item.value">

          </el-option>
        </el-select>
        <el-button class="mybutton">搜索</el-button>
      </el-col>
      <el-col :span="7" align="middle">
        hello
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
.el-col {
  border-radius: 4px;
}
.el-select{
  width: 80%;
  height: 80%;
  border-right: 0;
}
.my-select{
  margin-right: 0;
  border-right: 0 solid #888888;
  border-top-left-radius: 16px;
  border-bottom-left-radius: 16px;
  vertical-align: middle;
  display: inline-block;
  height: 35px;
}
.mybutton {
  margin-left: 0;
  margin-top: -3px;
  border-left: 0 solid #888888;
  background-color: #e86363;
  width: 88px;
  border-top-right-radius: 16px;
  border-bottom-right-radius: 16px;
  display: inline-block;
  padding-top: 0;
  padding-bottom: 0;
  border-top: 0;
  border-bottom: 0;
  height: 32px;
}
.search{
  vertical-align: middle;
  display: inline-block;
}
.menu{
  height: 39px;
  border-bottom: 0;
}
</style>