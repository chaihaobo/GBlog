import api from "@/api";
import {AxiosInstance} from "axios";

export default class BaseController {
    private api: AxiosInstance;

    constructor() {
        this.api = api
    }

}

