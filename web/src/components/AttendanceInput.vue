<script setup>

import {ref} from "vue";

const matricNumber = ref('')
const name = ref('')

function submitInfo(){
  fetch('http://localhost:8080/attendance', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify({
      matricNumber: matricNumber.value,
      name: name.value
    })
  })
}
</script>

<template>
<div class="my-5  card w-96 bg-base-100 shadow-xl">
  <form class="card-body" @submit.prevent="onSubmit">
    <h1 class="">Matric Number</h1>
    <input v-model="matricNumber" class="input bg-gray-50">
    <h1>Name</h1>
    <input v-model="name" class="input bg-gray-50">
    <button class="btn btn-primary" @click="submitInfo">Submit</button>

  </form>
</div>
</template>

<style scoped>

</style>