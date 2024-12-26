#include <iostream>
#include "queue.h"
using namespace std;


void createQueue(Queue &Q) {
    head(Q) = nill;
    tail(Q) = nill;
}

bool isEmpty(Queue Q) {
 return head(Q) == nill;
}

ElemQ* createElemQueue(string nama, int usia, string pekerjaan, int nomor_antrean) {
    ElemQ *P = new ElemQ;
    info(P).nama = nama;
    info(P).usia = usia;
    info(P).pekerjaan = pekerjaan;
    info(P).prioritas = (usia >= 60 || pekerjaan == "tenaga kesehatan");
    info(P).nomor_antrean = nomor_antrean;
    info(P).kondisi_darurat = false;
    next(P) = nill;
    return P;
}

void enqueue(Queue &Q, ElemQ *P) {
  if (isEmpty(Q)) {
    head(Q) = P;
    tail(Q) = P;
  } else if (info(P).prioritas) {
    if (!info(head(Q)).prioritas) {
      next(P) = head(Q);
      head(Q) = P;
    } else {
      ElemQ *X = head(Q);
      while (next(X) != nill && info(next(X)).prioritas) {
        X = next(X);
      }
      next(P) = next(X);
      next(X) = P;
      if (next(P) == nill) {
        tail(Q) = P;
      }
    }
  } else {
    next(tail(Q)) = P;
    tail(Q) = P;
  }
}

void dequeue(Queue & Q, ElemQ * & P) {
  if (isEmpty(Q)) {
    P = nill;
    cout << "Semua warga telah terlayani." <<endl;
  } else {
    P = head(Q);
    head(Q) = next(head(Q));
    if (head(Q) == nill) {
      tail(Q) = nill;
    }
    next(P) = nill;
  }
}

ElemQ* front(Queue Q) {
    return head(Q);
}

ElemQ* back(Queue Q) {
    return tail(Q);
}

int size(Queue Q) {
    int count = 0;
    ElemQ *temp = head(Q);
    while (temp != nill) {
        count++;
        temp = next(temp);
    }
    return count;
}

void printInfo(Queue Q) {
    if (isEmpty(Q)) {
        cout << "Antrean kosong." << endl;
        return;
    }

    cout << "Daftar Antrean:" << endl;
    ElemQ *P = head(Q);
    while (P != nill) {
        cout << "Nama: " << info(P).nama << endl;
        cout << "Usia: " << info(P).usia << endl;
        cout << "Pekerjaan: " << info(P).pekerjaan << endl;
        cout << "Prioritas: " << (info(P).prioritas ? "Ya" : "Tidak") << endl;
        cout << "Nomor Antrean: " << info(P).nomor_antrean << endl;
        cout << "------------------------" << endl;
        P = next(P);
    }
}

void serveQueue(Queue &Q) {
    if (isEmpty(Q)) {
        cout << "Antrean kosong." << endl;
        return;
    }

    int terlayani = 0;
    const int MAX = 100;

    while (!isEmpty(Q) && terlayani < MAX) {
        ElemQ *P = head(Q);
        dequeue(Q, P);
        if (P != nill) {
            cout << "Melayani warga:" << endl;
            cout << "Nama       : " << info(P).nama << endl;
            cout << "Usia       : " << info(P).usia << endl;
            cout << "Pekerjaan  : " << info(P).pekerjaan << endl;
            cout << "Prioritas  : " << (info(P).prioritas ? "Ya" : "Tidak") << endl;
            cout << "Vaksinasi berhasil." << endl;
            cout << "------------------------" << endl;
            terlayani++;
        }
    }
    if (terlayani < MAX) {
        cout << "Kapasitas telah penuh." << endl;
        if (!isEmpty(Q)){
            cout << "Warga yang belum terlayani diminta kembali besok." << endl;
        }
        return;
    }
}

void reassignQueue(Queue &Q) {
    if (isEmpty(Q)) {
        cout << "Antrean kosong." << endl;
        return;
    }

    Queue priorityQ, normalQ;
    createQueue(priorityQ);
    createQueue(normalQ);

    while (!isEmpty(Q)) {
        ElemQ *P;
        dequeue(Q, P);
        if (P != nill) {
            if (info(P).prioritas || info(P).kondisi_darurat) {
                enqueue(priorityQ, P);
            } else {
                enqueue(normalQ, P);
            }
        }
    }

    while (!isEmpty(priorityQ)) {
        ElemQ *P;
        dequeue(priorityQ, P);
        if (P != nill) {
            enqueue(Q, P);
        }
    }

    while (!isEmpty(normalQ)) {
        ElemQ *P;
        dequeue(normalQ, P);
        if (P != nill) {
            enqueue(Q, P);
        }
    }
}

void checkWaitingTime(Queue &Q, int waktu_sekarang) {
    if (isEmpty(Q)) {
        cout << "Antrean kosong." << endl;
        return;
    }

    Queue tempQ;
    createQueue(tempQ);

    while (!isEmpty(Q)) {
        ElemQ *P;
        dequeue(Q, P);
        if (P != nill) {
            int waiting_time = waktu_sekarang;
            if (waiting_time > 120) {
                info(P).prioritas = true;
            }
            enqueue(tempQ, P);
        }
    }
    Q = tempQ;
}

void emergencyHandle(Queue &Q, int nomor_antrean) {
    if (isEmpty(Q)) {
        cout << "Antrean kosong." << endl;
        return;
    }

    bool found = false;
    Queue tempQ;
    createQueue(tempQ);

    while (!isEmpty(Q)) {
        ElemQ *P;
        dequeue(Q, P);
        if (P != nill) {
            if (info(P).nomor_antrean == nomor_antrean) {
                info(P).kondisi_darurat = true;
                info(P).prioritas = true;
                found = true;
            }
            enqueue(tempQ, P);
        }
    }

    Q = tempQ;
    if (!found) {
        cout << "Warga dengan nomor antrean " << nomor_antrean << " tidak ditemukan." << endl;
    }
}

void updatePriority(Queue &Q) {
    if (isEmpty(Q)) {
        cout << "Antrean kosong." << endl;
        return;
    }

    Queue emergencyQ, priorityQ, normalQ;
    createQueue(emergencyQ);
    createQueue(priorityQ);
    createQueue(normalQ);

    int Time = 130;

    while (!isEmpty(Q)) {
        ElemQ *P;
        dequeue(Q, P);
        if (P != nill) {
            if (info(P).kondisi_darurat) {
                enqueue(emergencyQ, P);
            } else if (info(P).prioritas || Time > 120) {
                enqueue(priorityQ, P);
            } else {
                enqueue(normalQ, P);
            }
        }
    }

    while (!isEmpty(emergencyQ)) {
        ElemQ *P;
        dequeue(emergencyQ, P);
        if (P != nill) {
            enqueue(Q, P);
        }
    }

    while (!isEmpty(priorityQ)) {
        ElemQ *P;
        dequeue(priorityQ, P);
        if (P != nill) {
            enqueue(Q, P);
        }
    }

    while (!isEmpty(normalQ)) {
        ElemQ *P;
        dequeue(normalQ, P);
        if (P != nill) {
            enqueue(Q, P);
        }
    }
}


ElemQ* findAndRemove(Queue &Q, int nomor_antrean) {
    if (isEmpty(Q)) {
        cout << "Antrean kosong." << endl;
        return nill;
    }

    Queue tempQ;
    createQueue(tempQ);
    ElemQ *found = nill;

    while (!isEmpty(Q)) {
        ElemQ *P;
        dequeue(Q, P);
        if (P != nill) {
            if (info(P).nomor_antrean == nomor_antrean) {
                found = P;
            } else {
                enqueue(tempQ, P);
            }
        }
    }

    Q = tempQ;
    if (found == nill) {
        cout << "Warga dengan nomor antrean " << nomor_antrean << " tidak ditemukan dalam antrean." << endl;
    }
    return found;
}
