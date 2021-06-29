# service.eventmanager.live

## Event Types

### Raw-Event: Click
```json
{   
    "type": "int", // indicates what type of event
    "timestamp": "int64", 
    "device_id": "string", // most likly the devices IP address
    "target": "string", // clicked HTML element - if given HTML-Name-Tag else whatever if find lol
    "elapsed_time": "int64", // passed time since last click
    "current_url": "string" // URL clicked happened
}
```
### Raw-Event: URL-Change
```json
{   
    "type": "int", // indicates what type of event
    "timestamp": "int64", 
    "device_id": "string", // most likly the devices IP address
    "from": "string", // URL jumped from
    "to": "string", // URL jumped to
    "elapsed_time": "int64", // passed time on "from" URL
}
```

### Config-Based-Event: BTN-Time
```json 
{   
    "type": "int", // see above
    "timestamp": "int64", 
    "device_id": "string", // see above
    "target": "string", // HTML element triggered by
    "action": "string", // can be hover-then-leave or hover-then-clicked
    "elapsed_time": "int64", // passed time from click to action
}
```

### Config-Based-Event: Funnel-Change (onClick / URL-Change if part of funnel config)
```json 
{
    "type": "int", // see above
    "timestamp": "int64", 
    "device_id": "string", // see above
    "action": "string", // onClick || onUrlChange
    "entered": "int", // stage id
    "elapsed_time": "int64" 
}
```
