<script setup lang="ts">
import { ref } from 'vue'
import axios from 'axios'

const API_URL = 'http://localhost:4000/api'

const email = ref("")
const password = ref("")

const onSubmit = async () => {

 var options = {
  email: email.value,
  password: password.value
 }

 const response = await axios.post(`${API_URL}/login`, options).then(res => res.data).catch(error => console.error(`error posting to backend: ${error}`)
 )

 console.log(response);
 
}
</script>

<template>
 <input v-model="email" type="email">
 <input v-model="password" type="password">
 <button @click="onSubmit">Submit</button>
 <span>Not Authenticated!</span>
</template>