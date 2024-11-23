#include <iostream>
#include "list.h"
#include "menu.h"

using namespace std;

void menu_103032300152(List L){
    int pilihan;
    string pil;
    cout << "======MENU=======" << endl;
    cout << "1. Menambah N data baru" << endl;
    cout << "2. Menampilkan semua data" << endl;
    cout << "3. Menampilkan nama terpendek" << endl;
    cout << "4. Menampilkan nama yang berawalan dari huruf X" << endl;
    cout << "0. Exit" << endl;
    cout << "Masukkan menu : ";
    cin >> pilihan;
    cout << endl;
    if (pilihan == 1){
        Pilihan1_103032300152(L);
    }else if (pilihan == 2){
        show_103032300152(L);
        cout << "Kembali ke menu utama? (Y/N) : ";
        cin >> pil;
        if (pil == "Y"){
            menu_103032300152(L);
        }else{
            exit(0);
        }
    }else if (pilihan == 3){
        address terpendek;
        terpendek = shortestName_103032300152(L);
        cout << "Nama Terpendek : " << info(terpendek) << endl;
        cout << "Kembali ke menu utama? (Y/N) : ";
        cin >> pil;
        if (pil == "Y"){
            menu_103032300152(L);
        }else{
            exit(0);
        }
    }else if (pilihan == 4){
        int K;
        char c;
        cout << "Berapa data yang anda ingin liat : ";
        cin >> K;
        cout << "Huruf pertama apa yang anda ingin liat : ";
        cin >> c;
        showFirstKNameC_103032300152(L, K, c);
        cout << "Kembali ke menu utama? (Y/N) : ";
        cin >> pil;
        if (pil == "Y"){
            menu_103032300152(L);
        }else{
            exit(0);
        }
    }else {
        cout << "ANDA TELAH KELUAR DARI PROGRAM" << endl;
        exit(0);
    }
}

void Pilihan1_103032300152(List L){
    int N = 0;
    string pilihan;
    infotype data;
    cout << "Jumlah data yang akan ditambahkan : ";
    cin >> N;
    for (int i = 1; i <= N ; i++){
        cout << "Masukkan data baru : ";
        cin >> data;
        insertLast_103032300152(L, allocate_103032300152(data));
    }
    cout << "Kembali ke menu utama? (Y/N) : ";
    cin >> pilihan;
    if (pilihan == "Y"){
        menu_103032300152(L);
    }else{
        exit(0);
    }
}

