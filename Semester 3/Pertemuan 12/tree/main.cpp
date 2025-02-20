#include <iostream>
#include "tree.h"

using namespace std;

int main() {
    adrNode root = NULL;

    int x[9] = {5, 3, 9, 10, 4, 7, 1, 8, 6};

    /* Tampilkan isi dari array */
    printf("============================================================");
    printf("\n");
    for (int i = 0; i < 9; i++) {
        cout << x[i] << " ";
    }

    /* 1. Tambahkan setiap elemen array x kedalam BST secara berurutan */
    /* sehingga dihasilkan BST seperti Gambar 1, gunakan looping */
    for (int i = 0; i < 9; i++) {
        insertNode_103032300152(root, newNode_103032300152(x[i]));
    }

    /* 2. Tampilkan node dari BST secara Pre-Order */
    printf("\n");
    printf("\nPre Order\t\t: ");
    printPreOrder_103032300152(root);

    /* 3. Tampilkan keturunan dari node 9*/
    printf("\n");
    printf("\nDescendent of Node 9\t: ");
    printDescendant_103032300152(root, 9);

    /* 4. Tampilkan total info semua node pada BST */
    printf("\n");
    printf("\nSum of BST Info\t\t: ");
    cout << sumNode_103032300152(root);

    /* 5. Tampilkan banyaknya daun dari BST */
    printf("\nNumber of Leaves\t: ");
    cout << countLeaves_103032300152(root);

    /* 6. Tampilkan Tinggi dari Tree */
    printf("\nHeight of Tree\t\t: ");
    cout  << heightTree_103032300152(root) - 1 << endl;
    cout << "============================================================" << endl;

    return 0;
}

