---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "postman_environment Data Source - terraform-provider-postman"
subcategory: ""
description: |-
  
---

# postman_environment (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `id` (String) The environment's ID.

### Optional

- `is_public` (Boolean) If true, the environment is public.
- `owner` (String) The environment owner's ID.
- `workspace` (String) The environment's workspace ID. If not specified, the default workspace is used.

### Read-Only

- `created_at` (String) The date and time at which the environment was created.
- `name` (String) The environment's name.
- `updated_at` (String) The date and time at which the environment was last updated.

