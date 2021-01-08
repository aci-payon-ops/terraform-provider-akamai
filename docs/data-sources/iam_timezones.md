---
layout: "akamai"
page_title: "Akamai: akamai_iam_timezones"
subcategory: "IAM"
description: |-
 IAM Timeout Policies
---

# akamai_iam_timezones

Use `akamai_iam_timezones` all time zones Akamai supports. Time zones are in ISO 8601 format. Use the values from this operation to set the timeZone for a user. Administrators use this operation to set a user’s time zone. Users who modify it need to run View time zones for a user profile. The default time zone is GMT.

## Example usage

Basic usage:

```hcl
data "akamai_iam_timezones" "timezones" {
}

output "supported_timezones" {
  value = data.akamai_iam_timezones.timezones
}
```

## Argument reference

There are no arguments for this data source.

## Attributes reference

These attributes are returned:

* `timezone` — The time zone ID.
* `description` - The description of a time zone, including the GMT +/-.
* `offset` - The time zone offset from GMT.
* `posix` - The time zone posix.

[API Reference](https://developer.akamai.com/api/core_features/identity_management_user_admin/v2.html#getadmintimezones)