#ifndef LIST_H_INCLUDED
#define LIST_H_INCLUDED
#include <iostream>
using namespace std;
struct akun{
    string username;
    string password;
};
typedef akun infotype;
typedef struct elmlist *address;
struct elmlist{
    address next;
    address prev;
    infotype info;
};
struct List{
    address first;
    address last;
};
void createList(List &L);
address createNewElm(infotype akun);
void InsertLast(List &L, address p);
address findAkun(List L, string username);
void signUp(List &L, infotype akun);
void deleteFirst(List &L, address &p);
void deleteAfter(address q, address &p);
void deleteLast(List &L, address &p);
void removeAkun(List &L, string username);
void show(List L);
#endif // LIST_H_INCLUDED
