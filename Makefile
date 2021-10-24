LESS_DIR = ./static/css/
LESS_ENTRYPOINT = $(LESS_DIR)/application.less
CSS_FILE = ./static/css/main.css
GO_FILES = $(wildcard *.go)
LESS_SRC = $(wildcard $(LESS_DIR)/*.less)

all: build

build: $(GO_FILES)
	go build -o bin/planetocd .
	go build -o bin/translate ./translate
	go build -o bin/convert ./convert

.PHONY: run-local
run-local:
	./run-local.sh

css: $(CSS_FILE)

$(CSS_FILE): $(LESS_SRC)
	lessc $(LESS_ENTRYPOINT) > $(CSS_FILE)

clean:
	-rm -f $(CSS_FILE)
	-rm -f bin/planetocd
	-rm -f bin/translate
	-rm -f bin/convert
