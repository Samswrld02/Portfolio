<script setup>
import StackCard from './StackCard.vue';

const TestSkill = {
    url: "/favicon/icon.png",
    title: "Go",
    description: "CONCURRENCY & SPEED"
}

</script>

<template>

    <section id="skills">

        <div
            class="px-10 flex flex-col gap-8 border-b border-t border-b-[#978085]/20 border-t-[#978085]/20 py-10 bg-black">
            <h2 id="skillsHeader" class="text-3xl text-white">
                Core technology stack
                <p class="text-[#978085] text-sm">The foundation of my engineering workflow.</p>
            </h2>
            <!-- skillcard grid -->
            <div class="grid md:grid-cols-5 grid-cols-2 gap-2 ">
                <StackCard class="text-white" :skill="TestSkill"></StackCard>
                <StackCard class="text-white" :skill="TestSkill"></StackCard>
                <StackCard class="text-white" :skill="TestSkill"></StackCard>
                <StackCard class="text-white" :skill="TestSkill"></StackCard>
                <StackCard class="text-white" :skill="TestSkill"></StackCard>

            </div>
        </div>
    </section>
</template>