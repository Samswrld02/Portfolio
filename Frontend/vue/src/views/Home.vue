<script setup>
import Footer from '@/components/Footer.vue';
import Hero from '@/components/Hero.vue';
import Stack from '@/components/Stack.vue';
</script>

<template>
    <div class="flex flex-col">
        <Hero class="bg-linear-to-r from-black via-[#9d1a3b]/12 to-black/10 py-30" />
    </div>
    <Stack />
    <Footer />
</template>