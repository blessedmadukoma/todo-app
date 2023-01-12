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
 <input v-model="email" type="email" placeholder="Email">
 <input v-model="password" type="password" placeholder="Password">
 <button @click="onSubmit">Submit</button>
 <span v-if="authenticated">Authenticated!</span>
 <span v-else>Not Authenticated!</span>
</template>