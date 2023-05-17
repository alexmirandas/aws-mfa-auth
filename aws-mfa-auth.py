import argparse
import os
import boto3

def get_temp_credentials(profile_name, token_code, mfa_serial):
    # Create a session with the specified profile
    session = boto3.Session(profile_name=profile_name)

    # Get temporary credentials using the STS client and MFA token code
    client = session.client('sts')
    response = client.get_session_token(
        DurationSeconds=3600,
        SerialNumber=mfa_serial,
        TokenCode=token_code
    )

    # Return the temporary credentials
    return response['Credentials']

if __name__ == '__main__':
    # Parse command line arguments
    parser = argparse.ArgumentParser(description='Get temporary AWS credentials using MFA')
    parser.add_argument('--profile', required=True, help='Name of the AWS CLI profile to use')
    parser.add_argument('--token', required=True, help='MFA token code')
    parser.add_argument('--mfa-serial', required=True, help='ARN of the MFA device')
    args = parser.parse_args()

    # Get temporary credentials
    temp_creds = get_temp_credentials(args.profile, args.token, args.mfa_serial)

    # Set environment variables with the temporary credentials
    os.environ['AWS_ACCESS_KEY_ID'] = temp_creds['AccessKeyId']
    os.environ['AWS_SECRET_ACCESS_KEY'] = temp_creds['SecretAccessKey']
    os.environ['AWS_SESSION_TOKEN'] = temp_creds['SessionToken']

    # Print a message indicating that the environment variables have been set
    print("Temporary credentials obtained with MFA. Environment variables have been set.")

    # Print the temporary credentials
    print(f"AWS_ACCESS_KEY_ID={temp_creds['AccessKeyId']}")
    print(f"AWS_SECRET_ACCESS_KEY={temp_creds['SecretAccessKey']}")
    print(f"AWS_SESSION_TOKEN={temp_creds['SessionToken']}")
