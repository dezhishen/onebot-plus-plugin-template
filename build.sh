
#export CGO_ENABLED=1
plugin_name=$1
mkdir -p ./plugins
if [ -z "$plugin_name" ]; then
    read -p "Enter the plugin name: " plugin_name
    if [ -z "$plugin_name" ]; then
        echo "Plugin name is empty"
        exit 1
    fi
fi
export GOARCH=amd64
export GOOS=windows
go build -o ./plugins/${plugin_name}-windows-amd64.exe ./main.go

export GOARM=7
export GOARCH=arm64
export GOOS=linux

go build -o ./plugins/${plugin_name}-linux-arm64v7 ./main.go

export GOARCH=amd64
export GOOS=darwin
go build -o ./plugins/${plugin_name}-linux-amd64 ./main.go

