<template>
  <div class="home">
    ok-{{name}}-{{this.$i18n.login}}
  </div>
</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';
import {mapState,mapActions} from 'pinia';
import {useMainStore} from '@/store';
import {useLanguageStore} from '@/store/i18n';


@Component({
  components: {},
  computed: {
    ...mapState(useMainStore, ["name"]),
  },
  methods:{
    ...mapActions(useLanguageStore,["CHANGE_LANG"])
  },
})
export default class HomeView extends Vue {

}
</script>
