<template>
    <div id="order">
      <div id="cartbox">
        <div id="cartheader">
          <div style="margin-left: 1.2rem;font-weight: bolder;">ご注文（全部{{ }}）</div>
        </div>
        <div id="cartbody">
          <div id="title">
            <div class="isdeleted"><el-checkbox label="全選択" size="large" @change="HandleAllSeleted" v-model="selecttag"/></div>
            <div class="info">商品情報</div>
            <div class="price">単価</div>
            <div class="amount">数量</div>
            <div class="fee">費用</div>
            <div class="time">購入時間</div>
          </div>
          <div class="content" v-for="item in items">
            <div class="isdeleted"><el-checkbox  size="large" style="margin-left: -3rem;" @change="HandleSeleted(item)" v-model="item.select"  /></div>
            <div v-for="sub in item.OrderedItems" style="margin-left: 8rem;">
            <div class="info" style="padding: 1rem;">
              <div class="productshow">
                <el-image :src="sub.Product.Url" lazy loading="eager" fit="fill" style="max-width: 7rem;"></el-image>
                <div class="Name">{{sub.Product.Name}}</div>
              </div>
          </div>
            <div class="price">{{sub.Product.Price}}</div>
            <div class="amount"><el-input-number v-model="sub.Count" :min="1" :max="50" size="small" style="width: 5rem;" disabled /></div>
            <div class="fee">{{sub.Product.Price*sub.Count}}</div>
            <div class="time"></div>
            </div>
            <div>
            <div class="fee" style="margin-left: 63rem;">{{item.TotalPrice}}</div>
            <div class="time" style="margin-left: 3rem;">{{item.CreatedAt}}</div>
        </div>
          </div>
        </div>
        <div id="cartfooter">
          <div id="deleteall" @click="DeleteSelectedItems">キャンセル</div>
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
.content{
  cursor: pointer;
  width: 100%;
  display: flex;
  flex-direction: column;
  border-bottom: 0.1rem solid #e6e3e4;
  background-color: #f5f5f5;
}
.content div{
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
.time{
  width: 10rem;
}
#title{
  width: 100%;
  display: flex;
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
  #order{
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
  updateorder()
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
    selected_items.value.set(item.ID,item.ID) 
  }else{
    selected_items.value.delete(item.ID)
  }
  console.log(selected_items)
}


const updateorder= () => {
  axios.get('/api/order/get')
.then((response) => {
  if(response.status==200){
    tag=false
    items.value=response.data.items
    console.log(items)
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
  var IDs=Array()
  for (let item of selected_items.value.values()) {
    IDs.push(item)
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
      axios.delete('/api/order/delete',{
    params: {
      orderIds: IDs
  },
  paramsSerializer: params => {
    return qs.stringify(params, {arrayFormat: 'repeat'})
  }
} )
.then((response) => {
  if(response.status==201){
    updateorder()
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
    selected_items.value.set(items.value[index].ID,items.value[index].ID) 
  }
  }else{
    for (let index = 0; index < items.value.length; index++) {
    items.value[index].select=ref(false)
    selected_items.value.delete(items.value[index].ID)
  }
  }
}

  </script>
  