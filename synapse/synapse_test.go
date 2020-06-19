// // Copyright (c) Microsoft and contributors.  All rights reserved.
// //
// // This source code is licensed under the MIT license found in the
// // LICENSE file in the root directory of this source tree.

package synapse

import (
	"flag"
	"log"
	"os"
	"testing"

	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/config"
	"github.com/Azure-Samples/azure-sdk-for-go-samples/internal/util"
)

// TestMain sets up the environment and initiates tests.
func TestMain(m *testing.M) {
	var err error
	err = config.ParseEnvironment()
	if err != nil {
		log.Fatalf("failed to parse env: %v\n", err)
	}
	err = config.AddFlags()
	if err != nil {
		log.Fatalf("failed to parse env: %v\n", err)
	}
	flag.Parse()

	code := m.Run()
	os.Exit(code)
}

func ExampleCreateHadoopCluster() {
	// rgName := config.GenerateGroupName("SynapseWorkspaceExample")
	// config.SetGroupName(rgName)

	// ctx, cancel := context.WithTimeout(context.Background(), 6000*time.Second)
	// defer cancel()
	// defer resources.Cleanup(ctx)

	// _, err := resources.CreateGroup(ctx, rgName)
	// if err != nil {
	// 	util.LogAndPanic(err)
	// }
	// util.PrintAndLog("created resource group")

	// clusterName := strings.ToLower(config.AppendRandomSuffix("exsynapseworkspace"))
	// _, err = CreateWorkspace(rgName, clusterName)
	// if err != nil {
	// 	util.LogAndPanic(err)
	// }
	// util.PrintAndLog("created workspace")

	err := getMonitoring()
	// _, err := getSparkClient()
	//err := getAccessControl()
	if err != nil {
		util.LogAndPanic(err)
	}

	// Output:
	// created resource group
	// creating synapse workspace
	// waiting for synapse workspace to finish deploying, this will take a while...
	// created workspace
}
