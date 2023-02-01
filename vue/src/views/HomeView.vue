<template>
  <div class="home">
    <StatusTable/>
  </div>
</template>

<script>
// @ is an alias to /src
import StatusTable from "@/components/StatusTable.vue";

export default {
  name: 'HomeView',
  components: {
    StatusTable
  }
}
</script>