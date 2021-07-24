import platform from 'platform';
import axios from 'axios';
import $ from 'jquery';

enum EVENT {
    RAW_CLICK = 0,
    RAW_URL,
    BTN_TIME,
    FUNNEL_CHANGE,
}

enum LISTENER {
    CLICK = "click",
    HOVER = "mouseover",
}

export class DataKraken {   
    private API_WS = "ws://192.168.0.232:8004/api/v1/open?" 
    private URL_TIMEOUT_RATE: number = 1000
    private URL_TIME: number = new Date().getTime()
    private CURRENT_URL: string = history.state.current

    private LAST_CLICK: number = new Date().getTime()
    private BTN_DEFS: Array<any> = []
    private STAGES: Array<any> = []

    private CURRENT_STAGE_LABEL: string = ""
    private CURRENT_STAGE_ID: number = -1

    private WS_TICKET: string = ""

    private WEB_SOCKET: any

    // TODO: struct how the init of the class must look like when to init the session, the web socket
    // the event listener
    constructor(app_token: string) {
        this.sayHello(app_token).then(ok => {
            if (!ok)
                return

            this.WEB_SOCKET = this.open(this.WS_TICKET)
            if (this.WEB_SOCKET === null)
                return

            this.CURRENT_STAGE_LABEL = history.state.current
            this.attach(LISTENER.HOVER, this.onHover)
            this.attach(LISTENER.CLICK, this.onClick)
            this.urlListener()

        })

    }

    // sayHello initializes the client session passing basic client information to the
    // server. If a cookie is present it will get send along the request else the server
    // assigns a new cookie (also indicating that the client is new). The session start is handled
    // server-side. If the authentication succeeds the response will hole the web-socket ticket to establish 
    // the web-socket connection further, the response holds meat-data such as button-definitions.
    // If the authentication fails or the server fails respond (including re-tries) the function returns a false
    // indicating to not do anything further.
    private async sayHello(token: string): Promise<any> {
        const opts = {
            headers: {
                "x-datalab-token": token
            },
            // withCredentials: true,
        }
        const resp: any = await axios.get("http://192.168.0.232:8004/api/v1/hello", opts)
        
        if (resp.status != 200)
            return false
        this.WS_TICKET = resp?.data?.ticket
        this.STAGES = resp?.data?.meta?.stages
        this.BTN_DEFS = resp?.data?.meta?.btn_defs
        return true
    }

    private open(ticket: string): any {
        const deviceInfo = this.getDevice()
        const URL_PARAMS = "ticket="+ticket+"&ref="+this.getReferrer()+"&os_name="+deviceInfo.OS?.name+"&os_vs="+deviceInfo.OS?.version+"&device="+deviceInfo.device+"&browser="+deviceInfo.browser
        const ws = new WebSocket(this.API_WS+URL_PARAMS)
        ws.onerror = function(err: any){
            console.log(err)
        }
        ws.onmessage = function(msg: any){
            console.log(msg)
        }
        return ws
    }

    // functions for events
    // events:
    //      - element clicked         [done]
    //      - URL change              [done]
    //      - hover time over element [done]
    //      - time per URL            [done]
    //      - referrer page           [done]
    //      - device info             [done]

    // attach adds a given event and function to the root document and binding
    // the function to "this"
    private attach(event_name: string, fn: any) {
        console.log("Attaching: ", event_name)
        document.addEventListener(event_name, fn.bind(this))
    }
    // getReferrer returns the page this one was referenced by.
    // If it is an empty string it returns null
    private getReferrer(): any {
        if (document.referrer === "")
            return null
        return document.referrer
    }

    // getCampaign returns the value of the URL-Query("campaign") if not present returns null
    private getCampaign(): any {
        const url = new URL(document.location.href);
        const params = new URLSearchParams(url.search.slice(1));
        
        if (params.has("campaign")) {
            return params.get("campaign")
        }
        return null
    }

    // urlListener periodically checks if the url has changed. If so it captures the
    // prevues URL and the current URL along with the time passed in-between.
    // raw-url => data_point: any = {
    //     type: EVENT.RAW_URL,
    //     timestamp: int64,
    //     from: history.state.back,
    //     to: history.state.current,
    //     elapsed_time: elapsed,
    // }
    // funnel-change => data_point: any = {
    //     type: EVENT.FUNNEL_CHANGE,
    //     timestamp: int64,
    //     leaving: this.CURRENT_STAGE_LABEL,
    //     entered: history.state.current,
    //     elapsed_time: elapsed,
    // }
    private urlListener() {
        setInterval(()=>{
            if (this.CURRENT_URL == history.state.current) 
                return
            
            const elapsed: number = DataKraken.elapsed(new Date().getTime(), this.URL_TIME)
            this.URL_TIME = new Date().getTime()
            const data_point: any = {
                type: EVENT.RAW_URL,
                timestamp: new Date().getTime(),
                from: history.state.back,
                to: history.state.current,
                elapsed_time: elapsed,
            }

            const id_if_stage: number = this.isStageRelevant(1, null)
            if (id_if_stage !== -1) {
                console.log("!!!Stage-Change!!!")
                const data_point: any = {
                    type: EVENT.FUNNEL_CHANGE,
                    timestamp: new Date().getTime(),
                    from_stage_label: this.CURRENT_STAGE_LABEL,
                    from_stage: "" + this.CURRENT_STAGE_ID + "",
                    to_stage_label: history.state.current,
                    to_stage: "" + id_if_stage + "",
                    elapsed_time: elapsed,
                }
                this.WEB_SOCKET.send(JSON.stringify(data_point))
                this.CURRENT_STAGE_LABEL = history.state.current
                this.CURRENT_STAGE_ID = id_if_stage
            }
                
            // create url-change event

            this.WEB_SOCKET.send(JSON.stringify(data_point))
            this.CURRENT_URL = history.state.current
        }, this.URL_TIMEOUT_RATE)
    }

    // onClick captures any click event
    // raw-click => data_point: any = {
    //     type: EVENT.RAW_CLICK,
    //     timestamp: int64,
    //     target: target,
    //     elapsed_time: elapsed,
    //     current_url: URL,
    // }
    // funnel change => data_point: any = {
    //     type: EVENT.FUNNEL_CHANGE,
    //     timestamp: int64,
    //     leaving: this.CURRENT_STAGE_LABEL,
    //     entered: target,
    //     elapsed_time: elapsed,
    // }
    private onClick(event: any) {
        const target: string = this.buildXPath(event.srcElement)
        if (target === undefined || target === "") {
            return
        }
        const elapsed: number = DataKraken.elapsed(new Date().getTime(), this.LAST_CLICK)
        const URL: string = history.state.current
        const data_point: any = {
            type: EVENT.RAW_CLICK,
            timestamp: new Date().getTime(),
            target: target,
            elapsed_time: elapsed,
            current_url: URL,
        }
        // send funnel change
        const id_if_stage: number = this.isStageRelevant(2, event)
        if (id_if_stage !== -1) {
            console.log("!!!Stage-Change!!!")
            const data_point: any = {
                type: EVENT.FUNNEL_CHANGE,
                timestamp: new Date().getTime(),
                from_stage_label: this.CURRENT_STAGE_LABEL,
                from_stage: "" + this.CURRENT_STAGE_ID + "",
                to_stage_label: target,
                to_stage: "" + id_if_stage + "",
                elapsed_time: elapsed,
            }
            this.WEB_SOCKET.send(JSON.stringify(data_point))
            this.CURRENT_STAGE_LABEL = target
            this.CURRENT_STAGE_ID = id_if_stage
        }
            
        // create stage change event

        console.log("Element-Clicked: ", data_point, event)
        this.WEB_SOCKET.send(JSON.stringify(data_point))
        this.LAST_CLICK = new Date().getTime()
    }

    // onHover tracks the time a user hovers of a specified element (set in config in datalab app)
    // it attaches a follow-up event (onClick and onLeave) to denote the results of the user action
    // data_point: any = {
    //     type: EVENT.BTN_TIME,
    //     timestamp: int64,
    //     target: xpath,
    //     action: "hover-leave" || "hover-click",
    //     elapsed_time: elapsed,
    // }
    private onHover(event: any) {
        // lookup if target is listed as watcher
        const xpath: string = this.buildXPath(event.srcElement)
        let match: boolean = false
        for (let i = 0; i < this.BTN_DEFS.length; i++) {   
            if (this.BTN_DEFS[i]?.name === xpath) {
                match = true
            }     
        }
        if (!match)
            return
        
        const event_start: number = new Date().getTime()
        // only one follow-up event must be satisfied. After the "click" event
        // the "mouseleave" event must be ignored and vice-versa
        let taken: boolean = false
        // attach follow-up events
        event.target.addEventListener("click", (evt: any) => {
            // TODO: what does a click mean in data flow language
            if (taken)
                return
            taken =  true
            const elapsed: number = DataKraken.elapsed(new Date().getTime(), event_start)
            // ignore noise events
            if (elapsed <= 0)
                return
            const data_point: any = {
                type: EVENT.BTN_TIME,
                timestamp: new Date().getTime(),
                target: xpath,
                action: "hover-click",
                elapsed_time: elapsed,
            }
            console.log("hover-then-clicked: ", data_point)
            this.WEB_SOCKET.send(JSON.stringify(data_point))
        })

        event.target.addEventListener("mouseleave", (evt: any) => {
            // TODO: what does a leave mean in data flow language
            if (taken)
                return
            taken =  true
            const elapsed: number = DataKraken.elapsed(new Date().getTime(), event_start)
            // ignore noise events
            if (elapsed <= 0)
                return
            const data_point: any = {
                type: EVENT.BTN_TIME,
                timestamp: new Date().getTime(),
                target: xpath,
                action: "hover-leave",
                elapsed_time: elapsed,
            }
            console.log("hover-then-left: ", data_point)
            this.WEB_SOCKET.send(JSON.stringify(data_point))
        })
    } 

    // isStageRelevant checks if an event matches the stage critieria
    private isStageRelevant(type: number, evt: any): number {
        for (let i = 0; i < this.STAGES.length; i++) {
            if (this.STAGES[i]?.type === type && type === 1) { // match url pattern
                const url: string = history.state.current
        
                if (url === this.STAGES[i]?.transition && !this.STAGES[i]?.regex) {
                    return this.STAGES[i]?.id
                }
                if (url.search(this.STAGES[i]?.transition) == 0 && this.STAGES[i]?.regex) {
                    // match partial url with  regex
                    const match: boolean = this.regexMatch(url, this.STAGES[i]?.transition, this.STAGES[i]?.regex)
                    if (match) {
                        return this.STAGES[i]?.id
                    }
                    return -1
                }
            }
            if (this.STAGES[i]?.type === type && type === 2) { // element xpath match
                const xpath: string = this.buildXPath(evt?.srcElement)
                if (this.STAGES[i]?.transition !== xpath)
                    continue
                return this.STAGES[i]?.id
            }
        }
        return -1
    }

    private regexMatch(str: string, stage_url: string, regex: string): boolean {
        try {
            let re = new RegExp(stage_url+regex)
            const res: any = re.exec(str)
            if (res === null || res?.length === 0) {
                return false
            }
        } catch(err) {
            return false
        }
        return true
    }
    // getDevice captures the device information of the user
    // if device not mobile device will be "laptop/PC"
    // data-point: {
    //     browser,
    //     OS: {name, version},
    //     device
    // }
    private getDevice() {
        const browser: any = platform?.name
        const OS: any = {name: platform?.os?.family, version: platform?.os?.version}
        const device: any = platform.product === null ? "laptop/PC" : platform.product
        return {
            browser: browser,
            OS: OS,
            device: device,
        }
    }


    private  buildXPath(element: any) {
        let xpath = '';
        for ( ; element && element.nodeType == 1; element = element.parentNode )
        {
            let id:any = $(element.parentNode).children(element.tagName).index(element) + 1;
            id > 1 ? (id = '[' + id + ']') : (id = '');
            xpath = '/' + element.tagName.toLowerCase() + id + xpath;
        }
        return xpath;
    }

    // Event builds the event as it will be send to the web-socket
    private static Event(type: number, data: any): any {
        return {
            type: type,
            timestamp: new Date().getTime(),
            event: data,
        }
    }
    // elapsed computed the time difference between to UNIX time-stamps
    // difference in seconds
    private static elapsed(date_1: number, date_2: number): number {
        const tmp: number = Math.floor((date_1 - date_2))
        if (date_1 === 0 || date_2 === 0)
            return 0
        if (tmp <= 0) 
            return 0
        return Math.floor(tmp / 1000)
    }
}