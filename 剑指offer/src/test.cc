#include <iostream>

using namespace std;

void test(int* t) {
  cout << t << ' ' << *t << endl;
  int num = 4;
  t = &num;
  cout << t << ' ' << *t << endl;
}

int main() {
  int number = 3;
  int* p = &number;
  cout << p << " " << *p << endl;
  test(p);
  cout << p << " " << *p << endl;
  return 0;
}