LESS_DIR = ./static/css/
LESS_ENTRYPOINT = $(LESS_DIR)/application.less
CSS_FILE = ./static/css/main.css
less_src = $(wildcard $(LESS_DIR)/*.less)

all: run-local

.PHONY: run-local
run-local:
	./run-local.sh

css: $(CSS_FILE)

$(CSS_FILE): $(less_src)
	lessc $(LESS_ENTRYPOINT) > $(CSS_FILE)

clean:
	-rm -f $(CSS_FILE)
