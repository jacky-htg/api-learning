# Rebel Services

Context Part II.

We are now using middleware chains and a custom handler. We are also making greater use of context for passing values. As we continue to expand this program we will find relying on request.Context and request.WithContext to be annoying.

Tasks:
- Add ctx context.Context as the first argument in the api.Handler type
- Make the adapter function in libraries/api/app.go pass down the first context which was derived from r.Context().
- Anything else using r.Context() should instead use the passed ctx.

## File Changes :
- libraries/api/app.go
- middleware/error.go
- middleware/logger.go
- middleware/metrics.go
- middleware/auth.go
- controllers/access.go
- controllers/auths.go
- controllers/checks.go
- controllers/roles.go
- controllers/users.go

## New File :

## Adding Dependency :