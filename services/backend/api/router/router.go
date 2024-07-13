package router

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	s "students/api/service"
	"students/db"
	"students/limit"
)

type Router struct {
	routes map[string]map[string]http.HandlerFunc
}

func NewRouter(db *db.Queries) *Router {
	router := &Router{
		routes: make(map[string]map[string]http.HandlerFunc),
	}

	service := s.NewService(db)

	router.addRoute("GET", "/message/{id}", limit.RateLimiter(service.GetProduct).(http.HandlerFunc))
	router.addRoute("GET", "/message", limit.RateLimiter(service.ListAllProducts).(http.HandlerFunc))
	router.addRoute("POST", "/message", limit.RateLimiter(service.CreateProduct).(http.HandlerFunc))
	/*router.addRoute("PUT", "/update_message", limit.RateLimiter(service.UpdateProduct).(http.HandlerFunc))
	router.addRoute("DELETE", "/delete_message/{id}", limit.RateLimiter(service.DeleteProduct).(http.HandlerFunc)) */

	return router
}

func (r *Router) addRoute(method string, path string, handler http.HandlerFunc) {
	if r.routes[method] == nil {
		r.routes[method] = make(map[string]http.HandlerFunc)
	}
	r.routes[method][path] = handler
	fmt.Printf("New route added to the router: Method=%s, Path=%s\n", method, path)
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path

	fmt.Printf("Handling request: Method=%s, Path=%s\n", method, path)
	for route, handler := range r.routes[method] {
		if match, params := matchRoute(route, path); match {
			ctx := req.Context()
			for k, v := range params {
				fmt.Printf("Route parameter: %s = %s\n", k, v)
				ctx = context.WithValue(ctx, k, v)
			}
			fmt.Printf("Matched route: %s\n", route)
			fmt.Printf("Calling handler for Method=%s, Path=%s handler=%v ctx=%v\n", method, path, handler, ctx)
			handler.ServeHTTP(w, req.WithContext(ctx))
			return
		}
	}
	fmt.Printf("Route not found for Method=%s, Path=%s\n", method, path)
	http.NotFound(w, req)
}

func matchRoute(route, path string) (bool, map[string]string) {
	routeMtx := strings.Split(route, "/")
	pathMtx := strings.Split(path, "/")

	if len(routeMtx) != len(pathMtx) {
		return false, nil
	}

	params := make(map[string]string)
	for i := range routeMtx {
		if strings.HasPrefix(routeMtx[i], "{") && strings.HasSuffix(routeMtx[i], "}") {
			paramName := routeMtx[i][1 : len(routeMtx[i])-1]
			params[paramName] = pathMtx[i]
		} else if routeMtx[i] != pathMtx[i] {
			return false, nil
		}
	}
	return true, params
}