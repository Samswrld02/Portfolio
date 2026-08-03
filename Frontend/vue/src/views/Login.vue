<script setup>

import router from "@/router";
import api from "@/helpers/api";
import { ref } from "vue";

import LoginInput from "@/components/LoginInput.vue";

//formdata refs
const username = ref("")
const password = ref("")

const error = ref("")

const HandleLogin = async () => {
    const Form = new FormData()
    Form.append("username", username.value)
    Form.append("password", password.value)

    try {
        console.log(Form)
        console.log(username.value, password.value)
        const result = await api.post("/login", Form)
        // console.log(result)

        //login to router
        router.push({ name: "dashboard" })

    } catch (error) {
        console.log(error)
        // console.log(er.value)
        // console.error(err.message)
        // error.value = err.message

    }
}




</script>

<template>


    <div id="LoginContainer"
        class="max-w-lg  mx-auto flex flex-col items-center bg-[#0c0c0c] pb-20 border border-gray-300/10 rounded-sm">

        <!-- header section -->

        <div class="my-20 text-2xl flex flex-col items-center  text-white">
            <h1>System Access</h1>
            <h2 class="text-[18px]">Portfolio</h2>
        </div>

        <!-- input -->
        <div class="flex flex-col gap-5  w-2/3 text-white">
            <div class="flex flex-col">
                <label class="text-sm" for="">Username</label>
                <LoginInput v-model="username" placeholder="Username"></LoginInput>
            </div>

            <div class="flex flex-col">
                <label class="text-sm" for="">Password</label>
                <LoginInput type="password" v-model="password" placeholder="Password"></LoginInput>

            </div>

            <!-- button -->
            <div class="text-center mt-5 pb-10 border-b border-b-gray-200/20">
                <h2 @click="HandleLogin"
                    class="transition ease-in-out hover:bg-red-600 py-2 hover:border-gray-200/10 bg-[#ff5357] rounded-sm ">
                    Authenticate
                    →</h2>
            </div>

            <div>
                <p class="text-sm text-center text-gray-200/30">Connection encrypted</p>
            </div>
        </div>

    </div>




</template>
