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
    <div class="min-w-30/100 py-1 px-2 flex flex-nowrap">
      <select class="block min-h-full min-w-3/4 rounded-l-full px-2 my-select" >
        <option>hello</option>
      </select>
      <button
          class="min-h-full block min-w-2/10 border-l-2 bg-red-500 border-r-2 rounded-r-full"
      >搜索</button>
    </div>
    <div class="min-w-35/100 py-2">4</div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
const value =ref("")
const placeholder = ref("python")
const options = ref([{
  label: 'python',
  value: 'python'
}])
const search = () => {

}
</script>
<style scoped>
.my-select{
  outline: 0;
}
.my-select option {
  outline: 0;
}
</style>