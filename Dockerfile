from golang:1.23-bookworm as build
workdir /app
copy . .
run CGO_ENABLED=0 go build -o oombomb ./main.go

from scratch
copy --from=build /app/oombomb /oombomb
cmd ["/oombomb"]
