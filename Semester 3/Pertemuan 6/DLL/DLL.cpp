#include "DLL.h"

bool isEmpty(List L) {
    address p;
    p = first(L);
    if (p == NULL) {
        return true;
    } else {
        return false;
    }
}

void createList(List& L) {
    first(L) = NULL;
    last(L) = NULL;
}


address createNewElmt(infotype X){
    address p = new elmtList;
    info(p) = X;
    next(p) = NULL;
    prev(p) = NULL;
    return p;
}

void insertFirst(List &L, address P){
    if (isEmpty(L)) {
        first(L) = P;
        last(L) = P;
    } else{
        next(P) = first(L);
        prev(first(L)) = P;
        first(L) = P;
    }
}

void insertAfter(List &L, address Prec, address P) {
    prev(P) = Prec;
    next(P) = next(Prec);
    if (next(Prec) == NULL) {
        last(L) = P;
    }
    else {
        prev(next(Prec)) = P;
    }
    next(Prec) = P;
}

void insertLast(List &L, address P){
    if (isEmpty(L)){
        first(L) = P;
        last(L) = P;
    }else{
        prev(P) = last(L);
        next(last(L)) = P;
        last(L) = P;
    }
}

void deleteFirst(List &L, address &P) {
    if (!isEmpty(L)) {
        P = first(L);
        if (next(first(L)) != NULL) {
            first(L) = next(first(L));
            prev(first(L)) = NULL;
        }
        else {
            first(L) = NULL;
            last(L) = NULL;
        }
        next(P) = NULL;
    }
}

void deleteAfter(List &L, address Prec, address &P){
    P = next(Prec);
    next(Prec) = next(P);
    if (next(P) != NULL){
        prev(next(P)) = Prec;
    }else{
        last(L) = Prec;
    }
    prev(P) = NULL;
    next(P) = NULL;
}

void deleteLast(List &L, address &P) {
    if (!isEmpty(L)) {
        P = last(L);
        if (prev(P) != NULL) {
            last(L) = prev(P);
            next(last(L)) = NULL;
            prev(P) = NULL;
        }
        else {
            first(L) = NULL;
            last(L) = NULL;
        }
    }
}

void concat(List L1, List L2, List &L3){
    next(last(L1)) = first(L2);
    prev(first(L2)) = last(L1);
    first(L3) = first(L1);
    last(L3) = last(L2);
}


address findLagu(string Judul, List L) {
    address P = first(L);
    while (P != NULL) {
        if (info(P).song == Judul) {
            return P;
        }
        P = next(P);
    }
    return NULL;
}

void removeLagu(string Judul, List& L) {
    address P = findLagu(Judul, L);
    if (P == NULL) {
        cout << "Lagu tidak ditemukan." << endl;
    }
    else {
        if (P == first(L)) {
            deleteFirst(L, P);
            cout << "Lagu " << Judul << " telah dihapus menggunakan deleteFirst." << endl;
        }
        else if (P == last(L)) {
            deleteLast(L, P);
            cout << "Lagu " << Judul << " telah dihapus menggunakan deleteLast." << endl;
        }
        else {
            address Prec = prev(P);
            deleteAfter(L, Prec, P);
            cout << "Lagu " << Judul << " telah dihapus menggunakan deleteAfter." << endl;
        }
    }
}


void show(List L) {
    if (!isEmpty(L)) {
        address P = first(L);
        while (P != NULL) {
            cout << "Nama Band     : " << info(P).name << endl;
            cout << "Judul Lagu    : " << info(P).song << endl;
            cout << "-----------------------------" << endl;
            P = next(P);
        }
    }
    else {
        cout << "List kosong." << endl;
    }
}
