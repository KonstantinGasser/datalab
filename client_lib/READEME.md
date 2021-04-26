# Client Library: Data Flow
The idea of this `README` is to explain how the data collection on the client-side works. Further, the session and data life-cycle will be explained as well as the data format.

## Session Life-Cycle
STATE: INIT <br>
As soon as the `DataKraken` object gets instantiated the code will perform an API call to the data backend. As a result, the server will set a new `cookie` if non is present in the request along with a web-socket  `ticket` and app-specific meta-data. The `ticket` can be used to open the web-socket connection (`ttl` is 15sec). Based on the meta-data, all allowed `EventListener` will be attached to the `document` with their corresponding functions. <br><br>
STATE: LISTEN <br>
After calling `DataKraken.Listen()` the web-socket gets opened and the functions `onopen`, `onclose` and `onerror` will be assigned. `.Listen()` further changes the `DataKraken.CONN_STATE` to `SOCKET_EVENT.READY` - if `onerror` is not triggered. This will wake up the `EventListener` which can now start streaming events to the socket connection. If the connections gets closed (by closing the tab/browser, network errors) the `onclose` function fires a close-event to the socket connection, indicated that the session is over.<br><br>

STATE: CLOSING<br>
While the `DataKraken.CONN_STATE == SOCKET_EVENT.CLOSING` the `EventListener` will be detached from the document. Computing the session duration will happen on the server since the web-socket connection might not work any longer.

## Data  Life-Cycle

## Data Specifications
DATA: SESSION_RECORD<br>
```json
{
    "type": "start",
    "meta": {
        "device": {"os": "Macintosh", "browser": "Chrome"},
        "referrer": "https://www.google.com",
    }
}
```
DATA: EVENT_MOUSECLICK<br>
```json
{
    "type": 0,
    "timestamp" unix-timestamp,
    "event": {
        "X": pos-mouse-x,
        "Y": pos-mouse-y,
        "target": "css class | id | name"
    }
}
```
DATA: EVENT_MOUSEMOVE<br>
```json
{
    "type": 1,
    "timestamp" unix-timestamp,
    "event": {
        "X": pos-mouse-x,
        "Y": pos-mouse-y,
        "ellapsed": duration of no-pos-change
    }
}
```