package response

import (
	"github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
)

type ErrorResponse struct {
	Err error `json:"error"`
}

// WriteError wraps supplied error in order to return following JSON structure:
//
//	{
//		"error": {
//			"message": "string",
//			"code": "string"
//		}
//	}
func WriteError(res *restful.Response, status int, err error, contentType string) error {
	wrappedErr := ErrorResponse{err}

	if err := res.WriteHeaderAndJson(status, &wrappedErr, contentType); err != nil {
		zap.S().Errorf("unable to write error response: %w", err)
		return err
	}

	return nil
}
