export namespace main {
	
	export class MediaAsset {
	    id: string;
	    path: string;
	    thumbnail: string;
	
	    static createFrom(source: any = {}) {
	        return new MediaAsset(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.path = source["path"];
	        this.thumbnail = source["thumbnail"];
	    }
	}

}

