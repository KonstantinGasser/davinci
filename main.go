package main

import (
	"flag"
	_ "image/png"

	"github.com/KonstantinGasser/davinci/core/api"
	"github.com/KonstantinGasser/davinci/core/domain/matrix"
	matrixsvc "github.com/KonstantinGasser/davinci/core/domain/matrix/svc"
	"github.com/sirupsen/logrus"
)

func main() {

	rows := flag.Int("rows", 16, "number of rows of the matrix")
	cols := flag.Int("cols", 16, "number of columns of the matrix")
	flag.Parse()

	matrixLED, err := matrix.New(*rows, *cols)
	if err != nil {
		logrus.Fatalf("Could not connect to LED-Strip: %s", err.Error())
	}

	matrixSvc := matrixsvc.New(matrixLED)

	apiServer := api.New(matrixSvc)

	if err := apiServer.ListenAndServe(); err != nil {
		logrus.Fatalf("Could no start API-Server: %s", err.Error())
	}
}
