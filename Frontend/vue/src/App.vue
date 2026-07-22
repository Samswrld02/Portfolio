<script setup>
import { RouterView } from 'vue-router';
import Layout from './views/Layout.vue';


</script>

<template>

  <h1>You did it!</h1>
  <p class="text-orange-400 text-7xl">
    Visit <a href="https://vuejs.org/" target="_blank" rel="noopener">vuejs.org</a> to read the
    documentations
  </p>
  
  <component :is="$route.meta.layout">
    <RouterView></RouterView>
  </component>

</template>

<style scoped></style>
