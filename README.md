# go-wasm-example

Example of basic Go WebAssembly build and dev server using taskfile.

**Prerequisites**
* go 1.11+ (tested with 1.13) - https://golang.org/
* takfile 2.7+ - https://taskfile.dev

**Tasks**
* build - builds project into webassembly / creates dist folder containing web application
* run - runs build and starts simple web server on :8080 with dist folder
* clean - removes all files created by build/run
build and clean can have watch enabled by -w or --watch

**Example**

    > task run -w