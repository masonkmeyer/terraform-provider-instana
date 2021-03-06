# Custom Event Specification with System Rule Resource

Configuration of a custom event specification based on a system rule.

API Documentation: <https://instana.github.io/openapi/#operation/putCustomEventSpecification>

Custom event resources support `default_name_prefix` and `default_name_suffix`. The string will be appended automatically
to the name of the custom event.

## Example Usage

```hcl
resource "instana_custom_event_spec_system_rule" "example" {
  name            = "name"
  description     = "description"
  query           = "query"
  enabled         = true
  triggering      = true
  expiration_time = 60000

  rule_severity       = "warning"
  rule_system_rule_id = "system-rule-id"
}
```

## Argument Reference

* `name` - Required - The name of the custom event specification
* `description` - Required - The description text of the custom event specification
* `query` - Optional - The dynamic filter query for which the rule should be applied to
* `enabled` - Optional - Boolean flag if the rule should be enabled - default = true
* `triggering` - Optional - Boolean flag if the rule should trigger an incident - default = false
* `expiration_time` - Optional - The grace period in milliseconds until the issue is closed
* `rule_severity` - Required - The severity of the rule - allowed values: `warning`, `critical`
* `rule_system_rule_id` - Required - The id of the instana system rule of the given even