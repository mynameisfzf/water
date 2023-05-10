import {createApp} from 'vue'
import App from './App.vue'
import {createRouter,createWebHistory} from 'vue-router'

import SelectBack from './components/SelectBack.vue'
import SelectWater from './components/SelectWater.vue'
import Set from './components/Set.vue'
import Start from './components/Start.vue'

import store from './store/index.js'
const routes=[
    {path:'/',component:SelectBack,name:"home"},
    {path:'/water',component:SelectWater,name:"water"},
    {path:'/set',component:Set,name:"set"},
    {path:'/start',component:Start,name:"start"}
]

const router =createRouter({
    history:createWebHistory(),
    routes
})

const app=createApp(App)
app.use(store)
app.use(router)

app.mount('#app')
