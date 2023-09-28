<template>
    <div class="ctx">
        <div class="top">
            <span><label for="top">上边距：</label><input v-model="top" id="top" /> </span>
            <span><label for="left">左边距：</label><input v-model="left" id="left" /> </span>

            <span><label for="width">宽度：</label><input v-model="width" id="width" /> </span>
            <span><label for="height">高度：</label><input v-model="height" id="height" /> </span>
            <span class="next" @click="start">确定</span>
        </div>

        <div class="set" :style="backStyle">
            <img class="back" :src="'data:image/png;base64,' + store.state.backfileConfig.src" />
            
            <Vue3DraggableResizable :initW="width" :initH="height" :x="left" :y="top" :draggable="true" :resizable="true"
                @drag-end="dragend" @resize-end="resizeend">
               
                <img class="w" :src="'data:image/png;base64,' + store.state.waterfileConfig.src" />
                
            </Vue3DraggableResizable>
        </div>


    </div>
   <div class="modal" v-if="showModal">
      <div>{{rate}}%</div>
   </div>
</template>
<script setup>

import { computed ,ref} from 'vue'
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'
import {GetSetImage, SetOutDir, Start} from '../../wailsjs/go/main/App.js'
import {EventsOn,LogPrint} from '../../wailsjs/runtime/runtime.js'


const store = useStore()
const router=useRouter()

const rate=ref(0)
const showModal=ref(false)
EventsOn("starting",(ret)=>{
  //图片生成中
  LogPrint(ret)
  rate.value=ret
})
const top = computed({
    get() {
        return store.state.waterfileConfig.top
    },
    set(value) {
        console.log(value)
        store.commit("setWaterfileConfig", { top: value })
    }
})

const left = computed({
    get() {
        return store.state.waterfileConfig.left
    },
    set(value) {
        store.commit("setWaterfileConfig", { left: value })
    }
})

const width = computed({
    get() {
        return store.state.waterfileConfig.width
    },
    set(value) {
        store.commit("setWaterfileConfig", { width: value })
    }
})

const height = computed({
    get() {
        return store.state.waterfileConfig.height
    },
    set(value) {
        store.commit("setWaterfileConfig", { height: value })
    }
})

const waterStyle = computed(() => {

    return `width:${store.state.waterfileConfig.width}px;height:${store.state.waterfileConfig.height}px;top:${store.state.waterfileConfig.top}px;left:${store.state.waterfileConfig.left}px;`
})

const backStyle = computed(() => {

    return `width:${store.state.backfileConfig.width}px;height:${store.state.backfileConfig.height}px;`
})

const dragend=(data)=>{

    store.commit("setWaterfileConfig", { left: data.x, top: data.y })
}

const resizeend=(data)=>{

    store.commit("setWaterfileConfig", { left: data.x, top: data.y,width:data.w,height:data.h })
}

GetSetImage().then((ret) => {

    const { WaterFile, BackFile, BackWidth, BackHeight } = ret

    let width = 0
    let heigh = 0

    if (BackHeight > BackWidth) {
        width = 300
        heigh = BackHeight / BackWidth * width
    } else {
        heigh = 300
        width = BackWidth / BackHeight * heigh
    }

    store.commit("setBackfileConfig", {
        src: BackFile,
        width: width,
        height: heigh,
        realWidth: BackWidth,
        realHeight: BackHeight
    })

    store.commit("setWaterfileConfig", { src: WaterFile })

})

const start=()=>{
  //开始生成

  SetOutDir().then((dir)=>{
    if(!dir){
      return
    }
    showModal.value=true
    let rate=store.state.backfileConfig.realWidth/store.state.backfileConfig.width
    Start(dir,
        store.state.waterfileConfig.top,
        store.state.waterfileConfig.left,
        store.state.waterfileConfig.width,
        store.state.waterfileConfig.height,
        rate
    ).then(()=>{
        router.push({path:"/"})
    }).throw(()=>{
      showModal.value=false
    })
  })
}

</script>
<style scoped>
.ctx {
    width: 90%;
    margin: 20px auto;
}

.ctx .top {
    height: 50px;
    line-height: 50px;
    border-bottom: 1px solid #ffffff;
    display: flex;
    flex-direction: row;
    justify-content: space-around;

}

.ctx .top input {
    width: 80px;
    height: 20px;
    padding: 3px 5px;
    text-align: right;
    outline: none;


}

.images {
    display: flex;
    flex-direction: row;
    flex-wrap: wrap;

}

.images .item {
    margin: 10px;
    width: calc(100%/3 - 20px);
}

.images img {
    width: 100%;
    /* height: 100%; */
    object-fit: cover;
}

.set {
    position: relative;
    /* border:1px solid red; */
    margin: 30px auto;
}

.set .back {
    width: 100%;
    height: 100%;
}

.set .water {
    position: absolute;
}

.set .w{
    width: 100%;
    height: 100%;
}
.modal{
  width: 100%;
  height: 100%;
  position: fixed;
  left: 0;
  top: 0;
  background-color:rgba(0,0,0,0.7);
  display: flex;
  justify-content: center;
  align-items: center;
}
.modal div{
  font-size: 50px;
}
</style>
