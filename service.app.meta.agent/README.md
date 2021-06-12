# service.app.meta.agent

## Purpose
The `app.meta.agent` manages and maintains all the meta data involved for an `datalab app`.
Data includes the following:

- ID (random UUID used as `mongodb document id`)
- App-Name
- App-URL
- App-Description
- Owner Reference
- App-Hash (sha256 of `owner-domain/app-name`)
- Member Reference: Array<{`user-uuid`, `status`}>
- locked (indicates no further modification can be done until unlocked)

The service is structured in a `domain driven` approach. In the `internal` package all the service logic is written.
In this directory in the `apps package` is also the `interface` for the repository (here it will be implemented by a `mongoDB` client)
All `gRPC` related code can be found in the `cmd` directory. Any code related to handle `I/O` to other services either through `gRPC` or Kafka belongs in the `ports` directory. 


## Business Rules
- only app owner may invite new member
- to fetch data from an app the caller must either be owner or a member with status `apps.InviteAccepted`
- deleting an app can only be done by the owner
- locked apps can not be modified (the meta data can as of now not be changed anyway, however it must be considered if that changes in the future)

