<script setup>
import NavBar from '@/components/NavBar.vue';


</script>

<template>
    <div class="min-h-screen">
        <nav-bar></nav-bar>
        <main class="w-full">
            <slot />
        </main>
    </div>
</template>