Use this plugin for deplying an application to AWS OpsWorks. You can override
the default configuration with the following parameters:

* `access_key_id` - AWS access key ID
* `secret_access_key` - AWS secret access key
* `region` - AWS availability zone

## Example

The following is a sample configuration in your .drone.yml file:

```yaml
deploy:
  aws_opsworks:
    access_key_id:
    secret_access_key:
    region:
```
