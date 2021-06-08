const platform = require('platform');
const axios = require('axios');
enum EVENT {
    HOVER_THEN_CLICK = 0,
    HOVER_THEN_LEFT,
    URL_CHANGE,
    ELEMENT_CLICK,
    STAGE_CHANGE,
}

enum STAGE {
    DEFAULT = 0
}

export class DataKraken {    
    private URL_RATE: number = 1000
    private URL_TIME: number = new Date().getTime()
    private CURRENT_URL: string = history.state.current
    private URL_WATCH_LIST: Array<any> = []

    private LAST_CLICK: number = new Date().getTime()
    private BTN_DEFS: Array<string> =  []

    private WS_TICKET: string = ""

    private STAGE_DEFS: Array<any> = []
    private CURRENT_STAGE: STAGE = STAGE.DEFAULT
    private STAGE_TIME: number = 0

    // TODO: struct how the init of the class must look like when to init the session, the web socket
    // the event listener
    constructor(app_token: string) {
        // start counting how long user is in given stage
        // of funnel
        this.STAGE_TIME = new Date().getTime() 

        // initialize session, load meta data
        this.sayHello(app_token).then(data => {
            if (data === undefined || data === null)
                return
            for (let i = 0; i < data?.btn_defs?.length; i++) {
                this.BTN_DEFS.push(data?.btn_defs[i]?.btn_name)
            }
            this.STAGE_DEFS = data?.stages
            this.WS_TICKET = data?.ticket
            
             // apply meta-data to application
            this.attach("click", this.onClick)
            this.attach("mouseover", this.onHover)
            this.urlListener()
            for (let i = 0; i < this.STAGE_DEFS.length; i++) {
                this.attachWithStage(this.STAGE_DEFS[i].transition, "click", this.onStageEnterOnClick, this.STAGE_DEFS[i])
            }
        }
        );
    }

    // sayHello initializes the client session passing basic client information to the
    // server. If a cookie is present it will get send along the request else the server
    // assigns a new cookie (also indicating that the client is new). The session start is handled
    // server-side. If the authentication succeeds the response will hole the web-socket ticket to establish 
    // the web-socket connection further, the response holds meat-data such as button-definitions.
    // If the authentication fails or the server fails respond (including re-tries) the function returns a -1
    // indicating to not do anything further.
    private async sayHello(token: string): Promise<any> {
        // let success: boolean = true

        const opts = {
            headers: {
                "x-datalab-token": token
            }
        }
        const resp = await axios.post("http://localhost:8004/api/v1/hello", {
            referrer: this.getReferrer(),
            meta: this.getDevice(),
        }, opts)
        return resp?.data
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
        console.log("attaching event: ", event_name)
        document.addEventListener(event_name, fn.bind(this))
    }

    // attachWithStage attaches an event listener to a specific DOM element.   
    // it queries docuement.getElementByName() and selectec the first found instance
    // Thus, the "element" paramamter should be unique
    private attachWithStage(element_name: string, event_name: string, fn: any, stage: any) {
        let node: any = null
        switch(stage.trigger) {
            case 2: // react to onclick events
                console.log("attaching onclick-event with args: ", event_name, element_name)
                node = document.getElementById(element_name)
                node.args = stage; 
                node.addEventListener(event_name, fn.bind(this))
                break;
            case 1: // reacte to url events
                // once in the list the url_listener will take care about stage changes
                this.URL_WATCH_LIST.push(stage)
                break;
            default:
                return
        }
        
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
    // data-point: {
    //     from,
    //     to,
    //     elapsed
    // }
    private urlListener() {
        setInterval(()=>{
            if (this.CURRENT_URL == history.state.current) 
                return
            
             const elapsed: number = DataKraken.elapsed(new Date().getTime(), this.URL_TIME)
             this.URL_TIME = new Date().getTime()
             const data_point: any = DataKraken.Event(
                 EVENT.URL_CHANGE,
                 {
                     from: history.state.back,
                     to: history.state.current,
                     elapsed: elapsed,
                 })
            console.log("URL-CHANGE: ", data_point)
            
            for (let i = 0; i < this.URL_WATCH_LIST.length; i++) {
                if (this.URL_WATCH_LIST[i].transition === history.state.current) {
                    const data_point: any = DataKraken.Event(
                        EVENT.STAGE_CHANGE,
                        {
                            stage_name: this.URL_WATCH_LIST[i].name,
                            stage_id: this.URL_WATCH_LIST[i].id,
                            trigger: history.state.current,
                            previouse_stage: this.CURRENT_STAGE,
                            entered_stage: this.CURRENT_STAGE++,
                            elapsed: DataKraken.elapsed(new Date().getTime(), this.STAGE_TIME),  
                        }
                    )
                    console.log("New Stage (on-URL): ",data_point)
                    this.STAGE_TIME = new Date().getTime()
                }
            }
        this.CURRENT_URL = history.state.current
        }, this.URL_RATE)
    }

    // onClick captures any click event
    // data-point: {
    //     url,
    //     target,
    //     elapsed
    // }
    private onClick(event: any) {
        const target: string = event.target.name
        if (target === undefined || target === "")
            return
        const elapsed: number = DataKraken.elapsed(new Date().getTime(), this.LAST_CLICK)
        const URL: string = history.state.current
        const data_point: any = DataKraken.Event(
            EVENT.ELEMENT_CLICK,
            {
                url: URL,
                target: target,
                elapsed: elapsed,
            })
        console.log("Clicked: ", data_point)
        this.LAST_CLICK = new Date().getTime()
    }

    // onHover tracks the time a user hovers of a specified element (set in config in datalab app)
    // it attaches a follow-up event (onClick and onLeave) to denote the results of the user action
    // data-point: {
    //     target,
    //     elapsed
    // }
    private onHover(event: any) {
        // lookup if target is listed as watcher
        if (!this.BTN_DEFS.includes(event.target.name))
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
            const target: string = evt.target.name
            const data_point: any = DataKraken.Event(
                EVENT.HOVER_THEN_CLICK,
                {
                    target: target,
                    elapsed: elapsed
                })
            console.log("clicked: ", data_point)
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
            const target: string = evt.target.name
            const data_point: any = DataKraken.Event(
                EVENT.HOVER_THEN_LEFT, 
                {
                    target: target,
                    elapsed: elapsed
                })
            console.log("left: ", data_point)
        })
    } 

    // onStageEnterOnClick handels any jumps of users jumping from one funnel-stage to another
    // data-point: {
    //     trigger, -> what caused the stage change 
    //     previouse_stage, -> where did the user came from
    //     entered_stage, -> which stage did the user entered
    //     elapsed -> how long did the user stay on the stage before
    // }
    private onStageEnterOnClick(event: any) {
        const data_point: any = DataKraken.Event(
            EVENT.STAGE_CHANGE,
            {   
                stage_name: event?.target?.args.name,
                stage_id: event?.target?.args.id,
                trigger: event?.target?.name,
                previouse_stage: this.CURRENT_STAGE,
                entered_stage: event?.target?.args.id,
                elapsed: DataKraken.elapsed(new Date().getTime(), this.STAGE_TIME),
            }
        )
        console.log("New Stage (on-click): ",data_point)
        this.STAGE_TIME = new Date().getTime()
        this.CURRENT_STAGE = event?.target?.args.id
    }


    // getDevice captures the device information of the user
    // if device not mobile device will be "laptop/PC"
    // data-point: {
    //     browser,
    //     OS: {name, version},
    //     device
    // }
    private getDevice() {
        const browser: string = platform.name
        const OS: any = {name: platform.os.family, version: platform.os.version}
        const device: string = platform.product === null ? "laptop/PC" : platform.product
        return {
            browser: browser,
            OS: OS,
            device: device,
        }
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