#include "list.h"


int main()
{
    List L;
    address p = NULL;
    address q = NULL;
    akun data;
    createList(L);
    for (int i=0;i<3;i++){
        cout << "Masukkan akun ke-" << i+1 << ": ";
        cin >> data.username;
        cin >> data.password;
        cout << endl;
        signUp(L, data);

    }

    deleteFirst(L, p);
    cout << "Data setelah delete First: "<< endl;
    show(L);
    cout << "Masukkan username yang ingin dihapus: ";
    cin >> data.username;
    removeAkun(L, data.username);
    cout << endl << "Data setelah username tersebut dihapus: "<< endl;
    show(L);
    cout << "Data setelah delete Last: "<< endl;
    deleteLast(L,p);
    show(L);
}
