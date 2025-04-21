# Availability
(*Availability*)

## Overview

### Available Operations

* [Get](#get) - Check which operators and operations are currently available.

## Get

Check which operators and operations are currently available.

### Example Usage

```go
package main

import(
	"context"
	"os"
	paygo "github.com/nkwa/pay-go"
	"log"
)

func main() {
    ctx := context.Background()

    s := paygo.New(
        paygo.WithSecurity(os.Getenv("PAY_API_KEY_AUTH")),
    )

    res, err := s.Availability.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Availabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetAvailabilityResponse](../../models/operations/getavailabilityresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.HTTPError | 401                 | application/json    |
| apierrors.HTTPError | 500                 | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |