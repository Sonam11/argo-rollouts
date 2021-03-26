package kapcom

const (
	GroupName = "contour.heptio.com"

	IngressRouteKind     string = "IngressRoute"
	IngressRouteSingular string = "ingressroute"
	IngressRoutePlural   string = "ingressroutes"
	IngressRouteFullName string = IngressRoutePlural + "." + GroupName
)
