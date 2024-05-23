#include <stdio.h>

void move_cursor(int x, int y){
   printf("\033[%d;%dH",x,y);
}

int main() {
   move_cursor(0,0);
   printf("00");
   move_cursor(10,10);
   printf("1010");   
   return 0;
}
