#include<iostream>
#include<cstring>
#include<cstdio>
using namespace std;
bool map[6][6];
bool book[6][6];
int nextt[4][2] = {{0,1},{1,0},{0,-1},{-1,0}};
typedef struct step{
	int x;      
	int y;
	int f;   
} step;
step que[36];
void pp(int now){
	//printf("now:%d\r
",now);
	if(now==1)
		return;
	now = que[now].f;
	pp(now);	
	printf("(%d, %d)\r
",que[now].x-1,que[now].y-1); 
}
int main()
{
	int head = 1,tail = 1;
	for(int i = 1; i <= 5; i++)
		for(int j = 1; j <= 5; j++)
			scanf("%d",&map[i][j]);
	que[tail].x = 1;
	que[tail].y = 1;
	que[tail].f = 0;
	tail++;
	book[1][1] = 1;
	bool flag = 0;
	int tx,ty;
	while(head < tail){
		for(int i = 0; i < 3; i++){
			tx = que[head].x + nextt[i][0];
			ty = que[head].y + nextt[i][1];
			if(tx < 1||tx > 5||ty < 1||ty > 5)
				continue;
			if(map[tx][ty] == 0 && book[tx][ty] == 0){
				book[tx][ty] = 1;
				que[tail].x = tx;
				que[tail].y = ty;
				que[tail].f = head;
				tail++;
			}
			if(tx == 5&&ty == 5){
				flag = 1;
				break;
			}
		}
		if(flag) 
			break;
		head++;
	}
	
//	for(int i = 1; i <= 36; i++){
//		printf("x:%d y:%d f:%d\r
",que[i].x,que[i].y,que[i].f);
//	}
	
	
	pp(tail-1);
	printf("(4, 4)\r
");
	return 0;
}