package main

import (
	"context"
	"fmt"
	"os"

	"log/slog"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/backup"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	// Boilerplate init
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		logger.Error("Couldn't load default configuration. Have you set up your AWS account?")
		logger.Error("%v", err)
		return
	}
	backupClient := backup.NewFromConfig(sdkConfig)
	result, err := backupClient.ListProtectedResources(context.TODO(),
		&backup.ListProtectedResourcesInput{})
	if err != nil {
		logger.Error("ListProtectedResources error:%v", err)
	}
	for k, v := range result.Results {
		fmt.Println(k)

		logger.Info(
			"Report msg",
			"ResourceType", *v.ResourceType,
			"ResourceName", *v.ResourceName,
			"LastBackupTimee", *v.LastBackupTime,
			"ResourceArn", *v.ResourceArn)
	}
}
