<template>
  <el-table class="table"
      :data="Data"
      border
      style="width: 100%">
    <el-table-column
        fixed
        prop="user_id"
        label="用户编号"
        width="120">
    </el-table-column>
    <el-table-column
        prop="username"
        label="用户名"
        width="120">
    </el-table-column>
    <el-table-column
        prop="email"
        label="邮箱"
        width="150">
    </el-table-column>
    <el-table-column
        prop="sex"
        label="性别"
        width="120">
    </el-table-column>
  </el-table>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
        Data:null,
    }
  },
  async created(){
    const res = await axios.get("http://127.0.0.1:8000")
    this.Data = await res.data
  }
}
</script>

<style>
  .table{
    margin: auto;
  }
</style>