#include <iostream>
#include "MLL.h"

using namespace std;

int main()
{
    ListJadwal L;
    infotype data;
    int n;
    createListJadwal_103032300152(L);

    cout << "Berapa Jumlah data penerbangan yang anda ingin masukkan : ";
    cin >> n;
    for (int i = 1; i <= n; i++) {
        cout << "======= Data Ke " << i << " =======" <<endl;
        cout << "Kode       : ";
        cin >> data.kode;
        cout << "Jenis      : ";
        cin >> data.jenis;
        cout << "Tanggal    : ";
        cin >> data.tanggal;
        cout << "Waktu      : ";
        cin >> data.waktu;
        cout << "Asal       : ";
        cin >> data.asal;
        cout << "Tujuan     : ";
        cin >> data.tujuan;
        cout << "Kapasitas  : ";
        cin >> data.kapasitas;
        cout << "=========================" << endl;
        adr_jadwalP P = createElemenJadwal_103032300152(data);
        insertLastJ_103032300152(L, P);
    }
    cout << "Data Penerbangan :" << endl;
    ShowJadwal_103032300152(L);

    adr_jadwalP Q;
    deleteFirstJ_103032300152(L, Q);

    cout << "Data Penerbangan (setelah deleteFirst) : " << endl;
    ShowJadwal_103032300152(L);

    cout << "Pencarian Jadwal Tanggal 9 Desember 2022 : " << endl;
    string asal = "Surabaya";
    string tujuan = "Malang";
    string tanggal = "9-Desember-2022";
    adr_jadwalP cari = searchJ_103032300152(L, asal, tujuan, tanggal);
    if (cari == NULL){
        cout << "Data penerbangan tidak ditemukan." << endl;
    }else{
        cout << "Data penerbangan di temukan." << endl;
        cout << "Informasi Pernebangan : " << endl;
        cout << "________________________" << endl;
        cout << "Kode       : " << info(cari).kode << endl;
        cout << "Jenis      : " << info(cari).jenis << endl;
        cout << "Tanggal    : " << info(cari).tanggal << endl;
        cout << "Waktu      : " << info(cari).waktu << endl;
        cout << "Asal       : " << info(cari).asal << endl;
        cout << "Tujuan     : " << info(cari).tujuan << endl;
        cout << "Kapasitas  : " << info(cari).kapasitas << endl;
        cout << "________________________" << endl;
    }

    return 0;
}
