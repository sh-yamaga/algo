# /bin/bash
go mod download
go install github.com/spf13/cobra-cli@latest

chmod 600 ~/.ssh/id_ed25519
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_ed25519