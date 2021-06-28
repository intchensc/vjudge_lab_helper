#include<bits/stdc++.h>
using namespace std;
const int MAXN = 205;
int n,m;
int a[MAXN][MAXN];
int vis[MAXN][MAXN];
int vis_t[MAXN][MAXN];
int ans[MAXN][MAXN];
int nxy[4][2] = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};
struct node{
	int x;
	int y;
	int step;
};
int isIn(int x, int y){
	if(x > n || x < 1)
		return 0;
	if(y > m || y < 1)
		return 0;
	return 1;
}
int bfs(node dot, int flag){
	for(int i = 1; i <= n; i++)
	for(int j = 1; j <= m; j++)
		vis_t[i][j] = vis[i][j];
	queue<node> que;
	dot.step = 0;
	que.push(dot);
	int step = 0;
	while(!que.empty()){
		node t = que.front();
//		cout << "t:" << t.x <<" "<<t.y<<endl;
		

		if(a[t.x][t.y] == '@'){
			if(flag == 1)
				ans[t.x][t.y] = 0;
			ans[t.x][t.y] += t.step;
//			ma[make_pair(t.x, t.y)] += t.step;
//			cout <<"string:"<<t.step<<endl;
//			ma[ans] += step;
		}
		int flag = 0;
		for(int i = 0; i < 4; i++){
			int tx = t.x + nxy[i][0];
			int ty = t.y + nxy[i][1];
			if(isIn(tx,ty)&& !vis_t[tx][ty]){
//				cout <<"tx:"<<tx<<" ty:"<<ty<<" "<<step<<endl;
				step++;
				vis_t[tx][ty] = 1;
				node needIn;
				needIn.x = tx;
				needIn.y = ty;
				needIn.step = t.step+1;
				que.push(needIn);
			}
		}
		que.pop();
	}
}
int main(){
	node yf, mj;
	int len = 1;
	while(cin >> n >> m){
		memset(vis, 0, sizeof(vis));
		for(int i = 1; i <= n; i++)
		for(int j = 1; j <= m; j++)
			ans[i][j] = 1e8;
		len = 1;
		for(int i = 1; i <= n; i++){
			for(int j = 1; j <= m; j++){
				char t;
				cin >> t;
				if(t == 'Y'){
					yf.x = i;
					yf.y = j;
					vis[i][j] = 1;
				}else if(t == 'M'){
					mj.x = i;
					mj.y = j;
					vis[i][j] = 1;
				}else if(t == '#'){
					vis[i][j] = 1;
				}
				a[i][j] = t;
			}	
		}
		a[yf.x][yf.y] = '#';
		a[mj.x][mj.y] = '#';
//		cout << endl;
//		for(int i = 1; i <= n; i++){
//			for(int j = 1; j <= m; j++){
//				cout << ans[i][j] << " ";
//			}
//			cout << endl;
//		}
		
		bfs(yf,1);
//		cout << endl;
//		for(int i = 1; i <= n; i++){
//			for(int j = 1; j <= m; j++){
//				cout << ans[i][j] << " ";
//			}
//			cout << endl;
//		}
		bfs(mj,0);
//		cout << endl;
//		for(int i = 1; i <= n; i++){
//			for(int j = 1; j <= m; j++){
//				cout << ans[i][j] << " ";
//			}
//			cout << endl;
//		}
		int minn = 1e8;
		for(int i = 1; i <= n; i++)
		for(int j = 1; j <= m; j++){
			minn = min(minn, ans[i][j]);
		}
		cout << minn*11 << endl;
	}
	return 0;
}