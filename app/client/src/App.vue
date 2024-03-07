<template>
    <div class="common-layout">
      <el-container>
        <el-header id="header">
          <el-menu :default-active="current" mode="horizontal":ellipsis="false" router="true">
            <el-menu-item index="/">ホーム</el-menu-item>
            <div class="flex-grow" /> 
            <el-menu-item index="/cart" v-show="!isadmin">         
              <div class="el-icon-" id="shopping_cart"><ShoppingCart></ShoppingCart></div>
              <div>買い物かご</div>    
            </el-menu-item>
            <el-menu-item index="2" v-show="!isadmin">ご注文</el-menu-item>
            <el-menu-item index="/admin" v-show="isadmin">商品管理</el-menu-item>
            <el-sub-menu index="3" v-show="islogin"> 
              <template #title>会員メニュー</template>
              <el-menu-item index="3-1">会員情報</el-menu-item>
              <el-menu-item index="/log_out">ログアウト</el-menu-item>
            </el-sub-menu>
            <el-menu-item index="/login" v-show="!islogin">ログイン</el-menu-item>
            <el-menu-item index="/sign_up" v-show="!islogin">会員登録</el-menu-item>
          </el-menu>
        </el-header>
        <el-main id="cbody"><router-view></router-view></el-main>
      </el-container>
    </div>
  </template>
<script setup lang="ts">
import {computed, inject, onMounted, onUpdated, ref, type reactive } from 'vue';
import Cookies from 'js-cookie';
import axios from 'axios';
import router from '@/router';
import { global_state } from '@/store';

var gs=global_state()

var current=ref("0")
var  islogin=computed(()=>{
  return gs.islogin
})

var isadmin=computed(()=>{
  return gs.isadmin
})


onMounted(()=>{
  axios.get('/api/user/cookielogin')
.then((response) => {
  if(response.status==200){
    router.replace({ path:'/'})
    gs.islogin=true
    gs.username=response.data.username
    gs.isadmin=response.data.isadmin
    console.log(response.data)
  }
})
.catch((error) => {
  console.log("shibai ")
});
  
})
  

</script>
<style>
*{  
    padding: 0;
    margin: 0;
}
#cbody{
  background-color: #e1e5df;
}
</style>
<style scoped>
#header{
    background: color #f3f3f3;
}
.flex-grow {
  flex-grow: 1;
}
#shopping_cart{
  display: flex;
  align-items: center;
}
</style>
