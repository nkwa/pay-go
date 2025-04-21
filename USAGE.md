<!-- Start SDK Example Usage [usage] -->
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