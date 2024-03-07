<template>
    <div id="cart">

        cart
  </div>
    </template>
    
<style scoped>
  #cart{
    width: 93rem;
    display: flex;
    justify-content: center;
  }

</style>
    
  <script setup lang="ts">
  var product=ref({
    Url:"",
    Name:"",
    CategoryName:"",
    Desc:"",
    StockCount:"",
    Price:0,
    SKU:""
  })
  var num=ref(1)
  import { onMounted,ref} from 'vue';
  import axios from 'axios';
import router from '@/router';
import { global_state } from '@/store';



var gs=global_state()
onMounted(()=>{
 
})

  </script>
  