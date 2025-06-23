package application

import (
	"bosen/manifest"
	"bosen/pkg/runtime"
	"net/http"

	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	sglog "github.com/sourcegraph/log"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	// "go.opentelemetry.io/otel/trace"
)

type DiagnosticResource struct {
	log sglog.Logger
}

func NewDiagnosticResource(log sglog.Logger) *DiagnosticResource {
	return &DiagnosticResource{
		log: log.Scoped("diagnosticResource"),
	}
}

// func Trace(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
// 	// Configure the "http.route" for the HTTP instrumentation.
// 	attr := semconv.NewHTTPServer(nil).Route(route)
// 	path := httpRequest.URL.Path
// 	span := trace.SpanFromContext(req.Context())
// 	span.SetAttributes(attr)
//
// 	labeler, _ := LabelerFromContext(r.Context())
// 	labeler.Add(attr)
// 	chain.ProcessFilter(req, resp)
// }

func (r *DiagnosticResource) WebService() *restful.WebService {
	ws := new(restful.WebService).
		Path("/").
		Consumes(restful.MIME_OCTET).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/trace").To(r.trace).
		Doc("produce traceparent").
		Metadata(restfulspec.KeyOpenAPITags, []string{"opentelemetry", "trace"}).
		Writes(&manifest.Manifest{}).
		Returns(http.StatusOK, http.StatusText(http.StatusOK), &manifest.Manifest{}),
	)

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

	var fields []sglog.Field
	for key, value := range req.Request.Header {
		fields = append(fields, sglog.Strings(key, value))
	}

	r.log.Info("/manifest.json headers", sglog.Object("header", fields...))

	r.log.Info("diagnostic resource called")
	resp.WriteAsJson(manifest.Info())
}

func (r *DiagnosticResource) trace(req *restful.Request, resp *restful.Response) {
	// ctx, span := otel.Tracer(manifest.AppName).Start(req.Request.Context(), runtime.GetCurrentFunctionName())
	// defer span.End()

	// _span := trace.SpanFromContext(req.Request.Context())

	r.log.Info("/trace called")
	res, _ := otelhttp.Get(req.Request.Context(), "http://localhost:8080/manifest.json")

	defer res.Body.Close()

}

/* vim: set tabstop=4 softtabstop=4 shiftwidth=4 noexpandtab: */
