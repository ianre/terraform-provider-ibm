# Simplified Logs Router V1 Workflow - Direct to Fork

## Overview

Based on your feedback, we've simplified the workflow to generate Terraform provider code **directly into your public fork** of `terraform-provider-ibm`. This eliminates the intermediate copy step and streamlines the process.

## Simplified Workflow

### Two-Repo Approach (Recommended)

```
1. Update API Definition
   ↓
2. Generate Go SDK → logs-router-go-sdk (standalone repo)
   ↓
3. Generate Terraform → terraform-provider-ibm fork (YOUR FORK - direct generation)
   ↓
4. Test in fork
   ↓
5. Submit PR → IBM-Cloud/terraform-provider-ibm (public repo)
```

## Repository Roles

### 1. `/Users/ianre/go/src/github.com/IBM/logs-router-go-sdk`
- **Purpose**: Standalone Go SDK for Logs Router V1
- **Updated by**: SDK generation script
- **Used by**: Terraform provider (imported as dependency)

### 2. `/Users/ianre/go/src/github.com/ianre/terraform-provider-ibm`
- **Purpose**: Your fork of the public terraform-provider-ibm
- **Updated by**: Direct Terraform generation (NEW!)
- **Used for**: Testing and PR submission

### 3. `/Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router`
- **Purpose**: Development repo with helper scripts
- **Contains**: Generation scripts, documentation, utilities
- **Does NOT receive generated code** (optional testing environment)

## Complete Workflow Steps

### Step 1: Update API Definition
```bash
# Edit the API definition
vim /Users/ianre/go/src/github.ibm.com/cloud-api-docs/logs-router/logs-router-service-api.v1.json

# Add/update Terraform annotations if needed
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/add-terraform-annotations.sh
```

### Step 2: Generate Go SDK
```bash
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/generate-go-sdk.sh

# This generates into: /Users/ianre/go/src/github.com/IBM/logs-router-go-sdk
```

### Step 3: Generate Terraform Directly to Fork (NEW!)
```bash
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/generate-terraform-to-fork.sh

# This generates DIRECTLY into: /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm
```

### Step 4: Update Dependencies and Test
```bash
cd /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm

# Add/update the logs-router-go-sdk dependency
go get github.com/IBM/logs-router-go-sdk@latest
go mod tidy

# Build and test
make build
make testacc
```

### Step 5: Commit and Create PR
```bash
cd /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm

# Create a branch
git checkout -b add-logs-router-v1-support

# Review changes
git status
git diff

# Commit
git add .
git commit -m "Add Logs Router V1 Terraform provider support

- Add ibm_logs_router_tenant resource
- Add ibm_logs_router_tenants data source
- Add ibm_logs_router_targets data source
- Add examples and documentation"

# Push to your fork
git push origin add-logs-router-v1-support

# Create PR on GitHub to IBM-Cloud/terraform-provider-ibm
```

## New Script: generate-terraform-to-fork.sh

This new script generates Terraform code directly into your fork:

**Location**: `/Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router/scripts/generate-terraform-to-fork.sh`

**What it does**:
1. Validates API definition exists
2. Validates your fork exists
3. Checks git status (warns if uncommitted changes)
4. Generates Terraform code directly into your fork
5. Shows what files were generated/modified
6. Provides next steps

**Usage**:
```bash
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/generate-terraform-to-fork.sh
```

## Files Generated in Your Fork

When you run `generate-terraform-to-fork.sh`, it creates/updates:

```
/Users/ianre/go/src/github.com/ianre/terraform-provider-ibm/
├── ibm/service/ibmcloudlogsrouting/
│   ├── resource_ibm_logs_router_tenant.go
│   ├── resource_ibm_logs_router_tenant_test.go
│   ├── data_source_ibm_logs_router_tenants.go
│   ├── data_source_ibm_logs_router_tenants_test.go
│   ├── data_source_ibm_logs_router_targets.go
│   ├── data_source_ibm_logs_router_targets_test.go
│   └── README.md
├── website/docs/
│   ├── r/logs_router_tenant.html.markdown
│   └── d/
│       ├── logs_router_tenants.html.markdown
│       └── logs_router_targets.html.markdown
└── examples/ibm-ibm-cloud-logs-routing/
    ├── main.tf
    ├── outputs.tf
    ├── variables.tf
    ├── versions.tf
    └── README.md
```

## Comparison: Old vs New Workflow

### Old Workflow (3 Repos)
```
API → SDK → Dev Repo → Copy → Your Fork → PR
```
- More steps
- Manual copy required
- Dev repo receives generated code
- Extra validation step

### New Workflow (2 Repos)
```
API → SDK → Your Fork → PR
```
- Fewer steps
- Direct generation
- No manual copy
- Faster iteration

## Benefits of Direct Generation

1. **Simpler**: One less step in the workflow
2. **Faster**: No manual copying of files
3. **Cleaner**: Generated code goes directly where it's needed
4. **Less Error-Prone**: No risk of forgetting to copy files
5. **Git-Friendly**: All changes tracked in your fork from the start

## Optional: Using Dev Repo for Testing

You can still use the dev repo (`terraform-provider-logs-router`) for testing if you want:

```bash
# Generate to dev repo for testing
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/generate-terraform.sh

# Test in dev repo
make build
make testacc

# When satisfied, generate to fork
./scripts/generate-terraform-to-fork.sh
```

This gives you a safe testing environment before generating to your fork.

## Quick Reference Commands

### Full Workflow (API Change → PR)
```bash
# 1. Update API annotations
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/add-terraform-annotations.sh

# 2. Generate SDK
./scripts/generate-go-sdk.sh

# 3. Generate Terraform to fork
./scripts/generate-terraform-to-fork.sh

# 4. Test in fork
cd /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm
go get github.com/IBM/logs-router-go-sdk@latest
go mod tidy
make build
make testacc

# 5. Create PR
git checkout -b add-logs-router-v1-support
git add .
git commit -m "Add Logs Router V1 Terraform provider support"
git push origin add-logs-router-v1-support
```

### Regenerate After API Changes
```bash
# Just regenerate SDK and Terraform
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/generate-go-sdk.sh
./scripts/generate-terraform-to-fork.sh

# Test
cd /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm
go mod tidy
make build
```

## Summary

The simplified workflow:
- ✅ Generates directly to your fork
- ✅ Eliminates manual copy step
- ✅ Maintains all helper scripts in dev repo
- ✅ Keeps SDK in standalone repo
- ✅ Streamlines the PR process

This matches the spirit of the CloudEngineering wiki workflow while adapting it for a standalone SDK architecture.