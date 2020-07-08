package tests

import (
	"fmt"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/gruntwork-io/terratest/modules/terraform"
)

// To launch this test, please set up variables to deploy into EKS cluster
//
// export AWS_PROFILE=default
// export AWS_ACCESS_KEY=$(aws configure get aws_access_key_id --profile $AWS_PROFILE)
// export AWS_SECRET_ACCESS_KEY=$(aws configure get aws_secret_access_key --profile $AWS_PROFILE)
// export AWS_DEFAULT_REGION=$(aws configure get region --profile $AWS_PROFILE)

func TestTerraformS3PrivateBucket(t *testing.T) {
	t.Parallel()

	assert.NotEmpty(t, os.Getenv("AWS_PROFILE"), "AWS_PROFILE var expected")
	assert.NotEmpty(t, os.Getenv("AWS_ACCESS_KEY"), "AWS_ACCESS_KEY var expected")
	assert.NotEmpty(t, os.Getenv("AWS_SECRET_ACCESS_KEY"), "AWS_SECRET_ACCESS_KEY var expected")
	assert.NotEmpty(t, os.Getenv("AWS_DEFAULT_REGION"), "AWS_DEFAULT_REGION var expected")

	application := "app-" + strings.ToLower(random.UniqueId())
	awsRegion := os.Getenv("AWS_DEFAULT_REGION")

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "./../terraform/",
		// Variables to pass to our Terraform code using -var options
		Vars: map[string]interface{}{
			"aws_region":  awsRegion,
			"application": application,
		},
		// Environment variables to set when running Terraform
		EnvVars: map[string]string{},

		// Variables to pass to our Terraform code using -var-file options
		// VarFiles: []string{"./../vars/varfile.tfvars"},
		// Disable colors in Terraform commands so its easier to parse stdout/stderr
		NoColor: true,
	}

	// At the end of the test, run `terraform destroy` to clean up any resources that were created
	defer terraform.Destroy(t, terraformOptions)

	// This will run `terraform init` and `terraform apply` and fail the test if there are any errors
	terraform.InitAndApply(t, terraformOptions)
	fmt.Println(terraform.OutputAll(t, terraformOptions))

	assert.NotEmpty(t, terraform.Output(t, terraformOptions, "bucket_arn"))

	bucketID := terraform.Output(t, terraformOptions, "bucket_id")
	bucketRegion := terraform.Output(t, terraformOptions, "bucket_region")
	assert.NotEmpty(t, bucketID)
	assert.NotEmpty(t, bucketRegion)

	accessKey := terraform.Output(t, terraformOptions, "access_key")
	secretKey := terraform.Output(t, terraformOptions, "secret_key")
	assert.NotEmpty(t, accessKey)
	assert.NotEmpty(t, secretKey)

	// waiting
	time.Sleep(time.Second * 10)
	// Do test upload to the bucket - with new credentials
	uploadToS3(t, bucketID, bucketRegion, accessKey, secretKey)
}
