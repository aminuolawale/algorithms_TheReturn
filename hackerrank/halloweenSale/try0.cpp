#include <iostream>
#include <cmath>
using  namespace std;


int howManyGames(int p, int d, int m, int s) {
    // Return the number of games you can buy
    if(p >s)
        return 0;
    d = -d;
    int termCount = (m - p)/d +1;
    int theSum = 0.5 * termCount * (2 * p + (termCount-1) *d);
    if(theSum > s){
        int count = 0;
        int sum = 0;
        int k =s;
        while(sum < s){
            sum += k;
            k -=d;
            count ++;
        }
        return count;
    }
    int remainder = s - theSum;
    int rCount = remainder / m;
    return termCount + rCount;
}


int main(){
    cout << howManyGames(20, 3, 6, 30) << endl;

    return 0;
}