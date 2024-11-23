#ifndef DLL_H_INCLUDED
#define DLL_H_INCLUDED
#include <iostream>
#define first(L) L.first
#define last(L) L.last
#define next(P) P->next
#define info(P) P->info
#define prev(P) P->prev

using namespace std;


struct band {
    string name;
    string song;
};

typedef band infotype;
typedef struct elmtList* address;

struct elmtList {
    infotype info;
    address prev;
    address next;
};

struct List {
    address first;
    address last;
};

bool isEmpty(List L);
void createList(List &L);
address createNewElmt(infotype X);
void insertFirst(List &L, address P);
void insertAfter(List &L, address Prec, address P);
void insertLast(List &L, address P);
void deleteFirst(List &L, address &P);
void deleteAfter(List &L, address Prec, address &P);
void deleteLast(List &L, address &P);
void concat(List L1, List L2, List &L3);
address findLagu(string Judul, List L);
void removeLagu(string Judul, List &L);
void show(List L);

#endif // DLL_H_INCLUDED
