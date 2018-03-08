#include <iostream>
#include <string>
#include <unistd.h>
#include <stdio.h>
#include <vector>
#include <math.h>
#include <fstream>
#include <ctime>



namespace global {
  std::vector<size_t> primes;
  std::vector<size_t> base;
}



namespace io {
  namespace colors {
    std::string blue    = "\33[94m";
    std::string red     = "\033[91m";
    std::string white   = "\33[97m";
    std::string yellow  = "\33[93m";
    std::string magenta = "\033[1;35m";
    std::string green   = "\033[1;32m";
    std::string grey    = "\033[0m";
  }


  namespace headers {
    std::string input = colors::blue + " > " + colors::grey;
  }



  // Functions ----------------------------------------------------------------
  void printBanner();
  void printMenuOption(char index, std::string title);
  void getInput(char &input, std::string text);
  void getInput(size_t &input, std::string text);
  void printError(std::string text);
  void progressBar(size_t start, size_t end, size_t curr);
  void printPrimes(std::vector<size_t> primes);
  // --------------------------------------------------------------------------



  // --------------------------------------------------------------------------
  void clearScreen() {
    for (int i = 0; i < 100; i++) {
      std::cout << "\n";
    }
  }


  // --------------------------------------------------------------------------
  void printBanner() {
    printf("\n\n%sChoose option from menu:\n\n", colors::white.c_str());

    printMenuOption('F', "Find primes\n");
    printMenuOption('P', "Print primes");
    printMenuOption('S', "Save primes\n");
    printMenuOption('D', "Delete data");
    printMenuOption('C', "Clear screen");
    printMenuOption('E', "Exit Program");
  }


  // --------------------------------------------------------------------------
  void printMenuOption(char index, std::string title) {
    usleep(10000);
    printf("\t%s[%s%c%s] %s%s\n",
      colors::yellow.c_str(),
      colors::blue.c_str(),
      index,
      colors::yellow.c_str(),
      colors::grey.c_str(),
      title.c_str()
    );
  }


  // --------------------------------------------------------------------------
  void getInput(char &input, std::string text = "") {
    printf("\n%s%s > %s", colors::blue.c_str(), text.c_str(), colors::grey.c_str());
    std::cin >> input;

    if (std::cin.fail())  printError("could not read input");
    else                  input = (char) toupper(input);
  }


  // --------------------------------------------------------------------------
  void getInput(size_t &input, std::string text = "") {
    std::cout << colors::blue << text << headers::input;
    std::cin >> input;

    if (std::cin.fail()) printError("could not read input");
  }


  // --------------------------------------------------------------------------
  void printError(std::string text) {
    std::cout << colors::red << "\nERROR: " + text + " \n" << colors::white;
    std::cin.clear();
    std::cin.ignore(256,'\n');
    exit(1);
  }


  // --------------------------------------------------------------------------
  void progressBar(size_t start, size_t end, size_t curr) {
    int barWidth = 50;
    curr -= start;    end -= start;     start = 0;

    int pos = (curr*barWidth / end);

    if      (curr <  end)   std::cout << io::colors::magenta << "|";
    else if (curr == end)   std::cout << io::colors::green << "\a|";
    else                    std::cout << io::colors::red << "ERROR! |";

    for (size_t i = 0; i < barWidth; ++i) {
        if (i < pos)  std::cout << "=";
        else          std::cout << " ";
    }
    std::cout << "|" << (curr * 100 / end) << "%   \r";
    std::cout.flush();
  }


  // --------------------------------------------------------------------------
  void printPrimes() {
    std::cout << colors::white << std::endl;
    for (size_t i = 0; i < global::primes.size(); ++i) {
      usleep(1000);
      std::cout << global::primes.at(i) << std::endl;
    }
  }
}



// Functions ------------------------------------------------------------------
void                init();
void                clearScreen();
bool                parseInterval(size_t start, size_t end);
void                findPrimes();
void                extendBase(size_t end);
std::vector<size_t> getPrimes(size_t start, size_t end);
bool                search(size_t number);
void                updatePrimesFile();
void                updateBaseFile();
void                removeFiles();
void                removeFile(std::string filename);
void                updateBaseFromFile();
void                divmod(int a, int b, int &div, int &mod);
// ----------------------------------------------------------------------------




// ----------------------------------------------------------------------------
int main() {
  io::clearScreen();
  init();
  updateBaseFromFile();

  for (char input; input != 'E';) {
    io::printBanner();
    io::getInput(input);

    switch(input) {
      case 'F': {
        findPrimes();
        break;
      }

      case 'P': {
        io::printPrimes();
        break;
      }

      case 'S': {
        updatePrimesFile();
        break;
      }

      case 'D': {
        std::cout << std::endl;
        init();
        removeFiles();
        std::cout << io::colors::green << "Cache cleared!" <<std::endl;
        break;
      }

      case 'C': {
        io::clearScreen();
        break;
      }

      default: {
        std::cout << io::colors::red << "\nPlease select a valid option\n";
        break;
      }
    }
  }

  io::clearScreen();

  return 0;
}


// ----------------------------------------------------------------------------
void init() {
  std::cout << io::colors::yellow << "initializing data...\n";
  global::primes.clear();   global::base.clear();
  global::base.push_back(2);
  global::base.push_back(3);
  global::base.push_back(5);
  global::base.push_back(7);
}


// ----------------------------------------------------------------------------
bool parseInterval(size_t start, size_t end) {
  if (start > end) {
    io::printError("invalid interval");
    return false;
  }
  return true;
}


// ----------------------------------------------------------------------------
void findPrimes() {
  size_t start, end;

  std::cout << io::colors::white << "\nFind Primes:\n";
  io::getInput(start, "From"); io::getInput(end, "To  ");

  std::time_t startTime = std::time(0);

  if (parseInterval(start, end)) {
    size_t threshold = ceil(sqrt(end));
    if (threshold > global::base.at(global::base.size()-1)) {
      std::cout << io::colors::grey << "\nExtending prime base...\n";
      extendBase(threshold + ceil(sqrt(threshold)));
    }

    std::cout << io::colors::white << "\n\nSearching for prime numbers...\n";


    std::vector<size_t> primes = getPrimes(start, end);

    int hour, min, sec;
    divmod((int)(std::time(0) - startTime), 60, min, sec);
    divmod(min, 60, hour, min);



    printf("%s\n\nFound %s%i%s prime numbers in %s%ih %im %is %s\n",
      io::colors::white.c_str(),
      io::colors::green.c_str(),
      int(primes.size()),
      io::colors::white.c_str(),
      io::colors::yellow.c_str(),
      hour, min, sec,
      io::colors::white.c_str()
    );

    global::primes = primes;
  }
}


// ----------------------------------------------------------------------------
void extendBase(size_t end) {
  size_t start = global::base.at(global::base.size()-1) + 2;
  std::vector<size_t> extension = getPrimes(start, end);

  global::base.insert(global::base.end(), extension.begin(), extension.end());

  updateBaseFile();
}


// ----------------------------------------------------------------------------
std::vector<size_t> getPrimes(size_t start, size_t end) {
  std::vector<size_t> result;
  size_t curr      = start,
         threshold = ceil(sqrt(end)),
         tick      = 0,
         step      = ceil((end-start)/200)+1;

  if (threshold > global::base.at(global::base.size()-1))
    extendBase(threshold);

  if (curr < 3) {
    result.push_back(2); curr = 3;
  } else if (curr % 2 == 0) {
    ++curr;
  }

  while(curr < end) {
    if ( search(curr) )     result.push_back(curr);

    if ( tick % step == 0 ) io::progressBar(start, end, curr);

    curr += 2; tick++;
  }
  io::progressBar(0, 100, 100);

  return result;
}


// ----------------------------------------------------------------------------
bool search(size_t number) {
  size_t cap = ceil(sqrt(number)), i = 0;

  for (; i < global::base.size(); ++i) {
    if (number == global::base.at(i))               return true;
    else if (fmod(number, global::base.at(i)) == 0) return false;

    if (global::base.at(i) > cap) break;
  }

  return true;
}


// ----------------------------------------------------------------------------
void updatePrimesFile() {
  std::ofstream file;
  size_t end  = global::primes.size(),
         step = ceil(end/200)+1;

  file.open("primes.txt");

  std::cout << io::colors::white << "\nSaving primes...\n";

  for (size_t i = 0; i < end; ++i) {
    if (i % step == 0) {
      io::progressBar(0, end, i);
    }
    file << global::primes.at(i) << " ";
  }
  io::progressBar(0, 100, 100);

  file.close();
}


// ----------------------------------------------------------------------------
void updateBaseFile() {
  std::ofstream file;
  size_t end  = global::base.size();

  file.open("base.txt");

  for (size_t i = 0; i < end; ++i) {
    file << global::base.at(i) << ",";
  }
}


// ----------------------------------------------------------------------------
void removeFiles() {
  std::cout << io::colors::yellow << "Removing files:\n";
  removeFile("primes.txt");
  removeFile("base.txt");
  std::cout << "\n";
}


// ----------------------------------------------------------------------------
void removeFile(std::string filename) {
  std::ifstream file(filename.c_str());

  if ((bool)file) {
    if( remove(filename.c_str()) != 0 ) io::printError("could not delete file");

    std::cout << io::colors::yellow << "\t" << filename << "\n";
  }

}


// ----------------------------------------------------------------------------
void updateBaseFromFile() {
  std::ifstream file ("base.txt");

  if (file.is_open()) {
    std::cout << io::colors::yellow << "updating prime base...\n";

    std::string buffer = "";

    while (getline(file, buffer, ',')) {
      global::base.push_back(stod(buffer));
    }
  }
}


// ----------------------------------------------------------------------------
void divmod(int a, int b, int &div, int &mod) {
  div = a / b;
  mod = a % b;
}
