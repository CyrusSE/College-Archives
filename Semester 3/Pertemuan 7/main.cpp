#include <iostream>
#include "stack.h"
using namespace std;

int main()
{
    stack S;
    createStack_103032300152(S);

    push_103032300152(S, 'I');
    push_103032300152(S, 'R');
    push_103032300152(S, 'I');
    push_103032300152(S, 'D');
    push_103032300152(S, 'A');
    push_103032300152(S, 'Y');
    push_103032300152(S, 'A');
    push_103032300152(S, 'C');
    push_103032300152(S, 'R');
    push_103032300152(S, 'E');
    push_103032300152(S, 'P');



    cout << "Data Sebelum di POP :" << endl;
    printInfo_103032300152(S);

    pop_103032300152(S);
    pop_103032300152(S);
    pop_103032300152(S);
    pop_103032300152(S);
    pop_103032300152(S);
    pop_103032300152(S);
    pop_103032300152(S);

    cout << "Data Setelah di POP :" << endl;
    printInfo_103032300152(S);
}
