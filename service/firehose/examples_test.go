// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package firehose_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/session"
	"github.com/tily/sdk-go/service/firehose"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleFirehose_CreateDeliveryStream() {
	sess := session.Must(session.NewSession())

	svc := firehose.New(sess)

	params := &firehose.CreateDeliveryStreamInput{
		DeliveryStreamName: aws.String("DeliveryStreamName"), // Required
		ElasticsearchDestinationConfiguration: &firehose.ElasticsearchDestinationConfiguration{
			DomainARN: aws.String("ElasticsearchDomainARN"), // Required
			IndexName: aws.String("ElasticsearchIndexName"), // Required
			RoleARN:   aws.String("RoleARN"),                // Required
			S3Configuration: &firehose.S3DestinationConfiguration{ // Required
				BucketARN: aws.String("BucketARN"), // Required
				RoleARN:   aws.String("RoleARN"),   // Required
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix: aws.String("Prefix"),
			},
			TypeName: aws.String("ElasticsearchTypeName"), // Required
			BufferingHints: &firehose.ElasticsearchBufferingHints{
				IntervalInSeconds: aws.Int64(1),
				SizeInMBs:         aws.Int64(1),
			},
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			IndexRotationPeriod: aws.String("ElasticsearchIndexRotationPeriod"),
			ProcessingConfiguration: &firehose.ProcessingConfiguration{
				Enabled: aws.Bool(true),
				Processors: []*firehose.Processor{
					{ // Required
						Type: aws.String("ProcessorType"), // Required
						Parameters: []*firehose.ProcessorParameter{
							{ // Required
								ParameterName:  aws.String("ProcessorParameterName"),  // Required
								ParameterValue: aws.String("ProcessorParameterValue"), // Required
							},
							// More values...
						},
					},
					// More values...
				},
			},
			RetryOptions: &firehose.ElasticsearchRetryOptions{
				DurationInSeconds: aws.Int64(1),
			},
			S3BackupMode: aws.String("ElasticsearchS3BackupMode"),
		},
		ExtendedS3DestinationConfiguration: &firehose.ExtendedS3DestinationConfiguration{
			BucketARN: aws.String("BucketARN"), // Required
			RoleARN:   aws.String("RoleARN"),   // Required
			BufferingHints: &firehose.BufferingHints{
				IntervalInSeconds: aws.Int64(1),
				SizeInMBs:         aws.Int64(1),
			},
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			CompressionFormat: aws.String("CompressionFormat"),
			EncryptionConfiguration: &firehose.EncryptionConfiguration{
				KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
					AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
				},
				NoEncryptionConfig: aws.String("NoEncryptionConfig"),
			},
			Prefix: aws.String("Prefix"),
			ProcessingConfiguration: &firehose.ProcessingConfiguration{
				Enabled: aws.Bool(true),
				Processors: []*firehose.Processor{
					{ // Required
						Type: aws.String("ProcessorType"), // Required
						Parameters: []*firehose.ProcessorParameter{
							{ // Required
								ParameterName:  aws.String("ProcessorParameterName"),  // Required
								ParameterValue: aws.String("ProcessorParameterValue"), // Required
							},
							// More values...
						},
					},
					// More values...
				},
			},
			S3BackupConfiguration: &firehose.S3DestinationConfiguration{
				BucketARN: aws.String("BucketARN"), // Required
				RoleARN:   aws.String("RoleARN"),   // Required
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix: aws.String("Prefix"),
			},
			S3BackupMode: aws.String("S3BackupMode"),
		},
		RedshiftDestinationConfiguration: &firehose.RedshiftDestinationConfiguration{
			ClusterJDBCURL: aws.String("ClusterJDBCURL"), // Required
			CopyCommand: &firehose.CopyCommand{ // Required
				DataTableName:    aws.String("DataTableName"), // Required
				CopyOptions:      aws.String("CopyOptions"),
				DataTableColumns: aws.String("DataTableColumns"),
			},
			Password: aws.String("Password"), // Required
			RoleARN:  aws.String("RoleARN"),  // Required
			S3Configuration: &firehose.S3DestinationConfiguration{ // Required
				BucketARN: aws.String("BucketARN"), // Required
				RoleARN:   aws.String("RoleARN"),   // Required
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix: aws.String("Prefix"),
			},
			Username: aws.String("Username"), // Required
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			ProcessingConfiguration: &firehose.ProcessingConfiguration{
				Enabled: aws.Bool(true),
				Processors: []*firehose.Processor{
					{ // Required
						Type: aws.String("ProcessorType"), // Required
						Parameters: []*firehose.ProcessorParameter{
							{ // Required
								ParameterName:  aws.String("ProcessorParameterName"),  // Required
								ParameterValue: aws.String("ProcessorParameterValue"), // Required
							},
							// More values...
						},
					},
					// More values...
				},
			},
			RetryOptions: &firehose.RedshiftRetryOptions{
				DurationInSeconds: aws.Int64(1),
			},
			S3BackupConfiguration: &firehose.S3DestinationConfiguration{
				BucketARN: aws.String("BucketARN"), // Required
				RoleARN:   aws.String("RoleARN"),   // Required
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix: aws.String("Prefix"),
			},
			S3BackupMode: aws.String("RedshiftS3BackupMode"),
		},
		S3DestinationConfiguration: &firehose.S3DestinationConfiguration{
			BucketARN: aws.String("BucketARN"), // Required
			RoleARN:   aws.String("RoleARN"),   // Required
			BufferingHints: &firehose.BufferingHints{
				IntervalInSeconds: aws.Int64(1),
				SizeInMBs:         aws.Int64(1),
			},
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			CompressionFormat: aws.String("CompressionFormat"),
			EncryptionConfiguration: &firehose.EncryptionConfiguration{
				KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
					AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
				},
				NoEncryptionConfig: aws.String("NoEncryptionConfig"),
			},
			Prefix: aws.String("Prefix"),
		},
	}
	resp, err := svc.CreateDeliveryStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleFirehose_DeleteDeliveryStream() {
	sess := session.Must(session.NewSession())

	svc := firehose.New(sess)

	params := &firehose.DeleteDeliveryStreamInput{
		DeliveryStreamName: aws.String("DeliveryStreamName"), // Required
	}
	resp, err := svc.DeleteDeliveryStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleFirehose_DescribeDeliveryStream() {
	sess := session.Must(session.NewSession())

	svc := firehose.New(sess)

	params := &firehose.DescribeDeliveryStreamInput{
		DeliveryStreamName:          aws.String("DeliveryStreamName"), // Required
		ExclusiveStartDestinationId: aws.String("DestinationId"),
		Limit: aws.Int64(1),
	}
	resp, err := svc.DescribeDeliveryStream(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleFirehose_ListDeliveryStreams() {
	sess := session.Must(session.NewSession())

	svc := firehose.New(sess)

	params := &firehose.ListDeliveryStreamsInput{
		ExclusiveStartDeliveryStreamName: aws.String("DeliveryStreamName"),
		Limit: aws.Int64(1),
	}
	resp, err := svc.ListDeliveryStreams(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleFirehose_PutRecord() {
	sess := session.Must(session.NewSession())

	svc := firehose.New(sess)

	params := &firehose.PutRecordInput{
		DeliveryStreamName: aws.String("DeliveryStreamName"), // Required
		Record: &firehose.Record{ // Required
			Data: []byte("PAYLOAD"), // Required
		},
	}
	resp, err := svc.PutRecord(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleFirehose_PutRecordBatch() {
	sess := session.Must(session.NewSession())

	svc := firehose.New(sess)

	params := &firehose.PutRecordBatchInput{
		DeliveryStreamName: aws.String("DeliveryStreamName"), // Required
		Records: []*firehose.Record{ // Required
			{ // Required
				Data: []byte("PAYLOAD"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.PutRecordBatch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleFirehose_UpdateDestination() {
	sess := session.Must(session.NewSession())

	svc := firehose.New(sess)

	params := &firehose.UpdateDestinationInput{
		CurrentDeliveryStreamVersionId: aws.String("DeliveryStreamVersionId"), // Required
		DeliveryStreamName:             aws.String("DeliveryStreamName"),      // Required
		DestinationId:                  aws.String("DestinationId"),           // Required
		ElasticsearchDestinationUpdate: &firehose.ElasticsearchDestinationUpdate{
			BufferingHints: &firehose.ElasticsearchBufferingHints{
				IntervalInSeconds: aws.Int64(1),
				SizeInMBs:         aws.Int64(1),
			},
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			DomainARN:           aws.String("ElasticsearchDomainARN"),
			IndexName:           aws.String("ElasticsearchIndexName"),
			IndexRotationPeriod: aws.String("ElasticsearchIndexRotationPeriod"),
			ProcessingConfiguration: &firehose.ProcessingConfiguration{
				Enabled: aws.Bool(true),
				Processors: []*firehose.Processor{
					{ // Required
						Type: aws.String("ProcessorType"), // Required
						Parameters: []*firehose.ProcessorParameter{
							{ // Required
								ParameterName:  aws.String("ProcessorParameterName"),  // Required
								ParameterValue: aws.String("ProcessorParameterValue"), // Required
							},
							// More values...
						},
					},
					// More values...
				},
			},
			RetryOptions: &firehose.ElasticsearchRetryOptions{
				DurationInSeconds: aws.Int64(1),
			},
			RoleARN: aws.String("RoleARN"),
			S3Update: &firehose.S3DestinationUpdate{
				BucketARN: aws.String("BucketARN"),
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix:  aws.String("Prefix"),
				RoleARN: aws.String("RoleARN"),
			},
			TypeName: aws.String("ElasticsearchTypeName"),
		},
		ExtendedS3DestinationUpdate: &firehose.ExtendedS3DestinationUpdate{
			BucketARN: aws.String("BucketARN"),
			BufferingHints: &firehose.BufferingHints{
				IntervalInSeconds: aws.Int64(1),
				SizeInMBs:         aws.Int64(1),
			},
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			CompressionFormat: aws.String("CompressionFormat"),
			EncryptionConfiguration: &firehose.EncryptionConfiguration{
				KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
					AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
				},
				NoEncryptionConfig: aws.String("NoEncryptionConfig"),
			},
			Prefix: aws.String("Prefix"),
			ProcessingConfiguration: &firehose.ProcessingConfiguration{
				Enabled: aws.Bool(true),
				Processors: []*firehose.Processor{
					{ // Required
						Type: aws.String("ProcessorType"), // Required
						Parameters: []*firehose.ProcessorParameter{
							{ // Required
								ParameterName:  aws.String("ProcessorParameterName"),  // Required
								ParameterValue: aws.String("ProcessorParameterValue"), // Required
							},
							// More values...
						},
					},
					// More values...
				},
			},
			RoleARN:      aws.String("RoleARN"),
			S3BackupMode: aws.String("S3BackupMode"),
			S3BackupUpdate: &firehose.S3DestinationUpdate{
				BucketARN: aws.String("BucketARN"),
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix:  aws.String("Prefix"),
				RoleARN: aws.String("RoleARN"),
			},
		},
		RedshiftDestinationUpdate: &firehose.RedshiftDestinationUpdate{
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			ClusterJDBCURL: aws.String("ClusterJDBCURL"),
			CopyCommand: &firehose.CopyCommand{
				DataTableName:    aws.String("DataTableName"), // Required
				CopyOptions:      aws.String("CopyOptions"),
				DataTableColumns: aws.String("DataTableColumns"),
			},
			Password: aws.String("Password"),
			ProcessingConfiguration: &firehose.ProcessingConfiguration{
				Enabled: aws.Bool(true),
				Processors: []*firehose.Processor{
					{ // Required
						Type: aws.String("ProcessorType"), // Required
						Parameters: []*firehose.ProcessorParameter{
							{ // Required
								ParameterName:  aws.String("ProcessorParameterName"),  // Required
								ParameterValue: aws.String("ProcessorParameterValue"), // Required
							},
							// More values...
						},
					},
					// More values...
				},
			},
			RetryOptions: &firehose.RedshiftRetryOptions{
				DurationInSeconds: aws.Int64(1),
			},
			RoleARN:      aws.String("RoleARN"),
			S3BackupMode: aws.String("RedshiftS3BackupMode"),
			S3BackupUpdate: &firehose.S3DestinationUpdate{
				BucketARN: aws.String("BucketARN"),
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix:  aws.String("Prefix"),
				RoleARN: aws.String("RoleARN"),
			},
			S3Update: &firehose.S3DestinationUpdate{
				BucketARN: aws.String("BucketARN"),
				BufferingHints: &firehose.BufferingHints{
					IntervalInSeconds: aws.Int64(1),
					SizeInMBs:         aws.Int64(1),
				},
				CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
					Enabled:       aws.Bool(true),
					LogGroupName:  aws.String("LogGroupName"),
					LogStreamName: aws.String("LogStreamName"),
				},
				CompressionFormat: aws.String("CompressionFormat"),
				EncryptionConfiguration: &firehose.EncryptionConfiguration{
					KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
						AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
					},
					NoEncryptionConfig: aws.String("NoEncryptionConfig"),
				},
				Prefix:  aws.String("Prefix"),
				RoleARN: aws.String("RoleARN"),
			},
			Username: aws.String("Username"),
		},
		S3DestinationUpdate: &firehose.S3DestinationUpdate{
			BucketARN: aws.String("BucketARN"),
			BufferingHints: &firehose.BufferingHints{
				IntervalInSeconds: aws.Int64(1),
				SizeInMBs:         aws.Int64(1),
			},
			CloudWatchLoggingOptions: &firehose.CloudWatchLoggingOptions{
				Enabled:       aws.Bool(true),
				LogGroupName:  aws.String("LogGroupName"),
				LogStreamName: aws.String("LogStreamName"),
			},
			CompressionFormat: aws.String("CompressionFormat"),
			EncryptionConfiguration: &firehose.EncryptionConfiguration{
				KMSEncryptionConfig: &firehose.KMSEncryptionConfig{
					AWSKMSKeyARN: aws.String("AWSKMSKeyARN"), // Required
				},
				NoEncryptionConfig: aws.String("NoEncryptionConfig"),
			},
			Prefix:  aws.String("Prefix"),
			RoleARN: aws.String("RoleARN"),
		},
	}
	resp, err := svc.UpdateDestination(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
