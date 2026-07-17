# Terraform Provider Infoblox

A Terraform provider with a single interface to manage resources across Infoblox backends — NIOS (on-prem) and UDDI (cloud). 

## Architecture

```
Terraform Resource (infoblox schema)
        │
        ▼
   Core Module
   ┌─────────────────────────────────┐
   │  Model    - Infoblox structs    │
   │  Mapper   - Field translations  │
   │  Service  - Backend dispatching │
   └─────────────────────────────────┘
        │
        ├──────────────┐
        ▼              ▼
   NIOS              UDDI
   (on-prem)         (cloud)
```

The core module provides common data structures, maps unified field names to backend-specific SDK fields, and routes each request to the active backend (transforming the response back to the unified model).

## Supported Backends

| Backend | Description | Status |
|---------|-------------|--------|
| NIOS | On-premise WAPI | Supported |
| UDDI | Cloud | Supported |

