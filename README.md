# Terraform Provider Unified (POC)

A unified Terraform provider that provides a single interface to manage resources across multiple Infoblox backends - NIOS (on-prem) and UDDI (cloud).

## Overview

This provider abstracts away backend differences, allowing users to write Terraform configurations once and deploy to different Infoblox environments by simply changing the backend configuration.

### Architecture

```
Terraform Resource (unified schema)
        │
        ▼
   Core Module
   ┌─────────────────────────────────┐
   │  Model    - Unified structs     │
   │  Mapper   - Field translations  │
   │  Service  - Backend dispatching │
   └─────────────────────────────────┘
        │
        ├──────────────┐
        ▼              ▼
   NIOS SDK       UDDI SDK
   (on-prem)       (cloud)
```

### Core Module

The core module handles:
- **Unified Model**: Common data structures for resources across backends
- **Field Mapping**: Translates unified field names to backend-specific SDK field names
- **Service Layer**: Routes requests to the appropriate backend and transforms responses back to the unified format

### How It Works

1. User defines a resource using the unified schema (e.g., `unified_dns_record_a`)
2. The provider determines the active backend from configuration (NIOS or UDDI)
3. Core service maps the unified request to backend-specific SDK types
4. Backend SDK makes the actual API call
5. Response is mapped back to the unified model

## Supported Backends

| Backend | Description | Status |
|---------|-------------|--------|
| NIOS | On-premise WAPI | Supported |
| UDDI | BloxOne Cloud | Supported |

## TODOs

The following items are pending for full implementation:

- **Function Calls**: NIOS WAPI function calls (like `nextavailableip`) are not supported in the current POC.

- **ExtAttrs Prerequisites**: For NIOS extensible attributes, we need to add pre-requisite validation in the provider.

- **NIOS Version Upgrade**: Current POC is built against NIOS 9.0.6. Need to upgrade to NIOS 9.1.0 for production use.

- **Import State**: Importing of resources are not supported in the current POC.
