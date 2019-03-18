package higo

import "net/http"

// IController .
type IController interface {
	Get(ctx Context) error

	Post(ctx Context) error

	Delete(ctx Context) error

	Head(ctx Context) error

	Options(ctx Context) error

	Patch(ctx Context) error

	Put(ctx Context) error
}

// Controller .
type Controller struct {
}

// Get .
func (c *Controller) Get(ctx Context) error {
	return NewHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
}

// Post .
func (c *Controller) Post(ctx Context) error {
	return NewHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
}

// Delete .
func (c *Controller) Delete(ctx Context) error {
	return NewHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
}

// Head .
func (c *Controller) Head(ctx Context) error {
	return NewHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
}

// Options .
func (c *Controller) Options(ctx Context) error {
	return NewHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
}

// Patch .
func (c *Controller) Patch(ctx Context) error {
	return NewHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
}

// Put .
func (c *Controller) Put(ctx Context) error {
	return NewHTTPError(http.StatusMethodNotAllowed, "Method Not Allowed")
}
