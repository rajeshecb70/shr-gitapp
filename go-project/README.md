## 1. Project Title
  
Create a Go project to retrieve the system's CPU usage, CPU load, and system uptime.

## 2. Project Description
  
This project is built using the Go programming language and runs in a Docker container. It utilizes multi-stage builds to minimize the image size and is deployed on an Azure server.

## 3. Prerequisites
  
- Docker (must be installed on local and azure server)
- Go language
- Make
- Azure Server (SSH msut be enabled with related application port)
- GitHub Secrets Configuration
    - `DOCKERHUB_USERNAME`: Your DockerHub username.
    - `DOCKERHUB_TOKEN`: Your DockerHub access token/password.
    - `AZURE_SERVER_USERNAME`: The username for your Azure server.
    - `AZURE_SERVER_PASSWORD`: The password for SSH access to the Azure server.
    - `AZURE_SERVER_IP`: The IP address of your Azure server.
    - `DOCKERHUB_REPONAME`: This is the docker reponame.

## 4. Setup & configuration
 Makefile Configuration
  
 ```
 # Clone the repository
 git clone https://github.com/rajeshecb70/go-project.git
 cd go-project/
 ```

 ```
 # Target to Check version. (must installed go in system)
 make version
 ```

 ```
 # Target to Install dependencies
 make install
 ```

 ```
 # Target to check linting
 make lint
 ```

 ```
 # Target to run tests
 make test
 ```

 ```
 # Target to create a build
 make build
 ```

 ```
 # Target to  build a docker image
 make docker-build
 ```

 ```
 # Target to login in to docker
 make docker-login
 ```

 ```
 # Target to push the docker image in the docker hub
 make docker-push
 ```

 ```
 # Target to run the service
 make docker-run
 ```
### 5. **GitHub Actions CI Pipeline**
   The project uses GitHub Actions to automate the CI/CD pipeline. The pipeline is defined in `.github/workflows/go-ci.yml`.

### 6. Expectations
- The following expectations should be met to complete this project.
  -  Proper tagging to docker image.✅
  - Makefile should have all targets as below.
    - To check the version.✅
    - To install go dependencies.✅
    - To check code linting.✅
    - To run the test.✅
    - To build the project.✅
    - To build the docker image.✅
    - To push the artifact to docker artifactory.✅
    - To run the docker image.✅
  - README.md file should be updated with instructions
    - To add pre-requisites for any existing tools that must already be installed (e.g., docker, make, etc)✅
    - To run different make targets and the order of execution.✅
  
### 7. Accessing the Application
  - Application end-points at the local system:
    - CPU usages: http://localhost:8080/cpu ✅
    - System uptime: http://localhost:8080/uptime ✅
    - System Load: http://localhost:8080/load ✅
  - Application end-points at Cloud server:
    - CPU usages: http://Server-IP:8080/cpu ✅
    - System uptime: http://Server-IP:8080/uptime ✅
    - System Load: http://Server-IP:8080/load ✅

### 8. Snapshots
  - Local-CPU usages: ![CPU usages](snapshots/cpu.png)
  - Local-System Load: ![System Load](snapshots/load.png)
  - Local-System uptime: ![System uptime](snapshots/uptime.png)
  - Go-pipeline: ![go-pipeline](snapshots/go-pipeline.png)
  - Deployment on Azure Server: ![Azure-deployment](snapshots/deployment_on_Azure.png)
  
  - Azure-CPU usages: ![Azure-CPU usages](snapshots/Azure_cpu.png)
  - Azure-System Load: ![Azure-System Load](snapshots/azure_load.png)
  - Azure-System uptime: ![Azure-System uptime](snapshots/azure_uptime.png)
