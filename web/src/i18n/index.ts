import Vue from 'vue'
import VueI18n from 'vue-i18n'

Vue.use(VueI18n)

const messages = {
    en: require('./en.json'),
    zh: require('./zh.json')
}


const i18n = new VueI18n({
    locale: "zh",
    messages
})

export default i18n