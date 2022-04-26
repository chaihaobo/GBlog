import Category from "@/model/Category";

export default class Article{
    private _id!:number;
    private _title!:string;
    private _content!:string;
    private _createTime!:string;
    private _updateTime!:string;
    private _category!:Category;


    get id(): number {
        return this._id;
    }

    set id(value: number) {
        this._id = value;
    }

    get title(): string {
        return this._title;
    }

    set title(value: string) {
        this._title = value;
    }

    get content(): string {
        return this._content;
    }

    set content(value: string) {
        this._content = value;
    }

    get createTime(): string {
        return this._createTime;
    }

    set createTime(value: string) {
        this._createTime = value;
    }

    get updateTime(): string {
        return this._updateTime;
    }

    set updateTime(value: string) {
        this._updateTime = value;
    }

    get category(): Category {
        return this._category;
    }

    set category(value: Category) {
        this._category = value;
    }
}