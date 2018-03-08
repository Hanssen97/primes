CC=g++

CFLAGS=-O2



all: compile run

compile:
	$(CC) $(CFLAGS) main.cpp -o primes.out

run:
	./primes.out

clean:
	rm -rf primes.out base.txt primes.txt
