<template>
  <div class="home">
    ok-{{name}}
  </div>
</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';
import {mapState} from 'pinia';
import {useMainStore} from '@/store';


@Component({
  components: {},
  computed: {
    ...mapState(useMainStore, ["name"]),
  }
})
export default class HomeView extends Vue {

}
</script>
