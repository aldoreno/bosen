package application

import (
	"bosen/manifest"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
)

type DiagnosticResource struct{}

func NewDiagnosticResource() *DiagnosticResource {
	return &DiagnosticResource{}
}

func (r *DiagnosticResource) WebService() *restful.WebService {
	ws := new(restful.WebService).
		Path("/").
		Consumes(restful.MIME_OCTET).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("manifest.json").To(r.manifest).
		Doc("get application manifest").
		Metadata(restfulspec.KeyOpenAPITags, []string{"application", "manifest"}).
		Writes(&manifest.Manifest{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), &manifest.Manifest{}),
	)

	return ws
}

func (r *DiagnosticResource) manifest(req *restful.Request, resp *restful.Response) {
	resp.WriteAsJson(manifest.Info())
}
