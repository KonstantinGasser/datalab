# service.eventmanager.live



# Data modeling for Funnel Aggregation with Cassandra
A `funnel` can consist out of `N` stages, where each `stage` represents one state in the `funnel`. 
The objective is to understand how many users (unique users) are in each `stage`.

## Approach #1
For the first approach, I am using a simple data schema where all extra meta-data (for the column family) has been ignored
but only focuses on the `partition key` and `clustering key`. 

### ***Table Schema***

![](git-resources/cassandra_approach_1_table.png)

Here the `partition key` is defined by the `stage-name`. The `clustering key` is defined by users ***UUID***.
This allows to `insert` users entering a given `stage` in a distinct way. Hence, a user will not be two times in the same `stage`.

### ***Query: get distinct count for stage X***
![](git-resources/cassandra_approach_1_query_1.png)

The result of this query shows that in `stage == "/home"` are three distinct users.

### ***Query: get distinct count for all stages with GROUP BY***
![](git-resources/cassandra_approach_1_query_2.png)

With this query all `stages` and their `distinct count` can be queried. However, as stated by the console output (`"Aggregation query used without partition key"`) we get an indication that the query might not perform good at scale


### Challenges 
Even-though this example represents a use-case from the problem statement, it ignores some relevant points. Firstly, the table will hold more than one `funnel definition` either from different `Apps` of the same organization or `Apps` from other organizations. However, by using the `stage` and `app-uuid` as `partition key` the query can still be efficient


## Event Definitions
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
