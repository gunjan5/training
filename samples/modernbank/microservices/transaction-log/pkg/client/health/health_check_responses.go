// Code generated by go-swagger; DO NOT EDIT.

package health

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// HealthCheckReader is a Reader for the HealthCheck structure.
type HealthCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HealthCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewHealthCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewHealthCheckOK creates a HealthCheckOK with default headers values
func NewHealthCheckOK() *HealthCheckOK {
	return &HealthCheckOK{}
}

/*HealthCheckOK handles this case with default header values.

OK
*/
type HealthCheckOK struct {
	/*Version of the microservice that responded
	 */
	Version string
}

func (o *HealthCheckOK) Error() string {
	return fmt.Sprintf("[POST /health][%d] healthCheckOK ", 200)
}

func (o *HealthCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header version
	o.Version = response.GetHeader("version")

	return nil
}
