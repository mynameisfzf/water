<template>
    <div class="ctx">
        <div class="z-flex">
             <div class="dir bot">{{ store.state.outdir }}</div>  <div class="btn" @click="select">浏览</div>
        </div>
        <div class="z-flex">
            <div class="dir"></div>  <div class="btn" @click="starter">开始</div>
        </div>
    </div>
</template>
<script setup>
 
 import {useStore} from 'vuex'
 import {SetOutDir} from '../../wailsjs/go/main/App.js'

const store =useStore()
const select=()=>{
    SetOutDir().then((dir)=>{
        store.commit("setOutdir",{value:dir})
    })

}
const starter=()=>{

}
</script>
<style scoped>
.ctx{
    width:90%;
    margin:20px auto;
    display: flex;
    flex-flow: column nowrap;
}
.ctx .z-flex{
    width:100%;
    height: 50vh;
    display: flex;
    flex-direction: row;
    justify-content: center;
    align-items: center;
    line-height: 50vh;
   
}

.ctx .bot{
    border-bottom: 1px solid #ffffff;
}

.ctx .dir{
    width: 400px;
    height: 40px;
    line-height: 40px;
    margin-right: 10px;  
    padding: 0 5px;
    text-align: left;
}

.ctx .border{
    border: 1px solid #d1d1d1;
}

.ctx .btn{
    width:80px;
    height: 40px;
    line-height: 40px;
    cursor: pointer;
    background-color: #409eff;
    border-radius: 5px;
    
}


</style>
