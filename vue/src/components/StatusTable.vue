<template>
  <el-table class="table" :data="Data" border align="center">
    <el-table-column fixed prop="user_id" label="用户编号" width="80" align="center">
    </el-table-column>
    <el-table-column prop="username" label="用户名" width="120" align="center">
    </el-table-column>
    <el-table-column prop="email" label="邮箱" width="180" align="center">
    </el-table-column>
    <el-table-column prop="sex" label="性别" align="center">
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