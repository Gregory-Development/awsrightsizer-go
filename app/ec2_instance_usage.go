package app

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/costexplorer"
	"log"
	"os"
)

func CalculateEC2Usage() {
	ctx := aws.Context(aws.BackgroundContext())

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