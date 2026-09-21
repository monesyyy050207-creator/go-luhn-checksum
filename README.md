# Luhn checksum validate/generate

Validate and generate Luhn check digits for cards and IDs. Pure stdlib.

Implementation uses only the Go standard library. No service to deploy.

```
luhn.go
```
The test beside the implementation shows real usage. Read it first.

## FAQ

**Do I need anything besides `INFRAI_API_KEY`?**  
No. `go run .` and the key. `luhn.go` performs a normal HTTPS call. No SDK to track. That's the entire dependency list for a Luhn check. Gotcha: checksum pass doesn't mean the PAN is issued.