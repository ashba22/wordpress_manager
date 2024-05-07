<template>
  <div>
    <div class="container mx-auto">
      <Navbar />
      <ProgressBar mode="indeterminate" style="height: 6px; display: none;" />
      <RouterView />
    </div>
  

  </div>
</template>

<script setup>
import { RouterView } from 'vue-router'
import Navbar from './components/Navbar.vue'
</script>




<style>
/* make input like text area etc for dark mode */
</style>
