# Unique Identity (uid)

A library and RESTful webservice creating cryptographically strong pseudo-random 64bit integers (uint64).

[![Build Status](https://travis-ci.org/ory-platform/uid.svg?branch=master)](https://travis-ci.org/ory-platform/uid)

**ATTENTION**  
JavaScript does not support 64 bit integers. Use the "idStr" field, if you're using JS.

*Collision probability:* [This](http://preshing.com/20110504/hash-collision-probabilities) is an excellent document explaining hash collision probability of n bits.

## Install

```
go get github.com/ory-platform/uid
```
## Usage

### Library

```go

import (
    "fmt"
    "github.com/ory-platform/uid"
)

func main() {
    fmt.Printf("%d", uid.NewUid())
}
```

### RESTful server

To start the server, do:
* Windows: `%GOPATH%\bin\uid`
* MacOS & Linux: `$GOPATH/bin/uid`

The following environment variables are available:

| Variable             | Default                           | Description                          |
| -------------------- | --------------------------------- | ------------------------------------ |
| PORT                 | `80`                              | Port the application listens on      |
| HOST                 | `null`                            | Host to listen on (null = all hosts) |


#### UID Collection [/uids]

##### Create a UID [POST]
+ Response 200 (application/json)

        {
            "apiVersion": "1.0",
            "id": "5171f1e0-caa1-430d-9153-91acf82529fd",
            "data": {
                "uid": 39667251790123496,
                "uidStr": "39667251790123496"
            }
        }
        
+ Response 500 (application/json)
        
        {
            "apiVersion": "1.0",
            "id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
            "error": {
                "code": 500,
                "message": "Some error message"
            }
        }

Have a look at the [API Docs on apiary](http://docs.oryplatformuidserver.apiary.io/).
