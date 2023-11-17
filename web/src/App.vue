<script setup>
import AttendanceList from "./components/AttendanceList.vue";
import AttendanceInput from "./components/AttendanceInput.vue";
</script>

<template>
  <nav class="navbar-center bg-neutral text-neutral-content p-2">
    <div class="flex justify-center  ">
        <router-link class="btn btn-ghost text-xl mx-1" to="/">Home</router-link>
        <router-link class="btn btn-ghost text-xl mx-1" to="/attendance">Attendance</router-link>
    </div>
  </nav>
  <div class="min-h-screen w-full flex justify-center items-center">
    <RouterView />
  </div>

</template>

<style scoped>
</style>
