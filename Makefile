GO = go
GOFLAGS = -ldflags="-s -w"
SOURCES = *.go
BINARY = main

.PHONY: all build time clean run

$(BINARY): $(SOURCES)
	$(GO) build $(GOFLAGS) $(BINARY)

run: $(BINARY)
	./$(BINARY)

time: $(BINARY)
	time ./$(BINARY)

clean:
	rm -f $(BINARY)

