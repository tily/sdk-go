// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package iotdataplane provides a client for AWS IoT Data Plane.
package iotdataplane

import (
	"github.com/tily/sdk-go/aws/awsutil"
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/private/protocol"
	"github.com/tily/sdk-go/private/protocol/restjson"
)

const opDeleteThingShadow = "DeleteThingShadow"

// DeleteThingShadowRequest generates a "aws/request.Request" representing the
// client's request for the DeleteThingShadow operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeleteThingShadow for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeleteThingShadow method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeleteThingShadowRequest method.
//    req, resp := client.DeleteThingShadowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/DeleteThingShadow
func (c *IoTDataPlane) DeleteThingShadowRequest(input *DeleteThingShadowInput) (req *request.Request, output *DeleteThingShadowOutput) {
	op := &request.Operation{
		Name:       opDeleteThingShadow,
		HTTPMethod: "DELETE",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &DeleteThingShadowInput{}
	}

	output = &DeleteThingShadowOutput{}
	req = c.newRequest(op, input, output)
	return
}

// DeleteThingShadow API operation for AWS IoT Data Plane.
//
// Deletes the thing shadow for the specified thing.
//
// For more information, see DeleteThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_DeleteThingShadow.html)
// in the AWS IoT Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS IoT Data Plane's
// API operation DeleteThingShadow for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   The specified resource does not exist.
//
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request is not valid.
//
//   * ErrCodeThrottlingException "ThrottlingException"
//   The rate exceeds the limit.
//
//   * ErrCodeUnauthorizedException "UnauthorizedException"
//   You are not authorized to perform this operation.
//
//   * ErrCodeServiceUnavailableException "ServiceUnavailableException"
//   The service is temporarily unavailable.
//
//   * ErrCodeInternalFailureException "InternalFailureException"
//   An unexpected error has occurred.
//
//   * ErrCodeMethodNotAllowedException "MethodNotAllowedException"
//   The specified combination of HTTP verb and URI is not supported.
//
//   * ErrCodeUnsupportedDocumentEncodingException "UnsupportedDocumentEncodingException"
//   The document encoding is not supported.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/DeleteThingShadow
func (c *IoTDataPlane) DeleteThingShadow(input *DeleteThingShadowInput) (*DeleteThingShadowOutput, error) {
	req, out := c.DeleteThingShadowRequest(input)
	err := req.Send()
	return out, err
}

const opGetThingShadow = "GetThingShadow"

// GetThingShadowRequest generates a "aws/request.Request" representing the
// client's request for the GetThingShadow operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See GetThingShadow for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the GetThingShadow method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the GetThingShadowRequest method.
//    req, resp := client.GetThingShadowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/GetThingShadow
func (c *IoTDataPlane) GetThingShadowRequest(input *GetThingShadowInput) (req *request.Request, output *GetThingShadowOutput) {
	op := &request.Operation{
		Name:       opGetThingShadow,
		HTTPMethod: "GET",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &GetThingShadowInput{}
	}

	output = &GetThingShadowOutput{}
	req = c.newRequest(op, input, output)
	return
}

// GetThingShadow API operation for AWS IoT Data Plane.
//
// Gets the thing shadow for the specified thing.
//
// For more information, see GetThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_GetThingShadow.html)
// in the AWS IoT Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS IoT Data Plane's
// API operation GetThingShadow for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request is not valid.
//
//   * ErrCodeResourceNotFoundException "ResourceNotFoundException"
//   The specified resource does not exist.
//
//   * ErrCodeThrottlingException "ThrottlingException"
//   The rate exceeds the limit.
//
//   * ErrCodeUnauthorizedException "UnauthorizedException"
//   You are not authorized to perform this operation.
//
//   * ErrCodeServiceUnavailableException "ServiceUnavailableException"
//   The service is temporarily unavailable.
//
//   * ErrCodeInternalFailureException "InternalFailureException"
//   An unexpected error has occurred.
//
//   * ErrCodeMethodNotAllowedException "MethodNotAllowedException"
//   The specified combination of HTTP verb and URI is not supported.
//
//   * ErrCodeUnsupportedDocumentEncodingException "UnsupportedDocumentEncodingException"
//   The document encoding is not supported.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/GetThingShadow
func (c *IoTDataPlane) GetThingShadow(input *GetThingShadowInput) (*GetThingShadowOutput, error) {
	req, out := c.GetThingShadowRequest(input)
	err := req.Send()
	return out, err
}

const opPublish = "Publish"

// PublishRequest generates a "aws/request.Request" representing the
// client's request for the Publish operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See Publish for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the Publish method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the PublishRequest method.
//    req, resp := client.PublishRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/Publish
func (c *IoTDataPlane) PublishRequest(input *PublishInput) (req *request.Request, output *PublishOutput) {
	op := &request.Operation{
		Name:       opPublish,
		HTTPMethod: "POST",
		HTTPPath:   "/topics/{topic}",
	}

	if input == nil {
		input = &PublishInput{}
	}

	output = &PublishOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(restjson.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// Publish API operation for AWS IoT Data Plane.
//
// Publishes state information.
//
// For more information, see HTTP Protocol (http://docs.aws.amazon.com/iot/latest/developerguide/protocols.html#http)
// in the AWS IoT Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS IoT Data Plane's
// API operation Publish for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeInternalFailureException "InternalFailureException"
//   An unexpected error has occurred.
//
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request is not valid.
//
//   * ErrCodeUnauthorizedException "UnauthorizedException"
//   You are not authorized to perform this operation.
//
//   * ErrCodeMethodNotAllowedException "MethodNotAllowedException"
//   The specified combination of HTTP verb and URI is not supported.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/Publish
func (c *IoTDataPlane) Publish(input *PublishInput) (*PublishOutput, error) {
	req, out := c.PublishRequest(input)
	err := req.Send()
	return out, err
}

const opUpdateThingShadow = "UpdateThingShadow"

// UpdateThingShadowRequest generates a "aws/request.Request" representing the
// client's request for the UpdateThingShadow operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See UpdateThingShadow for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the UpdateThingShadow method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the UpdateThingShadowRequest method.
//    req, resp := client.UpdateThingShadowRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/UpdateThingShadow
func (c *IoTDataPlane) UpdateThingShadowRequest(input *UpdateThingShadowInput) (req *request.Request, output *UpdateThingShadowOutput) {
	op := &request.Operation{
		Name:       opUpdateThingShadow,
		HTTPMethod: "POST",
		HTTPPath:   "/things/{thingName}/shadow",
	}

	if input == nil {
		input = &UpdateThingShadowInput{}
	}

	output = &UpdateThingShadowOutput{}
	req = c.newRequest(op, input, output)
	return
}

// UpdateThingShadow API operation for AWS IoT Data Plane.
//
// Updates the thing shadow for the specified thing.
//
// For more information, see UpdateThingShadow (http://docs.aws.amazon.com/iot/latest/developerguide/API_UpdateThingShadow.html)
// in the AWS IoT Developer Guide.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for AWS IoT Data Plane's
// API operation UpdateThingShadow for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeConflictException "ConflictException"
//   The specified version does not match the version of the document.
//
//   * ErrCodeRequestEntityTooLargeException "RequestEntityTooLargeException"
//   The payload exceeds the maximum size allowed.
//
//   * ErrCodeInvalidRequestException "InvalidRequestException"
//   The request is not valid.
//
//   * ErrCodeThrottlingException "ThrottlingException"
//   The rate exceeds the limit.
//
//   * ErrCodeUnauthorizedException "UnauthorizedException"
//   You are not authorized to perform this operation.
//
//   * ErrCodeServiceUnavailableException "ServiceUnavailableException"
//   The service is temporarily unavailable.
//
//   * ErrCodeInternalFailureException "InternalFailureException"
//   An unexpected error has occurred.
//
//   * ErrCodeMethodNotAllowedException "MethodNotAllowedException"
//   The specified combination of HTTP verb and URI is not supported.
//
//   * ErrCodeUnsupportedDocumentEncodingException "UnsupportedDocumentEncodingException"
//   The document encoding is not supported.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/UpdateThingShadow
func (c *IoTDataPlane) UpdateThingShadow(input *UpdateThingShadowInput) (*UpdateThingShadowOutput, error) {
	req, out := c.UpdateThingShadowRequest(input)
	err := req.Send()
	return out, err
}

// The input for the DeleteThingShadow operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/DeleteThingShadowRequest
type DeleteThingShadowInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteThingShadowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteThingShadowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteThingShadowInput"}
	if s.ThingName == nil {
		invalidParams.Add(request.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetThingName sets the ThingName field's value.
func (s *DeleteThingShadowInput) SetThingName(v string) *DeleteThingShadowInput {
	s.ThingName = &v
	return s
}

// The output from the DeleteThingShadow operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/DeleteThingShadowResponse
type DeleteThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	//
	// Payload is a required field
	Payload []byte `locationName:"payload" type:"blob" required:"true"`
}

// String returns the string representation
func (s DeleteThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteThingShadowOutput) GoString() string {
	return s.String()
}

// SetPayload sets the Payload field's value.
func (s *DeleteThingShadowOutput) SetPayload(v []byte) *DeleteThingShadowOutput {
	s.Payload = v
	return s
}

// The input for the GetThingShadow operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/GetThingShadowRequest
type GetThingShadowInput struct {
	_ struct{} `type:"structure"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s GetThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetThingShadowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *GetThingShadowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "GetThingShadowInput"}
	if s.ThingName == nil {
		invalidParams.Add(request.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetThingName sets the ThingName field's value.
func (s *GetThingShadowInput) SetThingName(v string) *GetThingShadowInput {
	s.ThingName = &v
	return s
}

// The output from the GetThingShadow operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/GetThingShadowResponse
type GetThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`
}

// String returns the string representation
func (s GetThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GetThingShadowOutput) GoString() string {
	return s.String()
}

// SetPayload sets the Payload field's value.
func (s *GetThingShadowOutput) SetPayload(v []byte) *GetThingShadowOutput {
	s.Payload = v
	return s
}

// The input for the Publish operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/PublishRequest
type PublishInput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`

	// The Quality of Service (QoS) level.
	Qos *int64 `location:"querystring" locationName:"qos" type:"integer"`

	// The name of the MQTT topic.
	//
	// Topic is a required field
	Topic *string `location:"uri" locationName:"topic" type:"string" required:"true"`
}

// String returns the string representation
func (s PublishInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PublishInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PublishInput"}
	if s.Topic == nil {
		invalidParams.Add(request.NewErrParamRequired("Topic"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPayload sets the Payload field's value.
func (s *PublishInput) SetPayload(v []byte) *PublishInput {
	s.Payload = v
	return s
}

// SetQos sets the Qos field's value.
func (s *PublishInput) SetQos(v int64) *PublishInput {
	s.Qos = &v
	return s
}

// SetTopic sets the Topic field's value.
func (s *PublishInput) SetTopic(v string) *PublishInput {
	s.Topic = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/PublishOutput
type PublishOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PublishOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishOutput) GoString() string {
	return s.String()
}

// The input for the UpdateThingShadow operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/UpdateThingShadowRequest
type UpdateThingShadowInput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	//
	// Payload is a required field
	Payload []byte `locationName:"payload" type:"blob" required:"true"`

	// The name of the thing.
	//
	// ThingName is a required field
	ThingName *string `location:"uri" locationName:"thingName" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s UpdateThingShadowInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateThingShadowInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *UpdateThingShadowInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "UpdateThingShadowInput"}
	if s.Payload == nil {
		invalidParams.Add(request.NewErrParamRequired("Payload"))
	}
	if s.ThingName == nil {
		invalidParams.Add(request.NewErrParamRequired("ThingName"))
	}
	if s.ThingName != nil && len(*s.ThingName) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("ThingName", 1))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetPayload sets the Payload field's value.
func (s *UpdateThingShadowInput) SetPayload(v []byte) *UpdateThingShadowInput {
	s.Payload = v
	return s
}

// SetThingName sets the ThingName field's value.
func (s *UpdateThingShadowInput) SetThingName(v string) *UpdateThingShadowInput {
	s.ThingName = &v
	return s
}

// The output from the UpdateThingShadow operation.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/iot-data-2015-05-28/UpdateThingShadowResponse
type UpdateThingShadowOutput struct {
	_ struct{} `type:"structure" payload:"Payload"`

	// The state information, in JSON format.
	Payload []byte `locationName:"payload" type:"blob"`
}

// String returns the string representation
func (s UpdateThingShadowOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s UpdateThingShadowOutput) GoString() string {
	return s.String()
}

// SetPayload sets the Payload field's value.
func (s *UpdateThingShadowOutput) SetPayload(v []byte) *UpdateThingShadowOutput {
	s.Payload = v
	return s
}
