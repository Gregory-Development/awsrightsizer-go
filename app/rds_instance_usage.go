package app

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/rds"
	"log"
	"os"
)

type Instance struct {
	Name string
	ID string
}

var (
	ctx = aws.Context(aws.BackgroundContext())
	Instances []Instance
)

func GatherRDSInformation() {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("AWS_REGION"))})
	if err != nil {
		log.Println(err)
	}

	svc := rds.New(sess)

	r, err := svc.DescribeDBInstancesWithContext(ctx, &rds.DescribeDBInstancesInput{})

}