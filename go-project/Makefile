# Load environment variables from .env file if it exists
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

# Variables
DOCKER_USER = rajeshecb70
DOCKER_IMAGE = goproject
TAG = 1.1.3

print_version:
	@echo $(TAG)
	@echo $(DOCKER_USER)
	@echo $(DOCKER_IMAGE)

# Check the Go version
version:
	@echo "Check the go version..."
	go version

# Install dependencies
install:
	@echo "Install the dependencies..."
	go mod download

# Check linting
lint:
	@echo "Check the linting..."
	golangci-lint run main.go main_test.go

# Run tests
test:
	@echo "Test the project..."
	go test -v

# Create the build
build:
	@echo "Build the project..."
	go build -o build/bin/go-project main.go


# Build and tag the Docker image
docker-build:
	@echo "Building and tagging Docker image..."
	docker build -t $(DOCKER_USER)/$(DOCKER_IMAGE):$(TAG) .


# Login to Docker Hub (Docker login detail mention in your .env file)
docker-login:
	@echo "Logging in to Docker Hub..."
	@docker login -u $(DOCKER_USERNAME) -p $(DOCKER_PASSWORD)

# Push the Docker image to Docker Hub
docker-push: docker-login
	@echo "Pushing Docker image to Docker Hub..."
	docker push $(DOCKER_USER)/$(DOCKER_IMAGE):$(TAG)

# Run the Docker container locally
docker-run:
	@echo "Running Docker container locally..."
	docker run -d -p 8080:8080 --name $(DOCKER_IMAGE) $(DOCKER_USER)/$(DOCKER_IMAGE):$(TAG)
