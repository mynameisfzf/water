import {createStore} from 'vuex'

export default createStore({
    state(){
        return {
            backfile:[],
            waterfile:[],
            routeIndex:0,

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
    },
 
   
})