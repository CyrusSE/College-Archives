#ifndef TREE_H_INCLUDED
#define TREE_H_INCLUDED

typedef int infotype;
typedef struct Node *adrNode;
typedef adrNode BinTree;

struct Node {
    infotype info;
    adrNode left;
    adrNode right;
};


adrNode newNode_103032300152(infotype x);
/* mengembalikan alamat dari suatu node hasil alokasi, dengan info adalah x
dan left dan right adalah NULL */

adrNode findNode_103032300152(adrNode root, infotype x);
/* mengembalikan alamat dari node yang memiliki info sama dengan x, atau NULL
apabila tidak ditemukan, Catatan: implementasikan secara REKURSIF */

void insertNode_103032300152(adrNode &root, adrNode p);
/* I.S. terdefinisi root dari BST (mungkin NULL), dan pointer p yang berisi
alamat node yang mau ditambahkan pada BST
 F.S. elemen yang ditunjuk oleh p ditambahkan sebagai node dari BST
Catatan: implementasikan secara REKURSIF */

void printPreOrder_103032300152(adrNode root);
/* I.S. terdefinisi root dari BST (mungkin NULL)
 F.S. menampilkan node dari BST secara Preorder atau dengan urutan root, left
dan right
Catatan: implementasikan secara REKURSIF */

void printDescendant_103032300152(adrNode root, infotype x);
/* I.S. terdefinisi suatu root dari BST (mungkin NULL), dan infotype x
 F.S. menampilkan subtree atau keturunan dari node yang memiliki info x
Catatan: implementasikan secara REKURSIF */

int sumNode_103032300152(adrNode root);
/* mengembalikan hasil penjumlahan semua info node yang terdapat pada BST
Catatan: implementasikan secara REKURSIF */

int countLeaves_103032300152(adrNode root);
/* mengembalikan banyaknya daun atau node yang tidak memiliki anak dari BST
Catatan: implementasikan secara REKURSIF */

int heightTree_103032300152(adrNode root);
/* mengembalikan banyaknya edge dari suatu root menuju daun yang terjauh.
Catatan: implementasikan secara REKURSIF */

#endif // TREE_H_INCLUDED
