#include<iostream>
using namespace std;
const int MAXN = 15;
int a[MAXN][MAXN];
int vis[MAXN];
int n,k,ans;
void dfs(int step, int tot){
	if(tot==0){
		ans++;
		return;
	}
	//this if mast under ^ if
	if(step==n+1)
		return;
	for(int i = 1; i <= n; i++){
		if(a[step][i] == 1 && !vis[i]){
			vis[i] = 1;
			dfs(step+1,tot-1);
			vis[i] = 0;
		}
	}
	//line step no put
	dfs(step+1,tot);
}
int main(){
	while(cin >> n >> k){
		if(n == -1 || k == -1)
			break;
		ans = 0;
		for(int i = 1; i <= n; i++)
		for(int j = 1; j <= n; j++){
			char t;
			cin >> t;
			t=='#'?a[i][j]=1:a[i][j]=0;
		}
		dfs(1,k);
		cout << ans << endl;
	}
	return 0;
}