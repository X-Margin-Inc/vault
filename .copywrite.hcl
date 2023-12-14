schema_version = 1

project {
  license        = "BUSL-1.1"
<<<<<<< HEAD
  copyright_year = 2024
=======
  copyright_year = 2023
>>>>>>> 4cb759cfc9 (fixed log)

  # (OPTIONAL) A list of globs that should not have copyright/license headers.
  # Supports doublestar glob patterns for more flexibility in defining which
  # files or folders should be ignored
  header_ignore = [
<<<<<<< HEAD
    "helper/pkcs7/**",
=======
    "builtin/credential/aws/pkcs7/**",
>>>>>>> 4cb759cfc9 (fixed log)
    "ui/node_modules/**",
    "enos/modules/k8s_deploy_vault/raft-config.hcl",
    "plugins/database/postgresql/scram/**",
  ]
}
