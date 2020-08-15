package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"golang.org/x/crypto/ssh"
)

func main() {
	lambda.Start(HandleRequest)
}

// HandleRequest execute the logic
func HandleRequest(ctx context.Context) (string, error) {
	log.Print("snapshot export event received")

	config := &ssh.ClientConfig{
		User: os.Getenv("NODE_USER"),
		Auth: []ssh.AuthMethod{
			getPublicKeyFromS3(),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	instanceIP := *getInstanceIP()

	conn, err := ssh.Dial("tcp", instanceIP+":22", config)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	runCommand("sudo systemctl start tezos-backup", conn)

	log.Print("snapshot service started")
	return "snapshot service started.", nil
}

func getPublicKeyFromS3() ssh.AuthMethod {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("S3_REGION")),
	},
	)

	downloader := s3manager.NewDownloader(sess)

	buf := aws.NewWriteAtBuffer([]byte{})
	S3LambdaKey := os.Getenv("S3_LAMBDA_KEY")
	_, err := downloader.Download(buf,
		&s3.GetObjectInput{
			Bucket: aws.String(os.Getenv("S3_BUCKET")),
			Key:    aws.String(S3LambdaKey),
		})
	if err != nil {
		panic(fmt.Sprintf("Unable to download %q, %v", S3LambdaKey, err))
	}

	signer, err := ssh.ParsePrivateKey(buf.Bytes())
	if err != nil {
		panic(fmt.Sprintf("Unable to parse key %q at bucket %q, %v", S3LambdaKey, os.Getenv("S3_BUCKET"), err))
	}

	log.Print("got private key from S3 bucket.")
	return ssh.PublicKeys(signer)
}

func getInstanceIP() *string {
	sess, _ := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("S3_REGION"))},
	)

	svc := ec2.New(sess)

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name: aws.String("instance-state-name"),
				Values: []*string{
					aws.String("running"),
				},
			},
			{
				Name: aws.String("tag:Name"),
				Values: []*string{
					aws.String(fmt.Sprintf("tzlink-%s", os.Getenv("NETWORK"))),
				},
			},
			{
				Name: aws.String("tag:Project"),
				Values: []*string{
					aws.String("tezos-link"),
				},
			},
			{
				Name: aws.String("tag:Accessibility"),
				Values: []*string{
					aws.String("public"),
				},
			},
			{
				Name: aws.String("tag:aws:autoscaling:groupName"),
				Values: []*string{
					aws.String(fmt.Sprintf("tzlink-%s-archive", os.Getenv("NETWORK"))),
				},
			},
			{
				Name: aws.String("tag:Mode"),
				Values: []*string{
					aws.String("archive"),
				},
			},
		}}

	res, err := svc.DescribeInstances(params)

	if err != nil {
		panic(fmt.Sprintf("Error when describe instances :  %v", err))
	}

	firstIP := res.Reservations[0].Instances[0].PrivateIpAddress

	log.Print(fmt.Sprintf("Running lambda on instance : %s", *firstIP))

	return firstIP
}

func runCommand(cmd string, conn *ssh.Client) {
	log.Println(fmt.Sprintf("running command %q:", cmd))
	sess, err := conn.NewSession()
	if err != nil {
		panic(err)
	}
	defer sess.Close()

	sessStdOut, err := sess.StdoutPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stdout, sessStdOut)

	sessStderr, err := sess.StderrPipe()
	if err != nil {
		panic(err)
	}
	go io.Copy(os.Stderr, sessStderr)

	err = sess.Run(cmd)
	if err != nil {
		panic(err)
	}
}
