#include <iostream>
#include "list.h"

using namespace std;

int main() {
    List L;
    createList(L);
    address P, Prec;
    int NIM[10];
    cout << "Masukkan NIM perdigit" << endl;
    for (int i = 0; i < 10; i++) {
        cout << "Digit " << i + 1 << " : ";
        cin >> NIM[i];
        P = allocate(NIM[i]);
        insertLast(L, P);
    }
    cout << "Isi list setelah insertLast:" << endl;
    printInfo(L);
    P = allocate(111);
    insertFirst(L, P);
    cout << "Isi list setelah insertFirst (111):" << endl;
    printInfo(L);

    int searchNIM;
    cout << "Masukkan NIM untuk dicari: ";
    cin >> searchNIM;
    P = searchInfo(L, searchNIM);
    if (P != NULL) {
        cout << "NIM ditemukan: " << info(P) << endl;
    } else {
        cout << "NIM tidak ditemukan" << endl;
    }

    deleteFirst(L, P);
    cout << "Isi list setelah delete first:" << endl;
    printInfo(L);

    deleteLast(L, P);
    cout << "Isi list setelah delete last:" << endl;
    printInfo(L);

    Prec = first(L);
    P = allocate(999);
    insertAfter(Prec, P);
    cout << "Isi list setelah insertAfter (999 after first element):" << endl;
    printInfo(L);

    Prec = first(L);
    deleteAfter(Prec, P);
    cout << "Isi list setelah deleteAfter (delete element after first):" << endl;
    printInfo(L);

    return 0;
}
