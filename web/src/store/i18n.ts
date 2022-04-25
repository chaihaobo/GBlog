import {defineStore} from 'pinia'
import i18n from '@/i18n/';

export interface langState {
    locales?: objType [],
    locale?: string,
}

interface objType {
    [name: string]: number | string | boolean;
}

export const useLanguageStore = defineStore({
    id: 'language',
    state: () => ({
        name: '超级管理员',
        locales: [
            {
                value: "en",
                label: "英文"
            },
            {
                value: "zh",
                label: "中文"
            }
        ],
        locale: "zh"

    }),
    getters: {
        getSelectLang: (state) => state.locales,
        getLang: (state) => state.locale,
    },
    actions: {
        async CHANGE_LANG(lang: string) {
            this.locale = lang;
            i18n.locale = lang;
        },
    },

})
