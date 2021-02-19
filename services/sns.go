package services

import (
	"aws-test/schema"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

func SendSNSMsg(response schema.Response) {
	sess := session.Must(session.NewSession())
	snsClient := sns.New(sess)

	responseBytes, err := json.Marshal(response)
	if err != nil {
		fmt.Println("There was error in converting object into json")
	}
 	responseStr := string(responseBytes)

	req, _ := snsClient.PublishRequest(&sns.PublishInput{
		TopicArn: aws.String("arn:aws:sns:us-east-1:api-test:ok"),
		Message: aws.String(responseStr),
		MessageStructure: aws.String("json"),
	})

	res := req.Send()
	log.Print(res)
}