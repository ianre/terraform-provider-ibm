# Logs Router V1 Workflow Clarification

## Understanding the Different Workflows

You've raised an important question about the workflow differences. Let me clarify what we've set up versus what V3 and the CloudEngineering wiki recommend.

## V3 Workflow (Platform Services)

### V3 Uses the "Shared SDK" Approach

**Key Characteristics:**
- API Package: `github.com/IBM/platform-services-go-sdk`
- Uses a **shared SDK repository** that contains multiple services
- Generates directly into a **terraform-provider-template** fork
- Template is a lightweight development/testing environment
- After testing, code is merged into the main `terraform-provider-ibm` repo

**V3 Workflow:**
```
1. Update API definition in cloud-api-docs
2. Generate Go SDK into platform-services-go-sdk (shared repo)
3. Generate Terraform into terraform-provider-template (test environment)
4. Test in template environment
5. Merge tested code into terraform-provider-ibm (main repo)
6. Submit PR to public terraform-provider-ibm
```

**V3 Repository Structure:**
```
github.com/IBM/platform-services-go-sdk/          # Shared SDK (multiple services)
  ├── logsroutingv3/                              # V3 service package
  ├── atrackerv3/                                 # Another service
  └── ...                                         # Other platform services

github.ibm.com/CloudEngineering/terraform-provider-template/  # Test environment
  └── ibm/service/logsroutingv3/                  # Generated Terraform code

github.com/IBM-Cloud/terraform-provider-ibm/      # Main public repo
  └── ibm/service/logsroutingv3/                  # Final merged code
```

## V1 Workflow (Standalone SDK) - What We've Set Up

### V1 Uses the "Standalone SDK" Approach

**Key Characteristics:**
- API Package: `github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0`
- Uses a **dedicated standalone SDK repository** (not shared)
- We've created a **custom development repo** that mimics the terraform-provider-template structure
- This repo serves the same purpose as terraform-provider-template but is specific to logs-router

**V1 Workflow (What We Proposed):**
```
1. Update API definition in cloud-api-docs
2. Generate Go SDK into logs-router-go-sdk (standalone repo)
3. Generate Terraform into terraform-provider-logs-router (custom dev/test environment)
4. Test in custom development environment
5. Copy tested code to terraform-provider-ibm fork (your personal fork)
6. Submit PR to public terraform-provider-ibm
```

**V1 Repository Structure (Our Setup):**
```
github.com/IBM/logs-router-go-sdk/                # Standalone SDK (single service)
  └── ibmcloudlogsroutingv0/                      # V1 service package

github.ibm.com/Observability/terraform-provider-logs-router/  # Custom dev/test environment
  └── ibm/service/ibmcloudlogsrouting/            # Generated Terraform code

github.com/ianre/terraform-provider-ibm/          # Your fork for PRs
  └── ibm/service/ibmcloudlogsrouting/            # Code copied from dev repo

github.com/IBM-Cloud/terraform-provider-ibm/      # Main public repo (PR target)
  └── ibm/service/ibmcloudlogsrouting/            # Final merged code
```

## Key Differences

| Aspect | V3 (Platform Services) | V1 (Standalone) - Our Setup |
|--------|------------------------|------------------------------|
| **SDK Repository** | Shared (platform-services-go-sdk) | Standalone (logs-router-go-sdk) |
| **SDK Package Path** | `github.com/IBM/platform-services-go-sdk` | `github.com/IBM/logs-router-go-sdk/ibmcloudlogsroutingv0` |
| **Test Environment** | terraform-provider-template (official) | terraform-provider-logs-router (custom) |
| **Test Repo Owner** | CloudEngineering | Observability |
| **Generation Target** | Template repo | Custom dev repo |
| **Merge Process** | Template → terraform-provider-ibm | Custom dev → Your fork → terraform-provider-ibm |

## Why We Set It Up This Way

### Reasons for Custom Development Repo

1. **Standalone SDK**: V1 uses a standalone SDK, not the shared platform-services-go-sdk
2. **Team Ownership**: The Observability team owns the logs-router service
3. **Iterative Development**: Allows testing API changes before submitting to public repo
4. **Isolation**: Keeps V1 development separate from V3 and other services
5. **Flexibility**: Can customize the development environment for logs-router specific needs

### The Custom Repo Serves the Same Purpose as terraform-provider-template

The `terraform-provider-logs-router` repo we set up is functionally equivalent to `terraform-provider-template`:
- ✅ Contains the Terraform provider framework
- ✅ Allows generation and testing of Terraform code
- ✅ Provides isolation for development
- ✅ Enables iterative API definition changes
- ✅ Code is tested here before moving to public repo

## Detailed V1 Workflow Steps

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

### Step 3: Generate Terraform Provider
```bash
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
./scripts/generate-terraform.sh

# This generates into: /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router
```

### Step 4: Test in Development Environment
```bash
cd /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router

# Update go.mod to use local SDK
echo 'replace github.com/IBM/logs-router-go-sdk => /Users/ianre/go/src/github.com/IBM/logs-router-go-sdk' >> go.mod
go mod tidy

# Build and test
make build
make testacc
```

### Step 5: Copy to Your Fork (When Ready)
```bash
# Copy generated files from dev repo to your fork
cp -r /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router/ibm/service/ibmcloudlogsrouting \
      /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm/ibm/service/

cp -r /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router/website/docs/r/logs_router_* \
      /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm/website/docs/r/

cp -r /Users/ianre/go/src/github.ibm.com/Observability/terraform-provider-logs-router/website/docs/d/logs_router_* \
      /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm/website/docs/d/

# Update go.mod in your fork
cd /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm
# Add dependency on logs-router-go-sdk
go get github.com/IBM/logs-router-go-sdk@latest
go mod tidy
```

### Step 6: Submit PR
```bash
cd /Users/ianre/go/src/github.com/ianre/terraform-provider-ibm
git checkout -b add-logs-router-v1-support
git add .
git commit -m "Add Logs Router V1 Terraform provider support"
git push origin add-logs-router-v1-support

# Create PR to github.com/IBM-Cloud/terraform-provider-ibm
```

## Comparison with CloudEngineering Wiki Workflow

### Wiki Recommends (for Platform Services):
1. Use terraform-provider-template
2. Generate into template
3. Test in template
4. Merge to terraform-provider-ibm

### We Set Up (for Standalone Service):
1. Use custom terraform-provider-logs-router repo
2. Generate into custom repo
3. Test in custom repo
4. Copy to your fork
5. PR to terraform-provider-ibm

**Both approaches achieve the same goal**: Test generated code in isolation before merging to the main public repo.

## Why Not Use terraform-provider-template for V1?

We could have used terraform-provider-template, but:

1. **SDK Package Difference**: Template expects platform-services-go-sdk, V1 uses logs-router-go-sdk
2. **Team Ownership**: Observability team prefers their own repo
3. **Service Isolation**: V1 and V3 are different APIs for the same service
4. **Customization**: Custom repo allows logs-router specific configurations
5. **Long-term Maintenance**: Easier to maintain separate development environments

## Summary

**What we've set up is valid and follows IBM's principles**, just adapted for a standalone SDK:

✅ **Isolated development environment** (terraform-provider-logs-router instead of terraform-provider-template)
✅ **Generate and test before merging** (same principle as wiki)
✅ **Proper SDK dependency management** (standalone SDK instead of shared)
✅ **Clear workflow for API changes** (documented in scripts and guides)
✅ **PR process to public repo** (same as wiki recommends)

The key difference is:
- **V3**: Shared SDK → Template → Public Repo
- **V1**: Standalone SDK → Custom Dev Repo → Your Fork → Public Repo

Both are valid approaches. V1's approach adds one extra step (your fork) but provides more control and isolation for the Observability team's development process.