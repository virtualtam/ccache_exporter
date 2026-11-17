schema_version = 1

project {
  license          = "MIT"
  copyright_holder = "VirtualTam"

  header_ignore = [
    # SCM
    ".git/**",

    # IDEs
    ".idea/**",

    # Docker Compose
    "docker-compose*.yml",

    # Promu
    ".promu.yml",

    # Yamllint
    ".yamllint.yml",
    ".github/.yamllint.yml",

    # Docker Compose service configuration
    "docker/**"
  ]
}
