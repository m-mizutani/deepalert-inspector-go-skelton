# deepalert-inspector-go-skelton

## Deploy

Prepare a config file.

```go:config.json
{
    "StackName": "deepalert-inspector-test",
    "Region": "ap-northeast-1",
    "CodeS3Bucket": "mizutani-security-log.mgmt",
    "CodeS3Prefix": "functions",
    "SecretArn": "arn:aws:secretsmanager:ap-northeast-1:783957204773:secret:deepalert-test/sample-UXkOdC",
    "DeepAlertStackName": "deepalert-test"
}
```

Then, run

```bash
$ env STACK_CONFIG=config.json make deploy
```