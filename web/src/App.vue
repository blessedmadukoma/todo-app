<script setup lang="ts">
import { ref } from 'vue';
import Login from './pages/Login.vue';
import Register from './pages/Register.vue';

let currentRoute = ref("login");

const changeRoute = (path: string) => {
  console.log(path);

  currentRoute.value = path;
}

</script>

<template>

  <router-view></router-view>

  <!-- <button @click="changeRoute('register')">Register</button>
  <button @click="changeRoute('login')">Login</button>
  <br>
  <Login v-if="currentRoute === 'login'" />
  <Register v-else /> -->
</template>

<style scoped>
html,
body {
  width: 100%;
  height: 100%;
}

#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  width: 100%;
  height: 100%;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

* {
  font-family: "Source Sans Pro"
}
</style>
