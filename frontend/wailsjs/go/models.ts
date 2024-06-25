export namespace entity {
	
	export class Config {
	
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	
	    }
	}
	export class LogInfo {
	    id: string;
	    server_id: string;
	    exec_date: string;
	    command: string;
	    time: string;
	
	    static createFrom(source: any = {}) {
	        return new LogInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.server_id = source["server_id"];
	        this.exec_date = source["exec_date"];
	        this.command = source["command"];
	        this.time = source["time"];
	    }
	}
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
	export class SSHConfig {
	    id: string;
	    server: string;
	    port: number;
	    username: string;
	    password: string;
	    name: string;
	    edit: boolean;
	    type: number;
	
	    static createFrom(source: any = {}) {
	        return new SSHConfig(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.server = source["server"];
	        this.port = source["port"];
	        this.username = source["username"];
	        this.password = source["password"];
	        this.name = source["name"];
	        this.edit = source["edit"];
	        this.type = source["type"];
	    }
	}

}

