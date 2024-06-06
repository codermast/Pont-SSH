export namespace service {
	
	export class SSHConfig {
	    server: string;
	    port: number;
	    username: string;
	    password: string;
	    key_path: string;
	
	    static createFrom(source: any = {}) {
	        return new SSHConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.server = source["server"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.key_path = source["key_path"];
	    }
	}

}

export namespace utils {
	
	export class Result {
	    code: number;
	    data: any;
	    msg: string;
	
	    static createFrom(source: any = {}) {
	        return new Result(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.code = source["code"];
	        this.data = source["data"];
	        this.msg = source["msg"];
	    }
	}

}

