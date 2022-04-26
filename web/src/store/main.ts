import {defineStore} from 'pinia'
import categoryController from "@/controller/category-controller";
import Category from "@/model/Category";
import Article from "@/model/Article";
import articleController from "@/controller/article-controller";

export type RootState = {
    categoryList: Category[];
    currentCategoryId: number;
    currentCategoryArticleList: Article[];
};

export const useMainStore = defineStore({
    id: 'main',
    state: () => ({
        categoryList: [],
        currentCategoryId: 0,
        currentCategoryArticleList: []
    }) as RootState,
    actions: {
        async loadCategoryList() {
            this.categoryList = await categoryController.categoryList()
        },
        async loadArticleByCategoryId(categoryId: number) {
            this.currentCategoryId = categoryId
            this.currentCategoryArticleList = await articleController.loadArticleList(categoryId);
        }

    }
})
