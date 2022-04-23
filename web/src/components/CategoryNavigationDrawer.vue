<template>
  <v-navigation-drawer width="500px" app>
    <!-- -->
    <v-list dense>
      <v-list-item-group v-model="item" color="primary">
        <v-list-item
            v-for="(item, i) in items"
            :key="i"
        >
          <v-list-item-icon>
            <v-icon v-text="item.icon"></v-icon>
          </v-list-item-icon>
          <v-list-item-content>
            <v-list-item-title v-text="item.text"></v-list-item-title>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>

  </v-navigation-drawer>

</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';

@Component({
  components: {},
})
export default class HomeView extends Vue {
  public item = 1;
  public items = [
    {text: 'Real-Time', icon: 'mdi-clock'},
    {text: 'Audience', icon: 'mdi-account'},
    {text: 'Conversions', icon: 'mdi-flag'},
  ];

  public created(): void {
    this.$http.get("/api/category/list")
        .then(value => {
          console.log(value)
        })
    console.log(this.$http)
    console.log('created');
  }


}
</script>