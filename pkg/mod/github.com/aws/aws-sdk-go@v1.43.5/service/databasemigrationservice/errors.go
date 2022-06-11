// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedFault for service response error code
	// "AccessDeniedFault".
	//
	// DMS was denied access to the endpoint. Check that the role is correctly configured.
	ErrCodeAccessDeniedFault = "AccessDeniedFault"

	// ErrCodeInsufficientResourceCapacityFault for service response error code
	// "InsufficientResourceCapacityFault".
	//
	// There are not enough resources allocated to the database migration.
	ErrCodeInsufficientResourceCapacityFault = "InsufficientResourceCapacityFault"

	// ErrCodeInvalidCertificateFault for service response error code
	// "InvalidCertificateFault".
	//
	// The certificate was not valid.
	ErrCodeInvalidCertificateFault = "InvalidCertificateFault"

	// ErrCodeInvalidResourceStateFault for service response error code
	// "InvalidResourceStateFault".
	//
	// The resource is in a state that prevents it from being used for database
	// migration.
	ErrCodeInvalidResourceStateFault = "InvalidResourceStateFault"

	// ErrCodeInvalidSubnet for service response error code
	// "InvalidSubnet".
	//
	// The subnet provided is invalid.
	ErrCodeInvalidSubnet = "InvalidSubnet"

	// ErrCodeKMSAccessDeniedFault for service response error code
	// "KMSAccessDeniedFault".
	//
	// The ciphertext references a key that doesn't exist or that the DMS account
	// doesn't have access to.
	ErrCodeKMSAccessDeniedFault = "KMSAccessDeniedFault"

	// ErrCodeKMSDisabledFault for service response error code
	// "KMSDisabledFault".
	//
	// The specified KMS key isn't enabled.
	ErrCodeKMSDisabledFault = "KMSDisabledFault"

	// ErrCodeKMSFault for service response error code
	// "KMSFault".
	//
	// An Key Management Service (KMS) error is preventing access to KMS.
	ErrCodeKMSFault = "KMSFault"

	// ErrCodeKMSInvalidStateFault for service response error code
	// "KMSInvalidStateFault".
	//
	// The state of the specified KMS resource isn't valid for this request.
	ErrCodeKMSInvalidStateFault = "KMSInvalidStateFault"

	// ErrCodeKMSKeyNotAccessibleFault for service response error code
	// "KMSKeyNotAccessibleFault".
	//
	// DMS cannot access the KMS key.
	ErrCodeKMSKeyNotAccessibleFault = "KMSKeyNotAccessibleFault"

	// ErrCodeKMSNotFoundFault for service response error code
	// "KMSNotFoundFault".
	//
	// The specified KMS entity or resource can't be found.
	ErrCodeKMSNotFoundFault = "KMSNotFoundFault"

	// ErrCodeKMSThrottlingFault for service response error code
	// "KMSThrottlingFault".
	//
	// This request triggered KMS request throttling.
	ErrCodeKMSThrottlingFault = "KMSThrottlingFault"

	// ErrCodeReplicationSubnetGroupDoesNotCoverEnoughAZs for service response error code
	// "ReplicationSubnetGroupDoesNotCoverEnoughAZs".
	//
	// The replication subnet group does not cover enough Availability Zones (AZs).
	// Edit the replication subnet group and add more AZs.
	ErrCodeReplicationSubnetGroupDoesNotCoverEnoughAZs = "ReplicationSubnetGroupDoesNotCoverEnoughAZs"

	// ErrCodeResourceAlreadyExistsFault for service response error code
	// "ResourceAlreadyExistsFault".
	//
	// The resource you are attempting to create already exists.
	ErrCodeResourceAlreadyExistsFault = "ResourceAlreadyExistsFault"

	// ErrCodeResourceNotFoundFault for service response error code
	// "ResourceNotFoundFault".
	//
	// The resource could not be found.
	ErrCodeResourceNotFoundFault = "ResourceNotFoundFault"

	// ErrCodeResourceQuotaExceededFault for service response error code
	// "ResourceQuotaExceededFault".
	//
	// The quota for this resource quota has been exceeded.
	ErrCodeResourceQuotaExceededFault = "ResourceQuotaExceededFault"

	// ErrCodeS3AccessDeniedFault for service response error code
	// "S3AccessDeniedFault".
	//
	// Insufficient privileges are preventing access to an Amazon S3 object.
	ErrCodeS3AccessDeniedFault = "S3AccessDeniedFault"

	// ErrCodeS3ResourceNotFoundFault for service response error code
	// "S3ResourceNotFoundFault".
	//
	// A specified Amazon S3 bucket, bucket folder, or other object can't be found.
	ErrCodeS3ResourceNotFoundFault = "S3ResourceNotFoundFault"

	// ErrCodeSNSInvalidTopicFault for service response error code
	// "SNSInvalidTopicFault".
	//
	// The SNS topic is invalid.
	ErrCodeSNSInvalidTopicFault = "SNSInvalidTopicFault"

	// ErrCodeSNSNoAuthorizationFault for service response error code
	// "SNSNoAuthorizationFault".
	//
	// You are not authorized for the SNS subscription.
	ErrCodeSNSNoAuthorizationFault = "SNSNoAuthorizationFault"

	// ErrCodeStorageQuotaExceededFault for service response error code
	// "StorageQuotaExceededFault".
	//
	// The storage quota has been exceeded.
	ErrCodeStorageQuotaExceededFault = "StorageQuotaExceededFault"

	// ErrCodeSubnetAlreadyInUse for service response error code
	// "SubnetAlreadyInUse".
	//
	// The specified subnet is already in use.
	ErrCodeSubnetAlreadyInUse = "SubnetAlreadyInUse"

	// ErrCodeUpgradeDependencyFailureFault for service response error code
	// "UpgradeDependencyFailureFault".
	//
	// An upgrade dependency is preventing the database migration.
	ErrCodeUpgradeDependencyFailureFault = "UpgradeDependencyFailureFault"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedFault":                           newErrorAccessDeniedFault,
	"InsufficientResourceCapacityFault":           newErrorInsufficientResourceCapacityFault,
	"InvalidCertificateFault":                     newErrorInvalidCertificateFault,
	"InvalidResourceStateFault":                   newErrorInvalidResourceStateFault,
	"InvalidSubnet":                               newErrorInvalidSubnet,
	"KMSAccessDeniedFault":                        newErrorKMSAccessDeniedFault,
	"KMSDisabledFault":                            newErrorKMSDisabledFault,
	"KMSFault":                                    newErrorKMSFault,
	"KMSInvalidStateFault":                        newErrorKMSInvalidStateFault,
	"KMSKeyNotAccessibleFault":                    newErrorKMSKeyNotAccessibleFault,
	"KMSNotFoundFault":                            newErrorKMSNotFoundFault,
	"KMSThrottlingFault":                          newErrorKMSThrottlingFault,
	"ReplicationSubnetGroupDoesNotCoverEnoughAZs": newErrorReplicationSubnetGroupDoesNotCoverEnoughAZs,
	"ResourceAlreadyExistsFault":                  newErrorResourceAlreadyExistsFault,
	"ResourceNotFoundFault":                       newErrorResourceNotFoundFault,
	"ResourceQuotaExceededFault":                  newErrorResourceQuotaExceededFault,
	"S3AccessDeniedFault":                         newErrorS3AccessDeniedFault,
	"S3ResourceNotFoundFault":                     newErrorS3ResourceNotFoundFault,
	"SNSInvalidTopicFault":                        newErrorSNSInvalidTopicFault,
	"SNSNoAuthorizationFault":                     newErrorSNSNoAuthorizationFault,
	"StorageQuotaExceededFault":                   newErrorStorageQuotaExceededFault,
	"SubnetAlreadyInUse":                          newErrorSubnetAlreadyInUse,
	"UpgradeDependencyFailureFault":               newErrorUpgradeDependencyFailureFault,
}
