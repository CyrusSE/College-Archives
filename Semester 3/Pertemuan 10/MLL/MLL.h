#ifndef MLL_H_INCLUDED
#define MLL_H_INCLUDED
#include <iostream>
#define info(p) (p)->info
#define next(p) (p)->next
#define first(L) (L).first

using namespace std;

struct jadwalPenerbangan {
    string kode;
    string jenis;
    string tanggal;
    string waktu;
    string asal;
    string tujuan;
    int kapasitas;
};

typedef struct elemenJadwal *adr_jadwalP;
typedef jadwalPenerbangan infotype;

struct elemenJadwal {
    infotype info;
    adr_jadwalP next;
};

struct ListJadwal {
    adr_jadwalP first;
};

void createListJadwal_103032300152(ListJadwal &L);
adr_jadwalP createElemenJadwal_103032300152(infotype x);
void insertLastJ_103032300152(ListJadwal &L, adr_jadwalP P);
void deleteFirstJ_103032300152(ListJadwal &L, adr_jadwalP &P);
void ShowJadwal_103032300152(ListJadwal L);
adr_jadwalP searchJ_103032300152(ListJadwal L, string dari, string ke, string tanggal);

#endif // MLL_H_INCLUDED
