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
  <button @click="changeRoute('register')">Register</button>
  <button @click="changeRoute('login')">Login</button>
  <br>
  <Login v-if="currentRoute === 'login' " />
  <Register v-else />
</template>

<style scoped>

</style>
