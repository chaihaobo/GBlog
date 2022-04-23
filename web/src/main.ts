import Vue from 'vue'
import App from './App.vue'
import router from './router'
import vuetify from './plugins/vuetify'
import _axios from "@/api/index"
import {createPinia,PiniaVuePlugin} from "pinia";


require("./mock/index")

Vue.config.productionTip = false
Vue.prototype.$http = _axios
Vue.use(PiniaVuePlugin)

const pinia = createPinia()
const app = new Vue({
    el:"#app",
    router,
    vuetify,
    pinia,
    render: h => h(App)
});
