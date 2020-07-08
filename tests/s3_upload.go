package tests

import (
	"bytes"
	"fmt"
	"io"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// ok. this doesn't work for some reason in tests (access denied is returned)
//
// requires a small delay for access key to start working
func uploadToS3(t *testing.T, bucketName string, amazonRegionS3 string, amazonAccountS3 string, amazonPasswordS3 string) {

	fmt.Printf("aws --profile default configure set aws_access_key_id \"%v\"\n", amazonAccountS3)
	fmt.Printf("aws --profile default configure set aws_secret_access_key \"%v\"\n", amazonPasswordS3)
	fmt.Printf("aws --profile default configure set region \"%v\"\n", amazonRegionS3)
	fmt.Printf("export AWS_BUCKET=%v\n", bucketName)
	// time.Sleep( 5* time.Minute)

	cli := s3.New(session.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(amazonAccountS3, amazonPasswordS3, ""),
		Region:      aws.String(amazonRegionS3),
	}))

	path := aws.String("/_test_touch")
	body := []byte(time.Now().UTC().String())
	output, err := cli.PutObject(&s3.PutObjectInput{
		Body:        newReadSeeker(body),
		Bucket:      aws.String(bucketName),
		ContentType: aws.String("text/plain"),
		Key:         path,
	})
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	fmt.Printf("%v\noutput %v\n", time.Now(), output)

	_, err = cli.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(bucketName),
		Key:    path,
	})
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
}

type readSeeker struct {
	io.Reader

	buf    []byte
	offset int
}

func newReadSeeker(buf []byte) io.ReadSeeker {
	return &readSeeker{
		Reader: bytes.NewReader(buf),
		buf:    buf,
	}
}

func (r *readSeeker) Seek(off int64, whence int) (int64, error) {
	offset := int(off)
	switch whence {
	case io.SeekStart:
		if offset < 0 || offset > len(r.buf) {
			return 0, io.EOF
		}
		r.offset = offset
		r.Reader = bytes.NewReader(r.buf[offset:])
	case io.SeekEnd:
		if offset < 0 || offset > len(r.buf) {
			return 0, io.EOF
		}
		r.offset = len(r.buf) - offset
		r.Reader = bytes.NewReader(r.buf[len(r.buf)-offset:])
	case io.SeekCurrent:
		if offset+r.offset > len(r.buf) ||
			offset+r.offset < 0 {
			return 0, io.EOF
		}
		r.offset = r.offset + offset
		r.Reader = bytes.NewReader(r.buf[r.offset:])
	default:
		panic("wrong whence arg")
	}
	return int64(r.offset), nil
}
