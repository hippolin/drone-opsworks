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
	build     string
	buildDate string
)

func main() {
	fmt.Printf("Drone AWS OpsWorks Plugin built at %s\n", buildDate)

	workspace := drone.Workspace{}
	repo := drone.Repo{}
	build := drone.Build{}
	vargs := Params{}

	plugin.Param("workspace", &workspace)
	plugin.Param("repo", &repo)
	plugin.Param("build", &build)
	plugin.Param("vargs", &vargs)
	plugin.MustParse()

	if len(vargs.AccessKeyID) == 0 {
		fmt.Println("Please provide an access key")

		os.Exit(1)
		return
	}

	if len(vargs.SecretAccessKey) == 0 {
		fmt.Println("Please provide a secret key")

		os.Exit(1)
		return
	}

	if len(vargs.Region) == 0 {
		fmt.Println("Please provide a region")

		os.Exit(1)
		return
	}

	svc := opsworks.New(
		session.New(&aws.Config{
			Region:      aws.String(vargs.Region),
			Credentials: credentials.NewStaticCredentials(vargs.AccessKeyID, vargs.SecretAccessKey, ""),
		}))

	resp, err := svc.CreateDeployment(
		&opsworks.CreateDeploymentInput{
			Command: &opsworks.DeploymentCommand{
				Name: aws.String("DeploymentCommandName"),
				Args: map[string][]*string{
					"Key": {
						aws.String("String"),
					},
				},
			},
			StackId:    aws.String("String"),
			AppId:      aws.String("String"),
			Comment:    aws.String("String"),
			CustomJson: aws.String("String"),
			InstanceIds: []*string{
				aws.String("String"),
			},
		})

	if err != nil {
		fmt.Println(err.Error())

		os.Exit(1)
		return
	}

	fmt.Println(resp)
}
