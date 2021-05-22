# datalab analysis platform for user activity data


# Client Library: Data Flow
The idea of this `README` is to explain how the data collection on the client-side works. Further, the session and data life-cycle will be explained as well as the data format.

## Session Life-Cycle

STATE: INIT <br>
- call to `/api/hello` to indicate session start with following data:
``` json
{   
    "session_start": "UNIX time-stamp",
    "referrer": "page current page was called from",
    "browser": "Chrome",
    "OS": "MacOS",
}
```
- pass `cookie` if present else server sets new cookie
- call-back returns `web-socket ticket` to connect to socket
- attach `Event-Listener` to document

STATE: CONNECT <br>
- connect with Web-Socket

STATE: LISTEN <br>
- `listen for events` -> `process event` -> `send to web-socket` -> `start over`

STATE: CLOSING <br>
- graceful: send `goodbye` to server
- forceful: conn interrupt -> server terminates session

## Client data we get
- `referrer` | #1, #3
- `device info` | #4
- `click` of element | #2, #1
- `X,Y` of mouse-movement (needs more thinking - what to do with the data??)
- `elpased time` mouse hovered over specific element | #4
- `URL change` | #1, #2
- `time on URL` | #4

## What to visualize?
- `Customer Journey` [1]
- `Funnel (conversion rate)` [2]
- `Compaign Tracking` [3]
- `Audience Info` [4]


## Data by event

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
EVENT: MOUSECLICK<br>
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
EVENT: URLCHANGE<br>
```json
{
    "type": 0,
    "timestamp" unix-timestamp,
    "event": {
        "elapsed" time-in-seconds,
        "next": "http://awesome.dev/next"
    }
}
```
EVENT: MOUSEHOVER<br>
```json
{
    "type": 0,
    "timestamp" unix-timestamp,
    "event": {
        "elapsed": duration of no-pos-change,
        "target": "css class | id | name"
    }
}
```
EVENT: MOUSEMOVE<br>
```json
{
    "type": 1,
    "timestamp" unix-timestamp,
    "event": {
        "X": pos-mouse-x,
        "Y": pos-mouse-y,
        "elapsed": duration of no-pos-change
    }
}
```

## Docker-Swarm deployment
### CI/CD Pipe
<!-- The swarm lives on a Raspberry-PI4 (linux/arm64) consisting out of one node.
Each service (api,app,user,token,frontend) have their own `Makefile` with the `deploy` target. `make deploy` cross-compilies the executable for `linux/arm64` and builds a docker image also with cross-compilation for `linux/arm64`. Docker cross-compilation is achieved with the `docker buildx build` tool from docker which allows to build images on your local machine for a different OS/Arch. After the build `deploy` pushes the image to the `datalab-registry.dev:5000/<image-name>:<git-commit-hash>` which lives within the `swarm`. From their services can pull the latest images. -->



## Service - DNS Table (some say they can see a pattern..not sure where??)
| Service               | swarm-name           | port in:out | credentials                |
|-----------------------|----------------------|-------------|----------------------------|
| gateway               | gateway              | 8080:8080   |                            |
| app                   | appservice           | 8003:8003   |                            |
| user                  | userservice          | 8001:8001   |                            |
| userauth              | userauth             | 8002:8002   |                            |
| config                | configservice        | 8005:8005   |                            |
| apptoken              | apptokenservice      | 8006:8006   |                            |
| frontend              | vuefrontend          | 80:80       |                            |
| mongo-app             | appstorage           | 27018:27017 | appstorage:secure          |
| monog-user            | userstorage          | 27017:27017 | userstorage:secure         |
| monog-config          | configstorage        | 27019:27017 | configstorage:secure       |
| monog-apptoken        | apptokenstorage      | 27020:27017 | apptokenstorage:secure     |
| monog-userauth        | userauthstorage      | 27021:27017 | userauthstorage:secure     |



# So fare...
 
![](git-resources/demo_img_1.png)

![](git-resources/demo_img_2.png)
