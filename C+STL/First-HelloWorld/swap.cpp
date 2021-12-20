#include<iostream>

using namespace std;

void swap2(int &a, int &b) 
{
    int c = a;
    a = b;
    b = c;
}

int main(int argc, char const *argv[])
{
    int a,b;
    cin >> a >> b;
    swap2(a, b);
    cout << a<<" "<<b;
    return 0;
}
