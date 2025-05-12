# FaasFlows Demo Functions

## Overview

This repository contains example serverless functions built using a custom OpenFaaS template (`golang-flow`) that implements the FaasFlows approach for reducing vendor lock-in and response time in serverless workflows.

## Prerequisites

### Required Infrastructure

1. **Kubernetes cluster** - A running Kubernetes cluster (local or remote)
    - Minikube, kind, or a managed Kubernetes service (EKS, GKE, AKS)
    - Properly configured `kubectl` with access to your cluster

2. **Docker** - For building and pushing function images
    - Docker Desktop or Docker Engine with Docker CLI
    - Access to Docker Hub or a private container registry

3. **Helm** (v3+) - For deploying OpenFaaS components
    - Properly configured Helm with access to your Kubernetes cluster

### Required OpenFaaS Components with FaasFlows Support

1. **Modified OpenFaaS components** with FaasFlows functionality:
    - Modified [`faas-gateway`](https://github.com/danenherdi/faas-gateway) (with flow routing support)
    - Modified [`faas-provider`](https://github.com/danenherdi/faas-provider) (with flow execution support)
    - Modified [`faas-netes`](https://github.com/danenherdi/faas-netes) (with Kubernetes integration for flows)

2. **FaasFlows Configuration**:
    - Properly configured [FaasFlows-Demo-Config](https://github.com/danenherdi/FaasFlows-Demo-Config) repository
    - Applied ConfigMap with DAG definitions

### Required Tools
1. **faas-cli** - OpenFaaS CLI tool
   ```bash
   curl -sL https://cli.openfaas.com | sudo sh
   ```
2. **Node.js** - For running test scripts 
   - Required for executing the migration and redirection tests
3. **Custom golang-flow Template** - For building functions
    ```bash
   faas-cli template pull https://github.com/danenherdi/golang-http-template
   
   # OR

    git clone https://github.com/danenherdi/golang-http-template.git
    
   # Copy the template to your local OpenFaaS templates directory
   mkdir -p ~/.openfaas/templates
   cp -r ~/golang-http-template/template/golang-flow ~/.openfaas/templates/
   
   # Verify the template is available
   faas-cli template list
   faas-cli new --list
   ```
   It should show:
    ```
    Languages available as templates:
   - golang-flow
   - golang-http
   - golang-middleware
    ```


## Setup and Deployment

### 1. Clone the repository

```bash
git clone https://github.com/danenherdi/FaasFlows-Demo-Functions.git
cd FaasFlows-Demo-Functions
```

### 2. Verify and Configure OpenFaaS Gateway
Ensure your OpenFaaS gateway with FaasFlows support is running and accessible.

```bash
# Set your OpenFaaS gateway URL
export OPENFAAS_URL=http://127.0.0.1:8080

# Login if authentication is enabled
faas-cli login --password <your-password>

# Verify connection
faas-cli list
```
### 3. Build all functions
To build all functions defined in the .yml files:
```bash
faas-cli build -f user-info.yml
faas-cli build -f ride-history.yml
faas-cli build -f ride-recommend.yml
faas-cli build -f last-ride.yml
faas-cli build -f homepage.yml
```

Alternatively, you can combine all function definitions into a single stack.yml and build them all at once:
```bash
faas-cli build -f stack.yml
```

### Push the Function Images (if using a remote registry)
If you're using a remote Docker registry (not required for local testing with Minikube):
```bash
faas-cli push -f user-info.yml
faas-cli push -f ride-history.yml
faas-cli push -f ride-recommend.yml
faas-cli push -f last-ride.yml
faas-cli push -f homepage.yml
```

### 5. Deploy the functions
Deploy to your OpenFaaS Gateway:
```bash
faas-cli deploy -f user-info.yml
faas-cli deploy -f ride-history.yml
faas-cli deploy -f ride-recommend.yml
faas-cli deploy -f last-ride.yml
faas-cli deploy -f homepage.yml
faas-cli deploy -f hello-world.yml
```
Each function will be available through the OpenFaaS Gateway.

## Function Overview
The demo functions implement a simple ride-sharing application workflow:
1. **homepage** - Entry point function that aggregates data from other functions
2. **user-info** - Provides user profile information
3. **last-ride** - Retrieves information about the user's most recent ride 
4. **ride-history** - Shows a history of past rides 
5. **ride-recommend** - Recommends rides based on user history and location 
6. **friends** - Shows friends who have used the service 
7. **hello-world** - A simple test function to verify the workflow functionality

## Testing

The repository includes several K6 test scripts to verify FaasFlows functionality and measure performance improvements. [K6](https://k6.io/) is a modern load testing tool that allows us to simulate various traffic patterns and measure key performance metrics.

### Setup K6

Before running tests, ensure K6 is installed:

```bash
# Linux
sudo gpg -k
sudo gpg --no-default-keyring --keyring /usr/share/keyrings/k6-archive-keyring.gpg --keyserver hkp://keyserver.ubuntu.com:80 --recv-keys C5AD17C747E3415A3642D57D77C6C491D6AC1D69
echo "deb [signed-by=/usr/share/keyrings/k6-archive-keyring.gpg] https://dl.k6.io/deb stable main" | sudo tee /etc/apt/sources.list.d/k6.list
sudo apt-get update
sudo apt-get install k6

# macOS
brew install k6

# Windows
choco install k6
```

### Running Migration Tests
Migration tests verify the effectiveness of the function cluster migration method in reducing vendor lock-in. These tests simulate the migration of functions between environments while measuring performance impact.

```bash
# Run migration test for all functions
k6 run faasflows-migration-test.js

# Run specific function migration tests
k6 run faasflows-migration-homepage-test.js
k6 run faasflows-migration-friends-test.js
k6 run faasflows-migration-ride-history-test.js
```

### Running Request Redirection Tests
Request redirection tests measure the effectiveness of the redirecting requests to the nearest function method in reducing response time.

```bash
# Run request redirection test
k6 run faasflows-redirection-test.js
```

## Key Features of golang-flow Template
1. **Directed Acyclic Graph (DAG) Support** - Defines function dependencies and execution order 
2. **Flow Input Handling** - Automatically processes input parameters from parent functions 
3. **Children Dependency Management** - Manages and injects responses from child functions 
4. **Caching Integration** - Supports caching function responses for improved performance 
5. **Third-Party Function Support** - Enables integration with external functions
