#include <iostream>
#include <vector>
#include <cmath>
#include <climits>
using namespace std;

ostream& operator << (ostream& out, vector<int>& v){
    out << '[';
    for(unsigned int i=0; i< v.size(); i++){
        out << v[i];
        if(i != v.size() -1)
            out << ", ";
    }
    out << " ]" << endl;
    return out;
}


int minPositive(vector <int> v){
    int result = INT_MAX;
    for(auto a:v){
        if( a>0 && a < result)
            result = a;
    }
    return result==INT_MAX?-1:result;
}


int minimumDistances(vector<int> a) {
    vector <int> dValues;
    bool found = false;
    for(unsigned int i=0; i< a.size(); i++){
        for(unsigned int j=1; j<a.size(); j++){
            if(a[i] == a[j]){
                found = true;
                dValues.push_back(j-i);
            }
        }

    }
    return  found?minPositive(dValues):-1;
}

int main(){
    vector <int>  v = {1,2,3,4,10};
    cout << minimumDistances(v)<< endl;
    return 0;
}