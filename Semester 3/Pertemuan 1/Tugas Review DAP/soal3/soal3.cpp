#include <iostream>

// Faisal Ihsan Santoso
// 103032300152
// IT-47-KHS
using namespace std;

int main(){
    float x, rata = 0, num = 0;
    cin >> x;
    while (x != -999){
        rata += x;
        num += 1;
        cin >> x;
    }
    if (num == 0){
        cout << 0;
    }else{
        cout << rata/num;
    }
    return 0;
}
