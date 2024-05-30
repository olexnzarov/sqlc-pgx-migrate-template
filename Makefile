SQLC_IMAGE=sqlc/sqlc:1.26.0

generate-sqlc:
	@docker run --rm -v $(PWD):/src -w /src $(SQLC_IMAGE) generate
	
cleanup-template:
	@rm -rf ./cmd
	@rm docker-compose.yaml
	@rm Dockerfile
	@rm README.md
	@rm LICENSE
	@rm -rf .git
	@git init
