<template>
   <div class="warp">
<!--    <div class="menu">-->
<!--      <router-link v-if="store.state.iswork===false" class="button" :class="{active:btnIndex===0}" @click="changeBtn(0)" to="/">背景图</router-link>-->
<!--      <router-link v-if="store.state.iswork===false" class="button" :class="{active:btnIndex===1}" @click="changeBtn(1)" to="/water">水印图</router-link>-->
<!--      <router-link v-if="store.state.iswork===false" class="button" :class="{active:btnIndex===2}" @click="changeBtn(2)" to="/set">模板</router-link>-->
<!--      <router-link class="button" :class="{active:btnIndex===3}" @click="changeBtn(3)" to="/start">生成</router-link>-->
<!--   </div>-->
     <div class="content">
        <router-view></router-view>
     </div>
   </div>
 
</template>


<script setup>
 // import {computed} from 'vue'
 // import {useStore} from 'vuex'
 //
 // const store=useStore()
 //
 // const btnIndex=computed(()=>store.state.routeIndex)
 //
 // const changeBtn=(index)=>{
 //     store.commit("setRouteIndex",{value:index})
 // }

</script>
<style scoped>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
.warp{
  height: 100vh;
  //display: grid;
  //grid-template-columns: 100px auto;
  //grid-template-rows: 100%;
  background:linear-gradient(to bottom right,#617087,#242831);
  //overflow-y:auto;
}
.menu{
  width:100px;
  height: 100%;
  border-right: 1px solid #475560;
  display:flex;
  flex-direction: column;
  justify-content: center;
  background-color: #495568;
  border-radius: 0 10px 10px;
}

.menu .button{
  width: 100px;
  height: 80px;
  margin: 0 auto;
  background-color: #3f4754;
  border-top: 1px solid #475560;
  border-bottom: 1px solid #475560;
  text-align: center;
  line-height: 80px;
  cursor: pointer;
  letter-spacing: 1px;
  color:#ffffff;
  text-decoration: none;

}
.menu .button:hover{
  background-color: #343b4c;
}
.menu .active{
  background-color: #343b4c;
}
.content{
  height: 100vh;
  overflow-y: auto;
 
}
</style>
