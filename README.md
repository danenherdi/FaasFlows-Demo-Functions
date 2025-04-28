# FaasFlows Demo Functions

This repository contains example functions built using a custom OpenFaaS template: **golang-flow**.

The examples demonstrate how to use FaasFlows to compose workflows across multiple functions.

---

## Prerequisites

- OpenFaaS Gateway is running (with FaasFlows support).
- Kubernetes cluster is available and configured (`kubectl` context set).
- `faas-cli` is installed and configured.
- Docker is installed and running (for local builds).

## Setup

### 1. Clone the repository

```bash
git clone https://github.com/danenherdi/FaasFlows-Demo-Functions.git
cd FaasFlows-Demo-Functions
```

### 2. Pull the custom template
Make sure to pull the golang-flow template:
```bash
faas-cli template pull https://github.com/danenherdi/golang-http-template
```
Note: Ensure that golang-flow appears when listing available templates:
```bash
faas-cli new --list
```
It should show:
```
Languages available as templates:
- golang-flow
- golang-http
- golang-middleware
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

Or build them all at once (if combined into one YAML later):
```bash
faas-cli build -f stack.yml
```

### 4. Push (optional, if you use private registry)
If needed, push to your Docker registry:
```bash
faas-cli push -f user-info.yml
faas-cli push -f ride-history.yml
faas-cli push -f ride-recommend.yml
faas-cli push -f last-ride.yml
faas-cli push -f homepage.yml
```
(If using local Docker and Kubernetes integration, you may skip push.)

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

---
## Notes
If you're using FaasFlows, ensure the faas-provider is configured to support eventing.

The gateway URL inside .yml files points to http://127.0.0.1:8080 — adjust if necessary for remote clusters.

Template golang-flow is built for event composition and context passing.
