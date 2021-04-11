export class DataKraken {
    // API_URL
    API_URL: string = "http://localhost:8004/live/whoami"

    // WS_URL refers to the web-socket endpoint to call
    WS_URL: string = "ws://localhost:8004/live/websocket"

    // INIT_PARAM refers to whether a cookie must be set by the server
    INIT_PARAM: string = "?init="

    // QUERY_PARAM refers to the access_token which is required
    // to connect to the socket
    TOKEN_PARAM: string = "?token="

    // MAX_RETRY refers to the amount of trys to open the socket
    MAX_RETRY: number = 5

    // TRIES refers to the current number of failed tries to
    // open the web-socket connection
    TRIES: number = 0

    // REF_PAGE refers to the web-page the current page was loaded from
    // example: 
    //          www.google.com  -> www.my-site.com
    //          www.youtube.com -> www.my-site.com
    REF_PAGE: string 

    // COOKIE_NAME refers to the key storing the cookie value
    COOKIE_NAME: string = "datalabs.identity"

    // COOKIE refers to the cookie set by the server
    COOKIE: string = null

    // conn refers to the open web-socket connection
    conn: any = null

    // meta refers to the meta data like which events to listen on
    meta: any = {}

    // ticket is the issued token by the web-socket server granting
    // access to connect to the socket
    ticket: string
    

    constructor(application_token: string) {
        // request access token for web-socket
        this.issueTicket(application_token, this.hasCookie()).then(resp => {
            this.meta = resp.meta
            this.ticket = resp.ticket
        })
        // looks up referrer
        this.REF_PAGE = this.referrer()
        // init web-socket connection
        this.open(this.ticket)
    }

    // Listen takes care of the events triggered by the user.
    // It manages the events received by the client and the events
    // send to the server
    public Listen() {}

    // issueTicket requests a ticket to connect to the web-socket
    // returning the permission on what data is allowed to collect
    // (event listener, etc.)
    private async issueTicket(token, isInit): Promise<any> {
        let options: any = {
            headers: {
                "x-datalabs-token": token,
            }
        }
        let URL = this.API_URL + this.INIT_PARAM + isInit

        const resp = await fetch(URL, options)
        if (resp.status != 200)
           // do something if fails
        
        return resp.json()
    }

    // open opens the web-socket connection with the access_token
    // on failure the function will try to open the connection until
    // the MAX_RETRY is exceeded
    private open(token: string) {
        if (this.conn != null) 
            return

        let URL: string = this.WS_URL
        try {
            // update try count
            this.TRIES++
            // open conn
            this.conn = new WebSocket(URL)
        } catch (error) {
            // exit early if tries exceeded
            if (this.TRIES > this.MAX_RETRY) 
                return
            // try again
            this.open(token)
        }
    }

    // TODO: logic to return correct cookie ;D
    private hasCookie(): string {
        let cookies = document.cookie.split(";")
        return cookies[0]
    }
    // attach attaches all the JavaScript EventListener to the
    // present document - all refers to the ones set in the access_token
    private attach() {

    }

    // referrer looks for the web-page the current page
    // was called from. If null or empty returns null
    private referrer(): string {
        let ref = document.referrer;
        if (ref === null || ref === "")
            return null
        return ref
    }
}
