export namespace dockerbackend {
	
	export class Port {
	    IP: string;
	    PublicPort: number;
	    PrivatePort: number;
	    Type: string;
	
	    static createFrom(source: any = {}) {
	        return new Port(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.IP = source["IP"];
	        this.PublicPort = source["PublicPort"];
	        this.PrivatePort = source["PrivatePort"];
	        this.Type = source["Type"];
	    }
	}
	export class Container {
	    ID: string;
	    Image: string;
	    Command: string;
	    // Go type: time
	    Created: any;
	    Status: string;
	    SizeMB: number;
	    Ports: Port[];
	    Names: string[];
	
	    static createFrom(source: any = {}) {
	        return new Container(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.Image = source["Image"];
	        this.Command = source["Command"];
	        this.Created = this.convertValues(source["Created"], null);
	        this.Status = source["Status"];
	        this.SizeMB = source["SizeMB"];
	        this.Ports = this.convertValues(source["Ports"], Port);
	        this.Names = source["Names"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Image {
	    ID: string;
	    RepoTags: string[];
	    RepoDigests: string[];
	    // Go type: time
	    Created: any;
	    Size: number;
	    SharedSize: number;
	    VirtualSize: number;
	    Labels: {[key: string]: string};
	
	    static createFrom(source: any = {}) {
	        return new Image(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ID = source["ID"];
	        this.RepoTags = source["RepoTags"];
	        this.RepoDigests = source["RepoDigests"];
	        this.Created = this.convertValues(source["Created"], null);
	        this.Size = source["Size"];
	        this.SharedSize = source["SharedSize"];
	        this.VirtualSize = source["VirtualSize"];
	        this.Labels = source["Labels"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}

}

