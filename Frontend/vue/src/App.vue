<script setup>
import { RouterView } from 'vue-router';
import Layout from './views/Layout.vue';


</script>

<template>

  <component :is="$route.meta.layout">
    <RouterView></RouterView>
  </component>

</template>

<style scoped></style>
