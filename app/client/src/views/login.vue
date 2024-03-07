<template>
  <div id="login">
<el-form :model="new_user" label-width="auto" style="max-width: 600px;width: 350px;" label-position="top">
  <el-form-item label="ユーザID">
    <el-input v-model="new_user.username" clearable/>
  </el-form-item>
  <el-form-item label="パスワード">
    <el-input v-model="new_user.password" show-password clearable/>
  </el-form-item>
  <el-form-item>
    <div id="sign_notice">
      <div @click="route_to_sign" style="cursor: pointer; text-decoration: underline;color: red;margin: 0 auto;">会員登録</div>
    </div>
  </el-form-item>
  <el-form-item>
    <div style="display: flex;justify-content: center;width:100%">
      <el-button type="primary" @click="onSubmit"   class="temp_btn" round>ログイン</el-button>
    </div>
  </el-form-item>
</el-form>
</div>
</template>

<style scoped>
#login{
  width: 90rem;
  display: flex;
  justify-content: center;
}
#sign_notice{
  display: flex;
  width: 100%;
}
</style>
<script lang="ts" setup>
import { computed, onMounted, reactive } from 'vue'
import { ref} from 'vue';

import axios from 'axios';
import { ElMessage, ElMessageBox, type Action } from 'element-plus';
import router from '@/router';
import { global_state } from '@/store';

axios.defaults.baseURL = ''; 
axios.defaults.timeout = 5000; 

// do not use same name with ref
var new_user = reactive({
username: '',
password: '',
})

function route_to_sign(){
  router.replace('/sign_up')
}


const onSubmit = () => {

axios.post('/api/user/login', new_user)
.then((response) => {
  if(response.status==200){
    router.replace({ path:'/'})
    var gs=global_state()
    gs.islogin=true
    gs.isadmin=response.data.isadmin
    ElMessageBox.alert('ログインしました。', 'ご注意', {
        confirmButtonText: 'OK',
    })
  }
})
.catch((error) => {
  ElMessageBox.alert('ユーザID・パスワードが一致しません', 'ご注意', {
        confirmButtonText: 'OK',
    })
  console.log("shibai ")
});

}

</script>