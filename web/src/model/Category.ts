export default class Category {
   private _id!:number;
   private _name!:number;
   private _createTime!:string;

   get id(): number {
      return this._id;
   }

   set id(value: number) {
      this._id = value;
   }

   get name(): number {
      return this._name;
   }

   set name(value: number) {
      this._name = value;
   }

   get createTime(): string {
      return this._createTime;
   }

   set createTime(value: string) {
      this._createTime = value;
   }
}