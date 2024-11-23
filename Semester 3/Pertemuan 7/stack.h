#ifndef STACK_H_INCLUDED
#define STACK_H_INCLUDED
#include <iostream>
#define top(S) (S).top
#define info(S) (S).info

using namespace std;
typedef char infotype;

struct stack{
    infotype info[15];
    int top;
} ;

void createStack_103032300152(stack &S);
bool isEmpty_103032300152(stack S);
bool isFull_103032300152(stack S);
void push_103032300152(stack &S, infotype x);
infotype pop_103032300152(stack &S);
void printInfo_103032300152(stack S);


#endif // STACK_H_INCLUDED
