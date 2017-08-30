import sys, time;

def progressBar (start, end, curr):
    end  -= start;
    curr -= start;
    start =     0;

    if (curr > end or curr < start):
        print '{0}'.format('\033[91m'),
        sys.stdout.write("\r|{0}{2}|{1} ".format(
            "="*50,
            "ERROR",
            ""
        ))
    elif (curr != end):
        print '{0}'.format('\033[1;35m'),
        sys.stdout.write("\r|{0}{2}|{1}% ".format(
            "="*((curr*50)/end),
            (curr*100)/end,
            " "*int(50-(((curr*50)/end)))
        ))
    else:
        print '{0}'.format('\033[1;32m'),
        sys.stdout.write("\r|{0}|100%".format(
            "="*50,
        ))

    sys.stdout.flush()

    time.sleep(0.00013);
