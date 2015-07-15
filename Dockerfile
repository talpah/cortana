FROM golang:1.3-onbuild
CMD ["go-wrapper", "run", "cortana.go"]
