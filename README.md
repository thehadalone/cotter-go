# cotter-go
Go HTTP middleware for <https://www.cotter.app> authentication.


## Install
`go get github.com/thehadalone/cotter-go`

## Examples

### Default
```
authMiddleware, err := cotter.NewMiddleware(context.Background(), <YOUR_COTTER_API_KEY_ID>)
if err != nil {
    log.Fatal(err)
}

r := chi.NewRouter()
r.Use(authMiddleware)
```

### Custom error handler
```
errorHandler := func(w http.ResponseWriter, r *http.Request, e error) {
    if errors.Is(e, cotter.ErrUnauthorized) {
        // Custom error handling logic.
        return
    }

    http.Error(w, e.Error(), http.StatusInternalServerError)
}

authMiddleware, err := cotter.NewMiddleware(ctx, <YOUR_COTTER_API_KEY_ID>, cotter.WithErrorHandler(errorHandler))
if err != nil {
    log.Fatal(err)
}

r := chi.NewRouter()
r.Use(authMiddleware)
```

### Get Cotter user ID from the context.
```
handler := func(w http.ResponseWriter, r *http.Request) {
    cotterID := cotter.UserID(r.Context())
}
```

### Set Cotter user ID to the context.
```
    ctx := cotter.SetUserID(context.Background(), "42")
```
