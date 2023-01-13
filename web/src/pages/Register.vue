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
 <div class="w-full h-full flex flex-col justify-center items-center">
  <h2>Nice to meet you!</h2>
  <div class="flex flex-col">
   <div class="text-sm font-bold mb-2 flex flex-col">
    <label for="firstName">First Name</label>
    <input v-model="firstName" id="firstName" name="firstName" type="text" class="p-1" />
   </div>
   <div class="text-sm font-bold mb-2 flex flex-col">
    <label for="lastName">Last Name</label>
    <input v-model="lastName" id="lastName" name="lastName" type="text" class="p-1" />
   </div>
   <div class="text-sm font-bold mb-2 flex flex-col">
    <label for="email">Email</label>
    <input v-model="email" id="email" name="email" type="email" class="p-1" />
   </div>
   <div class="text-sm font-bold mb-2 flex flex-col">
    <label for="password">Password</label>
    <input v-model="password" id="password" name="password" type="password" class="p-1" />
   </div>
   <div class="text-sm font-bold mb-3 flex flex-col">
    <label for="passwordConfirmation">Confirm Password</label>
    <input v-model="passwordConfirmation" id="passwordConfirmation" name="passwordConfirmation" type="password"
     class="p-1" />
   </div>
   <div class="flex justify-center">
    <button class="w-1/2 p-2 border-none bg-indigo-600 text-white rounded-md" @click="onSubmit">Sign Up</button>
   </div>
  </div>
 </div>
</template>