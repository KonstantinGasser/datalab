# Clickstream analysis platform for user activity data


<!-- # Exposé
## Problemstellung 
Konstruieren einer Clickstream Platform für StartUps. StartUps sollen in der Lage sein, durch die Nutzung der Platform, mehr Informationen über Ihr Produkt (App) zu erfahren.
Die Platform muss somit in der Lage sein, Daten von den Usern zu sammeln, dass kann zum Beispiel über eine Client-Library passieren, welche Event Listener an den Context hängt, und entstehende Events an einen Service weiterleitet. Neben dem Sammeln der Daten, muss das System Skalierbare sein. Eine Hochrechnung kann wie folgt aussehen:    
Bei 5 aktiven StartUps mit jeweils 10.000 aktiven Usern, muss das System 5 x 10.000 x avg(Events/pro Minute) verarbeiten können. Darüberhinaus besteht ein essentieller Teil darin, Datenschemata und Datenbanken so zu entwicklen, dass auch das Querying nach Daten performant gestaltet werden kann. Diese Aufgaben und Herausforderungen sollen Bestand der Bachelor Arbeit sein, in welcher an Hand diese Beispiels der Prozess und die Entwicklung der Software nach gegeben Paradigmen und Methoden beschrieben wird.
Über die Bachelor Arbeit hinaus, ist es ebenfalls eine Aufgabe, den StartUps die Daten zu präsentieren, was über eine eigen Plattform erfolgen wird.

## Erkenntnisinteresse
- Performance oriented WebSocket programing (I/O Multiplexing)
- Einsatz von NoSQL Datenbanken (logisches Einsetzen von Polyglot-Persistence)
- Architektur von Micor-Services und die dazugehörige Infrastruktur

Abstract: some test goes.   
here I guess -->


## Table of Contents
1. [Data Collection System](#Data-Collection-System)    
    1.1 [Architecture/Infrastructure]()     
    1.2 [Web-Socket based Event-Service]()

2. [Analysis Platform System](#Analysis-Platform-System)    
    2.1 [Architecture/Infrastructure]()     
    2.2 [API-Gateway]()     
    2.3 [User-Service]()    
    2.4 [Token-Service]()   
    2.5 [App-Service]()     
    2.6 [Analysis-Service]()    
    2.7 [Subscription-Service]()??? 

3. [Micro-Service communication over gRPC]()
4. [Kafka as central message-bus system]()

## Service Communication
### API-Gateway
uses REST-API to serve client requests
uses gRPC client to communicate with Micro-Services
### Micro-Services
uses gRPCs to communicate with different services 


## API-Service
***Responsibilities***
- REST-API
- serve all platform requests
- login
- register
- auth
- app
- ...
---

### User-Service
***Responsibilities***
- Creating new Users
- Providing data for authentication of user
- Managing user profile data (CRUD)

***Interfaces***
- Service.CreateUser
- Service.AuthUser
- Service.GetUser
- Service.UpdateUser
- Service.DestroyUser

***Dependencies***
- lorum
---


### App-Service
***Responsibilities***
- Register a new App (here app is an application of the user which can be registered in oder to collect data from the app)
- Managing App-Team: other users allowed to see the app data ([app-team permissions](#Configuration-Service)) 
  
***Interfaces***    
- Service.RegisterApp
- Service.DeleteApp
- Service.GetApp
- Service.UpdateApp
- Service.AddMember
- Service.RemoveMember  

***Dependencies***  
- [User-Service](#User-Service)
- [Configuration-Service](#Configuration-Service)
- [Token-Service](#Token-Service)
---


### Configuration-Service
***Responsibilities***
- Keeping track of the configurations of an [App](#App-Service)
    - Rate-Limit for concurrent connection on the app
    - Permissions of app-team members
    - ...

***Interfaces***    
- Service.SetConfiguration
- Service.RemoveConfiguration
- Service.GetConfiguration
- Service.ApproveOperation  
  
***Dependencies***  
- None
---


### Token-Service
***Responsibilities***
- signing JSON-Web-Tokens (JWT) for authentication 
- verifying JWTs for authentication          
```json 
# AUTH-Token (JWT)
    {   
        Header {},
        User-ID,
        Exp,
        IAT,
        SUB,
        Signiture{}
     }
```
- signing App-Tokens (required for embedding client on App      
```json 
# App-Token
    {
        App-ID, 
        Access-Token,
        Meta{
            Rate-Limit, 
            Expiration,
        }
     }
```
    
***Interfaces***    
- Service.IssueAuthToken
- Service.VerifyAuthToken
- Service.IssueAppToken
- Service.VerifyAppToken   
  
***Dependencies***  
- None 
---


# Infrastructure cheat sheet

Services:
- APIGateway: :8080
- UserService: :8001
- TokenService: :8002
- AppService: :8003
(some say they can see a pattern not sure where..must change in future - ok for now)

Database server:
- UserService MongoDB: rasp-1:27017
- AppService MongoDB: rasp-1:27018
