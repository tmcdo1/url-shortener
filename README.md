# url-shortener
This is a basic go package that you can run and create your own shortened urls. There is a main package provided to get up and running with a premade server.

## Build
<code>go build main/main.go -o server</code>

## Dependencies
A <code>glide.yaml</code> file is provided to handle dependencies with glide. At the moment, there is one for parsing yaml that has not yet been used.

## Future Plans
- Include YAML support
- Include optional file definition for paths
- Include database support for paths
- Add a homepage