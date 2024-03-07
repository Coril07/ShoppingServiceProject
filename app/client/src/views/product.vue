<template>
    <div id="product">
        <div>
    <el-row :gutter="20">
        <el-col :span="12"><div class="img_part">
            <el-image :src="product.Url" lazy loading="lazy" fit="fill"></el-image>
        </div>
        </el-col>
        <el-col :span="12">
        <div style="display: flex; justify-content: space-around;">
            <div class="info_part">
            <el-form :model="product" label-width="auto" style="max-width: 600px">
                <el-form-item label="商品名">
                <el-input v-model="product.Name" disabled/>
                </el-form-item>
                <el-form-item label="種類">
                <el-input v-model="product.CategoryName" disabled/>
                </el-form-item>
                <el-form-item label="SKU">
                <el-input v-model="product.SKU" disabled/>
                </el-form-item>
                <el-form-item label="値段">
                <el-input v-model="product.Price" disabled/>
                </el-form-item>
                <el-form-item label="在庫数">
                <el-input v-model="product.StockCount" disabled/>
                </el-form-item>
                <el-form-item label="概要">
                <el-input v-model="product.Desc" disabled type="textarea" :rows="8"/>
                </el-form-item>
            </el-form>
        </div>
        <div style="display: flex; flex-direction: column; align-items: center;">
            <el-input-number v-model="num" :min="1" :max="100" />
            <div style="margin-top: 2rem;">
            <el-button type="success" @click="handleClick">買い物に入れて</el-button>
            <el-button type="info" @click="handleCancel">キャンセル</el-button>
            </div>


        </div>
    </div>
    </el-col>
    </el-row>
</div>
  </div>
    </template>
    
<style scoped>
  #product{
    width: 93rem;
    display: flex;
    justify-content: center;
  }
  .img_part{
    height: 27rem;
  }
  .info_part{
    height: 27rem;
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
import { ElMessage, ElMessageBox, type Action } from 'element-plus';

function add_to_cart(){
  var product=gs.seletproduct 
  console.log({sku:product.SKU,count:num.value})
   axios.post('/api/cart/item',{sku:product.SKU,count:num.value})
.then((response) => {
  if(response.status==201){
    ElMessageBox.alert('入れました', 'ご注意', {
        confirmButtonText: 'OK',
    })
    router.replace("/cart")
  }
})
.catch((error) => {
  ElMessageBox.alert(error.response.data.errorMessage, 'ご注意', {
        confirmButtonText: 'OK',
    })
})
}
const handleClick=()=>{
    add_to_cart()
}
const handleCancel=()=>{
    gs.seletproduct={
    Url:"",
    Name:"",
    CategoryName:"",
    Desc:"",
    StockCount:"",
    Price:0,
    SKU:""
  }
    router.replace("/")
}

var gs=global_state()
onMounted(()=>{
   var t=gs.seletproduct
    product.value=gs.seletproduct  
})
  </script>
  