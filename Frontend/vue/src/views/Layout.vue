<script setup>
import NavBar from '@/components/NavBar.vue';


</script>

<template>
    <div class="min-h-screen ">
        <nav-bar></nav-bar>
        <main class="max-w-7xl mx-auto">
            <slot />
        </main>
    </div>
</template>