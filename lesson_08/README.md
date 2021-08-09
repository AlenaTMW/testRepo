# Lesson 08 Docker

# Compile go app for linux 
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app1 .
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app2 .

# Build Docker image
docker build -t alenatmw/app1:latest .
docker build -t alenatmw/app2:latest .

# Push image to docker-hub
docker push alenatmw/app1
docker push alenatmw/app1

# Create user defined network
docker network create alenatmw

# Run container
docker run -d -it --rm --name app1 --network alenatmw -p 8080:8080 alenatmw/app1
docker run -d -it --rm --name app2 --network alenatmw alenatmw/app2

# or
docker-compose up
