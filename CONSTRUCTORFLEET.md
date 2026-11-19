# Constructorfleet Kit fork

Gitea owns review and merges: https://gitea.home.prettybaked.com/constructorfleet/kit.
GitHub is the public Go-module source mirror: https://github.com/constructorfleet/kit.
Upstream history, module paths and licenses are preserved.

Use named feature/fix branches and coherent commits. Open a Gitea pull request,
run the required Linux ACFS and quality checks, then squash merge to main.
Never push directly or force push to main, bypass checks, or overwrite tags.
Mirror merged main and immutable module tags to GitHub after acceptance.
Record changes and validation in CONSTRUCTORFLEET_RELEASE_NOTES.md.

The ACFS ownership patch supports Arcane QA issue #86. Existing-file updates
keep the original Unix UID/GID and mode while retaining atomic replacement.
An ownership or mode restoration failure leaves the original file untouched.
New files retain the caller's normal creation identity. Protocol 2 remains
compatible; `acfs version` advertises `preserve-update-owner` for safe callers.

Arcane builds the companion CLI from this module and packages it with each
manager/agent release. Its existing tools image supplies helper utilities;
updating the Go dependency alone does not update that image's bundled CLI.

Run `just format all`, then `just lint` on Linux and `cd acfs && go test -race ./...`.
The unchanged upstream macOS suite has symlink-mode/setuid failures and a FIFO
sync timeout; distinguish those from Linux helper acceptance.
