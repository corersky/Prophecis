// Code generated by go-swagger; DO NOT EDIT.

package auths

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// AdminUserCheckHandlerFunc turns a function with the right signature into a admin user check handler
type AdminUserCheckHandlerFunc func(AdminUserCheckParams) middleware.Responder

// Handle executing the request and returning a response
func (fn AdminUserCheckHandlerFunc) Handle(params AdminUserCheckParams) middleware.Responder {
	return fn(params)
}

// AdminUserCheckHandler interface for that can handle valid admin user check params
type AdminUserCheckHandler interface {
	Handle(AdminUserCheckParams) middleware.Responder
}

// NewAdminUserCheck creates a new http.Handler for the admin user check operation
func NewAdminUserCheck(ctx *middleware.Context, handler AdminUserCheckHandler) *AdminUserCheck {
	return &AdminUserCheck{Context: ctx, Handler: handler}
}

/*AdminUserCheck swagger:route GET /cc/v1/auth/access/users/{adminUsername}/users/{username} Auths adminUserCheck

auth by adminUsername and username.

Optional extended description in Markdown.

*/
type AdminUserCheck struct {
	Context *middleware.Context
	Handler AdminUserCheckHandler
}

func (o *AdminUserCheck) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewAdminUserCheckParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}