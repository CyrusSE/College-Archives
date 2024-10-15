#include <iostream>
#include "list.h"

using namespace std;

int main() {
    List L;
    createList(L);

    address P;

    int a, b, c;
    cout << "Masukkan angka pertama : ";
    cin >> a;

    P = allocate(a);
    insertFirst(L, P);
    printInfo(L);

    cout << "Masukkan angka kedua : ";
    cin >> b;

    P = allocate(b);
    insertFirst(L, P);
    printInfo(L);

    cout << "Masukkan angka ketiga : ";
    cin >> c;

    P = allocate(c);
    insertFirst(L, P);
    printInfo(L);

    return 0;
}
