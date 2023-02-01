<template>
  <div class="table-container">
    <StatusTable/>
  </div>
</template>

<script>
import StatusTable from "@/components/StatusTable.vue";
export default {
  name: 'HomeView',
  components: {
    StatusTable
  }
}
</script>

<style>
  .table-container{
    width: 50%;
  }
</style>