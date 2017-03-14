// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package rekognition_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/session"
	"github.com/tily/sdk-go/service/rekognition"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleRekognition_CompareFaces() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.CompareFacesInput{
		SourceImage: &rekognition.Image{ // Required
			Bytes: []byte("PAYLOAD"),
			S3Object: &rekognition.S3Object{
				Bucket:  aws.String("S3Bucket"),
				Name:    aws.String("S3ObjectName"),
				Version: aws.String("S3ObjectVersion"),
			},
		},
		TargetImage: &rekognition.Image{ // Required
			Bytes: []byte("PAYLOAD"),
			S3Object: &rekognition.S3Object{
				Bucket:  aws.String("S3Bucket"),
				Name:    aws.String("S3ObjectName"),
				Version: aws.String("S3ObjectVersion"),
			},
		},
		SimilarityThreshold: aws.Float64(1.0),
	}
	resp, err := svc.CompareFaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_CreateCollection() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.CreateCollectionInput{
		CollectionId: aws.String("CollectionId"), // Required
	}
	resp, err := svc.CreateCollection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_DeleteCollection() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.DeleteCollectionInput{
		CollectionId: aws.String("CollectionId"), // Required
	}
	resp, err := svc.DeleteCollection(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_DeleteFaces() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.DeleteFacesInput{
		CollectionId: aws.String("CollectionId"), // Required
		FaceIds: []*string{ // Required
			aws.String("FaceId"), // Required
			// More values...
		},
	}
	resp, err := svc.DeleteFaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_DetectFaces() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.DetectFacesInput{
		Image: &rekognition.Image{ // Required
			Bytes: []byte("PAYLOAD"),
			S3Object: &rekognition.S3Object{
				Bucket:  aws.String("S3Bucket"),
				Name:    aws.String("S3ObjectName"),
				Version: aws.String("S3ObjectVersion"),
			},
		},
		Attributes: []*string{
			aws.String("Attribute"), // Required
			// More values...
		},
	}
	resp, err := svc.DetectFaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_DetectLabels() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.DetectLabelsInput{
		Image: &rekognition.Image{ // Required
			Bytes: []byte("PAYLOAD"),
			S3Object: &rekognition.S3Object{
				Bucket:  aws.String("S3Bucket"),
				Name:    aws.String("S3ObjectName"),
				Version: aws.String("S3ObjectVersion"),
			},
		},
		MaxLabels:     aws.Int64(1),
		MinConfidence: aws.Float64(1.0),
	}
	resp, err := svc.DetectLabels(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_IndexFaces() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.IndexFacesInput{
		CollectionId: aws.String("CollectionId"), // Required
		Image: &rekognition.Image{ // Required
			Bytes: []byte("PAYLOAD"),
			S3Object: &rekognition.S3Object{
				Bucket:  aws.String("S3Bucket"),
				Name:    aws.String("S3ObjectName"),
				Version: aws.String("S3ObjectVersion"),
			},
		},
		DetectionAttributes: []*string{
			aws.String("Attribute"), // Required
			// More values...
		},
		ExternalImageId: aws.String("ExternalImageId"),
	}
	resp, err := svc.IndexFaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_ListCollections() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.ListCollectionsInput{
		MaxResults: aws.Int64(1),
		NextToken:  aws.String("PaginationToken"),
	}
	resp, err := svc.ListCollections(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_ListFaces() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.ListFacesInput{
		CollectionId: aws.String("CollectionId"), // Required
		MaxResults:   aws.Int64(1),
		NextToken:    aws.String("PaginationToken"),
	}
	resp, err := svc.ListFaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_SearchFaces() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.SearchFacesInput{
		CollectionId:       aws.String("CollectionId"), // Required
		FaceId:             aws.String("FaceId"),       // Required
		FaceMatchThreshold: aws.Float64(1.0),
		MaxFaces:           aws.Int64(1),
	}
	resp, err := svc.SearchFaces(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleRekognition_SearchFacesByImage() {
	sess := session.Must(session.NewSession())

	svc := rekognition.New(sess)

	params := &rekognition.SearchFacesByImageInput{
		CollectionId: aws.String("CollectionId"), // Required
		Image: &rekognition.Image{ // Required
			Bytes: []byte("PAYLOAD"),
			S3Object: &rekognition.S3Object{
				Bucket:  aws.String("S3Bucket"),
				Name:    aws.String("S3ObjectName"),
				Version: aws.String("S3ObjectVersion"),
			},
		},
		FaceMatchThreshold: aws.Float64(1.0),
		MaxFaces:           aws.Int64(1),
	}
	resp, err := svc.SearchFacesByImage(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
