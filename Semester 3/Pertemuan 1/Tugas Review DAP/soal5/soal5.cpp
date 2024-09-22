#include <iostream>

// Faisal Ihsan Santoso
// 103032300152
// IT-47-KHS
using namespace std;

const int NMAX = 51;

struct Mahasiswa{
    string NIM;
    string NAMA;
    int nilai;
};

Mahasiswa arrMahasiswa[NMAX];

int nilaiPertama(Mahasiswa T[], int N, string NIM){
    for (int i = 0; i < N; i++){
        if (T[i].NIM == NIM){
            return T[i].nilai;
        }
    }
    return -1;
}

void isiData(Mahasiswa *T, int *N){
    string NIM;
    string NAMA;
    int nilai;
    cin >> NIM;
    cin >> NAMA;
    cin >> nilai;
    while (NIM != "0" && NAMA != "0" && nilai != 0){
        T[*N].NIM = NIM;
        T[*N].NAMA = NAMA;
        T[*N].nilai = nilai;
        *N += 1;
        cin >> NIM;
        cin >> NAMA;
        cin >> nilai;
    }
}

void tampilData(Mahasiswa T[], int N){
    for (int i = 0 ; i < N ; i++){
        cout << T[i].NIM << endl;
        cout << T[i].NAMA << endl;
        cout << T[i].nilai << endl;
        cout << endl;
    }
}

int main(){
    int total = 0;
    string cariNIM;
    int x;

    isiData(arrMahasiswa, &total);
    tampilData(arrMahasiswa, total);

    cout << "Masukkan NIM yang ingin anda cari : ";
    cin >> cariNIM;

    x = nilaiPertama(arrMahasiswa, total, cariNIM);

    if (x == -1){
        cout << "Data tidak ditemukan." << endl;
    }else{
        cout << "Nilai Pertamanya adalah " << x << endl;
    }

    return 0;
}
