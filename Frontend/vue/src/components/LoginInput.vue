<script setup>
defineProps({
    placeholder: String,
})

const modelV = defineModel()

</script>

<template>
    <input v-model="modelV" class="bg-[#18181a] px-5 py-2 mt-2 rounded-sm   border border-gray-300/20"
        :placeholder="placeholder">
</template>