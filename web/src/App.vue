<template>
  <v-app>
    <v-navigation-drawer width="500px" app>
      <CategoryNavigationDrawer/>
    </v-navigation-drawer>

    <!-- Sizes your content based upon application components -->
    <v-main>

      <!-- Provides the application the proper gutter -->
      <v-container fluid>

        <!-- If using vue-router -->
        <router-view></router-view>
      </v-container>
    </v-main>

    <v-footer app>
      <!-- -->
    </v-footer>
  </v-app>
</template>

<script lang="ts">
import Vue from 'vue';
import CategoryNavigationDrawer from './components/CategoryNavigationDrawer.vue'

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
  }),
});
</script>
