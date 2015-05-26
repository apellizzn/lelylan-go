# Lelylan Go Package

Go client library for [Lelylan API](http://dev.lelylan.com)

## Introduction

#### What is Lelylan

[Lelylan](http://lelylan.com) makes it easy for developers to monitor and control all devices
in your house providing a simple and consistent REST API.


## Requirements


## Installation

Install using the Go command line

```
go get
```

## Getting started

### Import packages

```
import (
    "github.com/bradrydzewski/go.auth/oauth2"
    "github.com/apellizzn/lelylan-go/lelylan"
)
```

### Get an access token

First of all you need an access token to authoraze your requests in
Lelylan. To get one use the [OAuth2 package](https://github.com/bradrydzewski/go.auth/oauth2)
and if you are not used to OAuth2 concepts, take 10 minutes and read the
related documentation in the [dev center](http://dev.lelylan.com/api/oauth#language=node).



# Create a client
```
oauthClient := oauth2.Client {
        ClientId: os.Getenv("CLIENT_ID"),
    ClientSecret: os.Getenv("CLIENT_SECRET"),
    AccessTokenURL: "http://people.lelylan.com/oauth/token"}
```
# Require an access token with password credential flow

```
token, err := oauthClient.GrantTokenPassword(os.Getenv("USERNAME"), os.Getenv("PASSWORD"), "resources")
```
### Lelylan access

Once you have the access token you can access to the Lelylan API. The
following example shows how to print in the console a list of owned devices.


# Initialize Lelylan client
```
client := lelylan.Client{ Token: token }
```
# Get one device.
```
client.Device(<id here>)
fmt.Println(device.Uri) # get the device uri 
fmt.Println(device.Name) # get the device name
```

## Contributing

Fork the repo on github and send a pull requests with topic branches. Do not forget to
provide specs to your contribution.


### Running specs

* Work in progress

### Running locally

* Work in progress


### Coding guidelines

Follow [github](https://github.com/styleguide/) guidelines.


### Feedback

Use the [issue tracker](http://github.com/lelylan/lelylan-go/issues) for bugs.
[Mail](mailto:touch@lelylan.com) or [Tweet](http://twitter.com/lelylan) us for any idea that can improve the project.


### Links

* [GIT Repository](http://github.com/lelylan/lelylan-go)
* [Lelylan Dev Center](http://dev.lelylan.com)
* [Lelylan Site](http://lelylan.com)


## Authors

[Alberto Pellizzon](http://twitter.com/apellizz)


## Contributors

Special thanks to the following people for submitting patches.

## Copyright

Copyright (c) 2013 [Lelylan](http://lelylan.com).
See [LICENSE](https://github.com/lelylan/lelylan-rb/blob/master/LICENSE.md) for details.
