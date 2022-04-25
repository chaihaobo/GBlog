<template>
  <v-navigation-drawer
      v-model="show"
      app
      dark
      :width="200"
      :absolute="true"

  >
    <v-list>
      <v-list-item
          link
      >
        <v-list-item-content>
          <v-list-item-title>{{ $t("all-category") }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
      <v-list-item
          v-for="item in items"
          :key="item.id"
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
import {Component, Vue,Prop} from 'vue-property-decorator';
import categoryController from "@/controller/category-controller";
import Category from "@/model/Category";

@Component({
  components: {},
})
export default class HomeView extends Vue {
  public items: Category[] = [];
  @Prop({ default: true }) show!: boolean;

  public async created() {
    this.items = await categoryController.categoryList();
  }


}
</script>