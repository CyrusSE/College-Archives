#include <iostream>
#include "queue.h"
using namespace std;

int main()
{
    // Membuat antrean
    Queue Q;
    createQueue(Q);
    // Menambahkan beberapa elemen ke dalam queue
    ElemQ* P1 = createElemQueue("John Doe", 65, "lansia", 1);
    ElemQ* P2 = createElemQueue("Alice", 30, "tenaga kesehatan", 2);
    ElemQ* P3 = createElemQueue("Bob", 25, "pekerja", 3);
    ElemQ* P4 = createElemQueue("Charlie", 70, "pensiunan", 4);
    ElemQ* P5 = createElemQueue("David", 28, "pekerja", 5);
    enqueue(Q, P1);
    enqueue(Q, P2);
    enqueue(Q, P3);
    enqueue(Q, P4);
    enqueue(Q, P5);

    // Menampilkan isi queue
    cout << "Isi antrean awal:" << endl; printInfo(Q);
    // Melayani antrean
    cout << "\nMelakukan pelayanan pada antrean:" << endl;
    serveQueue(Q);
    // Memeriksa antrean setelah pelayanan
    cout << "\nIsi antrean setelah pelayanan:" << endl;
    printInfo(Q);
    // Menambahkan elemen baru untuk simulasi pengaturan ulang prioritas
    ElemQ* P6 = createElemQueue("Edward", 22, "pekerja", 6);
    enqueue(Q, P6);
    // Simulasi reassignQueue untuk mengatur ulang antrean ke prioritas
    cout << "\nMengatur ulang antrean berdasarkan prioritas:" << endl;
    reassignQueue(Q);
    printInfo(Q);

    // Simulasi kondisi warga yang menunggu lebih dari 2 jam
    cout << "\nMemeriksa waktu tunggu dan mengubah prioritas jika lebih dari 2 jam:" << endl;
    checkWaitingTime(Q, 130); // Asumsikan waktu sekarang 130 menit dari nomor antrean pertama
    printInfo(Q);
    // Menangani kondisi darurat untuk seorang warga
    cout << "\nMenangani kondisi darurat untuk warga dengan nomor antrean 5:" << endl;
    emergencyHandle(Q, 5);
    printInfo(Q);

    // Update prioritas warga setiap jam
    cout << "\nMengupdate prioritas antrean setiap jam:" << endl;
    updatePriority(Q);
    printInfo(Q);


    // Mencari dan menghapus warga dengan nomor antrean tertentu
    cout << "\nMenghapus warga dengan nomor antrean 3:" << endl;
    ElemQ* removedElem = findAndRemove(Q, 3);

    if (removedElem) {
        cout << "Warga yang dihapus: " << info(removedElem).nama << endl;
    }
    printInfo(Q);

    // Mengecek ukuran queue
    cout << "\nUkuran antrean saat ini: " << size(Q) << endl;
    return 0;

}
