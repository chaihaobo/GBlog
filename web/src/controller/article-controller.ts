import api from "@/api";
import apiConstant from '@/api/constant'
import Response from "@/model/Response";
import Article from "@/model/Article";

export default new class ArticleController {
    public async loadArticleList(categoryId: number): Promise<Article[]> {
        const axiosResponse = await api.get(apiConstant.articleList, {params: {categoryId}});
        const response = axiosResponse as unknown as Response<Article[]>;
        return response.data;
    }
}