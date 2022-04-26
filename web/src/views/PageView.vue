<template>
  <div>
    ss


  </div>
</template>

<script lang="ts">
import {Component, Vue} from "vue-property-decorator";
import {useMainStore} from "@/store/main";

@Component
export default class PageView extends Vue {
  public mainStore = useMainStore()

}

</script>
<style scoped>

</style>