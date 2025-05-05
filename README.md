# pay-go
Nkwa Pay SDK for Go

<!-- Start Summary [summary] -->
## Summary

Nkwa Pay API: Use this API to integrate mobile money across your payment flows, create and manage payments, collections, and disbursements.

Read the docs at https://docs.mynkwa.com/api-reference
<!-- End Summary [summary] -->

<!-- Start Table of Contents [toc] -->
## Table of Contents
<!-- $toc-max-depth=2 -->
* [pay-go](#pay-go)
  * [SDK Installation](#sdk-installation)
  * [SDK Example Usage](#sdk-example-usage)
  * [Authentication](#authentication)
  * [Available Resources and Operations](#available-resources-and-operations)
  * [Retries](#retries)
  * [Error Handling](#error-handling)
  * [Server Selection](#server-selection)
  * [Custom HTTP Client](#custom-http-client)

<!-- End Table of Contents [toc] -->

<!-- Start SDK Installation [installation] -->
## SDK Installation

To add the SDK as a dependency to your project:
```bash
go get github.com/nkwa/pay-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

```go
package main

import (
	"context"
	paygo "github.com/nkwa/pay-go"
	"log"
	"os"
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
<!-- End SDK Example Usage [usage] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type   | Scheme  | Environment Variable |
| ------------ | ------ | ------- | -------------------- |
| `APIKeyAuth` | apiKey | API key | `PAY_API_KEY_AUTH`   |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	paygo "github.com/nkwa/pay-go"
	"log"
	"os"
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
<!-- End Authentication [security] -->

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

<details open>
<summary>Available methods</summary>

### [Availability](docs/sdks/availability/README.md)

* [Get](docs/sdks/availability/README.md#get) - Check which operators and operations are currently available.


### [Payments](docs/sdks/payments/README.md)

* [GetByID](docs/sdks/payments/README.md#getbyid) - Get the payment (collection or disbursement) with the specified ID.
* [Collect](docs/sdks/payments/README.md#collect) - Collect a payment from a phone number.
* [Disburse](docs/sdks/payments/README.md#disburse) - Disburse a payment from your balance to a phone number.

</details>
<!-- End Available Resources and Operations [operations] -->

<!-- Start Retries [retries] -->
## Retries

Some of the endpoints in this SDK support retries. If you use the SDK without any configuration, it will fall back to the default retry strategy provided by the API. However, the default retry strategy can be overridden on a per-operation basis, or across the entire SDK.

To change the default retry strategy for a single API call, simply provide a `retry.Config` object to the call by using the `WithRetries` option:
```go
package main

import (
	"context"
	paygo "github.com/nkwa/pay-go"
	"github.com/nkwa/pay-go/retry"
	"log"
	"models/operations"
	"os"
)

func main() {
	ctx := context.Background()

	s := paygo.New(
		paygo.WithSecurity(os.Getenv("PAY_API_KEY_AUTH")),
	)

	res, err := s.Payments.GetByID(ctx, "96e9ed79-9fef-44a6-9435-0625338ca86a", operations.WithRetries(
		retry.Config{
			Strategy: "backoff",
			Backoff: &retry.BackoffStrategy{
				InitialInterval: 1,
				MaxInterval:     50,
				Exponent:        1.1,
				MaxElapsedTime:  100,
			},
			RetryConnectionErrors: false,
		}))
	if err != nil {
		log.Fatal(err)
	}
	if res.Payment != nil {
		// handle response
	}
}

```

If you'd like to override the default retry strategy for all operations that support retries, you can use the `WithRetryConfig` option at SDK initialization:
```go
package main

import (
	"context"
	paygo "github.com/nkwa/pay-go"
	"github.com/nkwa/pay-go/retry"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paygo.New(
		paygo.WithRetryConfig(
			retry.Config{
				Strategy: "backoff",
				Backoff: &retry.BackoffStrategy{
					InitialInterval: 1,
					MaxInterval:     50,
					Exponent:        1.1,
					MaxElapsedTime:  100,
				},
				RetryConnectionErrors: false,
			}),
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
<!-- End Retries [retries] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations. All operations return a response object or an error, they will never return both.

By Default, an API error will return `apierrors.APIError`. When custom error responses are specified for an operation, the SDK may also return their associated error. You can refer to respective *Errors* tables in SDK docs for more details on possible error types for each operation.

For example, the `GetByID` function may return the following errors:

| Error Type          | Status Code | Content Type     |
| ------------------- | ----------- | ---------------- |
| apierrors.HTTPError | 401, 404    | application/json |
| apierrors.HTTPError | 500         | application/json |
| apierrors.APIError  | 4XX, 5XX    | \*/\*            |

### Example

```go
package main

import (
	"context"
	"errors"
	paygo "github.com/nkwa/pay-go"
	"github.com/nkwa/pay-go/models/apierrors"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paygo.New(
		paygo.WithSecurity(os.Getenv("PAY_API_KEY_AUTH")),
	)

	res, err := s.Payments.GetByID(ctx, "96e9ed79-9fef-44a6-9435-0625338ca86a")
	if err != nil {

		var e *apierrors.HTTPError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.HTTPError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *apierrors.APIError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Override Server URL Per-Client

The default server can be overridden globally using the `WithServerURL(serverURL string)` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	paygo "github.com/nkwa/pay-go"
	"log"
	"os"
)

func main() {
	ctx := context.Background()

	s := paygo.New(
		paygo.WithServerURL("https://api.pay.staging.mynkwa.com"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->
