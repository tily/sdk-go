// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package dynamodb_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/session"
	"github.com/tily/sdk-go/service/dynamodb"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleDynamoDB_BatchGetItem() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.BatchGetItemInput{
		RequestItems: map[string]*dynamodb.KeysAndAttributes{ // Required
			"Key": { // Required
				Keys: []map[string]*dynamodb.AttributeValue{ // Required
					{ // Required
						"Key": { // Required
							B:    []byte("PAYLOAD"),
							BOOL: aws.Bool(true),
							BS: [][]byte{
								[]byte("PAYLOAD"), // Required
								// More values...
							},
							L: []*dynamodb.AttributeValue{
								{ // Required
								// Recursive values...
								},
								// More values...
							},
							M: map[string]*dynamodb.AttributeValue{
								"Key": { // Required
								// Recursive values...
								},
								// More values...
							},
							N: aws.String("NumberAttributeValue"),
							NS: []*string{
								aws.String("NumberAttributeValue"), // Required
								// More values...
							},
							NULL: aws.Bool(true),
							S:    aws.String("StringAttributeValue"),
							SS: []*string{
								aws.String("StringAttributeValue"), // Required
								// More values...
							},
						},
						// More values...
					},
					// More values...
				},
				AttributesToGet: []*string{
					aws.String("AttributeName"), // Required
					// More values...
				},
				ConsistentRead: aws.Bool(true),
				ExpressionAttributeNames: map[string]*string{
					"Key": aws.String("AttributeName"), // Required
					// More values...
				},
				ProjectionExpression: aws.String("ProjectionExpression"),
			},
			// More values...
		},
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
	}
	resp, err := svc.BatchGetItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_BatchWriteItem() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.BatchWriteItemInput{
		RequestItems: map[string][]*dynamodb.WriteRequest{ // Required
			"Key": { // Required
				{ // Required
					DeleteRequest: &dynamodb.DeleteRequest{
						Key: map[string]*dynamodb.AttributeValue{ // Required
							"Key": { // Required
								B:    []byte("PAYLOAD"),
								BOOL: aws.Bool(true),
								BS: [][]byte{
									[]byte("PAYLOAD"), // Required
									// More values...
								},
								L: []*dynamodb.AttributeValue{
									{ // Required
									// Recursive values...
									},
									// More values...
								},
								M: map[string]*dynamodb.AttributeValue{
									"Key": { // Required
									// Recursive values...
									},
									// More values...
								},
								N: aws.String("NumberAttributeValue"),
								NS: []*string{
									aws.String("NumberAttributeValue"), // Required
									// More values...
								},
								NULL: aws.Bool(true),
								S:    aws.String("StringAttributeValue"),
								SS: []*string{
									aws.String("StringAttributeValue"), // Required
									// More values...
								},
							},
							// More values...
						},
					},
					PutRequest: &dynamodb.PutRequest{
						Item: map[string]*dynamodb.AttributeValue{ // Required
							"Key": { // Required
								B:    []byte("PAYLOAD"),
								BOOL: aws.Bool(true),
								BS: [][]byte{
									[]byte("PAYLOAD"), // Required
									// More values...
								},
								L: []*dynamodb.AttributeValue{
									{ // Required
									// Recursive values...
									},
									// More values...
								},
								M: map[string]*dynamodb.AttributeValue{
									"Key": { // Required
									// Recursive values...
									},
									// More values...
								},
								N: aws.String("NumberAttributeValue"),
								NS: []*string{
									aws.String("NumberAttributeValue"), // Required
									// More values...
								},
								NULL: aws.Bool(true),
								S:    aws.String("StringAttributeValue"),
								SS: []*string{
									aws.String("StringAttributeValue"), // Required
									// More values...
								},
							},
							// More values...
						},
					},
				},
				// More values...
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
	}
	resp, err := svc.BatchWriteItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_CreateTable() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{ // Required
			{ // Required
				AttributeName: aws.String("KeySchemaAttributeName"), // Required
				AttributeType: aws.String("ScalarAttributeType"),    // Required
			},
			// More values...
		},
		KeySchema: []*dynamodb.KeySchemaElement{ // Required
			{ // Required
				AttributeName: aws.String("KeySchemaAttributeName"), // Required
				KeyType:       aws.String("KeyType"),                // Required
			},
			// More values...
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
			ReadCapacityUnits:  aws.Int64(1), // Required
			WriteCapacityUnits: aws.Int64(1), // Required
		},
		TableName: aws.String("TableName"), // Required
		GlobalSecondaryIndexes: []*dynamodb.GlobalSecondaryIndex{
			{ // Required
				IndexName: aws.String("IndexName"), // Required
				KeySchema: []*dynamodb.KeySchemaElement{ // Required
					{ // Required
						AttributeName: aws.String("KeySchemaAttributeName"), // Required
						KeyType:       aws.String("KeyType"),                // Required
					},
					// More values...
				},
				Projection: &dynamodb.Projection{ // Required
					NonKeyAttributes: []*string{
						aws.String("NonKeyAttributeName"), // Required
						// More values...
					},
					ProjectionType: aws.String("ProjectionType"),
				},
				ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
					ReadCapacityUnits:  aws.Int64(1), // Required
					WriteCapacityUnits: aws.Int64(1), // Required
				},
			},
			// More values...
		},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			{ // Required
				IndexName: aws.String("IndexName"), // Required
				KeySchema: []*dynamodb.KeySchemaElement{ // Required
					{ // Required
						AttributeName: aws.String("KeySchemaAttributeName"), // Required
						KeyType:       aws.String("KeyType"),                // Required
					},
					// More values...
				},
				Projection: &dynamodb.Projection{ // Required
					NonKeyAttributes: []*string{
						aws.String("NonKeyAttributeName"), // Required
						// More values...
					},
					ProjectionType: aws.String("ProjectionType"),
				},
			},
			// More values...
		},
		StreamSpecification: &dynamodb.StreamSpecification{
			StreamEnabled:  aws.Bool(true),
			StreamViewType: aws.String("StreamViewType"),
		},
	}
	resp, err := svc.CreateTable(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_DeleteItem() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{ // Required
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName:           aws.String("TableName"), // Required
		ConditionExpression: aws.String("ConditionExpression"),
		ConditionalOperator: aws.String("ConditionalOperator"),
		Expected: map[string]*dynamodb.ExpectedAttributeValue{
			"Key": { // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Bool(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: map[string]*dynamodb.AttributeValue{
							"Key": { // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Bool(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
				ComparisonOperator: aws.String("ComparisonOperator"),
				Exists:             aws.Bool(true),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Bool(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: map[string]*dynamodb.AttributeValue{
						"Key": { // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Bool(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ExpressionAttributeNames: map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
		ReturnValues:                aws.String("ReturnValue"),
	}
	resp, err := svc.DeleteItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_DeleteTable() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.DeleteTableInput{
		TableName: aws.String("TableName"), // Required
	}
	resp, err := svc.DeleteTable(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_DescribeLimits() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	var params *dynamodb.DescribeLimitsInput
	resp, err := svc.DescribeLimits(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_DescribeTable() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.DescribeTableInput{
		TableName: aws.String("TableName"), // Required
	}
	resp, err := svc.DescribeTable(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_DescribeTimeToLive() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.DescribeTimeToLiveInput{
		TableName: aws.String("TableName"), // Required
	}
	resp, err := svc.DescribeTimeToLive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_GetItem() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{ // Required
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName: aws.String("TableName"), // Required
		AttributesToGet: []*string{
			aws.String("AttributeName"), // Required
			// More values...
		},
		ConsistentRead: aws.Bool(true),
		ExpressionAttributeNames: map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ProjectionExpression:   aws.String("ProjectionExpression"),
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
	}
	resp, err := svc.GetItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_ListTables() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.ListTablesInput{
		ExclusiveStartTableName: aws.String("TableName"),
		Limit: aws.Int64(1),
	}
	resp, err := svc.ListTables(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_ListTagsOfResource() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.ListTagsOfResourceInput{
		ResourceArn: aws.String("ResourceArnString"), // Required
		NextToken:   aws.String("NextTokenString"),
	}
	resp, err := svc.ListTagsOfResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_PutItem() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{ // Required
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName:           aws.String("TableName"), // Required
		ConditionExpression: aws.String("ConditionExpression"),
		ConditionalOperator: aws.String("ConditionalOperator"),
		Expected: map[string]*dynamodb.ExpectedAttributeValue{
			"Key": { // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Bool(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: map[string]*dynamodb.AttributeValue{
							"Key": { // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Bool(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
				ComparisonOperator: aws.String("ComparisonOperator"),
				Exists:             aws.Bool(true),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Bool(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: map[string]*dynamodb.AttributeValue{
						"Key": { // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Bool(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ExpressionAttributeNames: map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
		ReturnValues:                aws.String("ReturnValue"),
	}
	resp, err := svc.PutItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_Query() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.QueryInput{
		TableName: aws.String("TableName"), // Required
		AttributesToGet: []*string{
			aws.String("AttributeName"), // Required
			// More values...
		},
		ConditionalOperator: aws.String("ConditionalOperator"),
		ConsistentRead:      aws.Bool(true),
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ExpressionAttributeNames: map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		FilterExpression:       aws.String("ConditionExpression"),
		IndexName:              aws.String("IndexName"),
		KeyConditionExpression: aws.String("KeyExpression"),
		KeyConditions: map[string]*dynamodb.Condition{
			"Key": { // Required
				ComparisonOperator: aws.String("ComparisonOperator"), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Bool(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: map[string]*dynamodb.AttributeValue{
							"Key": { // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Bool(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
			},
			// More values...
		},
		Limit:                aws.Int64(1),
		ProjectionExpression: aws.String("ProjectionExpression"),
		QueryFilter: map[string]*dynamodb.Condition{
			"Key": { // Required
				ComparisonOperator: aws.String("ComparisonOperator"), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Bool(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: map[string]*dynamodb.AttributeValue{
							"Key": { // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Bool(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
		ScanIndexForward:       aws.Bool(true),
		Select:                 aws.String("Select"),
	}
	resp, err := svc.Query(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_Scan() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String("TableName"), // Required
		AttributesToGet: []*string{
			aws.String("AttributeName"), // Required
			// More values...
		},
		ConditionalOperator: aws.String("ConditionalOperator"),
		ConsistentRead:      aws.Bool(true),
		ExclusiveStartKey: map[string]*dynamodb.AttributeValue{
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ExpressionAttributeNames: map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		FilterExpression:       aws.String("ConditionExpression"),
		IndexName:              aws.String("IndexName"),
		Limit:                  aws.Int64(1),
		ProjectionExpression:   aws.String("ProjectionExpression"),
		ReturnConsumedCapacity: aws.String("ReturnConsumedCapacity"),
		ScanFilter: map[string]*dynamodb.Condition{
			"Key": { // Required
				ComparisonOperator: aws.String("ComparisonOperator"), // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Bool(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: map[string]*dynamodb.AttributeValue{
							"Key": { // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Bool(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
			},
			// More values...
		},
		Segment:       aws.Int64(1),
		Select:        aws.String("Select"),
		TotalSegments: aws.Int64(1),
	}
	resp, err := svc.Scan(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_TagResource() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.TagResourceInput{
		ResourceArn: aws.String("ResourceArnString"), // Required
		Tags: []*dynamodb.Tag{ // Required
			{ // Required
				Key:   aws.String("TagKeyString"),   // Required
				Value: aws.String("TagValueString"), // Required
			},
			// More values...
		},
	}
	resp, err := svc.TagResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_UntagResource() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.UntagResourceInput{
		ResourceArn: aws.String("ResourceArnString"), // Required
		TagKeys: []*string{ // Required
			aws.String("TagKeyString"), // Required
			// More values...
		},
	}
	resp, err := svc.UntagResource(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_UpdateItem() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{ // Required
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		TableName: aws.String("TableName"), // Required
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			"Key": { // Required
				Action: aws.String("AttributeAction"),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Bool(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: map[string]*dynamodb.AttributeValue{
						"Key": { // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Bool(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ConditionExpression: aws.String("ConditionExpression"),
		ConditionalOperator: aws.String("ConditionalOperator"),
		Expected: map[string]*dynamodb.ExpectedAttributeValue{
			"Key": { // Required
				AttributeValueList: []*dynamodb.AttributeValue{
					{ // Required
						B:    []byte("PAYLOAD"),
						BOOL: aws.Bool(true),
						BS: [][]byte{
							[]byte("PAYLOAD"), // Required
							// More values...
						},
						L: []*dynamodb.AttributeValue{
							{ // Required
							// Recursive values...
							},
							// More values...
						},
						M: map[string]*dynamodb.AttributeValue{
							"Key": { // Required
							// Recursive values...
							},
							// More values...
						},
						N: aws.String("NumberAttributeValue"),
						NS: []*string{
							aws.String("NumberAttributeValue"), // Required
							// More values...
						},
						NULL: aws.Bool(true),
						S:    aws.String("StringAttributeValue"),
						SS: []*string{
							aws.String("StringAttributeValue"), // Required
							// More values...
						},
					},
					// More values...
				},
				ComparisonOperator: aws.String("ComparisonOperator"),
				Exists:             aws.Bool(true),
				Value: &dynamodb.AttributeValue{
					B:    []byte("PAYLOAD"),
					BOOL: aws.Bool(true),
					BS: [][]byte{
						[]byte("PAYLOAD"), // Required
						// More values...
					},
					L: []*dynamodb.AttributeValue{
						{ // Required
						// Recursive values...
						},
						// More values...
					},
					M: map[string]*dynamodb.AttributeValue{
						"Key": { // Required
						// Recursive values...
						},
						// More values...
					},
					N: aws.String("NumberAttributeValue"),
					NS: []*string{
						aws.String("NumberAttributeValue"), // Required
						// More values...
					},
					NULL: aws.Bool(true),
					S:    aws.String("StringAttributeValue"),
					SS: []*string{
						aws.String("StringAttributeValue"), // Required
						// More values...
					},
				},
			},
			// More values...
		},
		ExpressionAttributeNames: map[string]*string{
			"Key": aws.String("AttributeName"), // Required
			// More values...
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			"Key": { // Required
				B:    []byte("PAYLOAD"),
				BOOL: aws.Bool(true),
				BS: [][]byte{
					[]byte("PAYLOAD"), // Required
					// More values...
				},
				L: []*dynamodb.AttributeValue{
					{ // Required
					// Recursive values...
					},
					// More values...
				},
				M: map[string]*dynamodb.AttributeValue{
					"Key": { // Required
					// Recursive values...
					},
					// More values...
				},
				N: aws.String("NumberAttributeValue"),
				NS: []*string{
					aws.String("NumberAttributeValue"), // Required
					// More values...
				},
				NULL: aws.Bool(true),
				S:    aws.String("StringAttributeValue"),
				SS: []*string{
					aws.String("StringAttributeValue"), // Required
					// More values...
				},
			},
			// More values...
		},
		ReturnConsumedCapacity:      aws.String("ReturnConsumedCapacity"),
		ReturnItemCollectionMetrics: aws.String("ReturnItemCollectionMetrics"),
		ReturnValues:                aws.String("ReturnValue"),
		UpdateExpression:            aws.String("UpdateExpression"),
	}
	resp, err := svc.UpdateItem(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_UpdateTable() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.UpdateTableInput{
		TableName: aws.String("TableName"), // Required
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{ // Required
				AttributeName: aws.String("KeySchemaAttributeName"), // Required
				AttributeType: aws.String("ScalarAttributeType"),    // Required
			},
			// More values...
		},
		GlobalSecondaryIndexUpdates: []*dynamodb.GlobalSecondaryIndexUpdate{
			{ // Required
				Create: &dynamodb.CreateGlobalSecondaryIndexAction{
					IndexName: aws.String("IndexName"), // Required
					KeySchema: []*dynamodb.KeySchemaElement{ // Required
						{ // Required
							AttributeName: aws.String("KeySchemaAttributeName"), // Required
							KeyType:       aws.String("KeyType"),                // Required
						},
						// More values...
					},
					Projection: &dynamodb.Projection{ // Required
						NonKeyAttributes: []*string{
							aws.String("NonKeyAttributeName"), // Required
							// More values...
						},
						ProjectionType: aws.String("ProjectionType"),
					},
					ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
						ReadCapacityUnits:  aws.Int64(1), // Required
						WriteCapacityUnits: aws.Int64(1), // Required
					},
				},
				Delete: &dynamodb.DeleteGlobalSecondaryIndexAction{
					IndexName: aws.String("IndexName"), // Required
				},
				Update: &dynamodb.UpdateGlobalSecondaryIndexAction{
					IndexName: aws.String("IndexName"), // Required
					ProvisionedThroughput: &dynamodb.ProvisionedThroughput{ // Required
						ReadCapacityUnits:  aws.Int64(1), // Required
						WriteCapacityUnits: aws.Int64(1), // Required
					},
				},
			},
			// More values...
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1), // Required
			WriteCapacityUnits: aws.Int64(1), // Required
		},
		StreamSpecification: &dynamodb.StreamSpecification{
			StreamEnabled:  aws.Bool(true),
			StreamViewType: aws.String("StreamViewType"),
		},
	}
	resp, err := svc.UpdateTable(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExampleDynamoDB_UpdateTimeToLive() {
	sess := session.Must(session.NewSession())

	svc := dynamodb.New(sess)

	params := &dynamodb.UpdateTimeToLiveInput{
		TableName: aws.String("TableName"), // Required
		TimeToLiveSpecification: &dynamodb.TimeToLiveSpecification{ // Required
			AttributeName: aws.String("TimeToLiveAttributeName"), // Required
			Enabled:       aws.Bool(true),                        // Required
		},
	}
	resp, err := svc.UpdateTimeToLive(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
