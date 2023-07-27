//go:generate go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen -config cfg.yaml openapi3_1.yaml

package service1

import (
	// "fmt"
	"github.com/labstack/echo/v4"
	// "go_noodling/use_openapi_generated_server/clients/service1"
)

func (p *HealthReadResponse) HealthRead(ctx echo.Context) error {
	// TODO: Get client implementation into this file!!

	// c, err := service1.NewClient("http://localhost:8085")
	// fmt.Println(err)
	return ctx.JSON(p.StatusCode(), p.Body)
}
