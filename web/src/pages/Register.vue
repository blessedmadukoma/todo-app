<script setup lang="ts">
import { onMounted, ref } from 'vue'
import axios from 'axios'

const API_URL = 'http://localhost:4000/api'

const firstName = ref("")
const lastName = ref("")
const email = ref("")
const password = ref("")
const passwordConfirmation = ref("")
const authenticated = ref(false)

const registrationData = ref({})

const onSubmit = async () => {

 var options = {
  firstName: firstName.value,
  lastName: lastName.value,
  email: email.value,
  password: password.value,
  passwordConfirmation: passwordConfirmation.value
 }

 console.log(options);
 // return

 const response = await axios.post(`${API_URL}/register`, options).then(res => res.data).catch(error => console.error(`error posting to backend: ${error}`)
 )

 console.log(response.data);
 
 registrationData.value = response

 // store user data in local storage
 localStorage.setItem("token", response.data);
 authenticated.value = true

}
</script>

<template>
 <input v-model="firstName" type="text" placeholder="First Name">
 <input v-model="lastName" type="text" placeholder="Last Name">
 <input v-model="email" type="email" placeholder="Email">
 <input v-model="password" type="password" placeholder="Password">
 <input v-model="passwordConfirmation" type="password" placeholder="Confirm Password">
 <button @click="onSubmit">Submit</button>
 <span v-if="authenticated">Authenticated!</span>
 <span v-else>Not Authenticated!</span>
</template>