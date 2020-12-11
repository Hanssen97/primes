run:
	go run *.go

clean:
	rm -rf primes.txt

test:
	go test -v ./...