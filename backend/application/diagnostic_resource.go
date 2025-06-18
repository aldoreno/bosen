package application

import (
	"bosen/manifest"
	"bosen/pkg/runtime"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	sglog "github.com/sourcegraph/log"
	"go.opentelemetry.io/otel"
)

type DiagnosticResource struct {
	log sglog.Logger
}

func NewDiagnosticResource(log sglog.Logger) *DiagnosticResource {
	return &DiagnosticResource{
		log: log.Scoped("diagnosticResource", "provides build time information of the application"),
	}
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
	_, span := otel.Tracer(manifest.AppName).Start(req.Request.Context(), runtime.GetCurrentFunctionName())
	defer span.End()

	r.log.Info("diagnostic resource called")
	resp.WriteAsJson(manifest.Info())
}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
