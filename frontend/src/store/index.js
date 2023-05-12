import {createStore} from 'vuex'

export default createStore({
    state(){
        return {
            backfile:[],
            waterfile:[],
            routeIndex:0,
            outdir:"",
            iswork:false,
            backfileConfig:{
                src:"",
                width:300,
                height:300,
                realWidth:0,
                realHeight:0
            },
            waterfileConfig:{
                src:"",
                width:100,
                height:100,
                top:0,
                left:0
            }
        }
    },
    mutations:{
        setRouteIndex(state,payload){
            state.routeIndex=payload.value
        },
        setBackfile(state,payload){
            state.backfile=payload
        },
        setWaterfile(state,payload){
            state.waterfile=payload
        },
        setOutdir(state,payload){
            state.outdir=payload.value
        },
        setIswork(state,payload){
            state.iswork=payload.value
        },
        setBackfileConfig(state,payload){

            if(payload.src){
                state.backfileConfig.src=payload.src
            }

            if(payload.width>0){
                state.backfileConfig.width=payload.width
            }

            if(payload.height>0){
                state.backfileConfig.height=payload.height
            }

            if(payload.realWidth>0){
                state.backfileConfig.realWidth=payload.realWidth
            }

            if(payload.realHeight>0){
                state.backfileConfig.realHeight=payload.realHeight
            }
        },
        setWaterfileConfig(state,payload){
            if(payload.src){
                state.waterfileConfig.src=payload.src
            }

            if(payload.width>0){
                state.waterfileConfig.width=payload.width
            }

            if(payload.height>0){
                state.waterfileConfig.height=payload.height
            }
            if(payload.top>=0){
                state.waterfileConfig.top=payload.top
            }

            if(payload.left>=0){
                state.waterfileConfig.left=payload.left
            }
        }
    },
 
   
})