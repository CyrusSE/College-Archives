#include "list.h"

void createList(List &L){
    L.first = NULL;
    L.last = NULL;
}

address createNewElm(infotype akun){
    address p = new elmlist;
    p->next = NULL;
    p->prev = NULL;
    p->info = akun;
    return p;
}

void InsertLast(List &L, address p){
    if (L.first == NULL && L.last == NULL){
        L.first = p;
        L.last = p;
    } else {
        L.last->next = p;
        p->prev = L.last;
        L.last = p;
    }
}

address findAkun(List L, string username){
    if (L.first == NULL && L.last == NULL){
        return NULL;
    }else{
        address p = L.first;

        while (p!=NULL){
            if (p->info.username == username){
                return p;
            }
            p = p->next;
        }

        return NULL;
    }

}
void signUp(List &L, infotype akun){
    address p = findAkun(L,akun.username);

    if (p!=NULL){
        cout << "Account has been registered." << endl;
    }else{
        p = createNewElm(akun);
        InsertLast(L,p);
    }


}

void deleteFirst(List &L, address &p){
    p = L.first;
    L.first = p->next;
    p->next = NULL;
    L.first->prev = NULL;
}

void deleteAfter(address q, address &p){
    p = q->next;
    q->next = p->next;

    p->next->prev=q;

    p->next=NULL;
    p->prev=NULL;
}

void deleteLast(List &L, address &p){
    if (L.first->next == NULL){
        p = L.first;
        L.first = NULL;
        L.last = NULL;
    }else{
        p = L.last;
        L.last = p->prev;
        L.last->next = NULL;
        p->prev = NULL;
    }

}

void removeAkun(List &L, string username){
    if (L.first == NULL && L.last == NULL){
        cout << "List is empty.";
    }else{
        address q = findAkun(L, username);
        if (q == NULL){
            cout << "Account has not found." << endl;
        }else{
            if (q == L.first) {
                deleteFirst(L, q);
            } else if (q == L.last) {
                deleteLast(L, q);
            } else {
                deleteAfter(q->prev, q);
            }
            delete q;
        }
    }
}

void show(List L){
    if (L.first == NULL || L.last == NULL){
        cout << "List is empty.";
    }else{
        address p = L.first;
        while (p!=NULL){
            cout << "Username: " << p->info.username << endl << "Password: " << p->info.password << endl << endl;
            p = p->next;
        }
    }
}
