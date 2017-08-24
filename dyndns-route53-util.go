package main

import (
	"flag"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/route53"
	"io/ioutil"
	"net"
	"net/http"
)

func main() {

	var host, zoneId string
	flag.StringVar(&host, "host", "", "the host")
	flag.StringVar(&zoneId, "zoneId", "", "the zoneId")

	flag.Parse()

	ip, ipErr := getIP()
	if ipErr != nil {
		fmt.Println("failed to look up current dns,", ipErr)
		return
	}

	dns, dnsErr := getDNS(host)
	if dnsErr != nil {
		fmt.Println("failed to look up current dns,", dnsErr)
		return
	}

	if dns != ip {
		fmt.Println("Nothing to do here.", host, " is up to date :", ip)
		return
	}

	fmt.Println("Updating DNS.", host, " was: ", dns, " changing to: ", ip)
	update, updateErr := updateDNS(host, ip, zoneId)
	if updateErr == nil {
		fmt.Println(update)
	}
}

func getIP() (string, error) {
	resp, err := http.Get("http://ipecho.net/plain")
	if err != nil {
		fmt.Println("failed to look up current ip,", err)
		return "", err
	}
	defer resp.Body.Close()
	body_resp, body_err := ioutil.ReadAll(resp.Body)
	if body_err != nil {
		return "", body_err
	}

	return string(body_resp), nil
}

func getDNS(host string) (string, error) {
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Println("failed to look up current dns,", err)
		return "", err
	}

	return addrs[0], nil
}

func updateDNS(host string, newIp string, zoneId string) (string, error) {
	sess, err := session.NewSession()
	if err != nil {
		fmt.Println("failed to create session,", err)
		return "", err
	}

	svc := route53.New(sess)

	params := &route53.ChangeResourceRecordSetsInput{
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(host),
						Type: aws.String("A"),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(newIp),
							},
						},
						TTL: aws.Int64(86400),
					},
				},
			},
		},
		HostedZoneId: aws.String("/hostedzone/" + zoneId),
	}

	_, err = svc.ChangeResourceRecordSets(params)
	if err != nil {
		fmt.Println("failed to update dns,", err)
		return "", err
	}

	return "Success", nil
}
