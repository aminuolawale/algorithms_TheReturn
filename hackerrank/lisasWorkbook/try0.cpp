#include <iostream>
#include <vector>

using namespace std;

int workbook(int n, int k, vector<int> arr) {
    int count = 1;
    int page = 0;
    int chapter = 1;
    int nProblems = arr[chapter];
    int pages = nProblems/k;
    pages = nProblems % k > 0? pages+1:pages;
    page += pages;


}

int main(){
    return 0;
}