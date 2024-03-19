up:
	@docker compose -f ./deployments/compose.yml up -d
down:
	@docker compose -f ./deployments/compose.yml down
log:
	@docker compose -f ./deployments/compose.yml logs -f
pub:
	@go run ./cmd/pubsub/publisher/main.go
sub:
	@go run ./cmd/pubsub/subscriber/main.go
reply:
	@go run ./cmd/requestreply/replier/main.go
request:
	@go run ./cmd/requestreply/requester/main.go
