echo go-torch
go get github.com/uber/go-torch
# git clone git@github.com:brendangregg/FlameGraph.git
go-torch --binaryname bloomsky-api-go.test -b prof.cpu
open -a "google chrome" torch.svg 
