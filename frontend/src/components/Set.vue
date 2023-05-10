<template>
    <div class="ctx">
        <div class="top">
            <span><label for="top">上边距：</label><input v-model="top" id="top" /> </span>
            <span><label for="left">左边距：</label><input v-model="left" id="left" /> </span>

            <span><label for="width">宽度：</label><input v-model="width" id="width" /> </span>
            <span><label for="height">高度：</label><input v-model="height" id="height" /> </span>

        </div>
      
        <div> 
            <img id="back" />
            <img id="water" />
        </div>
       
      
    </div>
</template>
<script setup>
 
 import {ref} from 'vue'

 const top=ref(0)
 const left=ref(0)
 const width=ref(100)
 const height=ref(100)
 
 
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
    display: flex;
    flex-direction: row; 
    justify-content: space-around;

}
.ctx .top input{
   width:80px;
   height: 20px;
   padding: 3px 5px;
   text-align: right;
   outline: none;
    
   
}
.images{
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;
    
}
.images .item{
    margin: 10px;
    width: calc(100%/3 - 20px);
}
.images img{
    width:100%;
    /* height: 100%; */
    object-fit: cover;
}
</style>
