# dyndns-route53-util
Simple GO util that can update a route53 record set with the machine's public IP.

To build:
```
go get github.com/aws/aws-sdk-go/aws
go build dyndns-route53-util.go
```
To run:
```
AWS_ACCESS_KEY_ID=YOUR_ACCESS_KEY AWS_SECRET_ACCESS_KEY=your_secret_access_key /usr/local/bin/dyndns-route53-util -host=your_host -zoneId=YOUR_ZONE_ID
```
