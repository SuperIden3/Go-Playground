GO = go
GOFLAGS = -ldflags=""
SOURCES = *.go
BINARY = main

.PHONY: time debug run clean

$(BINARY): $(SOURCES)
	$(GO) build $(GOFLAGS) $(BINARY)

debug: $(BINARY)
	dlv exec --check-go-version=false ./$(BINARY)

run: $(BINARY)
	./$(BINARY)

time: $(BINARY)
	bash -c "time ./$(BINARY)"

clean:
	rm -f $(BINARY)

