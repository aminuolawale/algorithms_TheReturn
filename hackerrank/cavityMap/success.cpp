#include <iostream>
#include <vector>

using namespace std;


// golang giving issues, had to use cpp

int main(){
    return 0;
}


vector<string> cavityMap(vector<string> grid) {
    for(unsigned int i=1; i < grid.size()-1;i++){
        for(unsigned int j=1; j< grid.size()-1; j++){
            int c = grid[i][j];
            int f = grid[i][j+1];
            int b = grid[i][j-1];
            int bt = grid[i+1][j];
            int tp = grid[i-1][j];
            if(c>f && c > b && c > tp && c > bt){
                grid[i][j] = 'X';
            }
        }
    }
    return grid;

}