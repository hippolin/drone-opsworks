package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/opsworks"
	"github.com/drone/drone-go/drone"
	"github.com/drone/drone-go/plugin"
)

var (
	buildCommit string
)

func main() {
	fmt.Printf("Drone AWS OpsWorks Plugin built from %s\n", buildCommit)

	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if vargs.AccessKey == "" {
		fmt.Println("Please provide an access key id")
		os.Exit(1)
	}

	if vargs.SecretKey == "" {
		fmt.Println("Please provide a secret access key")
		os.Exit(1)
	}

	if aws.StringValue(vargs.Region) == "" {
		fmt.Println("Please provide a region")
		os.Exit(1)
	}

	if aws.StringValue(vargs.StackID) == "" {
		fmt.Println("Please provide a stack ID")
		os.Exit(1)
	}

	if aws.StringValue(vargs.Command) == "" {
		fmt.Println("Please provide a deploy command")
		os.Exit(1)
	}

	svc := opsworks.New(
		session.New(&aws.Config{
			Region: vargs.Region,
			Credentials: credentials.NewStaticCredentials(
				vargs.AccessKey,
				vargs.SecretKey,
				"",
			),
		}),
	)

	_, err := svc.CreateDeployment(
		&opsworks.CreateDeploymentInput{
			Command: &opsworks.DeploymentCommand{
				Name: vargs.Command,
				Args: vargs.Arguments,
			},
			StackId:     vargs.StackID,
			AppId:       vargs.AppID,
			Comment:     vargs.Comment,
			CustomJson:  vargs.CustomJSON,
			InstanceIds: vargs.Instances,
		},
	)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("Successfully deployed")
}
