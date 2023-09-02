export namespace store {
	
	export class Note {
	    id: number;
	    title: string;
	    body: string;
	    isDir: number;
	    parentId: number;
	    isFav: number;
	    createTimestamp: number;
	    updateTimestamp: number;
	
	    static createFrom(source: any = {}) {
	        return new Note(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.body = source["body"];
	        this.isDir = source["isDir"];
	        this.parentId = source["parentId"];
	        this.isFav = source["isFav"];
	        this.createTimestamp = source["createTimestamp"];
	        this.updateTimestamp = source["updateTimestamp"];
	    }
	}

}

