<script setup>
import { useRouter, useRoute } from 'vue-router'


const route = useRoute()

const router = useRouter()
const handleLogin = () => {
    router.push({ path: '/login' })
}

</script>

<template>
    <nav
        class="bg-[#060b17] sticky top-0 z-50 w-full flex flex-col md:flex-row items-center justify-between px-6 py-3 border-b border-b-[#7d5d66]">
        <div class="text-[#b9c1db] text-2xl font-bold">Backend_Portfolio</div>
        <div class="flex gap-8 text-[18px] text-[#7d5d66]">
            <RouterLink to="/" :class="[route.path == '/' ? 'border-b text-[#7d5d66] border-b-[#7d5d66]' : '']">
                Home
            </RouterLink>
            <RouterLink to="/projects"
                :class="[route.path == '/projects' ? 'border-b text-[#7d5d66] border-b-[#7d5d66]' : '']">
                Projects</RouterLink>
            <RouterLink :class="[route.path == '/skills' ? 'border-b text-[#7d5d66] border-b-test' : '']" to="/skills">
                Skills
            </RouterLink>
            <RouterLink :class="[route.path == '/contact' ? 'border-b  border-b-test' : '']" to="/contact">Contact
            </RouterLink>
        </div>
        <div class="flex gap-4">
            <div class="px-4 py-2 bg-[#80061c] text-white rounded-sm">Resume</div>

            <!-- avatar -->
            <div @click="handleLogin" class="avatar">
                <div class="w-6 rounded-full  hover:transition ease-in-out hover:scale-120">
                    <svg width="100%" height="100%" viewBox="0 0 24 24" fill="none" xmlns="http://www.w3.org/2000/svg">
                        <path
                            d="M20 21C20 19.6044 20 18.9067 19.8278 18.3389C19.44 17.0605 18.4395 16.06 17.1611 15.6722C16.5933 15.5 15.8956 15.5 14.5 15.5H9.5C8.10444 15.5 7.40665 15.5 6.83886 15.6722C5.56045 16.06 4.56004 17.0605 4.17224 18.3389C4 18.9067 4 19.6044 4 21M16.5 7.5C16.5 9.98528 14.4853 12 12 12C9.51472 12 7.5 9.98528 7.5 7.5C7.5 5.01472 9.51472 3 12 3C14.4853 3 16.5 5.01472 16.5 7.5Z"
                            stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round" />
                    </svg>
                </div>
            </div>
        </div>
    </nav>

</template>