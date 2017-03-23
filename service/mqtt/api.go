// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package mqtt provides a client for mqtt.
package mqtt

import (
	"github.com/tily/sdk-go/aws/awsutil"
	"github.com/tily/sdk-go/aws/request"
	"github.com/tily/sdk-go/private/protocol"
	"github.com/tily/sdk-go/private/protocol/query"
)

const opCreateMQTTInstance = "CreateMQTTInstance"

// CreateMQTTInstanceRequest generates a "aws/request.Request" representing the
// client's request for the CreateMQTTInstance operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See CreateMQTTInstance for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the CreateMQTTInstance method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the CreateMQTTInstanceRequest method.
//    req, resp := client.CreateMQTTInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTInstance
func (c *mqtt) CreateMQTTInstanceRequest(input *CreateMQTTInstanceInput) (req *request.Request, output *CreateMQTTInstanceOutput) {
	op := &request.Operation{
		Name:       opCreateMQTTInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMQTTInstanceInput{}
	}

	output = &CreateMQTTInstanceOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// CreateMQTTInstance API operation for mqtt.
//
// CreateMQTTInstance
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation CreateMQTTInstance for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTInstance
func (c *mqtt) CreateMQTTInstance(input *CreateMQTTInstanceInput) (*CreateMQTTInstanceOutput, error) {
	req, out := c.CreateMQTTInstanceRequest(input)
	err := req.Send()
	return out, err
}

const opCreateMQTTUser = "CreateMQTTUser"

// CreateMQTTUserRequest generates a "aws/request.Request" representing the
// client's request for the CreateMQTTUser operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See CreateMQTTUser for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the CreateMQTTUser method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the CreateMQTTUserRequest method.
//    req, resp := client.CreateMQTTUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTUser
func (c *mqtt) CreateMQTTUserRequest(input *CreateMQTTUserInput) (req *request.Request, output *CreateMQTTUserOutput) {
	op := &request.Operation{
		Name:       opCreateMQTTUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateMQTTUserInput{}
	}

	output = &CreateMQTTUserOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// CreateMQTTUser API operation for mqtt.
//
// CreateMQTTUser
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation CreateMQTTUser for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTUser
func (c *mqtt) CreateMQTTUser(input *CreateMQTTUserInput) (*CreateMQTTUserOutput, error) {
	req, out := c.CreateMQTTUserRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteMQTTInstance = "DeleteMQTTInstance"

// DeleteMQTTInstanceRequest generates a "aws/request.Request" representing the
// client's request for the DeleteMQTTInstance operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeleteMQTTInstance for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeleteMQTTInstance method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeleteMQTTInstanceRequest method.
//    req, resp := client.DeleteMQTTInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTInstance
func (c *mqtt) DeleteMQTTInstanceRequest(input *DeleteMQTTInstanceInput) (req *request.Request, output *DeleteMQTTInstanceOutput) {
	op := &request.Operation{
		Name:       opDeleteMQTTInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMQTTInstanceInput{}
	}

	output = &DeleteMQTTInstanceOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// DeleteMQTTInstance API operation for mqtt.
//
// DeleteMQTTInstance
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation DeleteMQTTInstance for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTInstance
func (c *mqtt) DeleteMQTTInstance(input *DeleteMQTTInstanceInput) (*DeleteMQTTInstanceOutput, error) {
	req, out := c.DeleteMQTTInstanceRequest(input)
	err := req.Send()
	return out, err
}

const opDeleteMQTTUser = "DeleteMQTTUser"

// DeleteMQTTUserRequest generates a "aws/request.Request" representing the
// client's request for the DeleteMQTTUser operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DeleteMQTTUser for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DeleteMQTTUser method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DeleteMQTTUserRequest method.
//    req, resp := client.DeleteMQTTUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTUser
func (c *mqtt) DeleteMQTTUserRequest(input *DeleteMQTTUserInput) (req *request.Request, output *DeleteMQTTUserOutput) {
	op := &request.Operation{
		Name:       opDeleteMQTTUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteMQTTUserInput{}
	}

	output = &DeleteMQTTUserOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// DeleteMQTTUser API operation for mqtt.
//
// DeleteMQTTUser
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation DeleteMQTTUser for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTUser
func (c *mqtt) DeleteMQTTUser(input *DeleteMQTTUserInput) (*DeleteMQTTUserOutput, error) {
	req, out := c.DeleteMQTTUserRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeMQTTInstances = "DescribeMQTTInstances"

// DescribeMQTTInstancesRequest generates a "aws/request.Request" representing the
// client's request for the DescribeMQTTInstances operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DescribeMQTTInstances for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DescribeMQTTInstances method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DescribeMQTTInstancesRequest method.
//    req, resp := client.DescribeMQTTInstancesRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTInstances
func (c *mqtt) DescribeMQTTInstancesRequest(input *DescribeMQTTInstancesInput) (req *request.Request, output *DescribeMQTTInstancesOutput) {
	op := &request.Operation{
		Name:       opDescribeMQTTInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMQTTInstancesInput{}
	}

	output = &DescribeMQTTInstancesOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// DescribeMQTTInstances API operation for mqtt.
//
// DescribeMQTTInstances
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation DescribeMQTTInstances for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTInstances
func (c *mqtt) DescribeMQTTInstances(input *DescribeMQTTInstancesInput) (*DescribeMQTTInstancesOutput, error) {
	req, out := c.DescribeMQTTInstancesRequest(input)
	err := req.Send()
	return out, err
}

const opDescribeMQTTUsers = "DescribeMQTTUsers"

// DescribeMQTTUsersRequest generates a "aws/request.Request" representing the
// client's request for the DescribeMQTTUsers operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See DescribeMQTTUsers for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the DescribeMQTTUsers method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the DescribeMQTTUsersRequest method.
//    req, resp := client.DescribeMQTTUsersRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTUsers
func (c *mqtt) DescribeMQTTUsersRequest(input *DescribeMQTTUsersInput) (req *request.Request, output *DescribeMQTTUsersOutput) {
	op := &request.Operation{
		Name:       opDescribeMQTTUsers,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeMQTTUsersInput{}
	}

	output = &DescribeMQTTUsersOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// DescribeMQTTUsers API operation for mqtt.
//
// DescribeMQTTUsers
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation DescribeMQTTUsers for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTUsers
func (c *mqtt) DescribeMQTTUsers(input *DescribeMQTTUsersInput) (*DescribeMQTTUsersOutput, error) {
	req, out := c.DescribeMQTTUsersRequest(input)
	err := req.Send()
	return out, err
}

const opModifyMQTTInstance = "ModifyMQTTInstance"

// ModifyMQTTInstanceRequest generates a "aws/request.Request" representing the
// client's request for the ModifyMQTTInstance operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ModifyMQTTInstance for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ModifyMQTTInstance method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ModifyMQTTInstanceRequest method.
//    req, resp := client.ModifyMQTTInstanceRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTInstance
func (c *mqtt) ModifyMQTTInstanceRequest(input *ModifyMQTTInstanceInput) (req *request.Request, output *ModifyMQTTInstanceOutput) {
	op := &request.Operation{
		Name:       opModifyMQTTInstance,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyMQTTInstanceInput{}
	}

	output = &ModifyMQTTInstanceOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// ModifyMQTTInstance API operation for mqtt.
//
// ModifyMQTTInstance
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation ModifyMQTTInstance for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTInstance
func (c *mqtt) ModifyMQTTInstance(input *ModifyMQTTInstanceInput) (*ModifyMQTTInstanceOutput, error) {
	req, out := c.ModifyMQTTInstanceRequest(input)
	err := req.Send()
	return out, err
}

const opModifyMQTTUser = "ModifyMQTTUser"

// ModifyMQTTUserRequest generates a "aws/request.Request" representing the
// client's request for the ModifyMQTTUser operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See ModifyMQTTUser for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the ModifyMQTTUser method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the ModifyMQTTUserRequest method.
//    req, resp := client.ModifyMQTTUserRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTUser
func (c *mqtt) ModifyMQTTUserRequest(input *ModifyMQTTUserInput) (req *request.Request, output *ModifyMQTTUserOutput) {
	op := &request.Operation{
		Name:       opModifyMQTTUser,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &ModifyMQTTUserInput{}
	}

	output = &ModifyMQTTUserOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// ModifyMQTTUser API operation for mqtt.
//
// ModifyMQTTUser
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation ModifyMQTTUser for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTUser
func (c *mqtt) ModifyMQTTUser(input *ModifyMQTTUserInput) (*ModifyMQTTUserOutput, error) {
	req, out := c.ModifyMQTTUserRequest(input)
	err := req.Send()
	return out, err
}

const opPublishMQTTMessage = "PublishMQTTMessage"

// PublishMQTTMessageRequest generates a "aws/request.Request" representing the
// client's request for the PublishMQTTMessage operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See PublishMQTTMessage for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the PublishMQTTMessage method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the PublishMQTTMessageRequest method.
//    req, resp := client.PublishMQTTMessageRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI//PublishMQTTMessage
func (c *mqtt) PublishMQTTMessageRequest(input *PublishMQTTMessageInput) (req *request.Request, output *PublishMQTTMessageOutput) {
	op := &request.Operation{
		Name:       opPublishMQTTMessage,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &PublishMQTTMessageInput{}
	}

	output = &PublishMQTTMessageOutput{}
	req = c.newRequest(op, input, output)
	req.Handlers.Unmarshal.Remove(query.UnmarshalHandler)
	req.Handlers.Unmarshal.PushBackNamed(protocol.UnmarshalDiscardBodyHandler)
	return
}

// PublishMQTTMessage API operation for mqtt.
//
// PublishMQTTMessage
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for mqtt's
// API operation PublishMQTTMessage for usage and error information.
// Please also see https://docs.aws.amazon.com/goto/WebAPI//PublishMQTTMessage
func (c *mqtt) PublishMQTTMessage(input *PublishMQTTMessageInput) (*PublishMQTTMessageOutput, error) {
	req, out := c.PublishMQTTMessageRequest(input)
	err := req.Send()
	return out, err
}

// CreateMQTTInstanceRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTInstanceRequest
type CreateMQTTInstanceInput struct {
	_ struct{} `type:"structure"`

	// String
	Description *string `type:"string"`

	// String
	MQTTInstanceClass *string `type:"string"`

	// String
	//
	// MQTTInstanceIdentifier is a required field
	MQTTInstanceIdentifier *string `type:"string" required:"true"`

	// String
	//
	// Password is a required field
	Password *string `type:"string" required:"true"`

	// String
	//
	// Username is a required field
	Username *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateMQTTInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMQTTInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMQTTInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateMQTTInstanceInput"}
	if s.MQTTInstanceIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("MQTTInstanceIdentifier"))
	}
	if s.Password == nil {
		invalidParams.Add(request.NewErrParamRequired("Password"))
	}
	if s.Username == nil {
		invalidParams.Add(request.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateMQTTInstanceInput) SetDescription(v string) *CreateMQTTInstanceInput {
	s.Description = &v
	return s
}

// SetMQTTInstanceClass sets the MQTTInstanceClass field's value.
func (s *CreateMQTTInstanceInput) SetMQTTInstanceClass(v string) *CreateMQTTInstanceInput {
	s.MQTTInstanceClass = &v
	return s
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *CreateMQTTInstanceInput) SetMQTTInstanceIdentifier(v string) *CreateMQTTInstanceInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *CreateMQTTInstanceInput) SetPassword(v string) *CreateMQTTInstanceInput {
	s.Password = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *CreateMQTTInstanceInput) SetUsername(v string) *CreateMQTTInstanceInput {
	s.Username = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTInstanceOutput
type CreateMQTTInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateMQTTInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMQTTInstanceOutput) GoString() string {
	return s.String()
}

// CreateMQTTUserRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTUserRequest
type CreateMQTTUserInput struct {
	_ struct{} `type:"structure"`

	// String
	Description *string `type:"string"`

	// String
	//
	// MQTTInstanceIdentifier is a required field
	MQTTInstanceIdentifier *string `type:"string" required:"true"`

	// String
	//
	// Password is a required field
	Password *string `type:"string" required:"true"`

	// String
	//
	// Username is a required field
	Username *string `type:"string" required:"true"`
}

// String returns the string representation
func (s CreateMQTTUserInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMQTTUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *CreateMQTTUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "CreateMQTTUserInput"}
	if s.MQTTInstanceIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("MQTTInstanceIdentifier"))
	}
	if s.Password == nil {
		invalidParams.Add(request.NewErrParamRequired("Password"))
	}
	if s.Username == nil {
		invalidParams.Add(request.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *CreateMQTTUserInput) SetDescription(v string) *CreateMQTTUserInput {
	s.Description = &v
	return s
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *CreateMQTTUserInput) SetMQTTInstanceIdentifier(v string) *CreateMQTTUserInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *CreateMQTTUserInput) SetPassword(v string) *CreateMQTTUserInput {
	s.Password = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *CreateMQTTUserInput) SetUsername(v string) *CreateMQTTUserInput {
	s.Username = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//CreateMQTTUserOutput
type CreateMQTTUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s CreateMQTTUserOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateMQTTUserOutput) GoString() string {
	return s.String()
}

// DeleteMQTTInstanceRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTInstanceRequest
type DeleteMQTTInstanceInput struct {
	_ struct{} `type:"structure"`

	// String
	//
	// MQTTInstanceIdentifier is a required field
	MQTTInstanceIdentifier *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMQTTInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMQTTInstanceInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMQTTInstanceInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteMQTTInstanceInput"}
	if s.MQTTInstanceIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("MQTTInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *DeleteMQTTInstanceInput) SetMQTTInstanceIdentifier(v string) *DeleteMQTTInstanceInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTInstanceOutput
type DeleteMQTTInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMQTTInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMQTTInstanceOutput) GoString() string {
	return s.String()
}

// DeleteMQTTUserRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTUserRequest
type DeleteMQTTUserInput struct {
	_ struct{} `type:"structure"`

	// String
	//
	// MQTTInstanceIdentifier is a required field
	MQTTInstanceIdentifier *string `type:"string" required:"true"`

	// String
	//
	// Username is a required field
	Username *string `type:"string" required:"true"`
}

// String returns the string representation
func (s DeleteMQTTUserInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMQTTUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DeleteMQTTUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DeleteMQTTUserInput"}
	if s.MQTTInstanceIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("MQTTInstanceIdentifier"))
	}
	if s.Username == nil {
		invalidParams.Add(request.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *DeleteMQTTUserInput) SetMQTTInstanceIdentifier(v string) *DeleteMQTTUserInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *DeleteMQTTUserInput) SetUsername(v string) *DeleteMQTTUserInput {
	s.Username = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//DeleteMQTTUserOutput
type DeleteMQTTUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DeleteMQTTUserOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteMQTTUserOutput) GoString() string {
	return s.String()
}

// DescribeMQTTInstancesRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTInstancesRequest
type DescribeMQTTInstancesInput struct {
	_ struct{} `type:"structure"`

	// String
	MQTTInstanceIdentifier *string `type:"string"`
}

// String returns the string representation
func (s DescribeMQTTInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeMQTTInstancesInput) GoString() string {
	return s.String()
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *DescribeMQTTInstancesInput) SetMQTTInstanceIdentifier(v string) *DescribeMQTTInstancesInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTInstancesOutput
type DescribeMQTTInstancesOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeMQTTInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeMQTTInstancesOutput) GoString() string {
	return s.String()
}

// DescribeMQTTUsersRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTUsersRequest
type DescribeMQTTUsersInput struct {
	_ struct{} `type:"structure"`

	// String
	//
	// MQTTInstanceIdentifier is a required field
	MQTTInstanceIdentifier *string `type:"string" required:"true"`

	// String
	Username *string `type:"string"`
}

// String returns the string representation
func (s DescribeMQTTUsersInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeMQTTUsersInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *DescribeMQTTUsersInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "DescribeMQTTUsersInput"}
	if s.MQTTInstanceIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("MQTTInstanceIdentifier"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *DescribeMQTTUsersInput) SetMQTTInstanceIdentifier(v string) *DescribeMQTTUsersInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *DescribeMQTTUsersInput) SetUsername(v string) *DescribeMQTTUsersInput {
	s.Username = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//DescribeMQTTUsersOutput
type DescribeMQTTUsersOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s DescribeMQTTUsersOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeMQTTUsersOutput) GoString() string {
	return s.String()
}

// ModifyMQTTInstanceRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTInstanceRequest
type ModifyMQTTInstanceInput struct {
	_ struct{} `type:"structure"`

	// String
	Description *string `type:"string"`

	// String
	MQTTInstanceClass *string `type:"string"`

	// String
	MQTTInstanceIdentifier *string `type:"string"`
}

// String returns the string representation
func (s ModifyMQTTInstanceInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyMQTTInstanceInput) GoString() string {
	return s.String()
}

// SetDescription sets the Description field's value.
func (s *ModifyMQTTInstanceInput) SetDescription(v string) *ModifyMQTTInstanceInput {
	s.Description = &v
	return s
}

// SetMQTTInstanceClass sets the MQTTInstanceClass field's value.
func (s *ModifyMQTTInstanceInput) SetMQTTInstanceClass(v string) *ModifyMQTTInstanceInput {
	s.MQTTInstanceClass = &v
	return s
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *ModifyMQTTInstanceInput) SetMQTTInstanceIdentifier(v string) *ModifyMQTTInstanceInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTInstanceOutput
type ModifyMQTTInstanceOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyMQTTInstanceOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyMQTTInstanceOutput) GoString() string {
	return s.String()
}

// ModifyMQTTUserRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTUserRequest
type ModifyMQTTUserInput struct {
	_ struct{} `type:"structure"`

	// String
	Description *string `type:"string"`

	// String
	//
	// MQTTInstanceIdentifier is a required field
	MQTTInstanceIdentifier *string `type:"string" required:"true"`

	// String
	Password *string `type:"string"`

	// String
	//
	// Username is a required field
	Username *string `type:"string" required:"true"`
}

// String returns the string representation
func (s ModifyMQTTUserInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyMQTTUserInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *ModifyMQTTUserInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "ModifyMQTTUserInput"}
	if s.MQTTInstanceIdentifier == nil {
		invalidParams.Add(request.NewErrParamRequired("MQTTInstanceIdentifier"))
	}
	if s.Username == nil {
		invalidParams.Add(request.NewErrParamRequired("Username"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetDescription sets the Description field's value.
func (s *ModifyMQTTUserInput) SetDescription(v string) *ModifyMQTTUserInput {
	s.Description = &v
	return s
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *ModifyMQTTUserInput) SetMQTTInstanceIdentifier(v string) *ModifyMQTTUserInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *ModifyMQTTUserInput) SetPassword(v string) *ModifyMQTTUserInput {
	s.Password = &v
	return s
}

// SetUsername sets the Username field's value.
func (s *ModifyMQTTUserInput) SetUsername(v string) *ModifyMQTTUserInput {
	s.Username = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//ModifyMQTTUserOutput
type ModifyMQTTUserOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s ModifyMQTTUserOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ModifyMQTTUserOutput) GoString() string {
	return s.String()
}

// PublishMQTTMessageRequest
// Please also see https://docs.aws.amazon.com/goto/WebAPI//PublishMQTTMessageRequest
type PublishMQTTMessageInput struct {
	_ struct{} `type:"structure"`

	// String
	MQTTInstanceIdentifier *string `type:"string"`

	// String
	Message *string `type:"string"`

	// String
	Password *string `type:"string"`

	// Integer
	Qos *int64 `type:"integer"`

	// Integer
	Retain *int64 `type:"integer"`

	// String
	Topic *string `type:"string"`
}

// String returns the string representation
func (s PublishMQTTMessageInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishMQTTMessageInput) GoString() string {
	return s.String()
}

// SetMQTTInstanceIdentifier sets the MQTTInstanceIdentifier field's value.
func (s *PublishMQTTMessageInput) SetMQTTInstanceIdentifier(v string) *PublishMQTTMessageInput {
	s.MQTTInstanceIdentifier = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *PublishMQTTMessageInput) SetMessage(v string) *PublishMQTTMessageInput {
	s.Message = &v
	return s
}

// SetPassword sets the Password field's value.
func (s *PublishMQTTMessageInput) SetPassword(v string) *PublishMQTTMessageInput {
	s.Password = &v
	return s
}

// SetQos sets the Qos field's value.
func (s *PublishMQTTMessageInput) SetQos(v int64) *PublishMQTTMessageInput {
	s.Qos = &v
	return s
}

// SetRetain sets the Retain field's value.
func (s *PublishMQTTMessageInput) SetRetain(v int64) *PublishMQTTMessageInput {
	s.Retain = &v
	return s
}

// SetTopic sets the Topic field's value.
func (s *PublishMQTTMessageInput) SetTopic(v string) *PublishMQTTMessageInput {
	s.Topic = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI//PublishMQTTMessageOutput
type PublishMQTTMessageOutput struct {
	_ struct{} `type:"structure"`
}

// String returns the string representation
func (s PublishMQTTMessageOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PublishMQTTMessageOutput) GoString() string {
	return s.String()
}
