package infrastructure

import (
	"encoding/json"
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
	dynamoSvc *dynamodb.DynamoDB
}

// NewDynamoDbRepo creates a new DynamoDb Repository
func NewDynamoDbRepo() *DynamoDbRepository {
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
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       aws.String("HASH"),
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
	}

	return &DynamoDbRepository{
		dynamoSvc: svc,
	}
}

// Store creates a new team in the in memory storage.
func (r *DynamoDbRepository) Store(team *domain.Team) string {
	team.ID = xid.New().String()

	av, err := dynamodbattribute.MarshalMap(team)

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("Teams"),
	}

	_, err = r.dynamoSvc.PutItem(input)

	if err != nil {
		fmt.Println("Got error calling CreateTable:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	return team.ID
}

// FindByID gets a team from the in memory store.
func (r *DynamoDbRepository) FindByID(teamID string) *domain.Team {
	filt := expression.Name("id").Equal(expression.Value(teamID))

	expr, _ := expression.NewBuilder().WithFilter(filt).Build()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("Teams"),
	}

	// Make the DynamoDB Query API call
	result, _ := r.dynamoSvc.Scan(params)

	if len(result.Items) > 0 {
		item := &domain.Team{}

		dynamodbattribute.UnmarshalMap(result.Items[0], &item)

		for _, res := range result.Items {
			for fieldName, playerRes := range res {
				if fieldName == "players" && playerRes.S != nil {
					s := *playerRes.S

					json.Unmarshal([]byte(s), &item.Players)
				}
			}
		}

		return item
	}

	return nil
}

// Update updates an existing team record
func (r *DynamoDbRepository) Update(team *domain.Team) *domain.Team {

	dynamoDbMappedPlayers, _ := json.Marshal(team.Players)

	updateString := string(dynamoDbMappedPlayers)

	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":n": {
				S: aws.String(team.Name),
			},
			":p": {
				S: aws.String(updateString),
			},
		},
		TableName: aws.String("Teams"),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(team.ID),
			},
		},
		ReturnValues:     aws.String("UPDATED_NEW"),
		UpdateExpression: aws.String("SET players = :p, teamName = :n"),
	}

	_, updateErr := r.dynamoSvc.UpdateItem(input)

	if updateErr != nil {

	}

	return team
}

// Search looks for a record in the database with the requiste search term
func (r *DynamoDbRepository) Search(searchTerm string) []domain.Team {
	teamsResponse := make([]domain.Team, 0)

	filt := expression.Name("teamName").Contains(searchTerm)

	expr, err := expression.NewBuilder().WithFilter(filt).Build()

	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		TableName:                 aws.String("Teams"),
	}

	// Make the DynamoDB Query API call
	result, err := r.dynamoSvc.Scan(params)

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
