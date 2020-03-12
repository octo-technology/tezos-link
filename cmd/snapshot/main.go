package main

import (
    "context"
    "fmt"
    "github.com/aws/aws-lambda-go/lambda"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/s3"
    "github.com/aws/aws-sdk-go/service/s3/s3manager"
    "golang.org/x/crypto/ssh"
    "io"
    "log"
    "os"
)

func main() {
    lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context) (string, error) {
    log.Print("Snapshot export event received.")

    config := &ssh.ClientConfig {
        User: "ec2-user",
        Auth: []ssh.AuthMethod{
            getPublicKeyFromS3(),
        },
        HostKeyCallback: ssh.InsecureIgnoreHostKey(),
    }

    conn, err := ssh.Dial("tcp", os.Getenv("NODE_IP"), config)
    if err != nil {
        panic(err)
    }
    defer conn.Close()

    runCommand("/usr/bin/whoami", conn)

    return "Snapshot done!", nil
}

func getPublicKeyFromS3() ssh.AuthMethod {
    sess, _ := session.NewSession(&aws.Config{
        Region: aws.String(os.Getenv("AWS_REGION"))},
    )

    downloader := s3manager.NewDownloader(sess)

    var buf []byte
    writer := aws.NewWriteAtBuffer(buf)
    S3Key := os.Getenv("S3_KEY")
    _, err := downloader.Download(writer,
        &s3.GetObjectInput{
            Bucket: aws.String(os.Getenv("S3_BUCKET")),
            Key:    aws.String(S3Key),
        })
    if err != nil {
        panic(fmt.Sprintf("Unable to download item %q, %v", S3Key, err))
    }

    signer, err := ssh.ParsePrivateKey(buf)
    if err != nil {
        panic(err)
    }

    log.Print("Got private key from S3 bucket.")
    return ssh.PublicKeys(signer)
}

func runCommand(cmd string, conn *ssh.Client) {
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
