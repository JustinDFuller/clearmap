test:
	@go test ./...;

test-watch:
	@reflex -s -- sh -c 'clear && $(MAKE) test';
