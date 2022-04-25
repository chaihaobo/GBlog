<template>
  <v-navigation-drawer
      absolute
      dark
      :src="imageUrl"
      width="100%"
      permanent
  >
    <v-list>
      <v-list-item
          v-for="([icon, text], i) in items"
          :key="i"
          link
      >
        <v-list-item-icon>
          <v-icon>{{ icon }}</v-icon>
        </v-list-item-icon>

        <v-list-item-content>
          <v-list-item-title>{{ text }}</v-list-item-title>
        </v-list-item-content>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>

</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';
import categoryController from "@/controller/category-controller";
import _axios from "@/api";
import Category from "@/model/Category";
@Component({
  components: {},
})
export default class HomeView extends Vue {
  public item = 1;
  public imageUrl = require('../../public/images/category-navigation-drawer.png')
  public items = [
    ['mdi-email', 'Inbox'],
    ['mdi-account-supervisor-circle', 'Supervisors'],
    ['mdi-clock-start', 'Clock-in'],
  ];
  public categoryList!:Category[];

  public async created() {
    let categories = await categoryController.categoryList();
    console.log(categories)


  }


}
</script>