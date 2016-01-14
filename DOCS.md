Use this plugin for deplying an application to OpsWorks. You can override the
default configuration with the following parameters:

* `access_key` - AWS access key ID
* `secret_key` - AWS secret access key
* `region` - AWS availability zone

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  opsworks:
    access_key: 970d28f4dd477bc184fbd10b376de753
    secret_key: 9c5785d3ece6a9cdefa42eb99b58986f9095ff1c
    region: us-east-1
```
