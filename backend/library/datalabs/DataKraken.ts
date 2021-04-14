enum SOCKET_STATE {
    READY = 0,
    IDLE,
    COLSED,
    FAILED,
}

enum EVENT_TYPE {
    CLICK = 0,
    MOUSEMOVE,
}

export class DataKraken {

    // CONN_STATE refers to the state of the web-socket connection
    private CONN_STATE: SOCKET_STATE = SOCKET_STATE.IDLE

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

    // PARAM_OS refers to the OS the user is running on
    private PARAM_OS: string = "os="

    // PARAM_BROWSER refers to the browser the client is using
    private PARAM_BROWSER: string = "browser="


    // COOKIE_NAME refers to the key storing the cooking value
    private COOKIE_NAME: string = "datalabs.identity"

    // TYPE_MOUSEMOVE refers to the event type "mousemove"
    private TYPE_MOUSEMOVE: EVENT_TYPE = EVENT_TYPE.MOUSEMOVE

    // TYPE_CLICK refers to the event type "click"
    private TYPE_CLICK: EVENT_TYPE = EVENT_TYPE.MOUSEMOVE

    // session refers to the current session data. Holding meta data and
    // allowed events and settings
    private session: any

    // device refers to the device data of the client
    private device: any

    // conn refers to the web-socket connection
    private conn: any = null

    constructor(app_token: string) {
        this.init(app_token)
    }

    // init takes care of acquiring the web-socket ticket and the initialization
    // of the meta data object
    private init(app_token: string): void {
        this.device = this.device_info().init()

        const data: any = this.initSession(app_token)

        this.session = data.session

        // attach all events as allowed by session permissions
        this.session.events.forEach(evt => {
            switch (evt) {
                case this.TYPE_CLICK:
                    this.attach("click", this.onMouseClick)
                    break
                case this.TYPE_MOUSEMOVE:
                    this.attach("mousemove", this.onMouseMove)
                    break
                default:
                    break
            }
        });


    }

    // Listen start the web-socket and handles onMessage, onClose events
    public Listen() {
        this.conn = this.open(this.session.ticket)
        this.conn.onopen = this.onOpen
        this.conn.onclose = this.onClose
        this.conn.onerror = this.onError
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
        if (this.CONN_STATE !== SOCKET_STATE.READY)
            return this.conn

        let queryTicket: string = this.PARAM_TICKET + ticket
        let queryReferrer: string = this.PARAM_REF + this.referrer()
        let queryCookie: string = this.PARAM_COOKIE + this.session.cookie
        let queryDevice: string = this.PARAM_BROWSER + this.device.browser.name + this.PARAM_OS + this.device.os.name

        let URL: string = this.WS_URL + "?" + queryTicket + "&" + queryReferrer + "&" + queryCookie + "&" + queryDevice
        var conn: any = new WebSocket(URL)

        return conn
    }

    // attach adds a new JavaScript EventListener to the current document
    // along with the passed function which will be executed on event received event
    private attach(evt_type: string, fn: (evt: any) => void) {
        document.addEventListener(evt_type, fn)
    }

        // ***** Functions for the EventListener *****

    // onMouseMove handles mouse-events like the position in (X,Y)
    // and sends them as a defined event to the web-socket server
    private onMouseMove(evt: any) {
        if (!this.isAlive())
            return

        let X: number = evt.clientX
        let Y: number = evt.clientY
        let point: string = this.KrakenEvent(this.TYPE_MOUSEMOVE, { "X": X, "Y": Y })
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
        if (!this.isAlive())
            return

        let clickX: number = evt.clientX
        let clickY: number = evt.clientY
        let targetName: string = evt.target.name
        let point: string = this.KrakenEvent(this.TYPE_CLICK, { "X": clickX, "Y": clickY, "target": targetName })
        try {
            // send over the wire 
            this.conn.send(point)
        } catch (error) {
            return
        }
    }

    // KrakenEvent transforms the data to a JSON-string representation
    private KrakenEvent(type: number, event: any): string {
        return JSON.stringify({
            "session": this.session.cookie,
            "timestamp": new Date().getTime(),
            "type": type,
            "event": event,
        })
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
        var cookiestring = RegExp(this.COOKIE_NAME + "=[^;]+").exec(document.cookie);
        // Return everything after the equal sign, or an empty string if the cookie name not found
        return decodeURIComponent(!!cookiestring ? cookiestring.toString().replace(/^[^=]+./, "") : "");
    }

    // hasCookie returns true if a cookie is already set
    private hasCookie(): boolean {
        var c = this.getCookie()
        if (c === null || c === "")
            return false
        return true
    }

    // **** Functions for WebSocket Events ****
    private onOpen(evt: any) {
        this.CONN_STATE = SOCKET_STATE.READY
    }
    private onClose(evt: any) {
        this.CONN_STATE = SOCKET_STATE.COLSED
    }
    private onError(evt: any) {
        this.CONN_STATE = SOCKET_STATE.FAILED
    }

    // isAlive tells if the socket is ready to accept events
    private isAlive(): boolean {
        return this.CONN_STATE === SOCKET_STATE.READY ? true : false
    }

    // device_info returns the all device info of the client that can be found
    private device_info(): any {
        var module = {
            options: [],
            header: [navigator.platform, navigator.userAgent, navigator.appVersion, navigator.vendor],
            dataos: [
                { name: 'Windows Phone', value: 'Windows Phone', version: 'OS' },
                { name: 'Windows', value: 'Win', version: 'NT' },
                { name: 'iPhone', value: 'iPhone', version: 'OS' },
                { name: 'iPad', value: 'iPad', version: 'OS' },
                { name: 'Kindle', value: 'Silk', version: 'Silk' },
                { name: 'Android', value: 'Android', version: 'Android' },
                { name: 'PlayBook', value: 'PlayBook', version: 'OS' },
                { name: 'BlackBerry', value: 'BlackBerry', version: '/' },
                { name: 'Macintosh', value: 'Mac', version: 'OS X' },
                { name: 'Linux', value: 'Linux', version: 'rv' },
                { name: 'Palm', value: 'Palm', version: 'PalmOS' }
            ],
            databrowser: [
                { name: 'Chrome', value: 'Chrome', version: 'Chrome' },
                { name: 'Firefox', value: 'Firefox', version: 'Firefox' },
                { name: 'Safari', value: 'Safari', version: 'Version' },
                { name: 'Internet Explorer', value: 'MSIE', version: 'MSIE' },
                { name: 'Opera', value: 'Opera', version: 'Opera' },
                { name: 'BlackBerry', value: 'CLDC', version: 'CLDC' },
                { name: 'Mozilla', value: 'Mozilla', version: 'Mozilla' }
            ],
            init: function () {
                var agent = this.header.join(' '),
                    os = this.matchItem(agent, this.dataos),
                    browser = this.matchItem(agent, this.databrowser);

                return { os: os, browser: browser };
            },
            matchItem: function (string, data) {
                var i = 0,
                    j = 0,
                    html = '',
                    regex,
                    regexv,
                    match,
                    matches,
                    version;

                for (i = 0; i < data.length; i += 1) {
                    regex = new RegExp(data[i].value, 'i');
                    match = regex.test(string);
                    if (match) {
                        regexv = new RegExp(data[i].version + '[- /:;]([\\d._]+)', 'i');
                        matches = string.match(regexv);
                        version = '';
                        if (matches) { if (matches[1]) { matches = matches[1]; } }
                        if (matches) {
                            matches = matches.split(/[._]+/);
                            for (j = 0; j < matches.length; j += 1) {
                                if (j === 0) {
                                    version += matches[j] + '.';
                                } else {
                                    version += matches[j];
                                }
                            }
                        } else {
                            version = '0';
                        }
                        return {
                            name: data[i].name,
                            version: parseFloat(version)
                        };
                    }
                }
                return { name: 'unknown', version: 0 };
            }
        }
        return module
    }

}

