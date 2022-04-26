<template>
  <div class="home">
    <v-card style="height: 850px;overflow: scroll"
            class="mx-auto"
    >
      <v-list>
        <template v-for="item in mainStore.currentCategoryArticleList">
          <v-list-item
              :key="item.id"
              @click="gotoPage(item.id)"
          >

            <v-list-item-content>
              <v-list-item-title v-html="item.title"></v-list-item-title>
              <v-list-item-subtitle v-html="item.createTime"></v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
          <v-divider :key="item.id"></v-divider>

        </template>
      </v-list>
    </v-card>


  </div>
</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';
import {mapActions} from 'pinia';
import {useMainStore} from '@/store/main';
import {useLanguageStore} from '@/store/i18n';
import CategoryNavigationDrawer from '@/components/CategoryNavigationDrawer.vue'


@Component({
  components: {
    CategoryNavigationDrawer
  },
  methods: {
    ...mapActions(useLanguageStore, ["CHANGE_LANG"])
  },
})
export default class HomeView extends Vue {
  public mainStore = useMainStore()

  public gotoPage(articleId: number) {
    this.$router.push({path: "/page", params: {id: articleId + ""}})
  }

}
</script>
<style scoped>


</style>