NAME=ral2jqlog

.PHONY:
clean:
	@rm -f $(NAME)

.PHONY:
lint:
	@go vet ./...
