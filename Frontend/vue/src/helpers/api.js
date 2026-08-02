import router from "@/router"
import axios from "axios"


const api = axios.create({ baseURL: "/api" })

api.interceptors.response.use(
    function (response) {
        if (response.data?.token) {
            localStorage.setItem("access_token", response.data.token)
        }
        return response
    },
    (error) => {
        if (error.response && error.response.status == 401) {
            localStorage.removeItem("access_token")
            router.push({name: "login"})
        }
        return Promise.reject(error)
    }
)

api.interceptors.request.use((config) => {
    const token = localStorage.getItem("access_token");
    if (token) {
        config.headers.set("Authorization", `Bearer ${token}`)
    }
    return config
})

export default api