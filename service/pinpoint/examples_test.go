// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package pinpoint_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/tily/sdk-go/aws"
	"github.com/tily/sdk-go/aws/session"
	"github.com/tily/sdk-go/service/pinpoint"
)

var _ time.Duration
var _ bytes.Buffer

func ExamplePinpoint_CreateCampaign() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.CreateCampaignInput{
		ApplicationId: aws.String("__string"), // Required
		WriteCampaignRequest: &pinpoint.WriteCampaignRequest{ // Required
			AdditionalTreatments: []*pinpoint.WriteTreatmentResource{
				{ // Required
					MessageConfiguration: &pinpoint.MessageConfiguration{
						APNSMessage: &pinpoint.Message{
							Action:       aws.String("Action"),
							Body:         aws.String("__string"),
							ImageIconUrl: aws.String("__string"),
							ImageUrl:     aws.String("__string"),
							JsonBody:     aws.String("__string"),
							MediaUrl:     aws.String("__string"),
							SilentPush:   aws.Bool(true),
							Title:        aws.String("__string"),
							Url:          aws.String("__string"),
						},
						DefaultMessage: &pinpoint.Message{
							Action:       aws.String("Action"),
							Body:         aws.String("__string"),
							ImageIconUrl: aws.String("__string"),
							ImageUrl:     aws.String("__string"),
							JsonBody:     aws.String("__string"),
							MediaUrl:     aws.String("__string"),
							SilentPush:   aws.Bool(true),
							Title:        aws.String("__string"),
							Url:          aws.String("__string"),
						},
						GCMMessage: &pinpoint.Message{
							Action:       aws.String("Action"),
							Body:         aws.String("__string"),
							ImageIconUrl: aws.String("__string"),
							ImageUrl:     aws.String("__string"),
							JsonBody:     aws.String("__string"),
							MediaUrl:     aws.String("__string"),
							SilentPush:   aws.Bool(true),
							Title:        aws.String("__string"),
							Url:          aws.String("__string"),
						},
					},
					Schedule: &pinpoint.Schedule{
						EndTime:     aws.String("__string"),
						Frequency:   aws.String("Frequency"),
						IsLocalTime: aws.Bool(true),
						QuietTime: &pinpoint.QuietTime{
							End:   aws.String("__string"),
							Start: aws.String("__string"),
						},
						StartTime: aws.String("__string"),
						Timezone:  aws.String("__string"),
					},
					SizePercent:          aws.Int64(1),
					TreatmentDescription: aws.String("__string"),
					TreatmentName:        aws.String("__string"),
				},
				// More values...
			},
			Description:    aws.String("__string"),
			HoldoutPercent: aws.Int64(1),
			IsPaused:       aws.Bool(true),
			Limits: &pinpoint.CampaignLimits{
				Daily: aws.Int64(1),
				Total: aws.Int64(1),
			},
			MessageConfiguration: &pinpoint.MessageConfiguration{
				APNSMessage: &pinpoint.Message{
					Action:       aws.String("Action"),
					Body:         aws.String("__string"),
					ImageIconUrl: aws.String("__string"),
					ImageUrl:     aws.String("__string"),
					JsonBody:     aws.String("__string"),
					MediaUrl:     aws.String("__string"),
					SilentPush:   aws.Bool(true),
					Title:        aws.String("__string"),
					Url:          aws.String("__string"),
				},
				DefaultMessage: &pinpoint.Message{
					Action:       aws.String("Action"),
					Body:         aws.String("__string"),
					ImageIconUrl: aws.String("__string"),
					ImageUrl:     aws.String("__string"),
					JsonBody:     aws.String("__string"),
					MediaUrl:     aws.String("__string"),
					SilentPush:   aws.Bool(true),
					Title:        aws.String("__string"),
					Url:          aws.String("__string"),
				},
				GCMMessage: &pinpoint.Message{
					Action:       aws.String("Action"),
					Body:         aws.String("__string"),
					ImageIconUrl: aws.String("__string"),
					ImageUrl:     aws.String("__string"),
					JsonBody:     aws.String("__string"),
					MediaUrl:     aws.String("__string"),
					SilentPush:   aws.Bool(true),
					Title:        aws.String("__string"),
					Url:          aws.String("__string"),
				},
			},
			Name: aws.String("__string"),
			Schedule: &pinpoint.Schedule{
				EndTime:     aws.String("__string"),
				Frequency:   aws.String("Frequency"),
				IsLocalTime: aws.Bool(true),
				QuietTime: &pinpoint.QuietTime{
					End:   aws.String("__string"),
					Start: aws.String("__string"),
				},
				StartTime: aws.String("__string"),
				Timezone:  aws.String("__string"),
			},
			SegmentId:            aws.String("__string"),
			SegmentVersion:       aws.Int64(1),
			TreatmentDescription: aws.String("__string"),
			TreatmentName:        aws.String("__string"),
		},
	}
	resp, err := svc.CreateCampaign(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_CreateImportJob() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.CreateImportJobInput{
		ApplicationId: aws.String("__string"), // Required
		ImportJobRequest: &pinpoint.ImportJobRequest{ // Required
			DefineSegment:     aws.Bool(true),
			ExternalId:        aws.String("__string"),
			Format:            aws.String("Format"),
			RegisterEndpoints: aws.Bool(true),
			RoleArn:           aws.String("__string"),
			S3Url:             aws.String("__string"),
			SegmentId:         aws.String("__string"),
			SegmentName:       aws.String("__string"),
		},
	}
	resp, err := svc.CreateImportJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_CreateSegment() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.CreateSegmentInput{
		ApplicationId: aws.String("__string"), // Required
		WriteSegmentRequest: &pinpoint.WriteSegmentRequest{ // Required
			Dimensions: &pinpoint.SegmentDimensions{
				Attributes: map[string]*pinpoint.AttributeDimension{
					"Key": { // Required
						AttributeType: aws.String("AttributeType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					// More values...
				},
				Behavior: &pinpoint.SegmentBehaviors{
					Recency: &pinpoint.RecencyDimension{
						Duration:    aws.String("Duration"),
						RecencyType: aws.String("RecencyType"),
					},
				},
				Demographic: &pinpoint.SegmentDemographics{
					AppVersion: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					DeviceType: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					Make: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					Model: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					Platform: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
				},
				Location: &pinpoint.SegmentLocation{
					Country: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
				},
			},
			Name: aws.String("__string"),
		},
	}
	resp, err := svc.CreateSegment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_DeleteApnsChannel() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.DeleteApnsChannelInput{
		ApplicationId: aws.String("__string"), // Required
	}
	resp, err := svc.DeleteApnsChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_DeleteCampaign() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.DeleteCampaignInput{
		ApplicationId: aws.String("__string"), // Required
		CampaignId:    aws.String("__string"), // Required
	}
	resp, err := svc.DeleteCampaign(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_DeleteGcmChannel() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.DeleteGcmChannelInput{
		ApplicationId: aws.String("__string"), // Required
	}
	resp, err := svc.DeleteGcmChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_DeleteSegment() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.DeleteSegmentInput{
		ApplicationId: aws.String("__string"), // Required
		SegmentId:     aws.String("__string"), // Required
	}
	resp, err := svc.DeleteSegment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetApnsChannel() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetApnsChannelInput{
		ApplicationId: aws.String("__string"), // Required
	}
	resp, err := svc.GetApnsChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetApplicationSettings() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetApplicationSettingsInput{
		ApplicationId: aws.String("__string"), // Required
	}
	resp, err := svc.GetApplicationSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetCampaign() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetCampaignInput{
		ApplicationId: aws.String("__string"), // Required
		CampaignId:    aws.String("__string"), // Required
	}
	resp, err := svc.GetCampaign(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetCampaignActivities() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetCampaignActivitiesInput{
		ApplicationId: aws.String("__string"), // Required
		CampaignId:    aws.String("__string"), // Required
		PageSize:      aws.String("__string"),
		Token:         aws.String("__string"),
	}
	resp, err := svc.GetCampaignActivities(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetCampaignVersion() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetCampaignVersionInput{
		ApplicationId: aws.String("__string"), // Required
		CampaignId:    aws.String("__string"), // Required
		Version:       aws.String("__string"), // Required
	}
	resp, err := svc.GetCampaignVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetCampaignVersions() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetCampaignVersionsInput{
		ApplicationId: aws.String("__string"), // Required
		CampaignId:    aws.String("__string"), // Required
		PageSize:      aws.String("__string"),
		Token:         aws.String("__string"),
	}
	resp, err := svc.GetCampaignVersions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetCampaigns() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetCampaignsInput{
		ApplicationId: aws.String("__string"), // Required
		PageSize:      aws.String("__string"),
		Token:         aws.String("__string"),
	}
	resp, err := svc.GetCampaigns(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetEndpoint() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetEndpointInput{
		ApplicationId: aws.String("__string"), // Required
		EndpointId:    aws.String("__string"), // Required
	}
	resp, err := svc.GetEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetGcmChannel() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetGcmChannelInput{
		ApplicationId: aws.String("__string"), // Required
	}
	resp, err := svc.GetGcmChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetImportJob() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetImportJobInput{
		ApplicationId: aws.String("__string"), // Required
		JobId:         aws.String("__string"), // Required
	}
	resp, err := svc.GetImportJob(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetImportJobs() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetImportJobsInput{
		ApplicationId: aws.String("__string"), // Required
		PageSize:      aws.String("__string"),
		Token:         aws.String("__string"),
	}
	resp, err := svc.GetImportJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetSegment() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetSegmentInput{
		ApplicationId: aws.String("__string"), // Required
		SegmentId:     aws.String("__string"), // Required
	}
	resp, err := svc.GetSegment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetSegmentImportJobs() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetSegmentImportJobsInput{
		ApplicationId: aws.String("__string"), // Required
		SegmentId:     aws.String("__string"), // Required
		PageSize:      aws.String("__string"),
		Token:         aws.String("__string"),
	}
	resp, err := svc.GetSegmentImportJobs(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetSegmentVersion() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetSegmentVersionInput{
		ApplicationId: aws.String("__string"), // Required
		SegmentId:     aws.String("__string"), // Required
		Version:       aws.String("__string"), // Required
	}
	resp, err := svc.GetSegmentVersion(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetSegmentVersions() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetSegmentVersionsInput{
		ApplicationId: aws.String("__string"), // Required
		SegmentId:     aws.String("__string"), // Required
		PageSize:      aws.String("__string"),
		Token:         aws.String("__string"),
	}
	resp, err := svc.GetSegmentVersions(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_GetSegments() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.GetSegmentsInput{
		ApplicationId: aws.String("__string"), // Required
		PageSize:      aws.String("__string"),
		Token:         aws.String("__string"),
	}
	resp, err := svc.GetSegments(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_UpdateApnsChannel() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.UpdateApnsChannelInput{
		APNSChannelRequest: &pinpoint.APNSChannelRequest{ // Required
			Certificate: aws.String("__string"),
			PrivateKey:  aws.String("__string"),
		},
		ApplicationId: aws.String("__string"), // Required
	}
	resp, err := svc.UpdateApnsChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_UpdateApplicationSettings() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.UpdateApplicationSettingsInput{
		ApplicationId: aws.String("__string"), // Required
		WriteApplicationSettingsRequest: &pinpoint.WriteApplicationSettingsRequest{ // Required
			Limits: &pinpoint.CampaignLimits{
				Daily: aws.Int64(1),
				Total: aws.Int64(1),
			},
			QuietTime: &pinpoint.QuietTime{
				End:   aws.String("__string"),
				Start: aws.String("__string"),
			},
		},
	}
	resp, err := svc.UpdateApplicationSettings(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_UpdateCampaign() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.UpdateCampaignInput{
		ApplicationId: aws.String("__string"), // Required
		CampaignId:    aws.String("__string"), // Required
		WriteCampaignRequest: &pinpoint.WriteCampaignRequest{ // Required
			AdditionalTreatments: []*pinpoint.WriteTreatmentResource{
				{ // Required
					MessageConfiguration: &pinpoint.MessageConfiguration{
						APNSMessage: &pinpoint.Message{
							Action:       aws.String("Action"),
							Body:         aws.String("__string"),
							ImageIconUrl: aws.String("__string"),
							ImageUrl:     aws.String("__string"),
							JsonBody:     aws.String("__string"),
							MediaUrl:     aws.String("__string"),
							SilentPush:   aws.Bool(true),
							Title:        aws.String("__string"),
							Url:          aws.String("__string"),
						},
						DefaultMessage: &pinpoint.Message{
							Action:       aws.String("Action"),
							Body:         aws.String("__string"),
							ImageIconUrl: aws.String("__string"),
							ImageUrl:     aws.String("__string"),
							JsonBody:     aws.String("__string"),
							MediaUrl:     aws.String("__string"),
							SilentPush:   aws.Bool(true),
							Title:        aws.String("__string"),
							Url:          aws.String("__string"),
						},
						GCMMessage: &pinpoint.Message{
							Action:       aws.String("Action"),
							Body:         aws.String("__string"),
							ImageIconUrl: aws.String("__string"),
							ImageUrl:     aws.String("__string"),
							JsonBody:     aws.String("__string"),
							MediaUrl:     aws.String("__string"),
							SilentPush:   aws.Bool(true),
							Title:        aws.String("__string"),
							Url:          aws.String("__string"),
						},
					},
					Schedule: &pinpoint.Schedule{
						EndTime:     aws.String("__string"),
						Frequency:   aws.String("Frequency"),
						IsLocalTime: aws.Bool(true),
						QuietTime: &pinpoint.QuietTime{
							End:   aws.String("__string"),
							Start: aws.String("__string"),
						},
						StartTime: aws.String("__string"),
						Timezone:  aws.String("__string"),
					},
					SizePercent:          aws.Int64(1),
					TreatmentDescription: aws.String("__string"),
					TreatmentName:        aws.String("__string"),
				},
				// More values...
			},
			Description:    aws.String("__string"),
			HoldoutPercent: aws.Int64(1),
			IsPaused:       aws.Bool(true),
			Limits: &pinpoint.CampaignLimits{
				Daily: aws.Int64(1),
				Total: aws.Int64(1),
			},
			MessageConfiguration: &pinpoint.MessageConfiguration{
				APNSMessage: &pinpoint.Message{
					Action:       aws.String("Action"),
					Body:         aws.String("__string"),
					ImageIconUrl: aws.String("__string"),
					ImageUrl:     aws.String("__string"),
					JsonBody:     aws.String("__string"),
					MediaUrl:     aws.String("__string"),
					SilentPush:   aws.Bool(true),
					Title:        aws.String("__string"),
					Url:          aws.String("__string"),
				},
				DefaultMessage: &pinpoint.Message{
					Action:       aws.String("Action"),
					Body:         aws.String("__string"),
					ImageIconUrl: aws.String("__string"),
					ImageUrl:     aws.String("__string"),
					JsonBody:     aws.String("__string"),
					MediaUrl:     aws.String("__string"),
					SilentPush:   aws.Bool(true),
					Title:        aws.String("__string"),
					Url:          aws.String("__string"),
				},
				GCMMessage: &pinpoint.Message{
					Action:       aws.String("Action"),
					Body:         aws.String("__string"),
					ImageIconUrl: aws.String("__string"),
					ImageUrl:     aws.String("__string"),
					JsonBody:     aws.String("__string"),
					MediaUrl:     aws.String("__string"),
					SilentPush:   aws.Bool(true),
					Title:        aws.String("__string"),
					Url:          aws.String("__string"),
				},
			},
			Name: aws.String("__string"),
			Schedule: &pinpoint.Schedule{
				EndTime:     aws.String("__string"),
				Frequency:   aws.String("Frequency"),
				IsLocalTime: aws.Bool(true),
				QuietTime: &pinpoint.QuietTime{
					End:   aws.String("__string"),
					Start: aws.String("__string"),
				},
				StartTime: aws.String("__string"),
				Timezone:  aws.String("__string"),
			},
			SegmentId:            aws.String("__string"),
			SegmentVersion:       aws.Int64(1),
			TreatmentDescription: aws.String("__string"),
			TreatmentName:        aws.String("__string"),
		},
	}
	resp, err := svc.UpdateCampaign(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_UpdateEndpoint() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.UpdateEndpointInput{
		ApplicationId: aws.String("__string"), // Required
		EndpointId:    aws.String("__string"), // Required
		EndpointRequest: &pinpoint.EndpointRequest{ // Required
			Address: aws.String("__string"),
			Attributes: map[string][]*string{
				"Key": { // Required
					aws.String("__string"), // Required
					// More values...
				},
				// More values...
			},
			ChannelType: aws.String("ChannelType"),
			Demographic: &pinpoint.EndpointDemographic{
				AppVersion:      aws.String("__string"),
				Locale:          aws.String("__string"),
				Make:            aws.String("__string"),
				Model:           aws.String("__string"),
				ModelVersion:    aws.String("__string"),
				Platform:        aws.String("__string"),
				PlatformVersion: aws.String("__string"),
				Timezone:        aws.String("__string"),
			},
			EffectiveDate:  aws.String("__string"),
			EndpointStatus: aws.String("__string"),
			Location: &pinpoint.EndpointLocation{
				City:       aws.String("__string"),
				Country:    aws.String("__string"),
				Latitude:   aws.Float64(1.0),
				Longitude:  aws.Float64(1.0),
				PostalCode: aws.String("__string"),
				Region:     aws.String("__string"),
			},
			Metrics: map[string]*float64{
				"Key": aws.Float64(1.0), // Required
				// More values...
			},
			OptOut:    aws.String("__string"),
			RequestId: aws.String("__string"),
			User: &pinpoint.EndpointUser{
				UserAttributes: map[string][]*string{
					"Key": { // Required
						aws.String("__string"), // Required
						// More values...
					},
					// More values...
				},
				UserId: aws.String("__string"),
			},
		},
	}
	resp, err := svc.UpdateEndpoint(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_UpdateEndpointsBatch() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.UpdateEndpointsBatchInput{
		ApplicationId: aws.String("__string"), // Required
		EndpointBatchRequest: &pinpoint.EndpointBatchRequest{ // Required
			Item: []*pinpoint.EndpointBatchItem{
				{ // Required
					Address: aws.String("__string"),
					Attributes: map[string][]*string{
						"Key": { // Required
							aws.String("__string"), // Required
							// More values...
						},
						// More values...
					},
					ChannelType: aws.String("ChannelType"),
					Demographic: &pinpoint.EndpointDemographic{
						AppVersion:      aws.String("__string"),
						Locale:          aws.String("__string"),
						Make:            aws.String("__string"),
						Model:           aws.String("__string"),
						ModelVersion:    aws.String("__string"),
						Platform:        aws.String("__string"),
						PlatformVersion: aws.String("__string"),
						Timezone:        aws.String("__string"),
					},
					EffectiveDate:  aws.String("__string"),
					EndpointStatus: aws.String("__string"),
					Id:             aws.String("__string"),
					Location: &pinpoint.EndpointLocation{
						City:       aws.String("__string"),
						Country:    aws.String("__string"),
						Latitude:   aws.Float64(1.0),
						Longitude:  aws.Float64(1.0),
						PostalCode: aws.String("__string"),
						Region:     aws.String("__string"),
					},
					Metrics: map[string]*float64{
						"Key": aws.Float64(1.0), // Required
						// More values...
					},
					OptOut:    aws.String("__string"),
					RequestId: aws.String("__string"),
					User: &pinpoint.EndpointUser{
						UserAttributes: map[string][]*string{
							"Key": { // Required
								aws.String("__string"), // Required
								// More values...
							},
							// More values...
						},
						UserId: aws.String("__string"),
					},
				},
				// More values...
			},
		},
	}
	resp, err := svc.UpdateEndpointsBatch(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_UpdateGcmChannel() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.UpdateGcmChannelInput{
		ApplicationId: aws.String("__string"), // Required
		GCMChannelRequest: &pinpoint.GCMChannelRequest{ // Required
			ApiKey: aws.String("__string"),
		},
	}
	resp, err := svc.UpdateGcmChannel(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}

func ExamplePinpoint_UpdateSegment() {
	sess := session.Must(session.NewSession())

	svc := pinpoint.New(sess)

	params := &pinpoint.UpdateSegmentInput{
		ApplicationId: aws.String("__string"), // Required
		SegmentId:     aws.String("__string"), // Required
		WriteSegmentRequest: &pinpoint.WriteSegmentRequest{ // Required
			Dimensions: &pinpoint.SegmentDimensions{
				Attributes: map[string]*pinpoint.AttributeDimension{
					"Key": { // Required
						AttributeType: aws.String("AttributeType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					// More values...
				},
				Behavior: &pinpoint.SegmentBehaviors{
					Recency: &pinpoint.RecencyDimension{
						Duration:    aws.String("Duration"),
						RecencyType: aws.String("RecencyType"),
					},
				},
				Demographic: &pinpoint.SegmentDemographics{
					AppVersion: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					DeviceType: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					Make: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					Model: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
					Platform: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
				},
				Location: &pinpoint.SegmentLocation{
					Country: &pinpoint.SetDimension{
						DimensionType: aws.String("DimensionType"),
						Values: []*string{
							aws.String("__string"), // Required
							// More values...
						},
					},
				},
			},
			Name: aws.String("__string"),
		},
	}
	resp, err := svc.UpdateSegment(params)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println(resp)
}
