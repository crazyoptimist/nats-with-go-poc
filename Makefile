# launch nats cluster using docker
up:
	@docker compose -f ./deployments/compose.yml up -d
down:
	@docker compose -f ./deployments/compose.yml down
log:
	@docker compose -f ./deployments/compose.yml logs -f

# pubsub
pub:
	@go run ./cmd/pubsub/publisher/main.go
sub:
	@go run ./cmd/pubsub/subscriber/main.go

# request-reply
reply:
	@go run ./cmd/requestreply/replier/main.go
request:
	@go run ./cmd/requestreply/requester/main.go

# limit based stream
createstream:
	@go run ./cmd/jetstream/createstream/main.go
produce:
	@go run ./cmd/jetstream/producer/main.go
consume:
	@go run ./cmd/jetstream/consumer/main.go
