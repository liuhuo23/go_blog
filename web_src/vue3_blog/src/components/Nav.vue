<template>
  <div>
    <el-row>
      <el-col :span="1" align="middle">
        <el-image src="src/assets/vue.svg" fit="contain"/>
      </el-col>
      <el-col :span="7" align="middle">
        hello
      </el-col>
      <el-col :span="9" align="middle">
        <el-select
            class="my-select"
            v-model="value"
            placeholder="选择"
            :remote-method="search"
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
.mybutton{
  border-left: 0;
}
</style>