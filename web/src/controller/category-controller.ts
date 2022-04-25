import api from "@/api";
import apiConstant from '@/api/constant'
import Category from "@/model/Category";
import Response from "@/model/Response";

export default new class CategoryController {

    public async categoryList(): Promise<Category[]> {
        const axiosResponse = await api.get(apiConstant.categoryList);
        const response = axiosResponse as unknown as Response<Category[]>;
        return response.data;
    }


}