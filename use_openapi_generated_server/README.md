[![Use OpenAPI Generated Server](https://github.com/jackson15j/go_noodling/actions/workflows/use_openapi_generated_server.yaml/badge.svg)](https://github.com/jackson15j/go_noodling/actions/workflows/use_openapi_generated_server.yaml)

# OpenAPI Codegen Experiment

Been playing with:
https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/README.md. The
`oapi-codegen` seems to be pretty smooth to work with.

* You can define where to generate sections of your YAML
  eg. `/api/server.cfg.yaml` and `/api/models.cfg.yaml`. The default is to
  generate Client/Server/Docs/Models all in one long file.
* You create your
  [implementation](https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/echo/api/petstore.go)
  based off the generated models and by referencing the function names of the
  generated server.
* `go generate ./...` will generate the `<server|models>.gen.go` files in the
  expected folders (folders must be created beforehand + `// go generate...`
  lines in the implementation files).
* Then routing/spinning up of the server is done in your
  [main](https://github.com/deepmap/oapi-codegen/blob/master/examples/petstore-expanded/echo/petstore.go)
  file.

I did a little play where:

* I got the code working in my repo.
* [Checked specmatic version](https://github.com/jackson15j/go_noodling/commit/d7b168d0831de4198f4572d6122c43e0ad13c733).
* [Made a non-backwards compatible change](https://github.com/jackson15j/go_noodling/commit/73dcd020aa1dc367b64bde9e6ec3c7eab00e3ab7).
* [Regenerated from V2 spec](https://github.com/jackson15j/go_noodling/commit/faff1f0b1e334f046a400c064694f7c2457af517).
* [Updated implementation](https://github.com/jackson15j/go_noodling/commit/85edba229d88ba2665a85defcfbcca2702e1d4b5).
* Tested the running App Server with Specmatic.
* Added CI that runs Specmatic tests against the App Server, with JUnit
  Reporting.

Overall, I think this can fly!
