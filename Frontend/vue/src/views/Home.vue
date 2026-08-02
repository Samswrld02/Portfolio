<script setup>
import Footer from '@/components/Footer.vue';
import Hero from '@/components/Hero.vue';
import NavBar from '@/components/NavBar.vue';
import Stack from '@/components/Stack.vue';
</script>
<template>
    <div class="flex flex-col">
        <NavBar class=""></NavBar>
        <Hero class="bg-linear-to-r from-[#050505] via-[#3b82f6]/10  to-[#ef3340]/10 py-30">
        </Hero>
    </div>
    <Stack></Stack>
    <Footer></Footer>

</template>