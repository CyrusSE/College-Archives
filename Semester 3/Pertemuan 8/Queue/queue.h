#ifndef QUEUE_H_INCLUDED
#define QUEUE_H_INCLUDED
using namespace std;

#define nill NULL
#define head(Q) (Q).head
#define tail(Q) (Q).tail
#define next(P) P->next
#define info(P) P->info

struct Infotype {
    string nama;
    int usia;
    string pekerjaan;
    bool prioritas;
    int nomor_antrean;
    bool kondisi_darurat;
};

struct ElemQ {
    Infotype info;
    ElemQ* next;
};

struct Queue {
    ElemQ* head;
    ElemQ* tail;
};

void createQueue(Queue &Q);
bool isEmpty(Queue Q);
ElemQ* createElemQueue(string nama, int usia, string pekerjaan, int nomor_antrean);
void enqueue(Queue &Q, ElemQ *P);
void dequeue(Queue &Q, ElemQ &P);
ElemQ* front(Queue Q);
ElemQ* back(Queue Q);
int size(Queue Q);

void printInfo(Queue Q);
void serveQueue(Queue &Q);
void reassignQueue(Queue &Q);
void checkWaitingTime(Queue &Q, int waktu_sekarang);
void emergencyHandle(Queue &Q, int nomor_antrean);
void updatePriority(Queue &Q);
ElemQ* findAndRemove(Queue &Q, int nomor_antrean);

#endif // QUEUE_H_INCLUDED
