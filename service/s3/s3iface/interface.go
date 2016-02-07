// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package s3iface provides an interface for the Amazon Simple Storage Service.
package s3iface

import (
	"github.com/somathor/aws-sdk-go/aws/request"
	"github.com/somathor/aws-sdk-go/service/s3"
)

// S3API is the interface type for s3.S3.
type S3API interface {
	AbortMultipartUploadRequest(*s3.AbortMultipartUploadInput) (*request.Request, *s3.AbortMultipartUploadOutput)

	AbortMultipartUpload(*s3.AbortMultipartUploadInput) (*s3.AbortMultipartUploadOutput, error)

	CompleteMultipartUploadRequest(*s3.CompleteMultipartUploadInput) (*request.Request, *s3.CompleteMultipartUploadOutput)

	CompleteMultipartUpload(*s3.CompleteMultipartUploadInput) (*s3.CompleteMultipartUploadOutput, error)

	CopyObjectRequest(*s3.CopyObjectInput) (*request.Request, *s3.CopyObjectOutput)

	CopyObject(*s3.CopyObjectInput) (*s3.CopyObjectOutput, error)

	CreateBucketRequest(*s3.CreateBucketInput) (*request.Request, *s3.CreateBucketOutput)

	CreateBucket(*s3.CreateBucketInput) (*s3.CreateBucketOutput, error)

	CreateMultipartUploadRequest(*s3.CreateMultipartUploadInput) (*request.Request, *s3.CreateMultipartUploadOutput)

	CreateMultipartUpload(*s3.CreateMultipartUploadInput) (*s3.CreateMultipartUploadOutput, error)

	DeleteBucketRequest(*s3.DeleteBucketInput) (*request.Request, *s3.DeleteBucketOutput)

	DeleteBucket(*s3.DeleteBucketInput) (*s3.DeleteBucketOutput, error)

	DeleteBucketCorsRequest(*s3.DeleteBucketCorsInput) (*request.Request, *s3.DeleteBucketCorsOutput)

	DeleteBucketCors(*s3.DeleteBucketCorsInput) (*s3.DeleteBucketCorsOutput, error)

	DeleteBucketLifecycleRequest(*s3.DeleteBucketLifecycleInput) (*request.Request, *s3.DeleteBucketLifecycleOutput)

	DeleteBucketLifecycle(*s3.DeleteBucketLifecycleInput) (*s3.DeleteBucketLifecycleOutput, error)

	DeleteBucketPolicyRequest(*s3.DeleteBucketPolicyInput) (*request.Request, *s3.DeleteBucketPolicyOutput)

	DeleteBucketPolicy(*s3.DeleteBucketPolicyInput) (*s3.DeleteBucketPolicyOutput, error)

	DeleteBucketReplicationRequest(*s3.DeleteBucketReplicationInput) (*request.Request, *s3.DeleteBucketReplicationOutput)

	DeleteBucketReplication(*s3.DeleteBucketReplicationInput) (*s3.DeleteBucketReplicationOutput, error)

	DeleteBucketTaggingRequest(*s3.DeleteBucketTaggingInput) (*request.Request, *s3.DeleteBucketTaggingOutput)

	DeleteBucketTagging(*s3.DeleteBucketTaggingInput) (*s3.DeleteBucketTaggingOutput, error)

	DeleteBucketWebsiteRequest(*s3.DeleteBucketWebsiteInput) (*request.Request, *s3.DeleteBucketWebsiteOutput)

	DeleteBucketWebsite(*s3.DeleteBucketWebsiteInput) (*s3.DeleteBucketWebsiteOutput, error)

	DeleteObjectRequest(*s3.DeleteObjectInput) (*request.Request, *s3.DeleteObjectOutput)

	DeleteObject(*s3.DeleteObjectInput) (*s3.DeleteObjectOutput, error)

	DeleteObjectsRequest(*s3.DeleteObjectsInput) (*request.Request, *s3.DeleteObjectsOutput)

	DeleteObjects(*s3.DeleteObjectsInput) (*s3.DeleteObjectsOutput, error)

	GetBucketAclRequest(*s3.GetBucketAclInput) (*request.Request, *s3.GetBucketAclOutput)

	GetBucketAcl(*s3.GetBucketAclInput) (*s3.GetBucketAclOutput, error)

	GetBucketCorsRequest(*s3.GetBucketCorsInput) (*request.Request, *s3.GetBucketCorsOutput)

	GetBucketCors(*s3.GetBucketCorsInput) (*s3.GetBucketCorsOutput, error)

	GetBucketLifecycleRequest(*s3.GetBucketLifecycleInput) (*request.Request, *s3.GetBucketLifecycleOutput)

	GetBucketLifecycle(*s3.GetBucketLifecycleInput) (*s3.GetBucketLifecycleOutput, error)

	GetBucketLifecycleConfigurationRequest(*s3.GetBucketLifecycleConfigurationInput) (*request.Request, *s3.GetBucketLifecycleConfigurationOutput)

	GetBucketLifecycleConfiguration(*s3.GetBucketLifecycleConfigurationInput) (*s3.GetBucketLifecycleConfigurationOutput, error)

	GetBucketLocationRequest(*s3.GetBucketLocationInput) (*request.Request, *s3.GetBucketLocationOutput)

	GetBucketLocation(*s3.GetBucketLocationInput) (*s3.GetBucketLocationOutput, error)

	GetBucketLoggingRequest(*s3.GetBucketLoggingInput) (*request.Request, *s3.GetBucketLoggingOutput)

	GetBucketLogging(*s3.GetBucketLoggingInput) (*s3.GetBucketLoggingOutput, error)

	GetBucketNotificationRequest(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfigurationDeprecated)

	GetBucketNotification(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfigurationDeprecated, error)

	GetBucketNotificationConfigurationRequest(*s3.GetBucketNotificationConfigurationRequest) (*request.Request, *s3.NotificationConfiguration)

	GetBucketNotificationConfiguration(*s3.GetBucketNotificationConfigurationRequest) (*s3.NotificationConfiguration, error)

	GetBucketPolicyRequest(*s3.GetBucketPolicyInput) (*request.Request, *s3.GetBucketPolicyOutput)

	GetBucketPolicy(*s3.GetBucketPolicyInput) (*s3.GetBucketPolicyOutput, error)

	GetBucketReplicationRequest(*s3.GetBucketReplicationInput) (*request.Request, *s3.GetBucketReplicationOutput)

	GetBucketReplication(*s3.GetBucketReplicationInput) (*s3.GetBucketReplicationOutput, error)

	GetBucketRequestPaymentRequest(*s3.GetBucketRequestPaymentInput) (*request.Request, *s3.GetBucketRequestPaymentOutput)

	GetBucketRequestPayment(*s3.GetBucketRequestPaymentInput) (*s3.GetBucketRequestPaymentOutput, error)

	GetBucketTaggingRequest(*s3.GetBucketTaggingInput) (*request.Request, *s3.GetBucketTaggingOutput)

	GetBucketTagging(*s3.GetBucketTaggingInput) (*s3.GetBucketTaggingOutput, error)

	GetBucketVersioningRequest(*s3.GetBucketVersioningInput) (*request.Request, *s3.GetBucketVersioningOutput)

	GetBucketVersioning(*s3.GetBucketVersioningInput) (*s3.GetBucketVersioningOutput, error)

	GetBucketWebsiteRequest(*s3.GetBucketWebsiteInput) (*request.Request, *s3.GetBucketWebsiteOutput)

	GetBucketWebsite(*s3.GetBucketWebsiteInput) (*s3.GetBucketWebsiteOutput, error)

	GetObjectRequest(*s3.GetObjectInput) (*request.Request, *s3.GetObjectOutput)

	GetObject(*s3.GetObjectInput) (*s3.GetObjectOutput, error)

	GetObjectAclRequest(*s3.GetObjectAclInput) (*request.Request, *s3.GetObjectAclOutput)

	GetObjectAcl(*s3.GetObjectAclInput) (*s3.GetObjectAclOutput, error)

	GetObjectTorrentRequest(*s3.GetObjectTorrentInput) (*request.Request, *s3.GetObjectTorrentOutput)

	GetObjectTorrent(*s3.GetObjectTorrentInput) (*s3.GetObjectTorrentOutput, error)

	HeadBucketRequest(*s3.HeadBucketInput) (*request.Request, *s3.HeadBucketOutput)

	HeadBucket(*s3.HeadBucketInput) (*s3.HeadBucketOutput, error)

	HeadObjectRequest(*s3.HeadObjectInput) (*request.Request, *s3.HeadObjectOutput)

	HeadObject(*s3.HeadObjectInput) (*s3.HeadObjectOutput, error)

	ListBucketsRequest(*s3.ListBucketsInput) (*request.Request, *s3.ListBucketsOutput)

	ListBuckets(*s3.ListBucketsInput) (*s3.ListBucketsOutput, error)

	ListMultipartUploadsRequest(*s3.ListMultipartUploadsInput) (*request.Request, *s3.ListMultipartUploadsOutput)

	ListMultipartUploads(*s3.ListMultipartUploadsInput) (*s3.ListMultipartUploadsOutput, error)

	ListMultipartUploadsPages(*s3.ListMultipartUploadsInput, func(*s3.ListMultipartUploadsOutput, bool) bool) error

	ListObjectVersionsRequest(*s3.ListObjectVersionsInput) (*request.Request, *s3.ListObjectVersionsOutput)

	ListObjectVersions(*s3.ListObjectVersionsInput) (*s3.ListObjectVersionsOutput, error)

	ListObjectVersionsPages(*s3.ListObjectVersionsInput, func(*s3.ListObjectVersionsOutput, bool) bool) error

	ListObjectsRequest(*s3.ListObjectsInput) (*request.Request, *s3.ListObjectsOutput)

	ListObjects(*s3.ListObjectsInput) (*s3.ListObjectsOutput, error)

	ListObjectsPages(*s3.ListObjectsInput, func(*s3.ListObjectsOutput, bool) bool) error

	ListPartsRequest(*s3.ListPartsInput) (*request.Request, *s3.ListPartsOutput)

	ListParts(*s3.ListPartsInput) (*s3.ListPartsOutput, error)

	ListPartsPages(*s3.ListPartsInput, func(*s3.ListPartsOutput, bool) bool) error

	PutBucketAclRequest(*s3.PutBucketAclInput) (*request.Request, *s3.PutBucketAclOutput)

	PutBucketAcl(*s3.PutBucketAclInput) (*s3.PutBucketAclOutput, error)

	PutBucketCorsRequest(*s3.PutBucketCorsInput) (*request.Request, *s3.PutBucketCorsOutput)

	PutBucketCors(*s3.PutBucketCorsInput) (*s3.PutBucketCorsOutput, error)

	PutBucketLifecycleRequest(*s3.PutBucketLifecycleInput) (*request.Request, *s3.PutBucketLifecycleOutput)

	PutBucketLifecycle(*s3.PutBucketLifecycleInput) (*s3.PutBucketLifecycleOutput, error)

	PutBucketLifecycleConfigurationRequest(*s3.PutBucketLifecycleConfigurationInput) (*request.Request, *s3.PutBucketLifecycleConfigurationOutput)

	PutBucketLifecycleConfiguration(*s3.PutBucketLifecycleConfigurationInput) (*s3.PutBucketLifecycleConfigurationOutput, error)

	PutBucketLoggingRequest(*s3.PutBucketLoggingInput) (*request.Request, *s3.PutBucketLoggingOutput)

	PutBucketLogging(*s3.PutBucketLoggingInput) (*s3.PutBucketLoggingOutput, error)

	PutBucketNotificationRequest(*s3.PutBucketNotificationInput) (*request.Request, *s3.PutBucketNotificationOutput)

	PutBucketNotification(*s3.PutBucketNotificationInput) (*s3.PutBucketNotificationOutput, error)

	PutBucketNotificationConfigurationRequest(*s3.PutBucketNotificationConfigurationInput) (*request.Request, *s3.PutBucketNotificationConfigurationOutput)

	PutBucketNotificationConfiguration(*s3.PutBucketNotificationConfigurationInput) (*s3.PutBucketNotificationConfigurationOutput, error)

	PutBucketPolicyRequest(*s3.PutBucketPolicyInput) (*request.Request, *s3.PutBucketPolicyOutput)

	PutBucketPolicy(*s3.PutBucketPolicyInput) (*s3.PutBucketPolicyOutput, error)

	PutBucketReplicationRequest(*s3.PutBucketReplicationInput) (*request.Request, *s3.PutBucketReplicationOutput)

	PutBucketReplication(*s3.PutBucketReplicationInput) (*s3.PutBucketReplicationOutput, error)

	PutBucketRequestPaymentRequest(*s3.PutBucketRequestPaymentInput) (*request.Request, *s3.PutBucketRequestPaymentOutput)

	PutBucketRequestPayment(*s3.PutBucketRequestPaymentInput) (*s3.PutBucketRequestPaymentOutput, error)

	PutBucketTaggingRequest(*s3.PutBucketTaggingInput) (*request.Request, *s3.PutBucketTaggingOutput)

	PutBucketTagging(*s3.PutBucketTaggingInput) (*s3.PutBucketTaggingOutput, error)

	PutBucketVersioningRequest(*s3.PutBucketVersioningInput) (*request.Request, *s3.PutBucketVersioningOutput)

	PutBucketVersioning(*s3.PutBucketVersioningInput) (*s3.PutBucketVersioningOutput, error)

	PutBucketWebsiteRequest(*s3.PutBucketWebsiteInput) (*request.Request, *s3.PutBucketWebsiteOutput)

	PutBucketWebsite(*s3.PutBucketWebsiteInput) (*s3.PutBucketWebsiteOutput, error)

	PutObjectRequest(*s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput)

	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)

	PutObjectAclRequest(*s3.PutObjectAclInput) (*request.Request, *s3.PutObjectAclOutput)

	PutObjectAcl(*s3.PutObjectAclInput) (*s3.PutObjectAclOutput, error)

	RestoreObjectRequest(*s3.RestoreObjectInput) (*request.Request, *s3.RestoreObjectOutput)

	RestoreObject(*s3.RestoreObjectInput) (*s3.RestoreObjectOutput, error)

	UploadPartRequest(*s3.UploadPartInput) (*request.Request, *s3.UploadPartOutput)

	UploadPart(*s3.UploadPartInput) (*s3.UploadPartOutput, error)

	UploadPartCopyRequest(*s3.UploadPartCopyInput) (*request.Request, *s3.UploadPartCopyOutput)

	UploadPartCopy(*s3.UploadPartCopyInput) (*s3.UploadPartCopyOutput, error)
}

var _ S3API = (*s3.S3)(nil)
