package infrastructure

import (
	"fmt"
	"os"
	"team-service/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/rs/xid"
)

// DynamoDbRepository is an implementation of the teams repository in dynamodb
type DynamoDbRepository struct {
}

// Store creates a new team in the in memory storage.
func (r *DynamoDbRepository) Store(team *domain.Team) string {
	team.ID = xid.New().String()

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})

	tableInput := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("teamName"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("teamName"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(10),
		},
		TableName: aws.String("Teams"),
	}

	_, tableErr := svc.CreateTable(tableInput)

	if tableErr != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(tableErr.Error())
		os.Exit(1)
	}

	av, err := dynamodbattribute.MarshalMap(team)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Teams"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return team.ID
}

// FindByID gets a team from the in memory store.
func (r *DynamoDbRepository) FindByID(teamID string) *domain.Team {

	return nil
}

// Update updates an existing team record
func (r *DynamoDbRepository) Update(team *domain.Team) *domain.Team {

	return team
}

// Search looks for a record in the database with the requiste search term
func (r *DynamoDbRepository) Search(searchTerm string) []domain.Team {
	teamsResponse := make([]domain.Team, 0)

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess, &aws.Config{Endpoint: aws.String("http://localhost:8000")})

	filt := expression.Name("teamName").Equal(expression.Value(searchTerm))

	// Get back the title, year, and rating
	proj := expression.NamesList(expression.Name("teamName"), expression.Name("id"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("Teams"),
	}

	// Make the DynamoDB Query API call
	result, err := svc.Scan(params)

	for _, i := range result.Items {
		item := domain.Team{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		teamsResponse = append(teamsResponse, item)

		if err != nil {
			fmt.Println("Got error unmarshalling:")
			fmt.Println(err.Error())
		}
	}

	return teamsResponse
}
