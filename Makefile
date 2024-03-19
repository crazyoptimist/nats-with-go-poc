up:
	@docker compose -f ./deployments/compose.yml up -d
down:
	@docker compose -f ./deployments/compose.yml down
log:
	@docker compose -f ./deployments/compose.yml logs -f
pub:
	@go run ./cmd/publisher/main.go
sub:
	@go run ./cmd/subscriber/main.go
reply:
	@go run ./cmd/replier/main.go
request:
	@go run ./cmd/requester/main.go
