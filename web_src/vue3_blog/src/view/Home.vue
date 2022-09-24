<template>
<div>
  hello
  {{count}}
  {{Mydata.message}}
  <el-button @click="getinfo()">点击</el-button>
</div>
</template>

<script setup>
import {computed, inject, ref} from "vue";
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