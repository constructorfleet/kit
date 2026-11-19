# acfs/v0.10.3-cf.1

- Preserve existing-file UID/GID before atomic publication by `apply update_file`.
- Refuse publication when original ownership or mode cannot be restored.
- Advertise the `preserve-update-owner` capability without changing protocol 2.
- Extend existing mode/version coverage to check ownership and atomic replacement.
- Add Gitea Linux ACFS and all-module quality checks for the Constructorfleet fork.

Tracking: https://gitea.home.prettybaked.com/constructorfleet/arcane/issues/86.
Validation and release acceptance are recorded in the fixing Gitea pull request.
