<template>
<!--博客列表item-->
  <div class="py-2">
<!--    标题-->
    <div class="col-span-3  text-start px-2 text-lg py-2">{{title}}</div>
<!--    图片 描述-->
    <div class="flex flex-nowrap py-2">
      <div class="row-span-2 min-w-2/10">
        <img src="@/assets/vue.svg" />
      </div>
      <div >
        <div class="col-span-2 text-start px-2 text-base">{{description}}</div>
        <div class="col-span-2 text-start px-2">3</div>
      </div>
    </div>
  </div>
</template>

<script setup>
defineProps({
  title: String,
  likes: Number,
  description: String,
})
</script>

<style scoped>

</style>