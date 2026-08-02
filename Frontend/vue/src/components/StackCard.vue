<script setup>
defineProps({
    skill: Object
})

</script>

<template>
    <div id="StackCard" class="bg-[#0c1220] py-10 border border-[#7d5d66]/50 rounded-sm hover:border-red-500">
        <div class="px-5">
            <img class="w-10 my-2" :src="skill.url ? skill.url : ''" id="icon" alt="icon">

            <div>
                <div id="title" class="text-xl">{{ skill.title }}</div>
                <p id="description" class="text-[11px]">{{ skill.description }}</p>
            </div>
        </div>
    </div>

</template>