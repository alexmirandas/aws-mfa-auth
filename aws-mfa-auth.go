package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
)

func getTempCredentials(profileName string, tokenCode string, mfaSerial string) (*sts.Credentials, error) {
	// Create a session with the specified profile
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: profileName,
	})

	if err != nil {
		return nil, err
	}

	// Get an STS client using the session
	svc := sts.New(sess)

	// Get temporary credentials using the STS client and MFA token code
	params := &sts.GetSessionTokenInput{
		DurationSeconds: aws.Int64(3600),
		SerialNumber:    aws.String(mfaSerial),
		TokenCode:       aws.String(tokenCode),
	}
	resp, err := svc.GetSessionToken(params)

	if err != nil {
		return nil, err
	}

	// Return the temporary credentials
	return resp.Credentials, nil
}

func main() {
	// Parse command line arguments
	profileName := flag.String("profile", "", "Name of the AWS CLI profile to use")
	tokenCode := flag.String("token", "", "MFA token code")
	mfaSerial := flag.String("mfa-serial", "", "ARN of the MFA device")
	flag.Parse()

	// Get temporary credentials
	tempCreds, err := getTempCredentials(*profileName, *tokenCode, *mfaSerial)

	if err != nil {
		fmt.Println(err)
		return
	}

	// Set environment variables with the temporary credentials
	os.Setenv("AWS_ACCESS_KEY_ID", *tempCreds.AccessKeyId)
	os.Setenv("AWS_SECRET_ACCESS_KEY", *tempCreds.SecretAccessKey)
	os.Setenv("AWS_SESSION_TOKEN", *tempCreds.SessionToken)

	// Print a message indicating that the environment variables have been set
	fmt.Println("Temporary credentials obtained with MFA. Environment variables have been set.")

	// Print the temporary credentials
	fmt.Printf("AWS_ACCESS_KEY_ID=%s\n", *tempCreds.AccessKeyId)
	fmt.Printf("AWS_SECRET_ACCESS_KEY=%s\n", *tempCreds.SecretAccessKey)
	fmt.Printf("AWS_SESSION_TOKEN=%s\n", *tempCreds.SessionToken)
}
