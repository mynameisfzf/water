import {createApp} from 'vue'
import App from './App.vue'
import {createRouter,createWebHistory} from 'vue-router'

import Vue3DraggableResizable from 'vue3-draggable-resizable'
//需引入默认样式
import 'vue3-draggable-resizable/dist/Vue3DraggableResizable.css'


import SelectBack from './components/SelectBack.vue'
import SelectWater from './components/SelectWater.vue'
import Set from './components/Set.vue'

import store from './store/index.js'
const routes=[
    {path:'/',component:SelectBack,name:"home"},
    {path:'/water',component:SelectWater,name:"water"},
    {path:'/set',component:Set,name:"set"},
]

const router =createRouter({
    history:createWebHistory(),
    routes
})

createApp(App).use(Vue3DraggableResizable).use(store).use(router).mount('#app')

 
