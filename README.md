# AutoApproveDemo

A deliberately small, boring Go service used purely as a sandbox for
LiveReview's **Auto-Approve Gallery** demo. Every pull request in this repo
is a synthetic scenario from LiveReview's auto-approve policy test suite,
reviewed and triaged here by `livereviewbot`.

This is not a real product — don't depend on anything in it.

## Local development

```bash
go build ./...
go test ./...
go run .
```

Requires Go 1.22+. No external services needed — `Store` is in-memory only.
