# Command Authentication Server
While writing commands for Floor Gang bot it's good to make sure that a given user has the
permissions to run a command if it's related to administrating. This program acts as a local
process which can be communicated to with RPC+HTTP. This allows any command to phone-in and see
if a certain member has the permissions to run a command.


## Setup
You need [Go](https://golang.org).
```shell script
$ go build
```

It will then create an executable. Run it once, fill out the configuration, and you're ready to
go. To communicate with this process use the [authclient](https://github.com/floor-gang/authclient)

## Versioning
`X.Y`
 - X: A major change has happened, please update your code.
 - Y: A patch was implemented.
