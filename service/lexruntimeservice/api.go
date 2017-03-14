// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package lexruntimeservice provides a client for Amazon Lex Runtime Service.
package lexruntimeservice

import (
	"github.com/tily/sdk-go/aws/awsutil"
	"github.com/tily/sdk-go/aws/request"
)

const opPostText = "PostText"

// PostTextRequest generates a "aws/request.Request" representing the
// client's request for the PostText operation. The "output" return
// value can be used to capture response data after the request's "Send" method
// is called.
//
// See PostText for usage and error information.
//
// Creating a request object using this method should be used when you want to inject
// custom logic into the request's lifecycle using a custom handler, or if you want to
// access properties on the request object before or after sending the request. If
// you just want the service response, call the PostText method directly
// instead.
//
// Note: You must call the "Send" method on the returned request object in order
// to execute the request.
//
//    // Example sending a request using the PostTextRequest method.
//    req, resp := client.PostTextRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostText
func (c *LexRuntimeService) PostTextRequest(input *PostTextInput) (req *request.Request, output *PostTextOutput) {
	op := &request.Operation{
		Name:       opPostText,
		HTTPMethod: "POST",
		HTTPPath:   "/bot/{botName}/alias/{botAlias}/user/{userId}/text",
	}

	if input == nil {
		input = &PostTextInput{}
	}

	output = &PostTextOutput{}
	req = c.newRequest(op, input, output)
	return
}

// PostText API operation for Amazon Lex Runtime Service.
//
// Sends user input text to Amazon Lex at runtime. Amazon Lex uses the machine
// learning model that the service built for the application to interpret user
// input.
//
// In response, Amazon Lex returns the next message to convey to the user (based
// on the context of the user interaction) and whether to expect a user response
// to the message (dialogState). For example, consider the following response
// messages:
//
//    * "What pizza toppings would you like?" – In this case, the dialogState
//    would be ElicitSlot (that is, a user response is expected).
//
//    * "Your order has been placed." – In this case, Amazon Lex returns one
//    of the following dialogState values depending on how the intent fulfillment
//    is configured (see fulfillmentActivity in CreateIntent):
//
// FulFilled – The intent fulfillment is configured through a Lambda function.
//
//
// ReadyForFulfilment – The intent's fulfillmentActivity is to simply return
//    the intent data back to the client application.
//
// Returns awserr.Error for service API and SDK errors. Use runtime type assertions
// with awserr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the AWS API reference guide for Amazon Lex Runtime Service's
// API operation PostText for usage and error information.
//
// Returned Error Codes:
//   * ErrCodeNotFoundException "NotFoundException"
//   Resource (such as the Amazon Lex bot or an alias) referred is not found.
//
//   * ErrCodeBadRequestException "BadRequestException"
//   Request validation failed, there is no usable message in the context, or
//   the bot build failed.
//
//   * ErrCodeLimitExceededException "LimitExceededException"
//
//   * ErrCodeInternalFailureException "InternalFailureException"
//   Internal service error. Retry the call.
//
//   * ErrCodeConflictException "ConflictException"
//   Two clients are using the same AWS account, Amazon Lex bot, and user ID.
//
//   * ErrCodeDependencyFailedException "DependencyFailedException"
//   One of the downstream dependencies, such as AWS Lambda or Amazon Polly, threw
//   an exception. For example, if Amazon Lex does not have sufficient permissions
//   to call a Lambda function which results in AWS Lambda throwing an exception.
//
//   * ErrCodeBadGatewayException "BadGatewayException"
//   Either the Amazon Lex bot is still building, or one of the dependent services
//   (Amazon Polly, AWS Lambda) failed with an internal service error.
//
//   * ErrCodeLoopDetectedException "LoopDetectedException"
//   Lambda fulfilment function returned DelegateDialogAction to Amazon Lex without
//   changing any slot values.
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostText
func (c *LexRuntimeService) PostText(input *PostTextInput) (*PostTextOutput, error) {
	req, out := c.PostTextRequest(input)
	err := req.Send()
	return out, err
}

// Represents an option to be shown on the client platform (Facebook, Slack,
// etc.)
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/Button
type Button struct {
	_ struct{} `type:"structure"`

	// Text visible to the user on the button.
	//
	// Text is a required field
	Text *string `locationName:"text" min:"1" type:"string" required:"true"`

	// Value sent to Amazon Lex when user clicks the button. For example, consider
	// button text "NYC". When the user clicks the button, the value sent can be
	// "New York City".
	//
	// Value is a required field
	Value *string `locationName:"value" min:"1" type:"string" required:"true"`
}

// String returns the string representation
func (s Button) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s Button) GoString() string {
	return s.String()
}

// SetText sets the Text field's value.
func (s *Button) SetText(v string) *Button {
	s.Text = &v
	return s
}

// SetValue sets the Value field's value.
func (s *Button) SetValue(v string) *Button {
	s.Value = &v
	return s
}

// Represents an option rendered to the user when a prompt is shown. It could
// be an image, a button, a link, or text.
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/GenericAttachment
type GenericAttachment struct {
	_ struct{} `type:"structure"`

	AttachmentLinkUrl *string `locationName:"attachmentLinkUrl" min:"1" type:"string"`

	// List of options to show to the user.
	Buttons []*Button `locationName:"buttons" type:"list"`

	// URL of an image that is displayed to the user.
	ImageUrl *string `locationName:"imageUrl" min:"1" type:"string"`

	// Subtitle shown below the title.
	SubTitle *string `locationName:"subTitle" min:"1" type:"string"`

	// Title of the option.
	Title *string `locationName:"title" min:"1" type:"string"`
}

// String returns the string representation
func (s GenericAttachment) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s GenericAttachment) GoString() string {
	return s.String()
}

// SetAttachmentLinkUrl sets the AttachmentLinkUrl field's value.
func (s *GenericAttachment) SetAttachmentLinkUrl(v string) *GenericAttachment {
	s.AttachmentLinkUrl = &v
	return s
}

// SetButtons sets the Buttons field's value.
func (s *GenericAttachment) SetButtons(v []*Button) *GenericAttachment {
	s.Buttons = v
	return s
}

// SetImageUrl sets the ImageUrl field's value.
func (s *GenericAttachment) SetImageUrl(v string) *GenericAttachment {
	s.ImageUrl = &v
	return s
}

// SetSubTitle sets the SubTitle field's value.
func (s *GenericAttachment) SetSubTitle(v string) *GenericAttachment {
	s.SubTitle = &v
	return s
}

// SetTitle sets the Title field's value.
func (s *GenericAttachment) SetTitle(v string) *GenericAttachment {
	s.Title = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostTextRequest
type PostTextInput struct {
	_ struct{} `type:"structure"`

	// Alias of the Amazon Lex bot.
	//
	// BotAlias is a required field
	BotAlias *string `location:"uri" locationName:"botAlias" type:"string" required:"true"`

	// Name of the Amazon Lex bot.
	//
	// BotName is a required field
	BotName *string `location:"uri" locationName:"botName" type:"string" required:"true"`

	// Text user entered (Amazon Lex interprets this text).
	//
	// InputText is a required field
	InputText *string `locationName:"inputText" min:"1" type:"string" required:"true"`

	// A session represents the dialog between a user and Amazon Lex. At runtime,
	// a client application can pass contextual information (session attributes)
	// in the request. For example, "FirstName" : "Joe". Amazon Lex passes these
	// session attributes to the AWS Lambda functions configured for the intent
	// (see dialogCodeHook and fulfillmentActivity.codeHook in CreateIntent).
	//
	// In your Lambda function, you can use the session attributes for customization.
	// Some examples are:
	//
	//    *  In a pizza ordering application, if you can pass user location as a
	//    session attribute (for example, "Location" : "111 Maple street"), your
	//    Lambda function might use this information to determine the closest pizzeria
	//    to place the order.
	//
	//    *  Use session attributes to personalize prompts. For example, you pass
	//    in user name as a session attribute ("FirstName" : "Joe"), you might configure
	//    subsequent prompts to refer to this attribute, as $session.FirstName".
	//    At runtime, Amazon Lex substitutes a real value when it generates a prompt,
	//    such as "Hello Joe, what would you like to order?"
	//
	// Amazon Lex does not persist session attributes.
	//
	//  If the intent is configured without a Lambda function to process the intent
	// (that is, the client application to process the intent), Amazon Lex simply
	// returns the session attributes back to the client application.
	//
	//  If the intent is configured with a Lambda function to process the intent,
	// Amazon Lex passes the incoming session attributes to the Lambda function.
	// The Lambda function must return these session attributes if you want Amazon
	// Lex to return them back to the client.
	SessionAttributes map[string]*string `locationName:"sessionAttributes" type:"map"`

	// User ID of your client application. Typically, each of your application users
	// should have a unique ID. Note the following considerations:
	//
	//    *  If you want a user to start a conversation on one mobile device and
	//    continue the conversation on another device, you might choose a user-specific
	//    identifier, such as a login or Amazon Cognito user ID (assuming your application
	//    is using Amazon Cognito).
	//
	//    *  If you want the same user to be able to have two independent conversations
	//    on two different devices, you might choose a device-specific identifier,
	//    such as device ID, or some globally unique identifier.
	//
	// UserId is a required field
	UserId *string `location:"uri" locationName:"userId" min:"2" type:"string" required:"true"`
}

// String returns the string representation
func (s PostTextInput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostTextInput) GoString() string {
	return s.String()
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *PostTextInput) Validate() error {
	invalidParams := request.ErrInvalidParams{Context: "PostTextInput"}
	if s.BotAlias == nil {
		invalidParams.Add(request.NewErrParamRequired("BotAlias"))
	}
	if s.BotName == nil {
		invalidParams.Add(request.NewErrParamRequired("BotName"))
	}
	if s.InputText == nil {
		invalidParams.Add(request.NewErrParamRequired("InputText"))
	}
	if s.InputText != nil && len(*s.InputText) < 1 {
		invalidParams.Add(request.NewErrParamMinLen("InputText", 1))
	}
	if s.UserId == nil {
		invalidParams.Add(request.NewErrParamRequired("UserId"))
	}
	if s.UserId != nil && len(*s.UserId) < 2 {
		invalidParams.Add(request.NewErrParamMinLen("UserId", 2))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetBotAlias sets the BotAlias field's value.
func (s *PostTextInput) SetBotAlias(v string) *PostTextInput {
	s.BotAlias = &v
	return s
}

// SetBotName sets the BotName field's value.
func (s *PostTextInput) SetBotName(v string) *PostTextInput {
	s.BotName = &v
	return s
}

// SetInputText sets the InputText field's value.
func (s *PostTextInput) SetInputText(v string) *PostTextInput {
	s.InputText = &v
	return s
}

// SetSessionAttributes sets the SessionAttributes field's value.
func (s *PostTextInput) SetSessionAttributes(v map[string]*string) *PostTextInput {
	s.SessionAttributes = v
	return s
}

// SetUserId sets the UserId field's value.
func (s *PostTextInput) SetUserId(v string) *PostTextInput {
	s.UserId = &v
	return s
}

// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/PostTextResponse
type PostTextOutput struct {
	_ struct{} `type:"structure"`

	// Represents the message type to be conveyed to the user. For example:
	//
	//    * ElicitIntent – Amazon Lex wants to elicit user intent. For example,
	//    Amazon Lex did not understand the first utterances such as "I want to
	//    order pizza", which indicates the OrderPizza intent. If Amazon Lex doesn't
	//    understand the intent, it returns this dialogState. Another example is
	//    when your intent is configured with a follow up prompt. For example, after
	//    OrderPizza intent is fulfilled, the intent might have a follow up prompt
	//    such as " Do you want to order a drink or desert?" In this case, Amazon
	//    Lex returns this dialogState.
	//
	//    * ConfirmIntent – Amazon Lex is expecting a yes/no response from the user
	//    indicating whether to go ahead and fulfill the intent (for example, OK
	//    to go ahead and order the pizza). In addition to a yes/no reply, the user
	//    might provide a response with additional slot information (either new
	//    slot information or changes to the existing slot values). For example,
	//    "Yes, but change to thick crust." Amazon Lex understands the additional
	//    information and updates the intent slots accordingly.
	//
	//  Consider another example. Before fulfilling an order, your application might
	//    prompt for confirmation such as "Do you want to place this pizza order?"
	//    A user might reply with "No, I want to order a drink." Amazon Lex recognizes
	//    the new OrderDrink intent.
	//
	//    * ElicitSlot – Amazon Lex is expecting a value of a slot for the current
	//    intent. For example, suppose Amazon Lex asks, "What size pizza would you
	//    like?" A user might reply with "Medium pepperoni pizza." Amazon Lex recognizes
	//    the size and the topping as the two separate slot values.
	//
	//    * Fulfilled – Conveys that the Lambda function has successfully fulfilled
	//    the intent. If Lambda function returns a statement/message to convey the
	//    fulfillment result, Amazon Lex passes this string to the client. If not,
	//    Amazon Lex looks for conclusionStatement that you configured for the intent.
	//
	//
	//  If both the Lambda function statement and the conclusionStatement are missing,
	//    Amazon Lex throws a bad request exception.
	//
	//    * ReadyForFulfillment – conveys that the client has to do the fulfillment
	//    work for the intent. This is the case when the current intent is configured
	//    with ReturnIntent as the fulfillmentActivity , where Amazon Lex returns
	//    this state to client.
	//
	//    * Failed – Conversation with the user failed. Some of the reasons for
	//    this dialogState are: after the configured number of attempts the user
	//    didn't provide an appropriate response, or the Lambda function failed
	//    to fulfill an intent.
	DialogState *string `locationName:"dialogState" type:"string" enum:"DialogState"`

	// Intent Amazon Lex inferred from the user input text. This is one of the intents
	// configured for the bot.
	IntentName *string `locationName:"intentName" type:"string"`

	// Prompt (or statement) to convey to the user. This is based on the application
	// configuration and context. For example, if Amazon Lex did not understand
	// the user intent, it sends the clarificationPrompt configured for the application.
	// In another example, if the intent requires confirmation before taking the
	// fulfillment action, it sends the confirmationPrompt. Suppose the Lambda function
	// successfully fulfilled the intent, and sent a message to convey to the user.
	// In that situation, Amazon Lex sends that message in the response.
	Message *string `locationName:"message" min:"1" type:"string"`

	// Represents the options that the user has to respond to the current prompt.
	// Amazon Lex sends this in the response only if the dialogState value indicates
	// that a user response is expected.
	ResponseCard *ResponseCard `locationName:"responseCard" type:"structure"`

	// Map of key value pairs representing the session specific context information.
	SessionAttributes map[string]*string `locationName:"sessionAttributes" type:"map"`

	// If dialogState value is ElicitSlot, returns the name of the slot for which
	// Amazon Lex is eliciting a value.
	SlotToElicit *string `locationName:"slotToElicit" type:"string"`

	// Intent slots (name/value pairs) Amazon Lex detected so far from the user
	// input in the conversation.
	Slots map[string]*string `locationName:"slots" type:"map"`
}

// String returns the string representation
func (s PostTextOutput) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s PostTextOutput) GoString() string {
	return s.String()
}

// SetDialogState sets the DialogState field's value.
func (s *PostTextOutput) SetDialogState(v string) *PostTextOutput {
	s.DialogState = &v
	return s
}

// SetIntentName sets the IntentName field's value.
func (s *PostTextOutput) SetIntentName(v string) *PostTextOutput {
	s.IntentName = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *PostTextOutput) SetMessage(v string) *PostTextOutput {
	s.Message = &v
	return s
}

// SetResponseCard sets the ResponseCard field's value.
func (s *PostTextOutput) SetResponseCard(v *ResponseCard) *PostTextOutput {
	s.ResponseCard = v
	return s
}

// SetSessionAttributes sets the SessionAttributes field's value.
func (s *PostTextOutput) SetSessionAttributes(v map[string]*string) *PostTextOutput {
	s.SessionAttributes = v
	return s
}

// SetSlotToElicit sets the SlotToElicit field's value.
func (s *PostTextOutput) SetSlotToElicit(v string) *PostTextOutput {
	s.SlotToElicit = &v
	return s
}

// SetSlots sets the Slots field's value.
func (s *PostTextOutput) SetSlots(v map[string]*string) *PostTextOutput {
	s.Slots = v
	return s
}

// If you configure a response card when creating your bots, Amazon Lex substitutes
// the session attributes and slot values available, and then returns it. The
// response card can also come from a Lambda function ( dialogCodeHook and fulfillmentActivity
// on an intent).
// Please also see https://docs.aws.amazon.com/goto/WebAPI/runtime.lex-2016-11-28/ResponseCard
type ResponseCard struct {
	_ struct{} `type:"structure"`

	// Content type of the response.
	ContentType *string `locationName:"contentType" type:"string" enum:"ContentType"`

	// An array of attachment objects representing options.
	GenericAttachments []*GenericAttachment `locationName:"genericAttachments" type:"list"`

	// Version of response card format.
	Version *string `locationName:"version" type:"string"`
}

// String returns the string representation
func (s ResponseCard) String() string {
	return awsutil.Prettify(s)
}

// GoString returns the string representation
func (s ResponseCard) GoString() string {
	return s.String()
}

// SetContentType sets the ContentType field's value.
func (s *ResponseCard) SetContentType(v string) *ResponseCard {
	s.ContentType = &v
	return s
}

// SetGenericAttachments sets the GenericAttachments field's value.
func (s *ResponseCard) SetGenericAttachments(v []*GenericAttachment) *ResponseCard {
	s.GenericAttachments = v
	return s
}

// SetVersion sets the Version field's value.
func (s *ResponseCard) SetVersion(v string) *ResponseCard {
	s.Version = &v
	return s
}

const (
	// ContentTypeApplicationVndAmazonawsCardGeneric is a ContentType enum value
	ContentTypeApplicationVndAmazonawsCardGeneric = "application/vnd.amazonaws.card.generic"
)

const (
	// DialogStateElicitIntent is a DialogState enum value
	DialogStateElicitIntent = "ElicitIntent"

	// DialogStateConfirmIntent is a DialogState enum value
	DialogStateConfirmIntent = "ConfirmIntent"

	// DialogStateElicitSlot is a DialogState enum value
	DialogStateElicitSlot = "ElicitSlot"

	// DialogStateFulfilled is a DialogState enum value
	DialogStateFulfilled = "Fulfilled"

	// DialogStateReadyForFulfillment is a DialogState enum value
	DialogStateReadyForFulfillment = "ReadyForFulfillment"

	// DialogStateFailed is a DialogState enum value
	DialogStateFailed = "Failed"
)
