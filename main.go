package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
)

func init() {
	fmt.Printf("init\n")
}

func main() {
	fmt.Printf("main\n")

	// Load session from shared config
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create new EC2 client
	svc := ec2.New(sess)
	resultRegions, err := svc.DescribeRegions(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	fmt.Println("resultRegions", resultRegions)

	for _, region := range resultRegions.Regions {
		printInfo(sess, *region.RegionName)
	}

	resultAvalZones, err := svc.DescribeAvailabilityZones(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("FINAL resultAvalZones", resultAvalZones.AvailabilityZones)

	/*
	sess1 := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("us-west-2"),
	}))


	fmt.Println("SDK:", aws.SDKName, aws.SDKVersion)


	resultAvalZones, err := svc.DescribeAvailabilityZones(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	fmt.Println("resultAvalZones", resultAvalZones.AvailabilityZones)

	resultVpcs, err := svc.DescribeVpcs(nil)
	fmt.Println("resultVpcs", resultVpcs)

	*/
}

func printInfo(sess *session.Session, region string) {
	println(region)
	svc := ec2.New(sess, aws.NewConfig().WithRegion(region))

	resultVpcs, err := svc.DescribeVpcs(nil)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	for _,vpc := range resultVpcs.Vpcs {

		fmt.Println("VPCS:", *vpc.VpcId)
	}

}