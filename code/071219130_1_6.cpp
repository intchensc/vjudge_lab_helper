#include <iostream>
#include <cstdio>
#include <cmath>
#include <cstring>
#include <algorithm>
#include <string>
#include <map>
#include <set>
#include <vector>
#include <queue>
#include <stack>
#define inf 100000000

using namespace std;

typedef long long ll;

int n, z, t;    // n表示输入的第一个数，z表示输入的第二个数，t表示案例数；

bool prime[10000];    //存素数表；

int p[10000];     //记录变化次数，bfs过程用；

int bfs()
{
    queue<int> q;
    q.push(n);
    p[n]=0;       //标记n;
    while(!q.empty())
    {
        int pa=q.front();
        q.pop();
        if(pa==z)  //判断是否已经搜索到 z 了；
            return p[z];
        int d[4];
        d[0]=(pa/10)*10;   //pa去个位数的值；
        d[1]=pa%10+(pa/100)*100;  //pa去十位数的值；
        d[2]=(pa/1000)*1000+pa%100;  //pa去百位数的值；
        d[3]=pa%1000;   //pa去千位数的值；
        int d2=1;
        for(int i=0;i<4;i++)
        {
            int j=0;
            if(i==3) //千位不为0处理；
                j=1;
            for(;j<10;j++)
            {
                int w=d[i]+j*d2;
                if(w>=1000&&w<10000&&prime[w]==0&&p[w]==-1)
                {
                    p[w]=p[pa]+1;
                    if(w==z)
                        return p[z];
                    q.push(w);
                }
            }
            d2=d2*10;
        }

    }
    return -1;
}

int main()
{
    prime[0]=1;
    prime[1]=1;
    for(int i=2;i*i<10000;i++)    //打素数表，表中0表示素数；
    {
        if(!prime[i])
        {
            for(int j=i*2;j<10000;j=j+i)
            {
                prime[j]=1;
            }
        }
    }
    scanf("%d",&t);
    for(int o=0;o<t;o++)
    {
        memset(p,-1,sizeof(p));   //对p初始化；
        int za=0;   //表示变化次数，输出用；
        scanf("%d %d",&n,&z);
        za=bfs();
        if(za==-1)  //没搜索到；
        {
            printf("Impossible\r
");
        }
        else   //搜索到了；
        {
            printf("%d\r
",za);
        }
    }
    return 0;
}
