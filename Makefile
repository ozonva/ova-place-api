GO_VERSION_SHORT:=$(shell echo `go version` | sed -E 's/.* go(.*) .*/\1/g')
ifneq ("1.16","$(shell printf "$(GO_VERSION_SHORT)\n1.16" | sort -V | head -1)")
$(error NEED GO VERSION >= 1.16. Found: $(GO_VERSION_SHORT))
endif

export GO111MODULE=on
export GOPROXY=https://proxy.golang.org|direct

PGV_VERSION:="v0.6.1"
GOOGLEAPIS_VERSION="master"
BUF_VERSION:="v0.51.0"

default: all

.PHONY: test
test: ## Run tests
	go test -race -p 1 -timeout 300s -coverprofile=.test_coverage.txt ./... && \
    	go tool cover -func=.test_coverage.txt | tail -n1 | awk '{print "Total test coverage: " $$3}'
	@rm .test_coverage.txt

.PHONY: lint
lint: ## Lint the source files
	golangci-lint run
	gosec -quiet ./...

.PHONY: fmt
fmt:
	gofmt -s -w .

.PHONY: goimports
goimports:
	goimports -w -local github.com/ozonva/ova-place-api ./..

all: generate build

.PHONY: vendor-proto
vendor-proto:
	$(eval THIRD_PARTY:=$(CURDIR)/third_party)
	@[ -f $(THIRD_PARTY)/validate/validate.proto ] || (mkdir -p $(THIRD_PARTY)/validate/ && curl -sSL0 https://raw.githubusercontent.com/envoyproxy/protoc-gen-validate/$(PGV_VERSION)/validate/validate.proto -o $(THIRD_PARTY)/validate/validate.proto)
	@[ -f $(THIRD_PARTY)/google/api/http.proto ] || (mkdir -p $(THIRD_PARTY)/google/api/ && curl -sSL0 https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_VERSION)/google/api/http.proto -o $(THIRD_PARTY)/google/api/http.proto)
	@[ -f $(THIRD_PARTY)/google/api/annotations.proto ] || (mkdir -p $(THIRD_PARTY)/google/api/ && curl -sSL0 https://raw.githubusercontent.com/googleapis/googleapis/$(GOOGLEAPIS_VERSION)/google/api/annotations.proto -o $(THIRD_PARTY)/google/api/annotations.proto)

GOBIN?=$(GOPATH)/bin
buf.work.yaml:
	@echo "version: v1\ndirectories:\n  - protos\n  - third_party\n" > $(CURDIR)/buf.work.yaml
buf.gen.yaml:
	@echo "version: v1\nplugins:\n  - name: go\n    out: .\n    opt: module=github.com/ozonva/ova-place-api\n  - name: go-grpc\n    out: .\n    opt: module=github.com/ozonva/ova-place-api\n  - name: grpc-gateway\n    out: .\n    opt: logtostderr=true,module=github.com/ozonva/ova-place-api\n  - name: validate\n    out: .\n    opt: lang=go,module=github.com/ozonva/ova-place-api\n  - name: openapiv2\n    out: swagger\n    opt: allow_merge=true,merge_file_name=api" > $(CURDIR)/buf.gen.yaml

.PHONY: generate
generate: vendor-proto buf.work.yaml buf.gen.yaml
	@command -v buf 2>&1 > /dev/null || (mkdir -p $(GOBIN) && curl -sSL0 https://github.com/bufbuild/buf/releases/download/$(BUF_VERSION)/buf-$(shell uname -s)-$(shell uname -m) -o $(GOBIN)/buf && chmod +x $(GOBIN)/buf)
	$(eval PROTOS:=$(CURDIR)/api/ova-place-api)
	@[ -f $(PROTOS)/buf.yaml ] || PATH=$(GOBIN):$(PATH) buf mod init --doc -o $(PROTOS)
	PATH=$(GOBIN):$(PATH) buf generate $(PROTOS)

.PHONY: deps
deps:
	@[ -f go.mod ] || go mod init github.com/ozonva/ova-place-api
	find . -name go.mod -exec bash -c 'pushd "$${1%go.mod}" && go mod tidy && popd' _ {} \;

.PHONY: build
build: deps
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o $(CURDIR)/bin $(CURDIR)/cmd/ova-place-api/main.go

bin-deps:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27.1
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0
	go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0
	go install github.com/envoyproxy/protoc-gen-validate@$(PGV_VERSION)

.PHONY: mocks
mocks:
	mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozonva/ova-place-api/internal/flusher Flusher
	mockgen -destination=./mocks/repo_mock.go -package=mocks github.com/ozonva/ova-place-api/internal/repo Repo
	mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozonva/ova-place-api/internal/producer Producer
	mockgen -destination=./mocks/counter_mock.go -package=mocks github.com/ozonva/ova-place-api/internal/metrics Counter
	mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozonva/ova-place-api/internal/saver Saver
