package httputils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/adroit-group/gote/pkg/logger"
	"github.com/stretchr/testify/assert"
)

type FailingEncoder struct{}

var _ json.Marshaler = FailingEncoder{}

func (f FailingEncoder) MarshalJSON() ([]byte, error) {
	return nil, errors.New("failed to marshal json")
}

func TestWriteJSONResponse(t *testing.T) {
	logger.SetupSlog("test", io.Discard)
	testCases := []struct {
		desc     string
		status   int
		response interface{}
		expected string
	}{
		{
			desc:   "error response",
			status: 500,
			response: ErrorResponse{
				Error:  "internal server error",
				Status: 500,
			},
			expected: "{\"error\":\"internal server error\",\"status\":500}\n",
		},
		{
			desc:     "json error 500",
			status:   500,
			response: FailingEncoder{},
			expected: "",
		},
	}
	for _, tC := range testCases {
		tC := tC
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			h := httptest.NewRecorder()

			WriteJSONResponse(h, tC.status, tC.response)

			assert.Equal(t, tC.expected, h.Body.String())
			assert.Equal(t, "application/json", h.Header().Get("Content-Type"))
			assert.Equal(t, tC.status, h.Code)
		})
	}
}
