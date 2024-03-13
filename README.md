# gnosis-safe-go-binding

This repository contains go-binding supporting calling specific contracts in ethereum network.

## Requirements

- [jq v1.7.1](https://jqlang.github.io/jq/download/)
- [abigen 1.13.5](https://geth.ethereum.org/docs/tools/abigen)

## Process

1. Build .abi

```
jq -r ".abi" contracts/contracts/Deps_V1_4_1.sol/CompatibilityFallbackHandler_SV1_4_1.json  > abi/CompatibilityFallbackHandler_SV1_4_1.abi
```

2. Build go-binding modules in [go-binding directory](/go-binding)

```
abigen --abi abi/CompatibilityFallbackHandler_SV1_4_1.abi --pkg compatibilityfallbackhandler --out go-binding/compatibility-fallback-handler-v1_4_0/CompatibilityFallbackHandler.go
```
