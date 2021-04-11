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
if `init=false` in query parameter.
