package response

import (
	"github.com/emicklei/go-restful/v3"
	"go.uber.org/zap"
)

type SuccessResponse struct {
	Data any `json:"data"`
}

// WriteSuccess wraps supplied data in order to return following JSON structure:
//
//	{
//		"data": {
//			"message": "string",
//			"code": "string"
//		}
//	}
func WriteSuccess(res *restful.Response, status int, data any, contentType string) error {
	wrappedData := SuccessResponse{data}

	if err := res.WriteHeaderAndJson(status, &wrappedData, contentType); err != nil {
		zap.S().Errorf("unable to write success response: %w", err)
		return err
	}

	return nil
}
