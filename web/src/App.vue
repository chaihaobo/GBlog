<template>
  <v-app>
    <v-app-bar
        app
        dark
        :clipped-left="true"
    >
    <CategoryNavigationDrawer/>


      <v-toolbar-title>{{ $t("contentTitle") }}</v-toolbar-title>
    </v-app-bar>
    <!-- Sizes your content based upon application components -->
    <v-main>

      <!-- Provides the application the proper gutter -->
      <v-container fluid>

        <!-- If using vue-router -->
        <router-view></router-view>
      </v-container>
    </v-main>

  </v-app>

</template>

<script lang="ts">
import Vue from 'vue';
import CategoryNavigationDrawer from '@/components/CategoryNavigationDrawer.vue'


export default Vue.extend({
  name: 'App',
  components: {
    CategoryNavigationDrawer
  },

  data: () => ({
    item: 1,
    items: [
      {text: 'Real-Time', icon: 'mdi-clock'},
      {text: 'Audience', icon: 'mdi-account'},
      {text: 'Conversions', icon: 'mdi-flag'},
    ],
    drawer: false,
  }),
});
</script>

<style>

body{
  height: 100%;
  width: 100%;
}
</style>
