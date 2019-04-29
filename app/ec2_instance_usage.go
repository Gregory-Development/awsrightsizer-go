package app

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"os"
)

type Instance struct {
	Name string
	ID string
}

var (
	Instances []Instance
	ctx = aws.Context(aws.BackgroundContext())
)

func GatherEC2Information() {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	if err != nil {
		log.Println(err)
	}

	svc := ec2.New(sess)

	r, err := svc.DescribeInstancesWithContext(ctx, &ec2.DescribeInstancesInput{})

}

func CalculateEC2Usage() {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	if err != nil {
		log.Println(err)
	}

	svc := costexplorer.New(sess)

	r, err := svc.GetCostAndUsageWithContext(ctx, &costexplorer.GetCostAndUsageInput{})
	if err != nil {
		log.Println(err)
	}

	print(r.String())
}