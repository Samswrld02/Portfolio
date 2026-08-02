<script setup></script>
<template>
    <main id="hero" class="">
        <div class="flex justify-center">
            <div id="hero-content" class="my-15 flex flex-col gap-10 max-w-5xl  text-[#7c8fb8]">

                <!-- header  -->
                <div class="text-[#7c8fb8] flex gap-2">
                    <span class="text-red-600 block">───</span>
                    <div>
                        SYSTEM ARCHITECT
                    </div>
                </div>

                <!-- main text -->
                <div class="text-5xl">Engineered for <span>Performance</span></div>
                <div> Specializing in high-concurrency systems and distributed architectures. Expertly navigating
                    the ecosystems of Golang and Laravel to build resilient, scalable backend solutions. </div>

                <!-- buttons -->
                <div id="buttons" class="text-white flex gap-3 max-w-1/2">
                    <div
                        class="hover:transition duration-200  ease-in-out  hover:scale-105  tr px-4 grow text-center  py-3 rounded-sm bg-red-700/60">
                        View
                        Projects → </div>
                    <div
                        class="hover:transition duration-200  ease-in-out  hover:scale-105  px-4 grow py-3 text-center rounded-sm border border-[#6d7da1] hover:bg-[#6d7da1]">
                        Get in
                        touch</div>
                </div>
            </div>
        </div>
    </main>
</template>