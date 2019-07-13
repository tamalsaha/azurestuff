package main

import (
	"fmt"
	"net/http"
	"github.com/Azure/go-autorest/tracing"
	"github.com/Azure/go-autorest/autorest"
)

func main() {
	fmt.Println(tracing.IsEnabled())

	fmt.Println(autorest.GetLocation(&http.Response{}))
}
