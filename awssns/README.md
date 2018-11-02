#util/awssns

Uses the AWS SNS service to send text messages to one or multiple phones.

You must specify your AWS credentials `AwsProfile` and `AwsRegion` to run this function.

Your AWS credentials should be at `~/.aws/credentials`, which should look something like:

```
[default]
aws_access_key_id=AKIAIOSFODNN7EXAMPLE
aws_secret_access_key=wJalrXUtnFEMI/K7MDENG/bPxRfiCYEXAMPLEKEY
```

The `AwsProfile` would be `default` in this case.

`AwsRegion` could be, for example, `us-east-1`, or any other region that supports SNS;

TODO: More info...
