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

	// oldAccessKey := os.Getenv("AWS_ACCESS_KEY")
	// oldSecretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	// os.Setenv("AWS_ACCESS_KEY", amazonAccountS3)
	// os.Setenv("AWS_SECRET_ACCESS_KEY", amazonPasswordS3)
	// defer func() {
	// 	os.Setenv("AWS_ACCESS_KEY", oldAccessKey)
	// 	os.Setenv("AWS_SECRET_ACCESS_KEY", oldSecretKey)
	// }()

	fmt.Printf("access key: %v\n", amazonAccountS3)
	fmt.Printf("secret key: %v\n", amazonPasswordS3)
	fmt.Printf("region: %v\n", amazonRegionS3)
	fmt.Printf("bucketName: %v\n", bucketName)

	cli := s3.New(session.New(&aws.Config{
		Credentials: credentials.NewStaticCredentials(amazonAccountS3, amazonPasswordS3, ""),
		Region:      aws.String(amazonRegionS3),
	}))

	body := []byte(time.Now().UTC().String())
	output, err := cli.PutObject(&s3.PutObjectInput{
		Body:        newReadSeeker(body),
		Bucket:      aws.String(bucketName),
		ContentType: aws.String("text/plain"),
		Key:         aws.String("/_test_touch"),
	})
	if err != nil {
		fmt.Println(err)
		t.Error(err)
	}
	fmt.Printf("%v\noutput %v\n", time.Now(), output)
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
