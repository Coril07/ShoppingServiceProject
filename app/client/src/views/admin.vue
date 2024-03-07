<template>
<div id="admin">
  <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
    <el-tab-pane label="catergory_m" name="catergory_m">
      <div id="catergory_m">
        <el-row :gutter="20">
          <el-col :span="4">
            <el-form :model="new_category">
              <el-form-item label="種類">
                <el-input v-model="new_category.name" clearable/>
               </el-form-item> 
              <el-form-item label="概要">
                <el-input v-model="new_category.desc" clearable/>
              </el-form-item>
              <el-form-item>
                <div style="display: flex;justify-content: center;width:100%">
                  <el-button type="primary" @click="onSubmit"   class="temp_btn" round>追加</el-button>
                </div>
              </el-form-item>
            </el-form>
          </el-col>
          <el-col :span="14">
            <div id="tablepart">
            <el-table :data="categories" style="width: 100%"  @selection-change="handleSelectionChange" border>
              <el-table-column type="selection" width="55" />
              <el-table-column prop="Name" label="種類" width="180" />
              <el-table-column prop="Desc" label="概要" width="180" />
              <el-table-column prop="related" label="関連商品" />
            </el-table>
            <div style="display: flex;justify-content: center;width:100%; margin-top: 1rem;">
                  <el-button type="primary" @click="onSubmit_del"   class="temp_btn" round>削除</el-button>
            </div>
            <div id="pagination">
              <el-pagination  background layout="prev, pager, next" :page-size="pagesize" :total="categorytotalcount" @current-change="handleCurrentChange"/>
            </div>
          </div>
          </el-col>
        </el-row>
      </div>
    </el-tab-pane>
    <el-tab-pane label="product_m" name="product_m">
      <div id="product_m">
        <el-row :gutter="20">
          <el-col :span="4">
            <el-form :model="new_product">
              <el-form-item label="商品名称">
                <el-input v-model="new_product.name" clearable/>
               </el-form-item> 
              <el-form-item label="概要">
                <el-input v-model="new_product.desc" clearable/>
              </el-form-item>
              <el-form-item label="値段">
                <el-input v-model="new_product.price" clearable/>
              </el-form-item>
              <el-form-item label="数量">
                <el-input v-model="new_product.count" clearable/>
              </el-form-item>
              <el-form-item label="種類選択">
                <el-select v-model="new_product.categoryName" placeholder="" size="large" style="width: 240px">
                  <el-option v-for="item in options" :key="item.Name" :label="item.Name":value="item.Name"/>
                </el-select>
              </el-form-item>
              <el-form-item>
                <div>選んだ商品に写真を付け足し</div>
                <div style="display: flex;justify-content: center;width:100%">
                  <el-select v-model="selectedpro" placeholder="" size="large" style="width: 240px">
                  <el-option v-for="item in option_pro" :key="item.Name" :label="item.Name":value="item.Name"/>
                  </el-select>
                </div>
              </el-form-item>
              <el-form-item>
                <div style="display: flex;justify-content: center;width:100%">
                  <el-button type="primary" @click="onSubmit_pro"   class="temp_btn" round>追加</el-button>
                </div>
              </el-form-item>
              <el-form-item>
                <div style="display: flex;justify-content: center;width:100%">
                  <el-upload ref="upload" class="upload-demo" :action="action" :auto-upload="false" :limit="1" :on-exceed="handleExceed" :http-request="uploadFile">
                    <template #trigger>
                      <el-button type="primary">select file</el-button>
                    </template>
                    <el-button class="ml-3" type="success" @click="submitUpload" style="margin-left: 2rem;">
                      upload
                    </el-button>
                    <template #tip>
                      <div class="el-upload__tip text-red">
                        limit 1 file, new file will cover the old file
                      </div>
                    </template>
                   </el-upload>
                </div>
              </el-form-item>
            </el-form>
          </el-col>
          <el-col :span="20">
            <el-table :data="products" style="width: 100%" @selection-change="handleSelectionChangepro" border>
              <el-table-column type="selection" width="55" />
              <el-table-column prop="Name" label="商品名" width="180" />
              <el-table-column prop="Desc" label="概要" width="180" />
              <el-table-column prop="Price" label="値段" width="180" />
              <el-table-column prop="StockCount" label="库存" width="180" />
              <el-table-column prop="SKU" label="SKU" width="180" />
              <el-table-column prop="CategoryName" label="種類" width="180" />
            </el-table>
            <div id="pagination2">
              <el-pagination  background layout="prev, pager, next" :page-size="pagesize" :total="producttotalcount" @current-change="handleCurrentChange2"/>
            </div>
            <div style="display: flex;justify-content: center;width:100%; margin-top: 1rem;">
                  <el-button type="primary" @click="onSubmit_del_pro"   class="temp_btn" round>削除</el-button>
            </div>
          </el-col>
        </el-row>
      </div>
    </el-tab-pane>
  </el-tabs>
</div>
</template>
    
<style scoped>
  #admin{
    width: 93rem;
  }
  #catergory_m{
    width: 93rem;

  }
  #product_m{
    width: 93rem;
  }
  #pagination,#pagination2{
    display: flex;
    justify-content: center;
    margin-top: 5rem;
  }
</style>
    
    
  <script setup lang="ts">
  import { onMounted,ref,reactive, toRaw} from 'vue';
  import axios from 'axios';
import router from '@/router';
import { global_state } from '@/store';
import type { TabsPaneContext} from 'element-plus'
import { ElMessage, ElMessageBox, type Action } from 'element-plus';
import qs from "qs"
import type { UploadInstance, UploadProps, UploadRawFile } from 'element-plus'
import { genFileId } from 'element-plus'

var selectedpro=ref("")
var action=ref("api/uploadproductimg")
const upload = ref<UploadInstance>()
const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)
}
const activeName = ref('catergory_m')

var pagesize=ref(10)
var categorypage_id=1
var categorytotalcount=ref(0)
var productpage_id=1
var producttotalcount=ref(0)

var new_category=ref({  
  name:"",
  desc:""
})
var selected_items=Array()
var selected_itemspro=Array()

var categories=ref(Array())


function updateCategoryTable(obj:{},tag:boolean){
  axios.get('/api/getcategoriesinpage',{params:obj})
.then((response) => {
  if(response.status==200){
    categories.value=response.data.items
    categorypage_id=response.data.page
    if(tag){
      categorytotalcount.value=response.data.totalCount
    }
  }
})
.catch((error) => {
  console.log(error)
})
}

// 在api.js中编写上传方法，并导出
const uploadFileRequest = (url, params) => {
  return axios({
    method: 'post',
    url: `${url}`,
    data: params,
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  });
}
const uploadFile=(params)=>{
        const file = params.file;
        var form = new FormData();
        form.append("file", file);

        form.append("Name", selectedpro.value);
        uploadFileRequest("/api/uploadproductimg",form).then(resp=>{
          if(resp && resp.status == 200){
            ElMessageBox.alert('upload終了。', 'ご注意', {
            confirmButtonText: 'OK',
          })
          }
        })
      }

const submitUpload = () => {
  upload.value!.submit()
}

const handleSelectionChange = (val: {}[]) => {
  selected_items=val
}
const handleSelectionChangepro= (val: {}[]) => {
  selected_itemspro=val
}
const onSubmit = () => {

axios.post('/api/addcategory', new_category.value)
.then((response) => {
  if(response.status==201){
    categories.value.push({name:new_category.value.name,desc:new_category.value.desc})
    new_category.value.name=""
    new_category.value.desc=""
    ElMessageBox.alert('追加が終わり。', 'ご注意', {
        confirmButtonText: 'OK',
    })
    updateCategoryTable({page:categorypage_id,pageSize:pagesize.value},true)
  }
})
.catch((error) => {
  ElMessageBox.alert('種類が重ねる', 'ご注意', {
        confirmButtonText: 'OK',
    })
})
}
const onSubmit_del_pro= () => {
  var t=Array()
  if(selected_itemspro.length!=0){
    for (let index = 0; index < selected_itemspro.length; index++) {
      const element = selected_itemspro[index].SKU;
      t.push(element)
    }
    axios.delete('/api/deleteselectedproducts',{
    params: {
      seleteditems: t
  },
  paramsSerializer: params => {
    return qs.stringify(params, {arrayFormat: 'repeat'})
  }
} )
.then((response) => {
  if(response.status==200){
    updateproductTable({page:productpage_id,pageSize:pagesize.value},true)
  }
})
.catch((error) => {
  ElMessageBox.alert('削除が失敗しました', 'ご注意', {
        confirmButtonText: 'OK',
    })
})
  }
  }
const onSubmit_del = () => {
  var t=Array()
  if(selected_items.length!=0){
    for (let index = 0; index < selected_items.length; index++) {
      const element = selected_items[index].Name;
      t.push(element)
    }
    axios.delete('/api/deleteselectedcategories',{
    params: {
      seleteditems: t
  },
  paramsSerializer: params => {
    return qs.stringify(params, {arrayFormat: 'repeat'})
  }
} )
.then((response) => {
  if(response.status==200){
    updateCategoryTable({page:categorypage_id,pageSize:pagesize.value},true)
  }
})
.catch((error) => {
  ElMessageBox.alert('削除が失敗しました', 'ご注意', {
        confirmButtonText: 'OK',
    })
})
  }
  }
var gs=global_state()
onMounted(()=>{
  updateCategoryTable({page:categorypage_id,pageSize:pagesize.value},true)
}
)

const handleCurrentChange = (val: number) => {
  updateCategoryTable({page:val,pageSize:pagesize.value},false)
}

const handleCurrentChange2 = (val: number) => {
  updateproductTable({page:val,pageSize:pagesize.value},false)
}
//........................................................................
var products=ref(Array())
var options=ref(Array())
var option_pro=ref(Array())
function updateproductTable(obj:{},tag:boolean){
  axios.get('/api/getproductsinpage',{params:obj})
.then((response) => {
  if(response.status==200){
    products.value=response.data.items
    productpage_id=response.data.page
    if(tag){
      producttotalcount.value=response.data.totalCount
    }
  }
})
.catch((error) => {
  console.log(error)
})

axios.get('/api/getallcategories')
.then((response) => {
  if(response.status==200){
    options.value=response.data.items
  }
})
.catch((error) => {
  console.log(error)
})

axios.get('/api/getallproducts')
.then((response) => {
  if(response.status==200){
    option_pro.value=response.data.items
  }
})
.catch((error) => {
  console.log(error)
})
}

var new_product=ref({  
  name:"",
  desc:"",
  price:0.0,
  count:0,
  categoryName:""
})

const onSubmit_pro = () => {
  axios.post('/api/createproduct', new_product.value)
.then((response) => {
  if(response.status==201){
    ElMessageBox.alert('完成', 'ご注意', {
        confirmButtonText: 'OK',
    })
    updateproductTable({page:productpage_id,pageSize:pagesize.value},true)
  }
})
.catch((error) => {
  console.log(error)
})
}
const handleClick = (tab: TabsPaneContext, event: Event) => {
  if (tab.props.name=="product_m"){
    updateproductTable({page:productpage_id,pageSize:pagesize.value},true)
}
}
const handlecatechange=()=>{

}



</script>

  