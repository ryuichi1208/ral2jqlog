NAME=ral2jqlog

.PHONY:
build:
	@go build -o $(NAME) -ldflags '-w -s' .

.PHONY:
clean:
	@rm -f $(NAME)

.PHONY:
lint:
	@go vet ./...
