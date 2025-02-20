#ifndef TREE_H_INCLUDED
#define TREE_H_INCLUDED
#include <iostream>
using namespace std;
typedef int infotype;
typedef struct elm *adrNode;

struct elm{
    adrNode right;
    adrNode left;
    infotype info;
};

adrNode newNode_103032300152(infotype x);
void insertNode_103032300152(adrNode &root, adrNode p);
void DeleteNode_103032300152(adrNode &root, adrNode &p);
void printInOrder_103032300152(adrNode root);
adrNode findMin_103032300152(adrNode root);
void printLeftMostNode_103032300152(adrNode root);
void printRightMostNode_103032300152(adrNode root);

#endif // TREE_H_INCLUDED
