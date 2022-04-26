<template>
  <v-navigation-drawer
      v-model="show"
      app
      dark
      :width="200"
      :absolute="true"

  >
    <v-list>
      <v-list-item @click="gotoArticleList(0)" link>
        <v-list-item-content>
          <v-list-item-title>{{ $t("all-category") }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item
          v-for="item in store.categoryList"
          :key="item.id"
          @click="gotoArticleList(item.id)"
          link
      >
        <v-list-item-content>
          <v-list-item-title>{{ item.name }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>

</template>

<script lang="ts">
import {Component, Vue, Prop} from 'vue-property-decorator';
import {useMainStore} from "@/store/main";
import categoryController from "@/controller/category-controller";
import Category from "@/model/Category";
import {mapState} from "pinia";

@Component({
  components: {},
})
export default class HomeView extends Vue {
  public items: Category[] = [];
  public store = useMainStore()
  @Prop({default: true}) show!: boolean;

  public async created() {
    await this.store.loadCategoryList()
  }

  public async gotoArticleList(categoryId: number) {
    if (this.$route.path != "/") {
      this.$router.push("/")
    }
    this.store.loadArticleByCategoryId(categoryId)
  }


}
</script>