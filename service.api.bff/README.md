# Service API BFF (backend-for-frontend)


## Responsibility
Die Aufgabe des `serivce.api.bff` ist es Client Request (von der WebApp) entgegenzunehmen und, wenn notwendig weitere calls zu N Services zu machen.

## Supported Endpoints
- [/api/v1/user/register](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_register_user.go)
- [/api/v1/user/login](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_login_user.go)
- [/api/v1/user/profile/update](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_update_profile.go)
- [/api/v1/user/profile](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_get_user.go)
- [/api/v1/user/colleagues](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_get_user.go)
- [/api/v1/app/create](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_create_app.go)
- [/api/v1/app](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_get_app.go)
- [/api/v1/app/all](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_get_app.go)
- [/api/v1/app/token/issue](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_issue_apptoken.go)
- [/api/v1/app/member](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_get_invitable_users.go)
- [/api/v1/app/config/update](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_update_appconfig.go)
- [/api/v1/app/invite](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_send_app_invite.go)
- [/api/v1/app/invite/reminder](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_send_app_invite.go)
- [/api/v1/app/invite/accept](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_send_app_invite.go)
- [/api/v1/app/member/invitable](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_get_invitable_users.go)
- [/api/v1/app/unlock](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/cmd/httpserver/api_unlock_app.go)

## Dependencies
Da `Kafka` aktuell keine Option ist hat das `api.bff` eine Dependency zu ALLEN Services und nutzt deren `gRPC` interface um Request durchzuführen.

## Workflow
API-Endpoints werden in der entsprechenden Datei `api_*.go` im `cmd/httpserver/` Directory angelegt und in der `main.go` verlinkt. Der `Api-Server` selbst hat Dependencies zu den Domain Modellen (`user`, `app`) mit welchen `http.HandlerFunc(s)` mit dem System interagieren können. Des Weiteren bietet der `Api-Server` diverse `middleware` Funktionen an, die vor einen Endpoint gesetzt werden können (siehe `cmd/httpserver/middleware.go`).
Für die Verarbeitung von JSON-Data werden keine `map[string]interfaces{}` als Datentype genutzt, um Type-Safety zu garantieren. Vielmehr stellt die entsprechende Domain `*Reqeust` und `*Response` structs zur Verfügung über die serialization und deserialization gehandhabt wird (siehe [apps-domain-request-response-types](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/internal/apps/apps.go)).


## Domains
Domains beschreiben einen Kontext (Zuständigkeit) in welchem verschiedene Services definiert werden können. Ein Service bietet dabei immer eine `type Service interface` an, das beschreibt, welche Funktionalität angeboten wird (siehe [apps-domain-service-collecting](https://github.com/KonstantinGasser/datalab/blob/main/service.api.bff/internal/apps/collecting/service.go)).

## Ports
Das package `ports` beinhaltet jeglichen Code, der gebraucht wird für Dependencies. Zum Beispiel wickelt der Service `api.bff` unter `ports/client` alle Interaktionen mit anderen Micro-Services via gRPC in diesem Directory ab (ein Kafka würde dann hier auch hingehören). 


## Open ToDos:
- [] move notification code in own domain with interfaces
- [] implement Kafka broker and domain once available
