#include<iostream>
#include<cstring>
#include<cstdio>
using namespace std;
bool map[101][101];
bool book[101][101];
int nextt[8][2] = {{0,1},{1,1},{1,0},{1,-1},{0,-1},{-1,-1},{-1,0},{-1,1}};
int m,n,count;
void dfs(int x, int y){
        book[x][y] = 1;
	for(int i = 0; i < 8; i++){
		int tx = x + nextt[i][0];
		int ty = y + nextt[i][1];
		if(tx<1 || tx>m || ty<1 || ty > n)
			continue;
		if(book[tx][ty] == 1 || map[tx][ty] == 1)
			continue;
			book[tx][ty] = 1;
		dfs(tx,ty);
	}
	return;	
}
int main(){
	while(scanf("%d %d",&m,&n)!=EOF&&m){
		memset(book,0,sizeof(book));
		char c;getchar();count = 0;
		for(int i = 1; i <= m; i++){
			for(int j = 1; j <= n; j++){
				c = getchar();
				map[i][j] = c!=64;
			}
		 	getchar();
		}
		
		
		for(int i = 1; i <= m; i++){
			for(int j = 1; j <= n; j++){
				if(map[i][j]==0&&!book[i][j]){
					count++;dfs(i,j);
				}		
			}	
		}	
		printf("%d\r
",count);	
	}
	return 0;
}