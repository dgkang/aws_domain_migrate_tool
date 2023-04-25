package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/route53domains"
)

var Region string = ""

func ListHostedZones(aws_access_key_id string, aws_secret_access_key string) {
	fmt.Printf("aws_access_key_id: %s,aws_secret_access_key: %s\n", aws_access_key_id, aws_secret_access_key)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, ""),
	}))

	route := route53.New(sess)

	in := &route53.ListHostedZonesInput{}
	if out, er := route.ListHostedZones(in); er != nil {
		fmt.Printf("%s\n", er.Error())
	} else {
		fmt.Printf("%s\n", out.GoString())
	}
}

func ListResourceRecords(aws_access_key_id string, aws_secret_access_key string, hosted_zone_id string) {
	fmt.Printf("aws_access_key_id: %s,aws_secret_access_key: %s hosted_zone_id: %s\n", aws_access_key_id, aws_secret_access_key, hosted_zone_id)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, ""),
	}))

	route := route53.New(sess)

	in := &route53.ListResourceRecordSetsInput{HostedZoneId: aws.String(hosted_zone_id)}
	if out, er := route.ListResourceRecordSets(in); er != nil {
		fmt.Printf("%s\n", er.Error())
	} else {
		fmt.Printf("%s\n", out.GoString())
	}
}

func ListDomain(aws_access_key_id string, aws_secret_access_key string) {
	fmt.Printf("aws_access_key_id: %s,aws_secret_access_key: %s\n", aws_access_key_id, aws_secret_access_key)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, ""),
	}))

	domains := route53domains.New(sess)

	in := &route53domains.ListDomainsInput{}
	if out, er := domains.ListDomains(in); er != nil {
		fmt.Printf("%s\n", er.Error())
	} else {
		fmt.Printf("%s\n", out.GoString())
	}

}

func TransferDomain(aws_access_key_id string, aws_secret_access_key string, account_id string, domain string) {

	fmt.Printf("aws_access_key_id: %s,aws_secret_access_key: %s,account_id: %s,domain: %s\n", aws_access_key_id, aws_secret_access_key, account_id, domain)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, ""),
	}))

	domains := route53domains.New(sess)

	in := &route53domains.TransferDomainToAnotherAwsAccountInput{AccountId: aws.String(account_id), DomainName: aws.String(domain)}
	if out, er := domains.TransferDomainToAnotherAwsAccount(in); er != nil {
		fmt.Printf("%s\n", er.Error())
	} else {
		fmt.Printf("%s\n", out.GoString())
	}
}

func AcceptDomain(aws_access_key_id string, aws_secret_access_key string, password string, domain string) {

	fmt.Printf("aws_access_key_id: %s,aws_secret_access_key: %s,domain: %s,password: %s\n", aws_access_key_id, aws_secret_access_key, domain, password)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, ""),
	}))

	domains := route53domains.New(sess)

	in := &route53domains.AcceptDomainTransferFromAnotherAwsAccountInput{DomainName: aws.String(domain), Password: aws.String(password)}
	if out, er := domains.AcceptDomainTransferFromAnotherAwsAccount(in); er != nil {
		fmt.Printf("%s\n", er.Error())
	} else {
		fmt.Printf("%s\n", out.GoString())
	}
}

func GetOperationStatus(aws_access_key_id string, aws_secret_access_key string, operate_id string) {

	fmt.Printf("aws_access_key_id: %s,aws_secret_access_key: %s,operate_id: %s\n", aws_access_key_id, aws_secret_access_key, operate_id)

	sess := session.Must(session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, ""),
	}))

	domains := route53domains.New(sess)

	detail := route53domains.GetOperationDetailInput{OperationId: aws.String(operate_id)}
	if out, er := domains.GetOperationDetail(&detail); er != nil {
		fmt.Printf("%s\n", er.Error())
	} else {
		fmt.Printf("%s\n", out.GoString())
	}

}

func main() {
	var aws_access_key_id, aws_secret_access_key, operate_id, account_id, password, domain, hosted_zone_id string
	var act string
	var help bool

	flag.StringVar(&act, "act", "", "aws operation (transfer|accept|query|list-domains|list-zones)")
	flag.StringVar(&aws_access_key_id, "id", "", "aws_access_key_id")
	flag.StringVar(&aws_secret_access_key, "key", "", "aws_secret_access_key")
	flag.StringVar(&operate_id, "optid", "", "Identifier for tracking the progress of the request")
	flag.StringVar(&account_id, "accid", "", "The account ID of the AWS")
	flag.StringVar(&password, "pwd", "", "The password that was returned by the TransferDomainToAnotherAwsAccount request")
	flag.StringVar(&domain, "domain", "", "domain name")
	flag.StringVar(&hosted_zone_id, "zoneid", "", "The ID of the hosted zone that contains the resource record")
	flag.StringVar(&Region, "region", endpoints.UsEast1RegionID, "The region to send requests to")

	flag.BoolVar(&help, "h", false, "help for usage")
	flag.Parse()

	if help {
		flag.PrintDefaults()
		return
	}

	switch act {
	case "transfer":
		TransferDomain(aws_access_key_id, aws_secret_access_key, account_id, domain)

	case "query":
		GetOperationStatus(aws_access_key_id, aws_secret_access_key, operate_id)

	case "accept":
		AcceptDomain(aws_access_key_id, aws_secret_access_key, password, domain)

	case "list-domains":
		ListDomain(aws_access_key_id, aws_secret_access_key)

	case "list-zones":
		ListHostedZones(aws_access_key_id, aws_secret_access_key)

	case "list-record":
		ListResourceRecords(aws_access_key_id, aws_secret_access_key, hosted_zone_id)

	default:
		fmt.Printf("error action,%s", act)

	}
}
