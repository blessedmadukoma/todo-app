<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'

const API_URL = 'http://localhost:4000/api'

const email = ref("")
const password = ref("")
const authenticated = ref(false)

onMounted(() => {
 const token = localStorage.getItem('token')

 authenticated.value = token != null;
})

const onSubmit = async () => {

 var options = {
  email: email.value,
  password: password.value
 }

 const response = await axios.post(`${API_URL}/login`, options).then(res => res.data).catch(error => console.error(`error posting to backend: ${error}`)
 )

 // console.log(response);

 // store user data in local storage
 localStorage.setItem("token", response.data);
 authenticated.value = true

}
</script>

<template>
 <div class="w-full h-full flex flex-col justify-center items-center">
  <h2>Welcome Back!</h2>
  <div class="flex flex-col">
   <div class="text-sm font-bold mb-2 flex flex-col">
    <label for="email">Email</label>
    <input v-model="email" id="email" name="email" type="email" class="p-1" />
   </div>
   <div class="text-sm font-bold mb-3 flex flex-col">
    <label for="password">Password</label>
    <input v-model="password" id="password" name="password" type="password" class="p-1" />
   </div>
   <div class="flex justify-end mb-3">
    <button class="font-bold border-none text-indigo-600 rounded-md">Forgot Password?</button>
   </div>
   <div class="flex justify-center">
    <button class="w-1/2 p-2 border-none bg-indigo-600 text-white rounded-md" @click="onSubmit">Submit</button>
   </div>
  </div>
 </div>
</template>