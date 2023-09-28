<template>
   <div class="warp">

     <div class="content">
        <router-view></router-view>
     </div>
   </div>
<!--    <div class="footer">-->
<!--            ysfzf@hotmial.com-->
<!--    </div>-->
</template>


<script setup>


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
.footer{
    position: fixed;
    bottom:0;
    text-align: center;
    width: 100%;
}
</style>
