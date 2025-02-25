go build -tags lambda.norpc -o build/bootstrap cmd/main.go
cd build
zip build.zip bootstrap
cd ..