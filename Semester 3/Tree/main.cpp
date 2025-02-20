#include <iostream>
#include "Tree.h"

using namespace std;

int main()
{
    int in[9] = {8,6,15,4,7,12,17,9,13};
    int del[9] = {8,9,12,13,15,17,7,6,4};
    adrNode root = NULL;
    for (int i = 0; i < 9; i++){
        insertNode_103032300152(root, newNode_103032300152(in[i]));
    }
    for (int i = 0; i <= 9; i++){
        if (root != NULL){
            printInOrder_103032300152(root);
            cout << endl;
            adrNode P = newNode_103032300152(del[i]);
            DeleteNode_103032300152(root, P);
        }else{
            cout << "(kosong)" << endl;
        }

    }
}
