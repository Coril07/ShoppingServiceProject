<template>
    <div id="cart">
      <div id="cartbox">
        <div id="cartheader">
          <div style="margin-left: 1.2rem;font-weight: bolder;">買い物かご（全部{{ }}）</div>
        </div>
        <div id="cartbody">
          <div id="title">
            <div class="isdeleted"><el-checkbox label="全選択" size="large" @change="HandleAllSeleted" v-model="selecttag"/></div>
            <div class="info">商品情報</div>
            <div class="price">単価</div>
            <div class="amount">数量</div>
            <div class="fee">費用</div>
            <div class="operate">操作</div>
          </div>
          <div id="content" v-for="item in items">
            <div class="isdeleted"><el-checkbox  size="large" style="margin-left: -3rem;" @change="HandleSeleted(item)" v-model="item.select" /></div>
            <div class="info" style="padding: 1rem;">
              <div class="productshow">
                <el-image :src="item.Product.Url" lazy loading="eager" fit="fill" style="max-width: 10rem;"></el-image>
                <div class="Name">{{item.Product.Name}}</div>
              </div>
          </div>
            <div class="price">{{item.Product.Price}}</div>
            <div class="amount"><el-input-number v-model="item.num" :min="1" :max="50" size="small" style="width: 5rem;" /></div>
            <div class="fee">{{item.Product.Price*item.num}}</div>
            <div class="operate">
              <div></div>
          </div>
          </div>
        </div>
        <div id="cartfooter">
          <div id="deleteall" @click="DeleteSelectedItems">削除</div>
          <div style="margin-left: 35rem; display: flex;align-items: center;width:40rem ;">
          <div id="selected">選んだ商品数：{{selectcount}}</div>
          <div id="totalfee" style="margin-left: 10rem;width: 10rem;">合計額：{{ totalfee }}円</div>
          <el-button type="primary" @click="HandlePay" style="margin-left: 6rem;">決算</el-button>
        </div>
        </div>
      </div>
  </div>
    </template>
    
<style scoped>
#title div{
  height: 3rem;
  display: flex;
  justify-content: center;
  align-items: center;
  font-weight: bolder;
}
#totalfee{
  
}
#selected{
  color: red;
}
#deleteall{
  cursor: pointer;
  margin-left: 2rem;
}
#deleteall:hover{
  color: red;
  text-decoration: underline;
}
#content{
  cursor: pointer;
}
#content div{
  display: flex;
  justify-content: center;
  align-items: center;
}
.productshow{
  display: flex;
  align-items: center;
}
.Name{
  margin-left: 1rem;
  max-width: 10rem;
  max-height: 6rem;
overflow: hidden;
text-overflow: ellipsis;
}
.Name:hover{
  color:red;
  text-decoration: underline;
}
.isdeleted{
  width: 10rem;
}
.info{
  width: 30rem;
}
.price{
  width: 10rem;
}
.amount{
  width: 10rem;
}
.fee{
  width: 10rem;
}
.operate{
  width: 10rem;
}
#title{
  width: 100%;
  display: flex;
}
#content{
  width: 100%;
  display: flex;
  border-bottom: 0.1rem solid #e6e3e4;
  background-color: #f5f5f5;
}
#cartbox{
  background-color:white;
  width: 80rem;
  border-radius: 1.5rem;
}
#cartheader{
  height: 4rem;
  display: flex;
  align-items: center;
  border-bottom: 0.1rem solid #e6e3e4;
}
#cartbody{
  min-height: 8rem;
  background-color: white;  
}
#cartfooter{
  height: 4rem;
  display: flex;
  align-items: center;
}
  #cart{
    width: 93rem;
    display: flex;
    justify-content: center;
  }

</style>
    
  <script setup lang="ts">
  import { computed, onMounted,ref} from 'vue';
  import axios from 'axios';
import router from '@/router';
import { global_state } from '@/store';
import { ElMessage, ElMessageBox, type Action } from 'element-plus';
import qs from "qs"

var items=ref(Array())
var selected_items=ref(new Map())
var selecttag=ref(false)
var tag=true

var gs=global_state()
onMounted(()=>{
  updatecart()
})

const totalfee=computed(()=>{
  var fee=0
  if(items.value.length!=0&&tag==true){
    for (let v of selected_items.value.values()) {
    var t=items.value[v.Ind].num*items.value[v.Ind].Product.Price
    fee+=t
    }
  }
  return fee
})

const selectcount=computed(()=>{
  var count=0
  if(items.value.length!=0&&tag==true){
    for (let v of selected_items.value.values()) {
    count+=1
    }
  }
  return count
})
const HandleSeleted = (item) => {
  if(item.select==true){
    selected_items.value.set(item.Product.Name,{SKU:item.Product.SKU,Ind:item.ind,PID:item.Product.ID}) 
  }else{
    selected_items.value.delete(item.Product.Name)
  }
}
const HandlePay = () => {
  var IDs=Array()
  for (let item of selected_items.value.values()) {
    IDs.push(item.PID)
    }
  ElMessageBox.confirm(
    '商品を購入しますか',
    'Warning',
    {
      confirmButtonText: 'はい',
      cancelButtonText: 'また考えさせてください',
      type: 'warning',
    }
  )
    .then(() => {
      axios.post('/api/order/pay',{PIDs:IDs})
.then((response) => {
  if(response.status==201){
    updatecart()
  }
})
.catch((error) => {
  ElMessageBox.alert('購入が失敗しました', 'ご注意', {
        confirmButtonText: 'OK',
    })
})
      ElMessage({
        type: 'success',
        message: 'completed',
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'canceled',
      })
    })
}

const updatecart= () => {
  axios.get('/api/cart/info')
.then((response) => {
  if(response.status==200){
    tag=false
    items.value=response.data
    for (let index = 0; index < items.value.length; index++) {
      items.value[index].select=ref(false)  
      items.value[index].num=ref(items.value[index].Count)  
      items.value[index].ind=index 
    }
    selected_items=ref(new Map())
    tag=true
  }
})
.catch((error) => {
  ElMessageBox.alert(error.response.data.errorMessage, 'ご注意', {
        confirmButtonText: 'OK',
    })
});
}

const DeleteSelectedItems = () => {
  var skus=Array()
  for (let item of selected_items.value.values()) {
    skus.push(item.SKU)
    }
  ElMessageBox.confirm(
    'そうしたら、選んだアイテムが買い物かごから削除されました。もう一度確認してください',
    'Warning',
    {
      confirmButtonText: 'はい',
      cancelButtonText: 'また考えさせてください',
      type: 'warning',
    }
  )
    .then(() => {
      axios.delete('/api/cart/delete',{
    params: {
      skus: skus
  },
  paramsSerializer: params => {
    return qs.stringify(params, {arrayFormat: 'repeat'})
  }
} )
.then((response) => {
  if(response.status==200){
    updatecart()
  }
})
.catch((error) => {
  ElMessageBox.alert('削除が失敗しました', 'ご注意', {
        confirmButtonText: 'OK',
    })
})
      ElMessage({
        type: 'success',
        message: 'Delete completed',
      })
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: 'Delete canceled',
      })
    })
}

const HandleAllSeleted = () => {
  if(selecttag.value==true){
    for (let index = 0; index < items.value.length; index++) {
    items.value[index].select=ref(true)
    selected_items.value.set(items.value[index].Product.Name,{SKU:items.value[index].Product.SKU,Ind:index,PID:items.value[index].Product.ID}) 
  }
  }else{
    for (let index = 0; index < items.value.length; index++) {
    items.value[index].select=ref(false)
    selected_items.value.delete(items.value[index].Product.Name)
  }
  }
}

  </script>
  