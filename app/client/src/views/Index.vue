<template>
  <div id="index">
<el-row :gutter="20">
  <el-col :span="6"><div class="cm_container"><el-image :src="logo_uri" style="width: 10rem;height: 10rem;min-width: 10rem;"></el-image></div></el-col>
    <el-col :span="12">
      <div id="cm_container">
        <el-carousel indicator-position="outside" id="cm" height="20rem">
          <el-carousel-item v-for="cm in cm_lsit" :key="cm.cm_id">
            <el-image :src="cm.url" fit="scale-down" style="width: 100%; height: 100%"></el-image>
          </el-carousel-item>
        </el-carousel>
      </div>
    </el-col>
    <el-col :span="6"><div class="cm_container" /></el-col>
  </el-row>
  <div id="category_container">
      <el-tabs v-model="activeName" tab-position="left" style="height: 46rem" class="demo-tabs" @tab-change="handlechange">
        <el-tab-pane :label="item.Name" :name="item.Name" v-for="item in options">
          <div id="productcontent">
        <div class="product_items" v-for="item in products" @click="handleproductclick(item)">
          <div class="card-header">
            <span>{{item.Name}}</span>
          </div>
          <el-image :src="item.Url" lazy loading="eager" fit="fill" class="proimg"></el-image>
          <div class="card-footer">{{item.Price}}円</div>
        </div>
          </div>
        </el-tab-pane>
      </el-tabs>
    </div>
    <div id="pagination">
          <el-pagination  background layout="prev, pager, next" :page-size="pageSize" :total="totalcount" @current-change="handlePageChange"/>
      </div>
</div>
</template>
  
  
  <style scoped>
  .proimg{
    height: 10rem;
  }
  #pagination{

    display: flex;
    justify-content: center;
  }
  .items{
    height: 14rem;
  }
.card-header{
  height: 3rem;
  border-bottom: 1px solid #01010132;
}
.card-footer{
  height: 2rem;
  border-top: 1px solid #01010132;
  overflow:hidden; 
text-overflow:ellipsis; 
white-space:nowrap; 
}
  #productcontent{
    display: flex;
    flex-wrap: wrap;
    align-items: top;
    height: 46rem;
  }
.product_items{
  cursor: pointer;
  text-decoration: none;
  color: black;
  margin-left: 0.3rem;
  width: 15rem;
    height: 15rem;
    display: flex;
    flex-direction: column;
    align-items: center;
    background-color: white;
    border-radius: 0.2rem;
  }
  #categorylink{
    width: 5rem;
    margin-top: 1rem;
    cursor: pointer;
    font-size: 0.8rem;
  }
  #categorylink:hover{
    color: red;
    text-decoration: underline;
  }
  .el-row {
    margin-bottom: 20px;
  }
  .el-row:last-child {
    margin-bottom: 0;
  }
  .el-col {
    border-radius: 4px;
    min-width: 15rem;
  }

#index{
  width: 93rem;
}
#category_container,#product_items{
    border-radius: 4px;
    min-height: 10rem;
  }
  #category_container{
    display: flex;
    flex-direction: column;
    align-items: top;
    background-color: #fff;
    width: 87rem;
    border-radius: 3rem;
  }
  #cm_container{
    display: flex;
    justify-content: center;
  }
  .cm_container{
    height: 20rem;
    display: flex;
    justify-content: left;
    align-items: center;
  }
  #cm{
    width: 50rem;
    min-width: 50rem;
  }
  </style>
  
  
<script setup lang="ts">

import { onMounted,ref,reactive, toRaw} from 'vue';
  import axios from 'axios';
import router from '@/router';
import { global_state } from '@/store';
import type { TabsPaneContext} from 'element-plus'
import { ElMessage, ElMessageBox, type Action } from 'element-plus';
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'
import { genFileId } from 'element-plus'

var cm_lsit=ref([{cm_id:1,url:"http://localhost:8080/static/cm01.jpg"},{cm_id:2,url:"http://localhost:8080/static/cm02.jpg"},{cm_id:3,url:"http://localhost:8080/static/cm03.jpg"},{cm_id:4,url:"http://localhost:8080/static/cm04.jpg"}])
var logo_uri=ref("http://localhost:8080/static/logo.png")
var activeName=ref("全部")

var page_id=1
var pageSize=ref(10)
var products=ref(Array())
var totalcount=ref(0)
var options=ref(Array())
var tag=false
var gs=global_state()

const handlePageChange = (val: number) => {
  if(tag==true){
    tag=false
  }else{
    var categoryname=activeName.value
  updatePage({page:val,pageSize:pageSize.value,Cname:categoryname},false)
  }
}

const handlechange=(pane: TabsPaneContext, ev: Event)=>{
  var categoryname=activeName.value
  if(page_id!=1){
    page_id=1
    tag=true
  }
 updatePage({page:page_id,pageSize:pageSize.value,Cname:categoryname},true)
}


function handleproductclick(obj :{}){
  router.replace("/product")
  gs.seletproduct=obj
}

function updatePage(obj:{},tag:boolean){
  axios.get('/api/getproductsinpage',{params:obj})
.then((response) => {
  if(response.status==200){
    console.log(response.data)
    products.value=response.data.items
    page_id=response.data.page
    if(tag){
      totalcount.value=response.data.totalCount
    }
  }
})
.catch((error) => {
  console.log(error)
})
}
function InitCategory(){
  axios.get('/api/getallcategories')
.then((response) => {
  if(response.status==200){
    response.data.items.splice(0, 0, {Name:"全部"})
    options.value=response.data.items
  }
})
.catch((error) => {
  console.log(error)
})
}
onMounted(()=>{
  InitCategory()
  updatePage({page:page_id,pageSize:pageSize.value},true)}
)

</script>
