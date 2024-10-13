#include <iostream>
#include "kerucut.h"

using namespace std;

int main()
{
    float r, t, s;
    cout << "Masukan jari-jari kerucut : ";
    cin >> r;
    cout << "Masukan tinggi kerucut : ";
    cin >> t;
    cout << "Volume : " << volume(r, t) << endl;
    cout << "Masukan sisi kerucut : ";
    cin >> s;
    cout << "Luas Permukaan : " << LP(r, s) << endl;
    return 0;
}
