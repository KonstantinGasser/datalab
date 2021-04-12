export default class DataKraken {

    // API_WHOAMI refers to the API verifying that the client is
    // authenticated and authorized to proceed
    private API_WHOAMI: string = "http://localhost:8004/live/whoami"
    
    // WS_URL refers to the API establishing the web-socket connection
    private WS_URL: string = "http://localhost:8004/live/websocket"
    
    // PARAM_TICKET refers to the WS-ticket required to establish a connection
    private PARAM_TICKET: string = "ticket="
    
    // PARAM_REF is part of the initial data for a user-record,
    // refers to the page this page was called from
    private PARAM_REF: string = "ref="

    // PARAM_COOKIE refers to the client-side cookie
    private PARAM_COOKIE: string = "cookie="

    // COOKIE_NAME refers to the key storing the cooking value
    private COOKIE_NAME: string = "datalabs.identity"

    // TYPE_MOUSEMOVE refers to the event type "mousemove"
    private TYPE_MOUSEMOVE: string = "event-mouse"

    // TYPE_CLICK refers to the event type "click"
    private TYPE_CLICK: string = "event-click"

    // session refers to the current session data. Holding meta data and
    // allowed events and settings
    private session: any

    // conn refers to the web-socket connection
    private conn: any = null
    
    constructor(app_token: string) {
        this.init(app_token)
    }

    // init takes care of acquiring the web-socket ticket and the initialization
    // of the meta data object
    private init(app_token: string): void {
        const data: any = this.initSession(app_token)
        
        this.session = data.session
        
        this.conn = this.open(data.ticket)
    }


    // initSession calls the ticket API to get a new ticket.
    // if present it sends the stored cookie - else the server will set
    // a new cookie for the client
    private initSession(app_token: string): any {
        let opts: any = {
            // will set cookie if present
            credentials: "same-origin",
            headers: {
                "x-datalabs-token": app_token,
            },
        }
        fetch(this.API_WHOAMI, opts).then(resp => resp.json()).then(data => {
            return data
        }).catch(err => {
            return null
        })
    }

    // open establishes the web-socket connection
    private open(ticket: string): any {
        if (this.conn === null)
            return this.conn
        
        let queryTicket: string = this.PARAM_TICKET + ticket
        let queryReferrer: string = this.PARAM_REF + this.referrer()
        let queryCookie: string = this.PARAM_COOKIE + this.session.cookie

        let URL: string = this.WS_URL + "?" + queryTicket + "&" + queryReferrer + "&" + queryCookie
        var conn: any = new WebSocket(URL)
        return conn
    }

    // attach adds a new JavaScript EventListener to the current document
    // along with the passed function which will be executed on event received event
    private attach(evt_type: string, fn: () => void) {
        document.addEventListener(evt_type, fn)
    }

    // referrer looks for the web-page the current page
    // was called from. If null or empty returns ""
    private referrer(): string {
        let ref = document.referrer;
        if (ref === null || ref === "")
            return ""
        return ref
    }

    // getCookie returns the stored cookie of the user or null
    private getCookie(): string {
        // Get name followed by anything except a semicolon
        var cookiestring=RegExp(this.COOKIE_NAME+"=[^;]+").exec(document.cookie);
        // Return everything after the equal sign, or an empty string if the cookie name not found
        return decodeURIComponent(!!cookiestring ? cookiestring.toString().replace(/^[^=]+./,"") : "");
    }

    // hasCookie returns true if a cookie is already set
    private hasCookie(): boolean {
        var c = this.getCookie()
        if ( c === null || c === "")
            return false
        return true
    }

    // ***** Functions for the EventListener *****

    // onMouseMove handles mouse-events like the position in (X,Y)
    // and sends them as a defined event to the web-socket server
    private onMouseMove(evt: any) {
        let X: number = evt.clientX
        let Y: number = evt.clientY
        let point: string = this.KrakenEvent(this.TYPE_MOUSEMOVE, {"X":X, "Y": Y})
        try {
            // send over the wire 
            this.conn.send(point)
        } catch (error) {
            return
        }
    }
    // onMouseClick handles mouse-click event: sampling the target-name 
    // along with its X,Y position
    private onMouseClick(evt: any) {
        let clickX: number = evt.clientX
        let clickY: number = evt.clientY
        let targetName: string = evt.target.name
        let point: string = this.KrakenEvent(this.TYPE_CLICK, {"X": clickX, "Y": clickY, "target": targetName})
        try {
            // send over the wire 
            this.conn.send(point)
        } catch (error) {
            return
        }
    }

    // KrakenEvent transforms the data to a JSON-string representation
    private KrakenEvent(type: string, event: any): string {
        return JSON.stringify({
            "session": this.session.cookie,
            "timestamp": new Date().getTime(),
            "type": type,
            "event": event,
        })
}
}

