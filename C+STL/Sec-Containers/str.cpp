// #include<iostream>
#include <bits/stdc++.h>

using namespace std;

int main(int argc, char const *argv[])
{
    // new string method-1
    string s = "string-1";
    cout << s << endl;
    // new string method-2
    string s_2(3, '6');
    cout << s_2 << endl;
    // input string
    string inputString;
    cin >> inputString;
    cout << inputString << endl;
    // input a line string
    getline(cin, inputString);
    cout << inputString << endl;

    printf("*************************\n");
    s_2 = "666";
    string s_3 = s_2 + s_2;
    cout << s_3 << endl;
    // string functions.
    string str = "test string function.";
    cout << "string: " << str << "; index of 2:" << str[2] << endl;
    cout << "string substr function result: " << str.substr(1) << endl;
    // printf("string substr function result: %s\n", str.substr(1));
    cout << "string: " << str << "; len: "<< str.length() << endl;
    cout << "string: " << str << "; find str: " << str.find("str") << endl;
    // cout << str << str.substr(5);
    return 0;
}
