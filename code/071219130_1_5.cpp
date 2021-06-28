#include<iostream>
#include<stdio.h>
#include<queue>
#include<string.h>
#include<algorithm>
#include<string>
#include<stack>
using namespace std;
typedef long long ll;
int  n;
void bfs(ll x)
{
	queue<ll> Q; 
	Q.push(1) ;
	while(!Q.empty())
	{
		ll u = Q.front() ;
		Q.pop() ;
		if(u%n == 0)
		{
			cout<<u<<endl;
			return ;
		}
		Q.push(u*10);
		Q.push(u*10+1);
	}
	return ;
} 
int main()
{
	while(cin>>n)
	{
		if(n==0)
		{
			break;
		}
		bfs(1);
		 
	} 
	
}