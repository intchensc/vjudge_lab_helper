#include<iostream>
#include<cstring>
#include<cstdio>
using namespace std;
int l,r,c;
char map[35][35][35];
bool book[35][35][35];
int fx[6] = {-1,1,0,0,0,0};
int fy[6] = {0,0,-1,1,0,0};
int fz[6] = {0,0,0,0,1,-1};
typedef struct step{
	int x;      
	int y;
	int z; 
	int s;
} step;
step que[43000];
step s,e;
//血的教训，变量名千万不能粗心写错了，一个bug改一天；
//地图不要忘记置0 
//坐标不能搞错了，又改了几个小时
int main(){
	while(cin>>l>>r>>c&&l&&r&&c){
		memset(book,0,sizeof(book));
		for(int i = 1; i <= l; i++){
			for(int j = 1; j <= r; j++){
				for(int k = 1; k <= c; k++){
					cin>>map[j][k][i];
					if(map[j][k][i] == 'S'){
						s.x = j; s.y = k; s.z = i;    //这里的坐标气死我也
					}
					else if(map[j][k][i] == 'E'){
						e.x = j; e.y = k; e.z = i;
					}
					map[j][k][i] = map[j][k][i]==35;
				}
			}
		}	
//		
//		for(int i = 1; i <= l; i++){
//			for(int j = 1; j <= r; j++){
//				for(int k = 1; k <= c; k++){
//					cout<<(int)map[j][k][i]<<" ";
//				}
//				cout<<endl;
//			}
//		}
		int head = 1,tail = 1;
		que[tail].x = s.x;
		que[tail].y = s.y;
		que[tail].z = s.z;
		que[tail].s = 0;
		tail++;
		book[s.x][s.y][s.z] = 1;
		bool flag = 0;
		int tx,ty,tz;
		while(head < tail){
			for(int i = 0; i < 6; i++){
				tx = que[head].x + fx[i];
				ty = que[head].y + fy[i];
				tz = que[head].z + fz[i];
				if(tx < 1||tx > r||ty < 1||ty > c||tz < 1||tz >l)
					continue;
				if(map[tx][ty][tz] == 0 && book[tx][ty][tz] == 0){
					book[tx][ty][tz] = 1;        //这里的坐标气我+1
					que[tail].x = tx;
					que[tail].y = ty;
					que[tail].z = tz;
					que[tail].s = que[head].s + 1;
					tail++;
				}
				if(tx == e.x&&ty == e.y&&tz == e.z){
					cout<<"Escaped in "<<que[tail-1].s<<" minute(s)."<<endl;
					flag = 1;
					break;
				}
			}
			if(flag) break;
			head++;
		}
			if(!flag) 
				cout<<"Trapped!"<<endl;
		
	}
	return 0;
}