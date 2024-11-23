#ifndef LIST_H_INCLUDED
#define LIST_H_INCLUDED
#include <iostream>
#define first(L) L.first
#define next(P) P->next
#define info(P) P->info

using namespace std;
typedef string infotype;
typedef struct elmlist *address;

struct elmlist{
    infotype info;
    address next;
};

struct List{
    address first;
};

void createList_103032300152(List &L);
address allocate_103032300152(infotype x);
void insertLast_103032300152(List &L, address P);
void show_103032300152(List L);
address shortestName_103032300152(List L);
void showFirstKNameC_103032300152(List L, int K, char c);

#endif // LIST_H_INCLUDED
