<template>
  <n-message-provider>
    <el-container class="h-screen">
      <el-header class="header" height="60px">
        <Nav />
      </el-header>
      <el-main>
        <router-view />
      </el-main>
    </el-container>
  </n-message-provider>
</template>
<script setup>
import Nav from './components/Nav.vue'
// This starter template is using Vue 3 <script setup> SFCs
// Check out https://vuejs.org/api/sfc-script-setup.html#script-setup
import HelloWorld from './components/HelloWorld.vue'
</script>

<style scoped>
.logo {
  height: 6em;
  padding: 1.5em;
  will-change: filter;
}
.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}
.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
.header{
  border-bottom: 1px solid #888888;
  padding-bottom: 0;
}
</style>
