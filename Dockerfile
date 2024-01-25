# Use the official Golang image as a base image
FROM golang:1.21.6

# Set the working directory inside the container
WORKDIR /app

# Copy the entire application directory to the container
COPY . .

# Run the go run command to execute the deployment script with environment variable substitution
CMD sh -c 'go run ./app/deploy_mate/Dockerfile -location=${LOCATION} -project=${PROJECT} -name=${NAME} -image=${IMAGE}'
