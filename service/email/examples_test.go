// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package email_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/session"
	"github.com/tily/sdk-go/service/email"
)

var _ time.Duration
var _ bytes.Buffer

func Exampleemail_DeleteIdentity() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.DeleteIdentityInput{
		Identity: aws.String("String"), // Required
	}
	resp, err := svc.DeleteIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_GetIdentityDkimAttributes() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.GetIdentityDkimAttributesInput{
		IdentitiesList: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityDkimAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_GetIdentityVerificationAttributes() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.GetIdentityVerificationAttributesInput{
		IdentitiesList: []*string{ // Required
			aws.String("String"), // Required
			// More values...
		},
	}
	resp, err := svc.GetIdentityVerificationAttributes(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_GetSendQuota() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	var params *email.GetSendQuotaInput
	resp, err := svc.GetSendQuota(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_GetSendStatistics() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	var params *email.GetSendStatisticsInput
	resp, err := svc.GetSendStatistics(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_ListIdentities() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	var params *email.ListIdentitiesInput
	resp, err := svc.ListIdentities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_SendEmail() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.SendEmailInput{
		Destination: &email.DestinationStruct{ // Required
			BccAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
			CcAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
			ToAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
		},
		Destination: &email.DestinationStruct{ // Required
			BccAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
			CcAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
			ToAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
		},
		Destination: &email.DestinationStruct{ // Required
			BccAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
			CcAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
			ToAddressesList: []*string{
				aws.String("String"), // Required
				// More values...
			},
		},
		Message: &email.MessageStruct{ // Required
			BodyStruct: &email.BodyStruct{
				HtmlStruct: &email.HtmlStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
				TextStruct: &email.TextStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
			},
			SubjectStruct: &email.SubjectStruct{
				Charset: aws.String("Charset"),
				Data:    aws.String("Data"),
			},
		},
		Message: &email.MessageStruct{ // Required
			BodyStruct: &email.BodyStruct{
				HtmlStruct: &email.HtmlStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
				TextStruct: &email.TextStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
			},
			SubjectStruct: &email.SubjectStruct{
				Charset: aws.String("Charset"),
				Data:    aws.String("Data"),
			},
		},
		Message: &email.MessageStruct{ // Required
			BodyStruct: &email.BodyStruct{
				HtmlStruct: &email.HtmlStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
				TextStruct: &email.TextStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
			},
			SubjectStruct: &email.SubjectStruct{
				Charset: aws.String("Charset"),
				Data:    aws.String("Data"),
			},
		},
		Message: &email.MessageStruct{ // Required
			BodyStruct: &email.BodyStruct{
				HtmlStruct: &email.HtmlStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
				TextStruct: &email.TextStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
			},
			SubjectStruct: &email.SubjectStruct{
				Charset: aws.String("Charset"),
				Data:    aws.String("Data"),
			},
		},
		Message: &email.MessageStruct{ // Required
			BodyStruct: &email.BodyStruct{
				HtmlStruct: &email.HtmlStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
				TextStruct: &email.TextStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
			},
			SubjectStruct: &email.SubjectStruct{
				Charset: aws.String("Charset"),
				Data:    aws.String("Data"),
			},
		},
		Message: &email.MessageStruct{ // Required
			BodyStruct: &email.BodyStruct{
				HtmlStruct: &email.HtmlStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
				TextStruct: &email.TextStruct{
					Charset: aws.String("Charset"),
					Data:    aws.String("Data"),
				},
			},
			SubjectStruct: &email.SubjectStruct{
				Charset: aws.String("Charset"),
				Data:    aws.String("Data"),
			},
		},
		Source: aws.String("String"), // Required
		ReplyToAddressesList: []*string{
			aws.String("String"), // Required
			// More values...
		},
		ReturnPath: aws.String("String"),
	}
	resp, err := svc.SendEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_SendRawEmail() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.SendRawEmailInput{
		RawMessage: &email.RawMessageStruct{ // Required
			Data: aws.String("Data"),
		},
		DestinationsList: []*string{
			aws.String("String"), // Required
			// More values...
		},
		Source: aws.String("String"),
	}
	resp, err := svc.SendRawEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_SetIdentityDkimEnabled() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.SetIdentityDkimEnabledInput{
		DkimEnabled: aws.String("String"), // Required
		Identity:    aws.String("String"), // Required
	}
	resp, err := svc.SetIdentityDkimEnabled(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_VerifyDomainDkim() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.VerifyDomainDkimInput{
		Domain: aws.String("String"), // Required
	}
	resp, err := svc.VerifyDomainDkim(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_VerifyDomainIdentity() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.VerifyDomainIdentityInput{
		Domain: aws.String("String"), // Required
	}
	resp, err := svc.VerifyDomainIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func Exampleemail_VerifyEmailIdentity() {
	sess := session.Must(session.NewSession())

	svc := email.New(sess)

	params := &email.VerifyEmailIdentityInput{
		EmailAddress: aws.String("String"), // Required
	}
	resp, err := svc.VerifyEmailIdentity(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
