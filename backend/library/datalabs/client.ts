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

    // SESSION_START refers to the time the session was started
    SESSION_START: number

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
        this.issueTicket(application_token)
        // set cookie in runtime to not query it all the time
        if (this.COOKIE === null || this.COOKIE === "")
            this.COOKIE = this.getCookie()

        // looks up referrer
        this.REF_PAGE = this.referrer()
        // init web-socket connection
        this.conn = this.open(this.ticket)
        // start session timer
        this.SESSION_START = new Date().getTime()
    }

    // Listen takes care of the events triggered by the user.
    // It manages the events received by the client and the events
    // send to the server
    public Listen() {}

    // issueTicket requests a ticket to connect to the web-socket
    // returning the permission on what data is allowed to collect
    // (event listener, etc.)
    private issueTicket(token) {
        let options: any = {
            credentials: "same-origin",
            headers: {
                "x-datalabs-token": token,
            }
        }
        fetch(this.API_URL, options).then(resp => {return resp.json()}).then(data => {
            this.meta = data.meta
            this.ticket = data.ticket
        }).catch(err => {
            console.log(err)
        })
    }

    // open opens the web-socket connection with the access_token
    // on failure the function will try to open the connection until
    // the MAX_RETRY is exceeded
    private open(token: string): any {
        if (this.conn != null) 
            return null
        if (token === null || token === "")
            return null

        let URL: string = this.WS_URL
        try {
            // update try count
            this.TRIES++
            // open conn
            return new WebSocket(URL)
        } catch (error) {
            // exit early if tries exceeded
            if (this.TRIES > this.MAX_RETRY) 
                return null
            // try again
            return this.open(token)
        }
    }


    // attach attaches all the JavaScript EventListener to the
    // present document - all refers to the ones set in the access_token
    private attach() {

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

    // referrer looks for the web-page the current page
    // was called from. If null or empty returns null
    private referrer(): string {
        let ref = document.referrer;
        if (ref === null || ref === "")
            return null
        return ref
    }
}
