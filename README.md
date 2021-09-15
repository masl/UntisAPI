# UnitsAPI
This project Provides a Golang wrapper for the WebUntis API. This Repo is still a work in progress. Therefore not all API features are avaliable. 

## Features
- Login / Logout
- Get all Rooms, Teachers and Subjects.
- Get teacher or student id from name.
- Get Timetable for specifig room.
- Get the current schoolyear.
- Get all schoolyears.

## How to install:

Run the go get command:
```
go get github.com/Stroby241/UnitsAPI
```
and then import it into your go project:
```Go
import "github.com/Stroby241/UntisAPI"
```

## How to use:
First login with your Untis user.
```Go
user := UntisAPI.NewUser([username], [password], [schoolname], [domain])
user.Login()
```

| input field | descibtion |
| --- | --- |
| [username] | your Units username. |
| [password] | your Units password. |
| [schoolname] | The name of your school. The name must be exactly the same as in WebUntis. |
| [domain] | The domain of your schools WebUntis. |

To find the schoolname and domian. Go to https://webuntis.com/  
Then select your school. Then you should be forwarded to a web site like this:
```
https://tipo.webuntis.com/WebUntis/?school=TBZ+Mitte+Bremen#/basic/login
```
In this example the domain is:
```
https://tipo.webuntis.com/
```
and the schoolname is:
```
TBZ Mitte Bremen
```
---
When your done with your Untis requests please logout.
```Go
user.Logout()
```
---
To get the current time and date in Untis time format do this:
```Go
untisDate := UntisAPI.ToUntisDate(time.Now())
untisTime := UntisAPI.ToUntisTime(time.Now())
```
To convert it back do this:
```Go
goDate := UnitsAPI.ToGoDate(untisDate)
goTime := UnitsAPI.ToGoTime(untisTime)
```

