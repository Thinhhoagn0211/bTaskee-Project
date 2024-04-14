.PHONY: install test clean build docs
# Run services
run:
	docker compose -f ./deployment/docker-compose.yml up -d

# Stop services
down:
	docker compose  -f ./deployment/docker-compose.yml down --remove-orphans

# Delete services
delete:
	docker compose  -f ./deployment/docker-compose.yml down --remove-orphans -v

# Build the Python package
build:
	docker buildx build -t booking_microservice:v1.0.0 -f ./deployment/Dockerfile.BookingService . 
	docker buildx build -t pricing_microservice:v1.0.0 -f ./deployment/Dockerfile.PricingService . 
	docker buildx build -t product_microservice:v1.0.0 -f ./deployment/Dockerfile.ProductService .
	docker buildx build -t sending_microservice:v1.0.0 -f ./deployment/Dockerfile.SendingService .