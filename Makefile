GO = go
GOFLAGS = -ldflags=""
SOURCES = *.go
BINARY = main

.PHONY: debug run clean

$(BINARY): $(SOURCES)
	$(GO) build $(GOFLAGS) $(BINARY)

debug: $(BINARY)
	dlv exec --check-go-version=false ./$(BINARY);

run: $(BINARY)
	./$(BINARY); \

clean:
	rm -f $(BINARY)

