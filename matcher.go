package main

import (
	"flag"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/cloudformation"
	pkill "github.com/leofigy/aws_resources"

	"github.com/leofigy/tfmatch/awspk"
)

const cloudc = "cloudformation"
const helpStackName = "stack to fetch     [required] "
const helpProfileName = "aws config profile [optional] "

var (
	stackName   string
	profileName string // optional
)

func init() {
	flag.StringVar(
		&stackName, "stack", "", helpStackName)

	flag.StringVar(
		&profileName, "profile", "", helpProfileName)
}

func main() {
	flag.Parse()

	if len(stackName) < 1 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	pkillFetch(loadAWSCFG())
}

func loadAWSCFG() *aws.Config {

	tentative := []external.Config{}

	if len(profileName) > 0 {
		tentative = append(tentative, external.WithSharedConfigProfile(profileName))
	}

	config, err := external.LoadDefaultAWSConfig(tentative...)
	if err != nil {
		log.Panicf("AWS2 load config error %+v", err)
	}

	return &config
}

func pkillFetch(myCfg *aws.Config) {

	pkill.SetConfig(*myCfg)

	cfConfig := &pkill.CloudFormationTypeConfig{
		StackStatus: []cloudformation.StackStatus{
			cloudformation.StackStatusCreateComplete,
		},

		StackName: stackName,
	}

	iCF := pkill.Relations[cloudc](pkill.Config())
	if err := iCF.Configure(*cfConfig); err != nil {
		panic(err)
	}

	iCF.GetServices()
	iCF.GetResources()
	iCF.GetResourcesDetail()

	awspk.InspectResources(iCF)
}
