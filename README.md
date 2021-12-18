# Execute SAM locally
cd build
GOARCH=amd64 GOOS=linux go build ../main.go
cd ..
sam local invoke FFprobeLambdaFunction -t template.yaml

# Included ffprobe executable
The included ffprobe executable was downloaded from https://ffbinaries.com/downloads
under linux-64
