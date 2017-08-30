#include <iostream>
#include <math.h>
#include <stdlib.h>
#include <string>
#include <fstream>
#include <vector>


std::vector<size_t> base;


bool search(size_t number);
void updateBase();




int main() {
  std::string buffer = "";
  updateBase();

  while (std::cin >> buffer) {
    if (buffer == "0") updateBase();
    else std::cout << search(std::stod(buffer)) << std::endl;
  }

  return 0;
}




void updateBase() {
  std::ifstream baseFile ("base.txt");

  if (baseFile.is_open()) {
    std::string buffer = "";

    while (getline(baseFile, buffer, ',')) {
      base.push_back(stod(buffer));
    }
  }
}




bool search(size_t number) {
  size_t cap = ceil(sqrt(number)), i = 0;

  for (; i < base.size(); ++i) {
    if (number == base.at(i))                 return true;
    else if (fmod(number, base.at(i)) == 0)   return false;

    if (base.at(i) > cap) break;
  }

  return true;
}
