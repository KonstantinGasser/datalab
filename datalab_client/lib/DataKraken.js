"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (_) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataKraken = void 0;
var platform_1 = __importDefault(require("platform"));
var axios_1 = __importDefault(require("axios"));
var jquery_1 = __importDefault(require("jquery"));
var EVENT;
(function (EVENT) {
    EVENT[EVENT["HOVER_THEN_CLICK"] = 0] = "HOVER_THEN_CLICK";
    EVENT[EVENT["HOVER_THEN_LEFT"] = 1] = "HOVER_THEN_LEFT";
    EVENT[EVENT["URL_CHANGE"] = 2] = "URL_CHANGE";
    EVENT[EVENT["ELEMENT_CLICK"] = 3] = "ELEMENT_CLICK";
})(EVENT || (EVENT = {}));
var LISTENER;
(function (LISTENER) {
    LISTENER["CLICK"] = "click";
    LISTENER["HOVER"] = "mouseover";
})(LISTENER || (LISTENER = {}));
var DataKraken = /** @class */ (function () {
    // TODO: struct how the init of the class must look like when to init the session, the web socket
    // the event listener
    function DataKraken(app_token) {
        var _this = this;
        this.API_WS = "ws://localhost:8004/api/v1/open?";
        this.URL_TIMEOUT_RATE = 1000;
        this.URL_TIME = new Date().getTime();
        this.CURRENT_URL = history.state.current;
        this.LAST_CLICK = new Date().getTime();
        this.BTN_DEFS = [];
        this.STAGES = [];
        this.WS_TICKET = "";
        this.sayHello(app_token).then(function (ok) {
            if (!ok)
                return;
            _this.WEB_SOCKET = _this.open(_this.WS_TICKET);
            if (_this.WEB_SOCKET === null)
                return;
            _this.attach(LISTENER.HOVER, _this.onHover);
            _this.attach(LISTENER.CLICK, _this.onClick);
            _this.urlListener();
        });
    }
    // sayHello initializes the client session passing basic client information to the
    // server. If a cookie is present it will get send along the request else the server
    // assigns a new cookie (also indicating that the client is new). The session start is handled
    // server-side. If the authentication succeeds the response will hole the web-socket ticket to establish 
    // the web-socket connection further, the response holds meat-data such as button-definitions.
    // If the authentication fails or the server fails respond (including re-tries) the function returns a false
    // indicating to not do anything further.
    DataKraken.prototype.sayHello = function (token) {
        var _a, _b, _c, _d, _e;
        return __awaiter(this, void 0, void 0, function () {
            var opts, resp;
            return __generator(this, function (_f) {
                switch (_f.label) {
                    case 0:
                        opts = {
                            headers: {
                                "x-datalab-token": token
                            },
                            // withCredentials: true,
                        };
                        return [4 /*yield*/, axios_1.default.get("http://localhost:8004/api/v1/hello", opts)];
                    case 1:
                        resp = _f.sent();
                        if (resp.status != 200)
                            return [2 /*return*/, false];
                        this.WS_TICKET = (_a = resp === null || resp === void 0 ? void 0 : resp.data) === null || _a === void 0 ? void 0 : _a.ticket;
                        this.STAGES = (_c = (_b = resp === null || resp === void 0 ? void 0 : resp.data) === null || _b === void 0 ? void 0 : _b.meta) === null || _c === void 0 ? void 0 : _c.stages;
                        this.BTN_DEFS = (_e = (_d = resp === null || resp === void 0 ? void 0 : resp.data) === null || _d === void 0 ? void 0 : _d.meta) === null || _e === void 0 ? void 0 : _e.btn_defs;
                        return [2 /*return*/, true];
                }
            });
        });
    };
    DataKraken.prototype.open = function (ticket) {
        var _a, _b;
        var deviceInfo = this.getDevice();
        var URL_PARAMS = "ticket=" + ticket + "&ref=" + this.getReferrer() + "&os_name=" + ((_a = deviceInfo.OS) === null || _a === void 0 ? void 0 : _a.name) + "&os_vs=" + ((_b = deviceInfo.OS) === null || _b === void 0 ? void 0 : _b.version) + "&device=" + deviceInfo.device + "&browser=" + deviceInfo.browser;
        var ws = new WebSocket(this.API_WS + URL_PARAMS);
        ws.onerror = function (err) {
            console.log(err);
        };
        ws.onmessage = function (msg) {
            console.log(msg);
        };
        return ws;
    };
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
    DataKraken.prototype.attach = function (event_name, fn) {
        console.log("Attaching: ", event_name);
        document.addEventListener(event_name, fn.bind(this));
    };
    // getReferrer returns the page this one was referenced by.
    // If it is an empty string it returns null
    DataKraken.prototype.getReferrer = function () {
        if (document.referrer === "")
            return null;
        return document.referrer;
    };
    // getCampaign returns the value of the URL-Query("campaign") if not present returns null
    DataKraken.prototype.getCampaign = function () {
        var url = new URL(document.location.href);
        var params = new URLSearchParams(url.search.slice(1));
        if (params.has("campaign")) {
            return params.get("campaign");
        }
        return null;
    };
    // urlListener periodically checks if the url has changed. If so it captures the
    // prevues URL and the current URL along with the time passed in-between.
    // {   
    //     "type": "int", // indicates what type of event
    //     "timestamp": "int64", 
    //     "from": "string", // URL jumped from
    //     "to": "string", // URL jumped to
    //     "elapsed_time": "int64", // passed time on "from" URL
    // }
    // TODO: check fort stage change -> including regex if found
    DataKraken.prototype.urlListener = function () {
        var _this = this;
        setInterval(function () {
            if (_this.CURRENT_URL == history.state.current)
                return;
            var elapsed = DataKraken.elapsed(new Date().getTime(), _this.URL_TIME);
            _this.URL_TIME = new Date().getTime();
            var data_point = {
                type: 1,
                timestamp: new Date().getTime(),
                from: history.state.back,
                to: history.state.current,
                elapsed_time: elapsed,
            };
            console.log(data_point);
            var isStage = _this.isStageRelevant(1, null);
            console.log("URL-CHANGE: ", isStage);
            _this.WEB_SOCKET.send(JSON.stringify(data_point));
            _this.CURRENT_URL = history.state.current;
        }, this.URL_TIMEOUT_RATE);
    };
    // onClick captures any click event
    // {   
    //     "type": "int", // indicates what type of event
    //     "timestamp": "int64", 
    //     "target": "string", // clicked HTML element - if given HTML-Name-Tag else whatever if find lol
    //     "elapsed_time": "int64", // passed time since last click
    //     "current_url": "string" // URL clicked happened
    // }
    // TODO: check for state change -> including regex if found
    DataKraken.prototype.onClick = function (event) {
        var target = this.buildXPath(event.srcElement);
        if (target === undefined || target === "") {
            console.log("Target undefined", event);
            return;
        }
        var elapsed = DataKraken.elapsed(new Date().getTime(), this.LAST_CLICK);
        var URL = history.state.current;
        var data_point = {
            type: 0,
            timestamp: new Date().getTime(),
            target: target,
            elapsed_time: elapsed,
            current_url: URL,
        };
        var isStage = this.isStageRelevant(2, event);
        console.log("CLICK-CHANGE: ", isStage);
        console.log("Clicked: ", data_point, event);
        this.WEB_SOCKET.send(JSON.stringify(data_point));
        this.LAST_CLICK = new Date().getTime();
    };
    // onHover tracks the time a user hovers of a specified element (set in config in datalab app)
    // it attaches a follow-up event (onClick and onLeave) to denote the results of the user action
    // data-point: {
    //     target,
    //     elapsed
    // }
    DataKraken.prototype.onHover = function (event) {
        var _a, _b;
        // lookup if target is listed as watcher
        var xpath = this.buildXPath(event.srcElement);
        var match = false;
        for (var i = 0; i < this.BTN_DEFS.length; i++) {
            if (((_a = this.BTN_DEFS[i]) === null || _a === void 0 ? void 0 : _a.name) === xpath) {
                console.log("want: " + xpath + " have: " + ((_b = this.BTN_DEFS[i]) === null || _b === void 0 ? void 0 : _b.name));
                match = true;
            }
        }
        if (!match)
            return;
        var event_start = new Date().getTime();
        // only one follow-up event must be satisfied. After the "click" event
        // the "mouseleave" event must be ignored and vice-versa
        var taken = false;
        // attach follow-up events
        event.target.addEventListener("click", function (evt) {
            // TODO: what does a click mean in data flow language
            if (taken)
                return;
            taken = true;
            var elapsed = DataKraken.elapsed(new Date().getTime(), event_start);
            // ignore noise events
            if (elapsed <= 0)
                return;
            var target = evt.target.name;
            var data_point = DataKraken.Event(EVENT.HOVER_THEN_CLICK, {
                target: target,
                elapsed: elapsed
            });
            console.log("clicked: ", data_point);
        });
        event.target.addEventListener("mouseleave", function (evt) {
            // TODO: what does a leave mean in data flow language
            if (taken)
                return;
            taken = true;
            var elapsed = DataKraken.elapsed(new Date().getTime(), event_start);
            // ignore noise events
            if (elapsed <= 0)
                return;
            var target = evt.target.name;
            var data_point = DataKraken.Event(EVENT.HOVER_THEN_LEFT, {
                target: target,
                elapsed: elapsed
            });
            console.log("left: ", data_point);
        });
    };
    // isStageRelevant checks if an event matches the stage critieria
    DataKraken.prototype.isStageRelevant = function (type, evt) {
        var _a, _b, _c, _d, _e, _f;
        for (var i = 0; i < this.STAGES.length; i++) {
            if (((_a = this.STAGES[i]) === null || _a === void 0 ? void 0 : _a.type) === type && type === 1) { // match url pattern
                var url = history.state.current;
                if (((_b = this.STAGES[i]) === null || _b === void 0 ? void 0 : _b.transition) === url) {
                    if ((_c = this.STAGES[i]) === null || _c === void 0 ? void 0 : _c.regex) {
                        if (!this.regexMatch(url, (_d = this.STAGES[i]) === null || _d === void 0 ? void 0 : _d.regex))
                            return false;
                        return true;
                    }
                    return true;
                }
            }
            if (((_e = this.STAGES[i]) === null || _e === void 0 ? void 0 : _e.type) === type && type === 2) { // element xpath match
                var xpath = this.buildXPath(evt === null || evt === void 0 ? void 0 : evt.srcElement);
                if (((_f = this.STAGES[i]) === null || _f === void 0 ? void 0 : _f.transition) !== xpath)
                    continue;
                return true;
            }
        }
        return false;
    };
    DataKraken.prototype.regexMatch = function (str, regex) {
        try {
            var re = new RegExp(regex);
            var res = re.exec(str);
            if ((res === null || res === void 0 ? void 0 : res.length) === 0) {
                return false;
            }
        }
        catch (err) {
            return false;
        }
        return true;
    };
    // getDevice captures the device information of the user
    // if device not mobile device will be "laptop/PC"
    // data-point: {
    //     browser,
    //     OS: {name, version},
    //     device
    // }
    DataKraken.prototype.getDevice = function () {
        var _a, _b;
        var browser = platform_1.default === null || platform_1.default === void 0 ? void 0 : platform_1.default.name;
        var OS = { name: (_a = platform_1.default === null || platform_1.default === void 0 ? void 0 : platform_1.default.os) === null || _a === void 0 ? void 0 : _a.family, version: (_b = platform_1.default === null || platform_1.default === void 0 ? void 0 : platform_1.default.os) === null || _b === void 0 ? void 0 : _b.version };
        var device = platform_1.default.product === null ? "laptop/PC" : platform_1.default.product;
        return {
            browser: browser,
            OS: OS,
            device: device,
        };
    };
    DataKraken.prototype.buildXPath = function (element) {
        var xpath = '';
        for (; element && element.nodeType == 1; element = element.parentNode) {
            var id = jquery_1.default(element.parentNode).children(element.tagName).index(element) + 1;
            id > 1 ? (id = '[' + id + ']') : (id = '');
            xpath = '/' + element.tagName.toLowerCase() + id + xpath;
        }
        return xpath;
    };
    // Event builds the event as it will be send to the web-socket
    DataKraken.Event = function (type, data) {
        return {
            type: type,
            timestamp: new Date().getTime(),
            event: data,
        };
    };
    // elapsed computed the time difference between to UNIX time-stamps
    // difference in seconds
    DataKraken.elapsed = function (date_1, date_2) {
        var tmp = Math.floor((date_1 - date_2));
        if (date_1 === 0 || date_2 === 0)
            return 0;
        if (tmp <= 0)
            return 0;
        return Math.floor(tmp / 1000);
    };
    return DataKraken;
}());
exports.DataKraken = DataKraken;
