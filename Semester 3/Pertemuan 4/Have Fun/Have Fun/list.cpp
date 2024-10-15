#include <iostream>
#include "list.h"

using namespace std;

void createList(List &L) {
    first(L) = NULL;
}

address allocate(infotype x) {
    address p = new elmlist;
    info(p) = x;
    next(p) = NULL;
    return p;
}

void insertFirst(List &L, address P) {
    next(P) = first(L);
    first(L) = P;
}

void insertLast(List &L, address P) {
    if (first(L) == NULL) {
        first(L) = P;
    } else {
        address last = first(L);
        while (next(last) != NULL) {
            last = next(last);
        }
        next(last) = P;
    }
}

void insertAfter(address Prec, address P) {
    if (Prec != NULL) {
        next(P) = next(Prec);
        next(Prec) = P;
    }
}

void deleteFirst(List &L, address &P) {
    if (first(L) == NULL) {
        P = NULL;
    } else {
        P = first(L);
        first(L) = next(P);
        next(P) = NULL;
    }
}

void deleteLast(List &L, address &P) {
    if (first(L) == NULL) {
        P = NULL;
    } else if (next(first(L)) == NULL) {
        P = first(L);
        first(L) = NULL;
    } else {
        address prev = NULL;
        P = first(L);
        while (next(P) != NULL) {
            prev = P;
            P = next(P);

        }
        next(prev) = NULL;
    }
}

void deleteAfter(address Prec, address &P) {
    if (Prec != NULL && next(Prec) != NULL) {
        P = next(Prec);
        next(Prec) = next(P);
        next(P) = NULL;
    }
}

address searchInfo(List L, infotype x) {
    address p = first(L);
    while (p != NULL && info(p) != x) {
        p = next(p);
    }
    return p;
}

//#void printInfo(List L) {
//#    address p = first(L);
//#    while (p != NULL) {
//#        cout << info(p) << "";
//#        p = next(p);
//#    }
//#    cout << endl;
//#}

void printInfo(List L) {
    if (L.first != NULL){
        address P = L.first;
        while (P != NULL){
            cout << P->info << ", ";
            P = next(P);
        }
    }

    cout << endl;
}
