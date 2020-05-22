#include <iostream>
#include <vector>
#include <sstream>

using namespace std;

void getElementToString_01(string input, vector<string> &output) {
        istringstream ss(input);

        while (ss.good()) {
          string substr;
          getline(ss, substr, '*');
          output.push_back(substr);
        }
}

void getElementToInt_02(string input, vector<int> &output) {
    stringstream ss(input);

    for (int i; ss >> i;) {
        output.push_back(i);    
        if (ss.peek() == '*')
            ss.ignore();
    }
}

// just for demo
int main() { 
  string input;
  vector<int> intOut;
  vector<string> stringOut;
  getline(cin, input);

  getElementToString_01(input, stringOut);
  getElementToInt_02(input, intOut);
}
