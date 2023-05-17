## To use
Run the desire file from the command line with the following arguments:

## Golang
```
go run get_temp_creds.go --profile <profile_name> --token <token_code> --mfa-serial <mfa_serial>
```

## Python
```
python get_temp_creds.py <profile_name> <token_code> <mfa_serial>
```


Replace <profile_name> with the name of your AWS CLI profile, <token_code> with your MFA token code, and <mfa_serial> with the ARN of your MFA device.

This script will print the temporary credentials to the console in the format of environment variables and export into your terminal or command prompt to set the credentials
