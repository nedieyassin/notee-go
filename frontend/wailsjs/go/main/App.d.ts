// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT
import {store} from '../models';

export function CreateNote(arg1:store.Note):Promise<number>;

export function DeleteNote(arg1:number):Promise<boolean>;

export function FavNote(arg1:number,arg2:number):Promise<boolean>;

export function GetNote(arg1:number):Promise<store.Note>;

export function GetNotes():Promise<Array<store.Note>>;

export function RenameNote(arg1:number,arg2:string):Promise<boolean>;

export function UpdateNote(arg1:number,arg2:string):Promise<boolean>;
