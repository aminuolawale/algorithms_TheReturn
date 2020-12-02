#include <iostream>

using namespace std;

int chocolateFeast(int n, int c, int m)
{
    // n = amount he has, c = cost of one chocolate, m= number of wrappers he can turn in for one chocolate
    int chocCount = 0;
    int initialCount = n / c;
    chocCount += initialCount;
    int remainder = n % c;
    int left = initialCount;
    if (initialCount < m)
        return chocCount;
    while (true)
    {
        int newChoc = left / m;
        chocCount += newChoc;
        int wrapperRemainder = left % m;
        left = newChoc + wrapperRemainder;
        if (left < m)
            return chocCount;
    }
    return chocCount;
}

int main()
{
    cout << chocolateFeast(6, 2, 2) << endl;
    return 0;
}