#include "Tree.h"

adrNode newNode_103032300152(infotype x){
    adrNode P = new elm;
    P->left = NULL;
    P->right = NULL;
    P->info = x;
    return P;
}

void insertNode_103032300152(adrNode &root, adrNode p){
    if (root == NULL){
        root = p;
    }else{
        if (p->info > root->info){
            insertNode_103032300152(root->right, p);
        }else{
            insertNode_103032300152(root->left, p);
        }
    }
}

void DeleteNode_103032300152(adrNode &root, adrNode &p){
    adrNode temp;
    if (root == NULL){
        cout << "Node tidak di temukan" << endl;
    }else if (p->info < root->info){
        DeleteNode_103032300152(root->left, p);
    }else if (p->info > root->info){
        DeleteNode_103032300152(root->right, p);
    }else{
        if (root->left == NULL && root->right == NULL){
            delete root;
            root = NULL;
        }else if (root->left == NULL){
            temp = root;
            root = root->right;
            delete temp;
        }else if (root->right == NULL){
            temp = root;
            root = root->left;
            delete temp;
        }else{
            temp = findMin_103032300152(root->right);
            root->info = temp->info;
            DeleteNode_103032300152(root->right, temp);
        }
    }
}

void printInOrder_103032300152(adrNode root){
    if (root != NULL){
        printInOrder_103032300152(root->left);
        cout << root->info << " ";
        printInOrder_103032300152(root->right);
    }
}

adrNode findMin_103032300152(adrNode root){
    if (root->left == NULL){
        return root;
    }else{
        return findMin_103032300152(root->left);
    }
}


void printLeftMostNode_103032300152(adrNode root) {
    if (root == NULL) {
        return;
    }
    if (root->left == NULL) {
        cout << root->info;
    }else {
        printLeftMostNode_103032300152(root->left);
    }
}

void printRightMostNode_103032300152(adrNode root) {
    if (root == NULL) {
        return;
    }
    if (root->right == NULL) {
        cout << root->info;
    }else {
        printRightMostNode_103032300152(root->right);
    }
}
