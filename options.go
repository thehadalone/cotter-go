package cotter

type options struct {
	errorHandler ErrorHandler
}

// Option for the middleware.
type Option func(*options)

// WithErrorHandler sets custom error handler for the middleware.
func WithErrorHandler(errorHandler ErrorHandler) Option {
	return func(o *options) {
		o.errorHandler = errorHandler
	}
}
