import api from "@/api";

export default new class CategoryController {

    public async categoryList() {
        const axiosResponse = await api.get("/category/list");
        console.log(axiosResponse);
        return axiosResponse
    }


}