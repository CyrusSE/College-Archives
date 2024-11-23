#include <iostream>
#include "stack.h"

using namespace std;

void createStack_103032300152(stack &S){
    top(S) = 0;
}

bool isEmpty_103032300152(stack S){
    if (top(S) == 0){
        return true;
    }else{
        return false;
    }
}

bool isFull_103032300152(stack S){
    if top(S) == 15{
        return true;
    }else{
        return false;
    }
}

void push_103032300152(stack &S, infotype x){
    if (isFull_103032300152(S) == false){
        top(S) = top(S) + 1;
        info(S)[top(S)] = x;
    }
}

infotype pop_103032300152(stack &S){
    infotype x = info(S)[top(S)];
    top(S) = top(S) - 1;
    return x;
}

void printInfo_103032300152(stack S){
    for (int i = top(S); i > 0; i--){
        cout << info(S)[i];
    }
    cout << endl;
}
