# Go MongoDB Processing

This guide walks you through setting up a Go program to process bank account data, 
reading from a JSON file, and updating a MongoDB database.

## Installation

### 1. Install Go

Download and install Go from [https://golang.org/dl/](https://golang.org/dl/). 
Ensure you set the Go binary directory to your system's PATH.

### 2. Install Docker

Download and install Docker Desktop from [Docker Desktop](https://www.docker.com/products/docker-desktop).

## Setting Up MongoDB

### 3. Pull MongoDB Docker Image

docker pull mongo

### 4. Create a MongoDB Container

docker run -d -p 27017:27017 --name my-mongo-container mongo

Verify that the container is running:

docker ps

Certainly! Below is the content formatted for a README.md file on GitHub:

markdown
Copy code
# Go MongoDB Processing

This guide walks you through setting up a Go program to process bank account data, reading from a JSON file, and updating a MongoDB database.

## Installation

### 1. Install Go

Download and install Go from [https://golang.org/dl/](https://golang.org/dl/). Ensure you set the Go binary directory to your system's PATH.

### 2. Install Docker

Download and install Docker Desktop from [Docker Desktop](https://www.docker.com/products/docker-desktop).

## Setting Up MongoDB

### 3. Pull MongoDB Docker Image

docker pull mongo

### 4. Create a MongoDB Container

docker run -d -p 27017:27017 --name my-mongo-container mongo

Verify that the container is running:

docker ps

### 5. Install MongoDB Compass

Download and install MongoDB Compass from MongoDB Compass.

### 6. Connect to MongoDB Container using MongoDB Compass
Open MongoDB Compass.
Click on "New Connection."
In the "Hostname" field, enter localhost.
Set the "Port" to 27017.
Click "Connect" to connect to the MongoDB container.

### 7. Create a Database and Collection using MongoDB Compass
Inside MongoDB Compass:
Click on the "Create Database" button.
Enter the database name as bank.
Create a collection named accounts within the bank database.
Prepare Data

### 8. Create Input JSON File
Create a file named input.json with sample account data:
[
  {"id": 1, "balance": 1000.0, "status": ""},
  {"id": 2, "balance": -500.0, "status": ""}
]

Go Application

### 9. Initialize Go Module and Write Program

Open a terminal or command prompt.
Navigate to the directory where you want to create your Go project.
Run the following command to initialize a Go module:

go mod init mymongoapp

Run the following command to download the dependencies mentioned in your Go program:

go get

Create a file named main.go and copy the provided Go program into it.

### 10. Run the Go Program
Open a terminal or command prompt.
Navigate to the directory containing your Go program.
Run the following command:

go run main.go

Verification
### 11. Verify Data using MongoDB Compass

Refresh the MongoDB Compass interface.
Navigate to the bank database and accounts collection.
You should see the processed data in the MongoDB collection.
