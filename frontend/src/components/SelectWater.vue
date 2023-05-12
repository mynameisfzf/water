<template>
    <div class="ctx">
        <div class="top">
            <span class="btn" @click="select" v-if="store.state.iswork===false">选择水印图</span>
        </div>
      
        <div class="images">
            <div v-for="(file,index) in store.state.waterfile" :key="index" class="item">
                <img :src="'data:image/png;base64,'+file"  />
            </div>
            
        </div>
       
      
    </div>
</template>
<script setup>
 
 import {useStore} from 'vuex'
 import {SelectWaterFiles,GetWaterFiles} from '../../wailsjs/go/main/App.js'

const store =useStore()
const select=()=>{
    SelectWaterFiles().then(()=>{
         GetWaterFiles().then((files)=>{
             store.commit("setWaterfile",files)
             
         })
    })

}
</script>
<style scoped>
.ctx{
    width:90%;
    margin:20px auto;
}
.ctx .top{
    height: 50px;
    line-height: 50px;
    border-bottom: 1px solid #ffffff;
    text-align: right;
}
.ctx .top span{
    display: inline-block;
    width:100px;
    height: 46px;
    line-height: 46px;
    cursor: pointer;
    margin-right: 20px;
    background-color: #409eff;
    border-radius: 5px;
    text-align: center;
}
.images{
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    
}
.images .item{
    margin: 10px;
    width: calc(100%/6 - 10px);
}
.images img{
    width:100%;
    /* height: 100%; */
    object-fit: cover;
}
</style>
