// This is an example of implementing the Pet Store from the OpenAPI documentation
// found at:
// https://github.com/OAI/OpenAPI-Specification/blob/master/examples/v3.0/petstore.yaml
//
// The code under api/petstore/ has been generated from that specification.
package main

import (
	// "encoding/json"
	"flag"
	"fmt"
	// "io/ioutil"
	"os"

	"go_noodling/use_openapi_generated_server/api"
	// "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/labstack/echo/v4"
	// echomiddleware "github.com/labstack/echo/v4/middleware"
	// "go_noodling/use_openapi_generated_server/clients/service1"
)

func main() {
	var port = flag.Int("port", 8083, "Port for test HTTP server")
	flag.Parse()

	swagger, err := api.GetSwagger()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error loading swagger spec\n: %s", err)
		os.Exit(1)
	}

	// Clear out the servers array in the swagger spec, that skips validating
	// that server names match. We don't know how this thing will be run.
	swagger.Servers = nil

	// Create an instance of our handler which satisfies the generated interface
	petStore := api.NewPetStore()

	// This is how you set up a basic Echo router
	e := echo.New()
	// // Log all requests
	// e.Use(echomiddleware.Logger())
	// // Use our validation middleware to check all requests against the
	// // OpenAPI schema.
	// e.Use(middleware.OapiRequestValidator(swagger))

	// We now register our petStore above as the handler for the interface
	api.RegisterHandlers(e, petStore)

	// // TODO:
	// //
	// // - Understand how to use Generated Client code.
	// // - Move out of `main.go`.
	// req, err := service1.NewHealthReadRequest("http://localhost:8085")
	// client, err := service1.NewClient("http://localhost:8085")
	// fmt.Println(err)
	// resp, err := client.Client.Do(req)
	// fmt.Println(err)
	// // https://pkg.go.dev/net/http@go1.20.6#Client.Do
	// defer resp.Body.Close()
	// fmt.Println(resp)
	// fmt.Println(resp.Body)
	// body, err := ioutil.ReadAll(resp.Body)
	// // https://stackoverflow.com/questions/64610203/golang-read-https-response-body
	// // Prints character codes.
	// fmt.Println(body, err)
	// // https://stackoverflow.com/questions/40632802/how-to-convert-byte-array-to-string-in-go
	// // Prints expected JSON messagge.
	// fmt.Println(string(body))
	// // https://pkg.go.dev/encoding/json
	// // No result. Guessing this is from the above `ioutil.ReadAll()`?
	// var health service1.HealthReadResponse
	// err = json.Unmarshal(body, &health)
	// fmt.Println(err)
	// fmt.Println(health)
	// var b []byte
	// i, err := resp.Body.Read(b)
	// fmt.Println(i, err)
	// fmt.Println(b)
	// // fmt.Println(client.HealthRead())
	// fmt.Println("--- Requests:")
	// fmt.Println(req)
	// // fmt.Println(req.GetBody())
	// fmt.Println(req.RequestURI)
	// fmt.Println(req.URL)
	// fmt.Println(req.Body)
	// fmt.Println(req.Response)
	// // fmt.Println(req.Response.Body)
	// fmt.Println(err)
	// And we serve HTTP until the world ends.
	e.Logger.Fatal(e.Start(fmt.Sprintf("0.0.0.0:%d", *port)))
}
