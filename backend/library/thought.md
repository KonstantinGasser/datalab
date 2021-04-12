## Connecting to the web-socket
Pre-Condition: presents of application-token referring to specific app.<br>
Steps:<br>
==== Cycle: START ====
client:{app_uuid}                -> api({app_uuid}) -> server: issues ticket (ttl=30sec)
client:{app_uuid,ticket,cookie}  -> open({ticket})  -> server: establish WS connection
client:{app_uuid,ticket,cookie}  -> close()         -> server: -
==== Cycle: END   ====

## Cookies
The call to acquire a web-socket ticket can either set a cookie if none present or skips it
if the request passes a cookie in the request.

## Interesting EventListener
- `click`
- `mousemove, pageX,pageY (coordiantes in respect to document)` : with time-outs thou
- `detail` (returns the number that indicates how many times the mouse was clicked)
- `newURL` (returns the URL of the document, after the hash has been changed)

## User Life-Cycle
1) connect with system: get cookie, open WS connection
2) init data: {
    referrer-page,
    session-start,
    identity (cookie),
}
3) Event-Loop: {
    identity (cookie),
    time-stamp,
    meta,
    event-type,
    event-value,
}
4) disconnect with system: handled server-side ()

## Can the lib aggregate data to do some work in prior?
Things like: 
- "hot areas" aggregating mouse movements
- aggregate summary per page (like /home, /product/my-product)
