package response

import (
	"github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
)

// ErrorResponse wraps supplied error in order to return following JSON structure:
//
//	{
//		"error": {
//			"message": "string",
//			"code": "string"
//		}
//	}
type ErrorResponse struct {
	Err error `json:"error"`
}

func WriteError(res *restful.Response, status int, err error, contentType string) error {
	wrappedErr := ErrorResponse{err}

	if _err := res.WriteHeaderAndJson(status, &wrappedErr, contentType); _err != nil {
		zap.S().Errorf("unable to write error response: %w", err)
		return _err
	}

	return nil
}
