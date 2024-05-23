// #pragma once
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <assert.h>
#include <time.h>
#include <locale.h>
#include <wchar.h>

#ifdef _WIN32
	#include <conio.h>
	#include <windows.h>
	#include "win_stuff.h"
    #define clear system("cls")
#elif (defined __unix__ || defined __APPLE__)
    #include <sys/time.h>
    #include <unistd.h>
    #include <termios.h>
    #include <sys/ioctl.h>
    #include "linux_stuff.h"
    #define clear system("clear")
/*#elif defined __APPLE__
    #include <sys/time.h>
    #include <unistd.h>
    #include <termios.h>
    #include "linux_stuff.h"
    #define clear system("clear")
*/
#endif

#define newline wprintf(L"\n")//putchar('\n')
// #define m(i,j,max_i) [i*max_i+j]
#define cell(i,j) ->m[i][j]
//"sprites"

#ifdef _WIN32 //sorry windows, but you're shit at this
#define BG 			(L' ')
#define top_u  		(L"══")//("==")
#define top_d  		(L"══")//("==")
#define side_l 		(L"║ ")//("| ")
#define side_r 		(L" ║")//(" |")
#define corner_ur 	(L"═╗")//("=\\")
#define corner_ul 	(L"╔═")//("/=")
#define corner_dr 	(L"═╝")//("=/")
#define corner_dl 	(L"╚═")//("\\=")

#define snake_up 	(L"/\\")//("/\\")
#define snake_down 	(L"/\\")//("\\/")
#define snake_left 	(L">#")//("<\%")
#define snake_right (L"#<")//("\%>")
#define snake_body 	(L"##")//("\%\%")
#define food 		(L":B")//(":B")
#else

#define BG 			(L' ')
#define top_u  		(L"══")//("==")
#define top_d  		(L"══")//("==")
#define side_l 		(L"║ ")//("| ")
#define side_r 		(L" ║")//(" |")
#define corner_ur 	(L"═╗")//("=\\")
#define corner_ul 	(L"╔═")//("/=")
#define corner_dr 	(L"═╝")//("=/")
#define corner_dl 	(L"╚═")//("\\=")

#define snake_up 	(L"ᒠᒍ")//("/\\")
#define snake_down 	(L"ᒖᒉ")//("\\/")
#define snake_left 	(L"ᑒ⠿")//("<\%")
#define snake_right (L"⠿ᑣ")//("\%>")
#define snake_body 	(L"⠿⠿")//("\%\%")
// #define food 		(L"හ୭")//(":B")
wchar_t food[][2] = {{L'A',L')'},{L'B',L')'},{L'C',L')'},{L'D',L')'},{L'E',L')'}};

#endif

#define IS_BG 0 //for checking the .code field inside the blocks
#define IS_SNAKE 1
#define IS_FOOD 2
#define IS_TEXT 3

// #define mSecond *1000000L //easier usage of nanosleep

// #define FLAG1 (int) 32 //for up to 32 char inputs and 10 flags in *out in the menus
// #define FLAG2 (int) 64
// #define FLAG3 (int) 128
// #define FLAG4 (int) 256
// #define FLAG5 (int) 512
// #define FLAG6 (int) 1024
// #define FLAG7 (int) 2048
// #define FLAG8 (int) 4096
// #define FLAG9 (int) 8192
// #define FLAG10 (int) 16384

typedef struct block{
	unsigned short int i,j;
	unsigned char code;
	wchar_t v[2];
}block;

block empty_block = {0,0,0,{BG,BG},};

typedef struct field{
	block** m; //matrice of blocks
	int i, j;
}field;

block block_new(char,int,int,wchar_t,wchar_t);
field* field_new(int,int);
void field_print(field *);
void field_free(field*);
void field_clear(field*);

block block_new(char code,int i,int j,wchar_t v0,wchar_t v1){
	block out;// = (block *) malloc(sizeof(block));
	out.code = code;
	out.v[0] = v0;
	out.v[1] = v1;
	out.i = i;
	out.j = j;
	return out;
}

field* field_new(int i,int j){
	field* f = (field *) malloc(sizeof(field));
	f->m = (block **) malloc(sizeof(block *)*i);
	for(int a = 0; a < i; a++){
		f->m[a] = (block *) malloc(sizeof(block)*j);
	}
	f->i = i;
	f->j = j;
    field_clear(f);
    return f;
}

void field_clear(field* f){
    for(int a = 0; a < f->i; a++){
        for(int b = 0; b < f->j; b++){
            f->m[a][b] = empty_block;
        }
    }
    for(int a = 0; a < f->i; a++){
        f cell(a,0).v[0] = side_l[0];
        f cell(a,0).v[1] = side_l[1];
        f cell(a,f->j-1).v[0] = side_r[0];
        f cell(a,f->j-1).v[1] = side_r[1];
    }
    for(int b = 0; b < f->j; b++){
        f cell(0,b).v[0] = top_u[0];
        f cell(0,b).v[1] = top_u[1];
        f cell(f->i-1,b).v[0] = top_d[0];
        f cell(f->i-1,b).v[1] = top_d[1];
    }
    f cell(0,0).v[0] = corner_ul[0]; f cell(0,0).v[1] = corner_ul[1];
    f cell(0,f->j-1).v[0] = corner_ur[0]; f cell(0,f->j-1).v[1] = corner_ur[1];
    f cell(f->i-1,0).v[0] = corner_dl[0]; f cell(f->i-1,0).v[1] = corner_dl[1];
    f cell(f->i-1,f->j-1).v[0] = corner_dr[0]; f cell(f->i-1,f->j-1).v[1] = corner_dr[1];
}


void field_print(field* f){
	clear;
	for(int i = 0; i < f->i; i++){
		for(int j = 0; j < f->j; j++){
			wprintf(L"%lc%lc",f->m[i][j].v[0],f->m[i][j].v[1]);
		}
		newline;
	}
}

void field_free(field* f){
	for(int i = 0; i < f->i;i++){
        free(f->m[i]);
	}
	free(f->m);
	return;
}

typedef struct question{
   wchar_t text[500];
   char answer;
}question;

question read_question(const char* filename){
   FILE* q_file = fopen(filename,"r");
   wchar_t tmp[500];
   question q;
   fgetws(tmp,500,q_file);
   for(int i = 0,j = 0; i < 500 && tmp[i]; i++,j++){
      q.text[j] = tmp[i];
      if(tmp[i] == L'\\' && tmp[i+1] == L'n'){
         i++;
         q.text[j] = L'\n';
      }
      if(!(tmp[i+1])){
        q.text[j] = L'\0';
      }
   }
   fgetws(tmp,500,q_file);
   q.answer = tmp[0]-L'0';
   return q;
}
