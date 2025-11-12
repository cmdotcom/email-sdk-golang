# CM Email Gateway API Golang package
Official API docs for Email Gateway API: https://developers.cm.com/messaging/docs/email-introduction

## Using the `emailgateway` package
[![Go Reference](https://pkg.go.dev/badge/github.com/cmdotcom/email-sdk-golang/emailgateway.svg)](https://pkg.go.dev/github.com/cmdotcom/email-sdk-golang/emailgateway)

### Installation
```shell
go get github.com/cmdotcom/email-sdk-golang
```

### Initialize client
```go
package main

import (
    "os"
	
    "github.com/cmdotcom/email-sdk-golang/emailgateway"
)

func main() {
	client, err := emailgateway.NewClient(emailgateway.Config{
		ProductToken:                 "your-product-token",
		DefaultTransactionalPriority: emailgateway.PriorityHigh,
	})
	if err != nil {
		panic(err)
	}
}
```

### Sending an email
```go

// Client initialization code

email := emailgateway.Email{
    From: emailgateway.Address{
        Name:  "CM.com",
        Email: "no-reply@cm.com",
    },
    To: []emailgateway.Address{
        {
            Name:  "Example Receiver",
            Email: "email@example.com",
        },
    },
    Subject:  "My first emailgateway email",
    HTMLBody: "<h1> Hello world! </h1>",
    TextBody: "email",
}

response, err := client.SendTransactionalEmail(email)
if err != nil {
	// Handle error
}
```
