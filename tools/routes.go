package tools

type RoutesTools struct {
	FileNames []string
}

func NewPackageRoutes() *RoutesTools {
	return &RoutesTools{}
}

// AddRoute adds a single route file name to the list.
func (r *RoutesTools) AddRoute(routeFileName string) *RoutesTools {
	r.FileNames = append(r.FileNames, routeFileName)
	return r
}

// AddRoutes adds multiple route file names to the list.
func (r *RoutesTools) AddRoutes(routeFileNames ...interface{}) *RoutesTools {
	for _, route := range routeFileNames {
		switch v := route.(type) {
		case string:
			r.FileNames = append(r.FileNames, v)
		case []string:
			r.FileNames = append(r.FileNames, v...)
		}
	}
	return r
}
