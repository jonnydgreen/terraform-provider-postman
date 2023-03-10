# Environments can be imported by specifying the ID of the workspace followed by the ID of the environment (separated by a comma)
terraform import postman_environment.example f358135f-180e-4331-b533-cb23a72b745a,8ce9271d-d4ea-458c-8aef-02cf08ca9558

# Alternatively, environments can be imported by just specifying the ID of the environment. In this case, the default workspace will be used.
terraform import postman_environment.example f358135f-180e-4331-b533-cb23a72b745a,8ce9271d-d4ea-458c-8aef-02cf08ca9558
