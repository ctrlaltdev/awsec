package iam

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
)

func ListUsers() (users []types.User) {
	result, err := client.ListUsers(context.TODO(), &iam.ListUsersInput{})

	if err != nil {
		log.Println(err)
	}

	users = append(users, result.Users...)

	for result.IsTruncated {
		result, err := client.ListUsers(context.TODO(), &iam.ListUsersInput{Marker: result.Marker})

		if err != nil {
			log.Println(err)
		}

		users = append(users, result.Users...)
	}

	return users
}

func ListAccessKeys(username *string) (keys []types.AccessKeyMetadata) {
	result, err := client.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{UserName: username})

	if err != nil {
		log.Println(err)
	}

	keys = append(keys, result.AccessKeyMetadata...)

	for result.IsTruncated {
		result, err := client.ListAccessKeys(context.TODO(), &iam.ListAccessKeysInput{UserName: username, Marker: result.Marker})

		if err != nil {
			log.Println(err)
		}

		keys = append(keys, result.AccessKeyMetadata...)
	}

	return keys
}
