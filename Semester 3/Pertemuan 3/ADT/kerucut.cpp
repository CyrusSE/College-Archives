#include <iostream>

using namespace std;

float volume(float r, float t){

    return (1.0 / 3.0) * 3.14 * r * r *t;
}

float LP(float r, float s){

    return 3.14 * r * (r + s);
}
