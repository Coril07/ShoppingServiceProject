<template>
    <div id="sign_up">
  <el-form :model="new_user" label-width="auto" style="max-width: 600px;width: 350px;" label-position="top">
    <el-form-item label="ユーザID">
      <el-input v-model="new_user.username" clearable/>
    </el-form-item>
    <el-form-item label="パスワード">
      <el-input v-model="new_user.password" show-password clearable/>
    </el-form-item>
    <el-form-item label="パスワード確認">
      <el-input v-model="new_user.password2" show-password clearable/>
    </el-form-item>
    <el-form-item label="メールアドレス">
      <el-input v-model="new_user.email" clearable/>
    </el-form-item>
    <el-form-item label="アドレス">
      <el-input v-model="df" clearable>
        <template #append>
            <div id="addressbox">
                <el-select v-model="address1" placeholder="都" style="width: 115px">
                    <el-option v-for="n in address1s" :label=n :value="n"/>
                </el-select>
            <el-input v-model="de" clearable>
            </el-input>
            <el-select v-model="address2" placeholder="市" style="width: 115px">
                <el-option v-for="n in address2s" :label=n :value="n"/>
            </el-select>
            </div>
      </template>
    </el-input>
    </el-form-item>
    <el-form-item label="アドレス2">
      <el-input v-model="address3"/>
    </el-form-item>
    <el-form-item label="性別">
      <el-radio-group v-model="new_user.gender">
        <el-radio label="男" />
        <el-radio label="女" />
      </el-radio-group>
    </el-form-item>
    <el-form-item label="誕生日">
        <el-date-picker v-model="new_user.birth" placeholder=""></el-date-picker>
    </el-form-item>
    <el-form-item>
      <el-button type="primary" @click="onSubmit"   class="temp_btn" round>同 意</el-button>
    </el-form-item>
  </el-form>
</div>
</template>

<style scoped>
#sign_up{
    width: 90rem;
    display: flex;
    justify-content: center;
}
.temp_btn{
    margin: 0 auto;
}
#addressbox{
    display: flex;
    justify-content: space-evenly;
}
</style>
<script lang="ts" setup>
import { computed, reactive } from 'vue'
import { ref} from 'vue';

import axios from 'axios';
import { ElMessage, ElMessageBox, type Action } from 'element-plus';
import router from '@/router';



axios.defaults.baseURL = ''; 
axios.defaults.timeout = 5000; 

const address1 = ref('')
const address2 = ref('')
const address3 = ref('')
const df = ref('')
const de = ref('')
const address1s = ref(["都","道","府","县"])
const address2s = ref(["市","町","村"])

// do not use same name with ref
var new_user = reactive({
  username: '',
  password: '',
  password2:'',
  gender: '',
  birth:'2000-03-03 15:26:44.374',
  url: '',
  email: '',
  address: ''
})

function isEmail(email:string) {
  const emailRegex = /^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$/;
  return emailRegex.test(email);
}
var address_res=computed(function () {return df.value+address1.value+de.value+address2.value+address3.value})
const onSubmit = () => {
  new_user.address=address_res.value
  
  axios.post('/api/user/sign', new_user)
  .then((response) => {
    if(response.status==201){
        ElMessageBox.alert('いまログインできる', 'サインイン完了', {
        confirmButtonText: 'OK',
        callback: (action: Action) => {
         ElMessage({
            type: 'info',
            message: `action: ${action}`,
        })
        },
    })
    router.replace('/')
    }
  })
  .catch((error) => {

    console.error(error);
  });

}

</script>