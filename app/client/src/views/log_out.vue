<template>
    <div id="log_out">

  </div>
    </template>
    
<style scoped>
  #log_out{
    width: 93rem;
  }

</style>
    
    
  <script setup lang="ts">
  import { onMounted,ref} from 'vue';
  import axios from 'axios';
import router from '@/router';
import { global_state } from '@/store';

var gs=global_state()
onMounted(()=>{
  axios.get('/api/user/log_out')
.then((response) => {
  if(response.status==200){
    router.replace({ path:'/'})
    gs.islogin=false
    gs.username=""
    gs.isadmin=false
  }
})
.catch((error) => {
  console.log("shibai ")
});
  
})
  
  </script>
  