import { Page, Route , Request, BrowserContext, Response} from "playwright";
import { DoorEntity } from "../entity";
import { EventEmitter } from 'events';
const axios = require('axios');

export function getUrlParameter(url: string) {
    const urlObj = new URL(url);
    return new URLSearchParams(urlObj.search);
}

export abstract class Monitor<T = any> {

    finishTag: boolean = false;
    timeout: number;
    key : string | undefined = undefined;
    eventEmitter: EventEmitter;
    waitResolve: (value: DoorEntity<T> | PromiseLike<DoorEntity<T>>) => void = () => {};
    waitPromise: Promise<DoorEntity<T>>;
    hadListen: boolean = false;
    allowRepeat: boolean = false;
    startTag: boolean = false;

    constructor(timeout: number = 60000){
        this.timeout = timeout;
        this.eventEmitter = new EventEmitter();
        this.waitPromise = new Promise<DoorEntity<T>>((resolve) => {
            this.waitResolve = resolve;
        });
    }

    setAllowRepeat(allowRepeat: boolean){
        this.allowRepeat = allowRepeat;
    }

    public close(){
        this.eventEmitter = new EventEmitter();
        this.waitResolve = () => {};
        this.startTag=false;
        this.hadListen =false;
        this.finishTag = false;
        this.setAllowRepeat(false)
    }
    
    protected getItemKey(params : URLSearchParams): string | undefined {
        return undefined;
    }

    getItemKeys(url : string): string | undefined {
        return this.getItemKey(getUrlParameter(url));
    }

    abstract getKey(): string;

    public async doMatchResponse(response: Response){
        const url = response.url();
        const method = response.request().method();
        const headers = await response.request().allHeaders();
        const resourceType = response.request().resourceType();
        return await this.isMatch(url, method, headers, resourceType);
    }

    async filter(url : string, resourceType : string, method: string, headers: {[key: string]: string;}){
        return false;
    }

    abstract isMatch(url: string, method: string, headers: {[key: string]: string;}, resourceType : string): Promise<boolean>;  

    setFinishTag(finishTag: boolean){
        if(this.allowRepeat){
            return;
        }
        this.finishTag = finishTag;
    }


    public async start(){
        if(this.startTag){
            return;
        }
        const monitor = this;
        this.setFinishTag(false);
        this.listenEvent();
        setTimeout(async () => {
            if(this.allowRepeat ||monitor.finishTag){
                return;
            }
            await monitor._doCallback(new DoorEntity(false, {} as T));
            if(!this.allowRepeat){
                monitor.setFinishTag(true)
            }
        }, this.timeout);
        this.startTag = true;
    }

    listenEvent(){
        if(this.allowRepeat && !this.hadListen){
            this.eventEmitter.on(this.getEventKey(), async (result: DoorEntity<T>) => {
                await this.waitResolve(result);
            });
            this.hadListen = true;
            return;
        }
        this.onceEnvet();
    }

    async _doCallback(doorEntity: DoorEntity<T>, request: Request | undefined = undefined, response : Response | undefined = undefined) : Promise<void>{
        try{
            await this.doCallback(doorEntity, request, response);
        }finally{
            this.eventEmitter.emit(this.getEventKey(), doorEntity);
        }
    }

    getEventKey(): string{
        return this.constructor.name + "_" + 'actionCompleted';
    }

    async doCallback(doorEntity: DoorEntity<T>, request: Request | undefined = undefined, response : Response | undefined = undefined) : Promise<void>{

    }

    async onceEnvet(){
        this.eventEmitter.once(this.getEventKey(), async (result: DoorEntity<T>) => {
            this.waitResolve(result);
        });
    }

    setTimeout(timeout: number){
        this.timeout = timeout;
    }

    async waitForAction(): Promise<DoorEntity<T>>{
        return await this.waitPromise;
    }

    isWait(){
        return true;
    }

}

export abstract class MonitorRequest<T> extends Monitor<T> {

    handler: undefined | ((request: Request | undefined, response: Response | undefined) => Promise<{} | undefined>) = undefined;

    setHandler(handler : (request: Request | undefined, response: Response | undefined) => Promise<{}|undefined>){
        this.handler = handler;
    }
}

  
export abstract class MonitorResponse<T> extends Monitor<T> {
    
    abstract getType(): string;

    public async getResponseData(response: Response): Promise<any>{
        const contentType = response.headers()['content-type'];
        if(contentType.includes('application/json')){
            return await response.json();
        }
        if (contentType && contentType.includes('text/html')) {
            return await response.text();
        }
        return await response.body();
    }

}

export abstract class MonitorChain<T> {

    monitors: Monitor<T>[] = [];

    data : {[key: string]: any} = {};

    constructor(){
        this.monitors = this.initMonitors();
    }

    abstract initMonitors(): Monitor<T>[];

    public getKey(){
        return this.constructor.name;
    }

    public append(monitor: Monitor<T>){
        this.monitors.push(monitor);
    }

    public async start(){
        for(const monitor of this.monitors){
            await monitor.start();
        }
    }

    abstract getType() :string;

    public getMonitors(): Monitor<T>[]{
        return this.monitors;
    }

    public async waitForAction(): Promise<DoorEntity<T>>{
        const doorData: {[key: string]: any} = {};
        let result = true;
        for(let monitor of this.monitors){
            if (!monitor.isWait()){
                continue;
            }
            const doorEntity = await monitor.waitForAction();
            if(!doorEntity.code){
                result = false;
            }
            doorData[monitor.getKey()] = doorEntity.data;
        }
        return new DoorEntity<T>(result, doorData as T);
    }

    protected getItemKey(params : URLSearchParams): string | undefined {
        return undefined;
    }

    getItemKeys(url : string): string | undefined {
        return this.getItemKey(getUrlParameter(url));
    }

}