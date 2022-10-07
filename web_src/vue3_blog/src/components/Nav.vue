<template>
  <div class="container-2xl mx-auto px-4 flex flex-nowrap min-h-full mb-0 justify-self-auto">
    <div class="min-w-1/10">
      <img src="@/assets/vue.svg" />
    </div>
    <div class="min-w-1/4  pt-0">
      <ol class="flex flex-nowrap space-x-6 min-h-full text-lg">
        <li class="block min-w-1/10 hover:bg-gray-300 pt-2 ">
          <router-link to="/" class="block min-h-full" active-class="border-b-2 border-red-600">博客</router-link>
        </li>
        <li class="block min-w-1/10  hover:bg-gray-300 pt-2" >
          <router-link to="/quession" class="block min-h-full" active-class="border-b-2 border-red-600">问答</router-link>
        </li>
        <li class="block min-w-1/10  hover:bg-gray-300 pt-2">
          <router-link to="/square" class="block min-h-full" active-class="border-b-2 border-red-600">广场</router-link>
        </li>
        <li class="block min-w-1/10  hover:bg-gray-300 pt-2">
          <router-link to="/bbs" class="block min-h-full" active-class="border-b-2 border-red-600">论坛</router-link>
        </li>
      </ol>
    </div>
    <div class="min-w-30/100 pt-2 px-2 flex flex-nowrap">
      <n-select
          class="min-h-full min-w-3/4 rounded-full px-2"
          v-model:value="value"
          filterable
          placeholder="搜索"
          :options="options"
          :loading="loading"
          clearable
          remote
          @search="handleSearch"
      >
      </n-select>
      <n-button
          class="min-w-2/10 border-l-2 bg-red-500 border-r-2 rounded-full"
          @click="btn_search"
      >搜索</n-button>
    </div>
    <div class="min-w-35/100 py-2 justify-end flex space-x-3">
      <div>
        <n-menu mode="horizontal" v-model:value="activeKey" :options="menuOptions"/>
      </div>
      <div class="py-0 my-0">
        <n-image
            width="40"
            class="rounded-full my-0 py-0"
            src="https://07akioni.oss-cn-beijing.aliyuncs.com/07akioni.jpeg"
            preview-disabled
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { IosAirplane } from '@vicons/ionicons4'
import {useMessage} from "naive-ui";
const value =ref("")
const placeholder = ref("python")
const options = ref([{
  label: 'python',
  value: 'python'
}])
const activeKey = ref("")
const loading = ref(false)
const message = useMessage()
const menuOptions = ref([
  {
    label: "用户名",
    key: "dance-dance-dance",
    children: [
      {
        type: "group",
        label: "人物",
        key: "people",
        children: [
          {
            label: "叙事者",
            key: "narrator",
          },
          {
            label: "羊男",
            key: "sheep-man",
          }
        ]
      },
      {
        label: "饮品",
        key: "beverage",
        children: [
          {
            label: "威士忌",
            key: "whisky"
          }
        ]
      },
      {
        label: "食物",
        key: "food",
        children: [
          {
            label: "三明治",
            key: "sandwich"
          }
        ]
      },
      {
        label: "过去增多，未来减少",
        key: "the-past-increases-the-future-recedes"
      }
    ]
  }
]);
const handleSearch = () => {
}
const btn_search = ()=>{
  message.info( "你正在搜索的是"+value.value, {keepAliveOnHover: true})
}
</script>
<style scoped>
</style>