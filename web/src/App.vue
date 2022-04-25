<template>
  <v-app>
    <CategoryNavigationDrawer :show="showDrawer"/>
    <v-app-bar
        app
        dark
    >
      <v-app-bar-nav-icon @click.stop="showDrawer = !showDrawer"></v-app-bar-nav-icon>
    </v-app-bar>
    <!-- Sizes your content based upon application components -->
    <v-main>

      <!-- Provides the application the proper gutter -->
      <v-container style="height: 100%;max-height: 100%" fluid>

        <!-- If using vue-router -->
        <router-view style="height: 100%;max-height: 100%;overflow-x: hidden;overflow-y: scroll"></router-view>
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
    showDrawer: true,
    drawer: false,
  }),
});
</script>

<style>

body {
  height: 100%;
  width: 100%;
}
</style>
