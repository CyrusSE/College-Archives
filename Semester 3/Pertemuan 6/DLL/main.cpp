#include <iostream>
#include "DLL.h"

using namespace std;

int main() {
    List L;
    createList(L);

    infotype band1 = {"Foster The People", "Imagination"};
    address p1 = createNewElmt(band1);
    insertFirst(L, p1);

    infotype band2 = {"The Corrs", "Radio"};
    address p2 = createNewElmt(band2);
    insertLast(L, p2);

    infotype band3 = {"Maroon 5", "Sugar"};
    address p3 = createNewElmt(band3);
    insertLast(L, p3);

    cout << "=== Daftar Lagu Awal ===" << endl;
    show(L);

    infotype band4 = {"Coldplay", "Yellow"};
    address p4 = createNewElmt(band4);
    insertAfter(L, p1, p4);

    cout << "\n=== Setelah Menambahkan Coldplay ===" << endl;
    show(L);

    cout << "\nMenghapus lagu 'Radio'..." << endl;
    removeLagu("Radio", L);

    cout << "\n=== Daftar Lagu Setelah Penghapusan ===" << endl;
    show(L);

    string searchSong = "Sugar";
    address found = findLagu(searchSong, L);
    if (found != NULL) {
        cout << "\nLagu '" << searchSong << "' ditemukan oleh band: " << info(found).name << endl;
    }
    else {
        cout << "\nLagu '" << searchSong << "' tidak ditemukan." << endl;
    }

    List L2, L3;
    createList(L2);

    infotype band5 = {"Imagine Dragons", "Believer"};
    address p5 = createNewElmt(band5);
    insertLast(L2, p5);

    infotype band6 = {"Linkin Park", "Numb"};
    address p6 = createNewElmt(band6);
    insertLast(L2, p6);

    cout << "\n=== Daftar Lagu List L2 ===" << endl;
    show(L2);

    concat(L, L2, L3);
    cout << "\n=== Daftar Lagu Setelah Menggabungkan L dan L2 ke L3 ===" << endl;
    show(L3);

    cout << "\nMenghapus lagu 'Imagination'..." << endl;
    removeLagu("Imagination", L3);
    cout << "\nMenghapus lagu 'Numb'..." << endl;
    removeLagu("Numb", L3);

    cout << "\n=== Daftar Lagu (List 3) Setelah Penghapusan ===" << endl;
    show(L3);
    return 0;
}
