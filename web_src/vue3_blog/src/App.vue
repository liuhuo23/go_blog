<template>
  <div class="nav-btn">
    <router-link to="/">Home</router-link>
    <router-link to="/about">about</router-link>
  </div>
  <router-view></router-view>
</template>

<script>
import HelloWorld from './components/HelloWorld.vue'

export default {
  name: 'App',
}
</script>
