# go-test-2024-09-20

This is a Go program which imports code generated from these proto files: [james-davidson-ev/proto-test-2024-09-20](https://github.com/james-davidson-ev/proto-test-2024-09-20)

The generated code is published in Go module "iw-mapviewer.cmh.platform-dev.evinternal.net/api" but this obviously not a source repository.

The trick is that when you **go get iw-mapviewer.cmh.platform-dev.evinternal.net/api@main** it will retrieve from [james-davidson-ev/test-2024-09-20](https://github.com/james-davidson-ev/test-2024-09-20) instead.

This works because if you GET https://iw-mapviewer.cmh.platform-dev.evinternal.net/api the response includes a special <meta> tag as per [Remote import paths](https://pkg.go.dev/cmd/go#hdr-Remote_import_paths)

## note

don't forget to `export GOPRIVATE=iw-mapviewer.cmh.platform-dev.evinternal.net`
