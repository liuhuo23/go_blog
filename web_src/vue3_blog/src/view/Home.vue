<template xmlns="">
<div class="container mx-auto flex">
  <n-list
      hoverable clickable
    class="min-w-3/4"
  >
    <n-list-item v-for="i in 10">
      <Blog :title="'博客'+i" description="这是一个很好的博客" />
    </n-list-item>
  </n-list>
</div>
</template>

<script setup>
import {computed, inject, ref} from "vue";
import '@/components/Blog.vue'
import { useStore } from 'vuex'
import {apiGetuserInfo} from "../apis/user.js";
const $axios = inject('$axios')
$axios.get("/api/api/v1/hello-world?name=12").then((data) => {
  console.log(data)
}).catch((err)=>{
  console.log(err.response)
})
const store  = useStore()
const count= computed(()=>store.state.count)
let Mydata = ref({})
const getinfo = function (){
  apiGetuserInfo("hello", "1232423").then((data)=>{
    Mydata.value = data
    console.log(data)
  }).catch((err)=>{
    console.log("get", err)
  })
}
</script>

<style scoped>

</style>