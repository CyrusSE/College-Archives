#include <iostream>
#include "tree.h"

using namespace std;

adrNode newNode_103032300152(infotype x) {
    adrNode N = new Node;
    N->info = x;
    N->left = NULL;
    N->right = NULL;
    return N;
}

adrNode findNode_103032300152(adrNode root, infotype x) {
    if (root == NULL){
        return NULL;
    }
    if (root->info == x){
        return root;
    }
    adrNode p = findNode_103032300152(root->left, x);
    if (p != NULL){
        return p;
    }
    return findNode_103032300152(root->right, x);
}

void insertNode_103032300152(adrNode &root, adrNode p) {
    if (root == NULL) {
        root = p;
    } else{
        if (p->info < root->info) {
            insertNode_103032300152(root->left, p);
        } else {
            insertNode_103032300152(root->right, p);
        }
    }
}

void printPreOrder_103032300152(adrNode root) {
    if (root == NULL){
        return;
    }
    cout << root->info << " ";
    printPreOrder_103032300152(root->left);
    printPreOrder_103032300152(root->right);
}

void printDescendant_103032300152(adrNode root, infotype x) {
    adrNode p = findNode_103032300152(root, x);
    if (p == NULL){
        cout << "Tidak ada Node." << endl;
    }
    printPreOrder_103032300152(p->left);
    printPreOrder_103032300152(p->right);
}

int sumNode_103032300152(adrNode root) {
    if (root == NULL) {
        return 0;
    }
    return root->info + sumNode_103032300152(root->left) + sumNode_103032300152(root->right);
}

int countLeaves_103032300152(adrNode root) {
    if (root == NULL){
        return 0;
    }
    if (root->left == NULL && root->right == NULL){
        return 1;
    }
    return countLeaves_103032300152(root->left) + countLeaves_103032300152(root->right);
}

int heightTree_103032300152(adrNode root) {
    if (root == NULL) {
        return 0;
    }
    int left_height = heightTree_103032300152(root->left);
    int right_height = heightTree_103032300152(root->right);
    return max(left_height, right_height) + 1;
}
