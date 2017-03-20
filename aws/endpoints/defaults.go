// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package endpoints

import (
	"regexp"
)

// Partition identifiers
const (
	NiftycloudPartitionID = "niftycloud" // NIFTYCloud Standard partition.
)

// NIFTYCloud Standard partition's regions.
const (
	JpEast1RegionID = "jp-east-1" // jp-east-1.
	JpEast2RegionID = "jp-east-2" // jp-east-2.
	JpWest1RegionID = "jp-west-1" // jp-west-1.
)

// Service identifiers
const (
	ComputingServiceID     = "computing"     // Computing.
	Ec2metadataServiceID   = "ec2metadata"   // Ec2metadata.
	EmailServiceID         = "email"         // Email.
	ObjectstorageServiceID = "objectstorage" // Objectstorage.
	RdbServiceID           = "rdb"           // Rdb.
	StorageServiceID       = "storage"       // Storage.
)

// DefaultResolver returns an Endpoint resolver that will be able
// to resolve endpoints for: NIFTYCloud Standard.
//
// Casting the return value of this func to a EnumPartitions will
// allow you to get a list of the partitions in the order the endpoints
// will be resolved in.
//
//    resolver := endpoints.DefaultResolver()
//    partitions := resolver.(endpoints.EnumPartitions).Partitions()
//    for _, p := range partitions {
//        // ... inspect partitions
//    }
func DefaultResolver() Resolver {
	return defaultPartitions
}

var defaultPartitions = partitions{
	niftycloudPartition,
}

// NiftycloudPartition returns the Resolver for NIFTYCloud Standard.
func NiftycloudPartition() Partition {
	return niftycloudPartition.Partition()
}

var niftycloudPartition = partition{
	ID:        "niftycloud",
	Name:      "NIFTYCloud Standard",
	DNSSuffix: "api.cloud.nifty.com",
	RegionRegex: regionRegex{
		Regexp: func() *regexp.Regexp {
			reg, _ := regexp.Compile("^(us|jp)\\-\\w+\\-\\d+$")
			return reg
		}(),
	},
	Defaults: endpoint{
		Hostname:          "{service}.{region}.{dnsSuffix}",
		Protocols:         []string{"https"},
		SignatureVersions: []string{"v4", "s3", "n2"},
	},
	Regions: regions{
		"jp-east-1": region{
			Description: "jp-east-1",
		},
		"jp-east-2": region{
			Description: "jp-east-2",
		},
		"jp-west-1": region{
			Description: "jp-west-1",
		},
	},
	Services: services{
		"computing": service{
			Defaults: endpoint{
				Protocols:         []string{"https"},
				SignatureVersions: []string{"n2"},
			},
			Endpoints: endpoints{
				"jp-east-1": endpoint{
					Hostname:          "east-1.cp.cloud.nifty.com",
					SignatureVersions: []string{"n2"},
				},
				"jp-east-2": endpoint{
					Hostname:          "east-2.cp.cloud.nifty.com",
					SignatureVersions: []string{"n2"},
				},
			},
		},
		"ec2metadata": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:  "169.254.169.254/latest",
					Protocols: []string{"http"},
				},
			},
		},
		"email": service{
			PartitionEndpoint: "aws-global",
			IsRegionalized:    boxedFalse,

			Endpoints: endpoints{
				"aws-global": endpoint{
					Hostname:          "ess.api.cloud.nifty.com",
					SignatureVersions: []string{"v4"},
					CredentialScope: credentialScope{
						Region: "east-1",
					},
				},
			},
		},
		"objectstorage": service{
			Defaults: endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3"},
			},
			Endpoints: endpoints{
				"jp-east-2": endpoint{
					Hostname:          "jp-east-2.os.cloud.nifty.com",
					SignatureVersions: []string{"s3"},
				},
			},
		},
		"rdb": service{

			Endpoints: endpoints{
				"jp-east-1": endpoint{},
				"jp-east-2": endpoint{},
			},
		},
		"storage": service{
			Defaults: endpoint{
				Protocols:         []string{"http", "https"},
				SignatureVersions: []string{"s3"},
			},
			Endpoints: endpoints{
				"jp-east-1": endpoint{
					Hostname:          "ncss.nifty.com",
					SignatureVersions: []string{"s3"},
				},
				"jp-west-1": endpoint{
					Hostname:          "west-1-ncss.nifty.com",
					SignatureVersions: []string{"s3"},
				},
			},
		},
	},
}
