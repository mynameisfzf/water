<template>
    <div class="ctx">
        <div>
            <div class="set-dir"><span class="dir">{{ store.state.outdir }}</span>  <span class="btn" @click="select">浏览</span></div>
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
</script>
<style scoped>
.ctx{
    width:90%;
    margin:20px auto;
}
.ctx>div{
    width:100%;
    height: 100vh;
    display: flex;
    flex-direction: row;
    justify-content: center;
    line-height: 100vh;
}

.ctx .set-dir{
    height: 40px;
    line-height: 40px;
    border: 1px solid red;
}

.ctx .dir{
    display: inline-block;
    width: 400px;
    height: 40px;
    line-height: 40px;
    margin-right: 10px;
    border: 1px solid #d1d1d1;
    padding: 0 5px;
    text-align: left;
}

.ctx .btn{
    display: inline-block;
    width:80px;
    height: 40px;
    line-height: 40px;
    cursor: pointer;
    background-color: #409eff;
    border-radius: 5px;
    /* margin-right: 20px; */
}
</style>
