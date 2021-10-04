package main

import (
	"github.com/KonstantinGasser/davinci/core/api"
	"github.com/sirupsen/logrus"
)

func main() {

	apiServer := api.New()

	if err := apiServer.ListenAndServe(); err != nil {
		logrus.Fatalf("Could no start API-Server: %s", err.Error())
	}
}
