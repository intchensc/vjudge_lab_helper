#include <bits/stdc++.h>
using namespace std;
int main()
{
    int a,b,c;
    while(cin>>a>>b>>c&&(a&&b&&c))
    {
        a /= __gcd(b,c);
        if(a&1)
            cout<<"NO"<<endl;
        else
            cout<<a-1<<endl;
    }
    return 0;
}