# Payments
(*Payments*)

## Overview

### Available Operations

* [GetByID](#getbyid) - Get the payment (collection or disbursement) with the specified ID.
* [Collect](#collect) - Collect a payment from a phone number.
* [Disburse](#disburse) - Disburse a payment from your balance to a phone number.

## GetByID

Get the payment (collection or disbursement) with the specified ID.

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

    res, err := s.Payments.GetByID(ctx, "96e9ed79-9fef-44a6-9435-0625338ca86a")
    if err != nil {
        log.Fatal(err)
    }
    if res.Payment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `id`                                                     | *string*                                                 | :heavy_check_mark:                                       | ID of payment to fetch                                   |
| `opts`                                                   | [][operations.Option](../../models/operations/option.md) | :heavy_minus_sign:                                       | The options for this request.                            |

### Response

**[*operations.GetPaymentsIDResponse](../../models/operations/getpaymentsidresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.HTTPError | 401, 404            | application/json    |
| apierrors.HTTPError | 500                 | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## Collect

Collect a payment from a phone number.

### Example Usage

```go
package main

import(
	"context"
	"os"
	paygo "github.com/nkwa/pay-go"
	"github.com/nkwa/pay-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := paygo.New(
        paygo.WithSecurity(os.Getenv("PAY_API_KEY_AUTH")),
    )

    res, err := s.Payments.Collect(ctx, components.PaymentRequest{
        Amount: 433642,
        PhoneNumber: "237650000000",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Payment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.PaymentRequest](../../models/components/paymentrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.PostCollectResponse](../../models/operations/postcollectresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.HTTPError | 400, 401, 403       | application/json    |
| apierrors.HTTPError | 500                 | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |

## Disburse

Disburse a payment from your balance to a phone number.

### Example Usage

```go
package main

import(
	"context"
	"os"
	paygo "github.com/nkwa/pay-go"
	"github.com/nkwa/pay-go/models/components"
	"log"
)

func main() {
    ctx := context.Background()

    s := paygo.New(
        paygo.WithSecurity(os.Getenv("PAY_API_KEY_AUTH")),
    )

    res, err := s.Payments.Disburse(ctx, components.PaymentRequest{
        Amount: 410119,
        PhoneNumber: "237650000000",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.Payment != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `request`                                                              | [components.PaymentRequest](../../models/components/paymentrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |
| `opts`                                                                 | [][operations.Option](../../models/operations/option.md)               | :heavy_minus_sign:                                                     | The options for this request.                                          |

### Response

**[*operations.PostDisburseResponse](../../models/operations/postdisburseresponse.md), error**

### Errors

| Error Type          | Status Code         | Content Type        |
| ------------------- | ------------------- | ------------------- |
| apierrors.HTTPError | 400, 401, 403       | application/json    |
| apierrors.HTTPError | 500                 | application/json    |
| apierrors.APIError  | 4XX, 5XX            | \*/\*               |