import Article from "@/model/Article";

export default class Category {
   private _id!:number;
   private _name!:string;
   private _createTime!:string;
   private _updateTime!:string;
   private _article!:Article[];


   get id(): number {
      return this._id;
   }

   set id(value: number) {
      this._id = value;
   }

   get name(): string {
      return this._name;
   }

   set name(value: string) {
      this._name = value;
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

   get article(): Article[] {
      return this._article;
   }

   set article(value: Article[]) {
      this._article = value;
   }
}