#include <iostream>
#include "list.h"


using namespace std;

void createList_103032300152(List &L){
    first(L) = NULL;
}

address allocate_103032300152(infotype x){
    address P = new elmlist;
    info(P) = x;
    next(P) = NULL;
    return P;
}

void insertLast_103032300152(List &L, address P){
    if (first(L) == NULL){
        first(L) = P;
    }else{
        address last = first(L);
        while (next(last) != NULL){
            last = next(last);
        }
        next(last) = P;
    }
}

void show_103032300152(List L){
    if (first(L) != NULL){
        address P = first(L);
        while (P != NULL){
            cout << "Data : " << info(P) << endl;
            P = next(P);
        }
    }else {
        cout << "List kosong." << endl;
    }
}

address shortestName_103032300152(List L) {
    address save = NULL;
    if (first(L) != NULL) {
        address P = first(L);
        save = P;
        while (P != NULL) {
            if (info(P).length() < info(save).length()) {
                save = P;
            }
            P = next(P);
        }
    } else {
        cout << "List kosong." << endl;
    }
    return save;
}

void showFirstKNameC_103032300152(List L, int K, char c) {
    int count = 0;
    if (first(L) != NULL) {
        address P = first(L);
        while (P != NULL && count < K) {
            if (info(P)[0] == c) {
                cout << "Data : " << info(P) << endl;
                count++;
            }
            P = next(P);
        }
        if (count == 0) {
            cout << "Tidak ada pengunjung dengan nama berawalan '" << c << "'" << endl;
        }
    } else {
        cout << "List kosong." << endl;
    }
}


