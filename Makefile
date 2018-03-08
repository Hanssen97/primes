CC=g++

CFLAGS=-O2



all: compile run

compile:
	$(CC) $(CFLAGS) main.cpp -o primes.out

run:
	./primes.out
