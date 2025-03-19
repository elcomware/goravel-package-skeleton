package tools

type PackageRoutes struct {
	RouteFileNames []string
}

// AddRoute adds a single route file name to the list.
func (r *PackageRoutes) AddRoute(routeFileName string) *PackageRoutes {
	r.RouteFileNames = append(r.RouteFileNames, routeFileName)
	return r
}

// AddRoutes adds multiple route file names to the list.
func (r *PackageRoutes) AddRoutes(routeFileNames ...interface{}) *PackageRoutes {
	for _, route := range routeFileNames {
		switch v := route.(type) {
		case string:
			r.RouteFileNames = append(r.RouteFileNames, v)
		case []string:
			r.RouteFileNames = append(r.RouteFileNames, v...)
		}
	}
	return r
}
