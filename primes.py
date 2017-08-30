import time, os, sys, logging, math, thread, threading, subprocess, pyio;

BLUE, RED, WHITE, YELLOW, MAGENTA, GREEN, END = '\33[94m', '\033[91m', '\33[97m', '\33[93m', '\033[1;35m', '\033[1;32m', '\033[0m'

primes, base, searcher = [], [2,3,5,7], None;

header = ('{0} > {1}'.format(BLUE, END));
f      = ('{0}From > {1}'.format(BLUE, END))
t      = ('{0}To   > {1}'.format(BLUE, END))



def parseArgs(start, end):
    if (start < 0 or end < 0):
        print '{0}\nInterval cannot be negative'.format(RED);
        return False;
    elif (start > end):
        print '{0}\nThe smallest number should be placed first'.format(RED);
        return False;

    return True;



def updateBaseFromFile():
    basefile            = open("base.txt", "r");
    text                = basefile.read();
    global base; base   = map(int, text.split(','));



def updateBaseFile(l):
    basefile = open("base.txt", "w");
    basefile.write(','.join(map(str, l)));



def setup():
    os.remove("pyio.pyc");

    if (not os.path.exists("searcher.out")):
        subprocess.call(["g++", "-o", "searcher.out", "searcher.cpp"]);

    if (not os.path.exists("base.txt")):
        updateBaseFile( base );
    else:
        updateBaseFromFile();

    global searcher; searcher = subprocess.Popen(
        ["./searcher.out"],
        bufsize=1,
        stdin=subprocess.PIPE,
        stdout=subprocess.PIPE,
        universal_newlines=True
    );



def extendBase(end):
    start = base[len(base)-1] + 2; # +2 since list is only odd numbers.

    base.extend(findPrimes(start, end));
    updateBaseFile(base);

    searcher.stdin.write('0\n'); # Updates prime base in external searcher.out



def findPrimes(start, end):
    curr    = start;
    ceil    = int(math.ceil(math.sqrt(end)));
    result  = [];
    tick    = 0;
    step    = int(math.ceil((end-start)/200))+1;

    if (ceil > base[len(base)-1]):
        extendBase(ceil)

    if (curr < 3):
        result.append(2);
        curr = 3;
    elif (curr % 2 == 0):
        curr += 1;

    while(curr < end):
        if (search(curr)):
            result.append(curr);

        if (tick % step == 0):
            pyio.progressBar(start, end, curr);

        curr += 2; tick += 1;

    pyio.progressBar(start, end, end);

    print '{0}\n'.format(WHITE),

    return result;



def search(number):
    searcher.stdin.write(str(number)+'\n');
    return int(searcher.stdout.readline().replace("\n", ""));



def printPrimes(primes):
    print '{0}'.format(GREEN);
    for prime in primes:
        time.sleep(0.00025)
        print(prime)
    print '{0}'.format(END);



def filePrimes(primes):
    f           = open('primes.txt', 'w')
    string      = '';
    curr, end   = 0, len(primes);
    step        = int(math.ceil(end/200)+1);

    print 'Saving primes to file...';

    while(curr < end):
        if (curr % step == 0):
            pyio.progressBar(0, end, curr);

        string = str(primes[curr])+' ';
        f.write(string)

        curr += 1;

    pyio.progressBar(0, end, end)

    print '{0}\a\n'.format(WHITE),



def optionBanner():
    print('\n{0}Choose option from menu:\n'.format(END))
    time.sleep(0.1)
    print('\t{0}[{1}1{2}]{3} Find primes\n').format(YELLOW, BLUE, YELLOW, WHITE)
    time.sleep(0.1)
    print('\t{0}[{1}2{2}]{3} Print primes').format(YELLOW, BLUE, YELLOW, WHITE)
    time.sleep(0.1)
    print('\t{0}[{1}3{2}]{3} Save primes').format(YELLOW, BLUE, YELLOW, WHITE)
    time.sleep(0.1)
    print('\n\t{0}[{1}M{2}]{3} Clear cache').format(YELLOW, BLUE, YELLOW, WHITE)
    time.sleep(0.1)
    print('\t{0}[{1}C{2}]{3} Clear screen').format(YELLOW, BLUE, YELLOW, WHITE)
    time.sleep(0.1)
    print('\t{0}[{1}E{2}]{3} Exit Program\n').format(YELLOW, BLUE, YELLOW, WHITE)




setup();

while True:
    optionBanner();
    choice = raw_input(header)


    if choice.upper() == 'E' or choice.upper() == 'EXIT': #--------------------
        break;


    elif choice == '1': #------------------------------------------------------
        print("\nFind primes: ");
        start   = int(raw_input(f));
        end     = int(raw_input(t));

        if (parseArgs(start, end)):
            startTime = time.time();

            ceil = int(math.ceil(math.sqrt(end)));
            if (ceil > base[len(base)-1]):
                print '\nExtending prime base...';
                extendBase(ceil+int(math.ceil(math.sqrt(ceil))));

            print '\nSearching for prime numbers...';
            primes = findPrimes(start, end);

            endTime = time.time();
            m, s = divmod(round(endTime-startTime), 60);
            h, m = divmod(m, 60);

            print '\aFound {0}{1} {2}primes in {0}{4}{2}h {0}{5}{2}m {0}{6}{2}s{3}'.format(YELLOW, len(primes), WHITE, END, int(h), int(m), int(s));


    elif choice == '2': #------------------------------------------------------
        if (len(primes) == 0):
            print '{0}\nThere are no primes{1}'.format(RED, END);
        else:
            printPrimes(primes);


    elif choice == '3': #------------------------------------------------------
        if (len(primes) == 0):
            print '{0}\nThere are no primes{1}'.format(RED, END);
        else:
            save = threading.Thread(target=filePrimes(primes))
            save.start();


    elif choice.upper() == 'M' or choice.upper() == 'MEMORY': #----------------
        primes, base = [], [2, 3, 5, 7];

        if (os.path.exists("searcher.out")):
            os.remove("searcher.out");
        if (os.path.exists("base.txt")):
            os.remove("base.txt");
        if (os.path.exists("primes.txt")):
            os.remove("primes.txt");

        print '\n{0}Cache cleared!{1}'.format(YELLOW, END);


    elif choice.upper() == 'C' or choice.upper() == 'CLEAR': #-----------------
        os.system("clear||cls")


    else: #--------------------------------------------------------------------
        print("\n{0}ERROR: Please select a valid option.{1}\n").format(RED, END)


searcher.kill();
