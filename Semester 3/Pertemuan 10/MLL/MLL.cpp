#include "MLL.h"
void createListJadwal_103032300152(ListJadwal &L) {
    first(L) = NULL;
}

adr_jadwalP createElemenJadwal_103032300152(infotype x) {
    adr_jadwalP p = new elemenJadwal;
    info(p) = x;
    next(p) = NULL;
    return p;
}

void insertLastJ_103032300152(ListJadwal &L, adr_jadwalP P) {
    if (first(L) == NULL) {
        first(L) = P;
    } else {
        adr_jadwalP q;
        q = first(L);
        while (next(q) != NULL) {
            q = next(q);
        }
        next(q) = P;
    }
}

void ShowJadwal_103032300152(ListJadwal L) {
    adr_jadwalP p;
    p = first(L);
    while (p != NULL) {
        cout << "________________________" << endl;
        cout << "Kode       : " << info(p).kode << endl;
        cout << "Jenis      : " << info(p).jenis << endl;
        cout << "Tanggal    : " << info(p).tanggal << endl;
        cout << "Waktu      : " << info(p).waktu << endl;
        cout << "Asal       : " << info(p).asal << endl;
        cout << "Tujuan     : " << info(p).tujuan << endl;
        cout << "Kapasitas  : " << info(p).kapasitas << endl;
        cout << "________________________" << endl;
        p = next(p);
    }
}

void deleteFirstJ_103032300152(ListJadwal &L, adr_jadwalP &P) {
    if (first(L) == NULL){
        P = NULL;
    }else if (next(first(L)) == NULL){
        P = first(L);
        first(L) = NULL;
    }else{
        P = first(L);
        first(L) = next(P);
        next(P) = NULL;
    }
}

adr_jadwalP searchJ_103032300152(ListJadwal L, string dari, string ke, string tanggal) {
    adr_jadwalP p;
    p = first(L);
    while (p != NULL) {
        if (info(p).asal == dari && info(p).tujuan == ke && info(p).tanggal == tanggal) {
            return p;
        } else {
            p = next(p);
        }
    }
    return p;
}
