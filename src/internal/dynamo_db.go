package internal

import (
	//"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/maximorero14/golang-api-test/src/models"
)

var dynamo *dynamodb.DynamoDB

// connectDynamo returns a dynamoDB client
func ConnectDynamo() (db *dynamodb.DynamoDB) {
	var region = "us-east-1"
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: &region,
	})))
}

func GetItem(id string) (activity models.Activity, err error) {
	dynamo = ConnectDynamo()
	var tableName = "merchands_activities"
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(id),
			},
		},
		TableName: &tableName,
	})

	if err != nil {
		return activity, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &activity)

	return activity, err

}

/*
// PutItem inserts the struct Person
func PutItem(activity models.Activity) error {
	var tableName = "merchands_activities"
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"ID": {
				N: aws.String(activity.ID),
			},
			"Type": {
				S: aws.String(activity.Type),
			},
			"UserId": {
				S: aws.String(activity.UserId),
			},
		},
		TableName: &tableName,
	})

	return err
}



// CreateTable creates a table
func CreateTable() error {
	_, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   &TableName,
	})

	return err
}


// UpdateItem updates the Person based on the Person.Id
func UpdateItem(activity Activity) error {
	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string{
			"#N": aws.String("Name"),
			"#W": aws.String("Website"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Name": {
				S: aws.String(person.Name),
			},
			":Website": {
				S: aws.String(person.Website),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
		},
		TableName:        &TableName,
		UpdateExpression: aws.String("SET #N = :Name, #W = :Website"),
	})

	return err
}

// DeleteItem deletes a Person based on the Person.Id
func DeleteItem(person Person) error {
	_, err := dynamo.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
		},
		TableName: &TableName,
	})

	return err
}

*/
