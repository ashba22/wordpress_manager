<template>
  <div>
    <Navbar />
    <RouterView />
  </div>
</template>

<script setup>
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
</script>




<style>
input, textarea, select, button {
  background-color: #1a202c;
  border-color: #2d3748;
  color: #e2e8f0;
}

</style>
