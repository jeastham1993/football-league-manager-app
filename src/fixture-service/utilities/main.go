package main


import (
	"strings"
	"time"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sns"
)

func main() {
	// Initialize the AWS session
	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: credentials.NewSharedCredentials("", "league-manager-sqs"),
	}))

	// Create new services for SQS and SNS
	sqsSvc := sqs.New(sess)
	snsSvc := sns.New(sess)

	requiredQueueName := "fixture-service-queue";
	requiredTopic := "leaguemanager-info-newteamcreated"

	queueURL := createAndSubscribeSqsQueueToSnsTopic(*sqsSvc, *snsSvc, requiredQueueName, requiredTopic)

	go checkMessages(*sqsSvc, queueURL)

	fmt.Scanln()
}


func checkMessages(sqsSvc sqs.SQS, queueURL string) {
    for ; ; {
		retrieveMessageRequest := sqs.ReceiveMessageInput{
			QueueUrl: &queueURL,
		}
	
		retrieveMessageResponse, _ := sqsSvc.ReceiveMessage(&retrieveMessageRequest)
	
		if len(retrieveMessageResponse.Messages) > 0 {

			processedReceiptHandles := make([]*sqs.DeleteMessageBatchRequestEntry, len(retrieveMessageResponse.Messages))
			
			for i, mess := range retrieveMessageResponse.Messages {
				fmt.Println(mess.String())
				
				processedReceiptHandles[i] = &sqs.DeleteMessageBatchRequestEntry{
					Id: mess.MessageId,
					ReceiptHandle: mess.ReceiptHandle,
				}
			}

			deleteMessageRequest := sqs.DeleteMessageBatchInput{
				QueueUrl: &queueURL,
				Entries: processedReceiptHandles,
			}

			_,err := sqsSvc.DeleteMessageBatch(&deleteMessageRequest)

			if err != nil {
				fmt.Println(err.Error())
			}
		}
	
		if len(retrieveMessageResponse.Messages) == 0 {
			fmt.Println(":(  I have no messages")
		}

        fmt.Printf("%v+\n", time.Now())
        time.Sleep(time.Minute)
    }
}

func convertQueueURLToARN(inputURL string) (string) {
	// Awfully bad string replace code to convert a SQS queue URL to an ARN
	queueARN := strings.Replace(strings.Replace(strings.Replace(inputURL, "https://sqs.", "arn:aws:sqs:", -1), ".amazonaws.com/", ":", -1), "/", ":", -1)

	return queueARN
}

func createAndSubscribeSqsQueueToSnsTopic(sqsSvc sqs.SQS, snsSvc sns.SNS, requiredQueueName, requiredTopic string) string {
	queueURL := ""

	// Create a new requst to list queues, first we will check to see if our required queue already exists
	listQueuesRequest := sqs.ListQueuesInput{}

	listQueueResults, _ := sqsSvc.ListQueues(&listQueuesRequest)

	for _, t := range listQueueResults.QueueUrls {
		// If one of the returned queue URL's contains the required name we need then break the loop
		if strings.Contains(*t, requiredQueueName) {
			queueURL = *t
			break
		}
	}

	// If, after checking existing queues, the URL is still empty then create the SQS queue.
	if queueURL == "" {
		createQueueInput := &sqs.CreateQueueInput{
			QueueName: &requiredQueueName,
		}

		createQueueResponse, err := sqsSvc.CreateQueue(createQueueInput)

		if err != nil {
			fmt.Println(err.Error())
		}

		if (createQueueResponse != nil) {
			queueURL = *createQueueResponse.QueueUrl

			fmt.Println(createQueueResponse.QueueUrl)
		}
	}

	// No way to retrieve the queue ARN through the SDK, manual string replace to generate the ARN
	queueARN := convertQueueURLToARN(queueURL)

	protocolName := "sqs"
	topicArn := ""

	listTopicsRequest := sns.ListTopicsInput{}

	// List all topics and loop through the results until we find a match
	allTopics, _ := snsSvc.ListTopics(&listTopicsRequest)

	for _, t := range allTopics.Topics {
		if strings.Contains(*t.TopicArn, requiredTopic) {
			topicArn = *t.TopicArn
			break
		}
	}

	// If the required topic is found, then create the subscription
	if topicArn != "" {
		subscibeQueueInput := sns.SubscribeInput{
			TopicArn: &topicArn,
			Protocol: &protocolName,
			Endpoint: &queueARN,
		}

		createSubRes, err := snsSvc.Subscribe(&subscibeQueueInput)

		if err != nil {
			fmt.Println(err.Error())
		}

		if (createSubRes != nil) {
			fmt.Println(createSubRes.SubscriptionArn)
		}
	}

	policyContent := "{\"Version\": \"2012-10-17\",  \"Id\": \"" + queueARN + "/SQSDefaultPolicy\",  \"Statement\": [    {     \"Sid\": \"Sid1580665629194\",      \"Effect\": \"Allow\",      \"Principal\": {        \"AWS\": \"*\"      },      \"Action\": \"SQS:SendMessage\",      \"Resource\": \"" + queueARN + "\",      \"Condition\": {        \"ArnEquals\": {         \"aws:SourceArn\": \"" + topicArn + "\"        }      }    }  ]}"

	attr := make(map[string]*string, 1)
	attr["Policy"] = &policyContent

	setQueueAttrInput := sqs.SetQueueAttributesInput{
		QueueUrl: &queueURL,
		Attributes: attr,
	}

	_, err := sqsSvc.SetQueueAttributes(&setQueueAttrInput)

	if err != nil {
		fmt.Println(err.Error())
	}

	return queueURL
}