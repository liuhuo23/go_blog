<template>
<div>
  hello
</div>
</template>

<script setup>
import {inject} from "vue";
const $axios = inject('$axios')
$axios.get("https://www.baidu.com").then((resp) => {
  console.log(resp.data)
}).catch((err) => {
  console.log(err)
})
</script>

<style scoped>

</style>