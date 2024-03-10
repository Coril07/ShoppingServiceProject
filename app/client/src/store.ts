import { defineStore } from "pinia"

export const global_state=defineStore({
    id:"global_state",
    state:()=>{
      return {islogin:false,username:"",activePath:"",isadmin:false,seletproduct:{Url:"",
      Name:"",
      CategoryName:"",
      Desc:"",
      StockCount:"",
      Price:0,
      SKU:""}}
    }
  
  })