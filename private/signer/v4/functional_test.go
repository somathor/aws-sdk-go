package v4_test

import (
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/somathor/aws-sdk-go/aws"
	"github.com/somathor/aws-sdk-go/awstesting/unit"
	"github.com/somathor/aws-sdk-go/service/s3"
)

func TestPresignHandler(t *testing.T) {
	svc := s3.New(unit.Session)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:             aws.String("bucket"),
		Key:                aws.String("key"),
		ContentDisposition: aws.String("a+b c$d"),
		ACL:                aws.String("public-read"),
	})
	req.Time = time.Unix(0, 0)
	urlstr, err := req.Presign(5 * time.Minute)

	assert.NoError(t, err)

	expectedDate := "19700101T000000Z"
	expectedHeaders := "content-disposition;host;x-amz-acl"
	expectedSig := "2d76a414208c0eac2a23ef9c834db9635ecd5a0fbb447a00ad191f82d854f55b"
	expectedCred := "AKID/19700101/mock-region/s3/aws4_request"

	u, _ := url.Parse(urlstr)
	urlQ := u.Query()
	assert.Equal(t, expectedSig, urlQ.Get("X-Amz-Signature"))
	assert.Equal(t, expectedCred, urlQ.Get("X-Amz-Credential"))
	assert.Equal(t, expectedHeaders, urlQ.Get("X-Amz-SignedHeaders"))
	assert.Equal(t, expectedDate, urlQ.Get("X-Amz-Date"))
	assert.Equal(t, "300", urlQ.Get("X-Amz-Expires"))

	assert.NotContains(t, urlstr, "+") // + encoded as %20
}

func TestPresignRequest(t *testing.T) {
	svc := s3.New(unit.Session)
	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket:             aws.String("bucket"),
		Key:                aws.String("key"),
		ContentDisposition: aws.String("a+b c$d"),
		ACL:                aws.String("public-read"),
	})
	req.Time = time.Unix(0, 0)
	urlstr, headers, err := req.PresignRequest(5 * time.Minute)

	assert.NoError(t, err)

	expectedDate := "19700101T000000Z"
	expectedHeaders := "content-disposition;host;x-amz-acl"
	expectedSig := "2d76a414208c0eac2a23ef9c834db9635ecd5a0fbb447a00ad191f82d854f55b"
	expectedCred := "AKID/19700101/mock-region/s3/aws4_request"
	expectedHeaderMap := http.Header{
		"x-amz-acl":           []string{"public-read"},
		"content-disposition": []string{"a+b c$d"},
	}

	u, _ := url.Parse(urlstr)
	urlQ := u.Query()
	assert.Equal(t, expectedSig, urlQ.Get("X-Amz-Signature"))
	assert.Equal(t, expectedCred, urlQ.Get("X-Amz-Credential"))
	assert.Equal(t, expectedHeaders, urlQ.Get("X-Amz-SignedHeaders"))
	assert.Equal(t, expectedDate, urlQ.Get("X-Amz-Date"))
	assert.Equal(t, expectedHeaderMap, headers)
	assert.Equal(t, "300", urlQ.Get("X-Amz-Expires"))

	assert.NotContains(t, urlstr, "+") // + encoded as %20
}
