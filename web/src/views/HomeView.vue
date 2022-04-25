<template>
  <div class="home">
      <v-card style="height: 850px;overflow: scroll"
              class="mx-auto"
      >
        <v-list>
          <template v-for="(item,i) in items">
            <v-list-item
                :key="item.title"
            >
              <v-list-item-avatar>
                <v-img :src="item.avatar"></v-img>
              </v-list-item-avatar>

              <v-list-item-content>
                <v-list-item-title v-html="item.title"></v-list-item-title>
                <v-list-item-subtitle v-html="item.subtitle"></v-list-item-subtitle>
              </v-list-item-content>
            </v-list-item>
            <v-divider :key="i"></v-divider>

          </template>
        </v-list>
      </v-card>


  </div>
</template>

<script lang="ts">
import {Component, Vue} from 'vue-property-decorator';
import {mapState,mapActions} from 'pinia';
import {useMainStore} from '@/store';
import {useLanguageStore} from '@/store/i18n';
import CategoryNavigationDrawer from '@/components/CategoryNavigationDrawer.vue'


@Component({
  components: {
    CategoryNavigationDrawer
  },
  computed: {
    ...mapState(useMainStore, ["name"]),
  },
  methods:{
    ...mapActions(useLanguageStore,["CHANGE_LANG"])
  },
})
export default class HomeView extends Vue {
  public items = [

    {
      avatar: 'https://cdn.vuetifyjs.com/images/lists/3.jpg',
      title: 'Oui oui',
      subtitle: "<span class='text--primary'>Sandra Adams</span> &mdash; Do you have Paris recommendations? Have you ever been?",
    },
    {
      avatar: 'https://cdn.vuetifyjs.com/images/lists/4.jpg',
      title: 'Birthday gift',
      subtitle: "<span class='text--primary'>Trevor Hansen</span> &mdash; Have any ideas about what we should get Heidi for her birthday?",
    },
    {
      avatar: 'https://cdn.vuetifyjs.com/images/lists/5.jpg',
      title: 'Recipe to try',
      subtitle: "<span class='text--primary'>Britta Holt</span> &mdash; We should eat this: Grate, Squash, Corn, and tomatillo Tacos.",
    },
  ]

}
</script>
<style scoped>


</style>