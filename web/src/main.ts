import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import _axios from "@/api/index"
import i18n from "@/i18n";

import {createPinia,PiniaVuePlugin} from "pinia";





Vue.config.productionTip = false
Vue.prototype.$http = _axios
Vue.use(PiniaVuePlugin)
const pinia = createPinia()

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore
new Vue({
    el:"#app",
    router,
    vuetify,
    pinia,
    i18n,
    render: h => h(App)
});