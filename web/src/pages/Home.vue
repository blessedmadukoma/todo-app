<script lang="ts" setup>
import axios from 'axios';
import { ref } from 'vue';

const API_URL = 'http://localhost:4000/api'

interface ITodoItem {
 title: string
}

const newTodoTitle = ref("")
const todoItems = ref<ITodoItem[]>([])


const onSubmit = async () => {
 const token = localStorage.getItem('token')

 const response = await axios.post(`${API_URL}/todos/`, {
  title: newTodoTitle.value
 },
  {
   headers: {
    Authorization: `Bearer ${token}`,
   }
  }
 )
  .then((res) => res.data)
  .catch((error) => console.error(`Error adding todo to database: ${error}`)
  )

 console.log(`Response: ${response}`);
 todoItems.value.push(
  {
   title: response.Title,
  }
 )
}

</script>
<template>

 <div>
  <input v-model="newTodoTitle" type="text" id="new-todo-title" name="new-todo-item" />
  <button @click="onSubmit">Create Todo</button>
 </div>

 <div>
  <p v-for="item in todoItems">{{ item.title }}</p>
 </div>

</template>