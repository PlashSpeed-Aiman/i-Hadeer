<script setup>
import {onMounted, onUnmounted, ref} from 'vue'

 let eventSource = null
  const count = ref(0)
  const increment = () => {
    count.value++
  }
  const matricNumbers = ref([])
  onMounted(() => {

     eventSource = new EventSource('http://localhost:8080/events?stream=students');

    eventSource.onmessage = (event) => {
      let data = JSON.parse(event.data)
      matricNumbers.value.push(data)
    }

    eventSource.onerror = (error) => {
      console.error('SSE error:', error);
    };
  })
function triggerFunc(){
  fetch('http://localhost:8080/ping', {
    method: 'GET',
  })
}
onUnmounted(() => {
  eventSource.close()
})
</script>

<template>
  <div>
    <div>

    </div>
    <table class="table">
      <thead>
      <tr>
        <th>Matric Number</th>
        <th>Name</th>
        <th>Time</th>

      </tr>
      </thead>
      <tbody>
      <tr v-for="matricNumber in matricNumbers" :key="matricNumber">
        <td>{{matricNumber.matricNumber}}</td>
        <td>{{matricNumber.name}}</td>
        <td>{{matricNumber.time}}</td>
      </tr>
      </tbody>
    </table>

  </div>

</template>

<style scoped>

</style>
