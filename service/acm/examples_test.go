// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package acm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/session"
	"github.com/tily/sdk-go/service/acm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleACM_AddTagsToCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.AddTagsToCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
		Tags: []*acm.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.AddTagsToCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_DeleteCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.DeleteCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DeleteCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_DescribeCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.DescribeCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.DescribeCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_GetCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.GetCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.GetCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_ImportCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.ImportCertificateInput{
		Certificate:      []byte("PAYLOAD"), // Required
		PrivateKey:       []byte("PAYLOAD"), // Required
		CertificateArn:   aws.String("Arn"),
		CertificateChain: []byte("PAYLOAD"),
	}
	resp, err := svc.ImportCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_ListCertificates() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.ListCertificatesInput{
		CertificateStatuses: []*string{
			aws.String("CertificateStatus"), // Required
			// More values...
		},
		MaxItems:  aws.Int64(1),
		NextToken: aws.String("NextToken"),
	}
	resp, err := svc.ListCertificates(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_ListTagsForCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.ListTagsForCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
	}
	resp, err := svc.ListTagsForCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_RemoveTagsFromCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.RemoveTagsFromCertificateInput{
		CertificateArn: aws.String("Arn"), // Required
		Tags: []*acm.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKey"), // Required
				Value: aws.String("TagValue"),
			},
			// More values...
		},
	}
	resp, err := svc.RemoveTagsFromCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_RequestCertificate() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.RequestCertificateInput{
		DomainName: aws.String("DomainNameString"), // Required
		DomainValidationOptions: []*acm.DomainValidationOption{
			{ // Required
				DomainName:       aws.String("DomainNameString"), // Required
				ValidationDomain: aws.String("DomainNameString"), // Required
			},
			// More values...
		},
		IdempotencyToken: aws.String("IdempotencyToken"),
		SubjectAlternativeNames: []*string{
			aws.String("DomainNameString"), // Required
			// More values...
		},
	}
	resp, err := svc.RequestCertificate(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleACM_ResendValidationEmail() {
	sess := session.Must(session.NewSession())

	svc := acm.New(sess)

	params := &acm.ResendValidationEmailInput{
		CertificateArn:   aws.String("Arn"),              // Required
		Domain:           aws.String("DomainNameString"), // Required
		ValidationDomain: aws.String("DomainNameString"), // Required
	}
	resp, err := svc.ResendValidationEmail(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
