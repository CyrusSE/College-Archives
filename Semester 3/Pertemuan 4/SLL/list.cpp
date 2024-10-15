#include <iostream>
#include "list.h"

using namespace std;

void createList(List &L){
    first(L) = NULL;
}

address allocate(infotype x){
    address p = new elmlist;
    info(p) = x;
    next(p) = NULL;
    return p;
}

void insertFirst(List &L, address P){
    next(P) = first(L);
    first(L) = P;
}

void printInfo(List L){
    address p = first(L);
    while (p != NULL){
        cout << info(p) << ", ";
        p = next(p);
    }
    cout << endl;
}
