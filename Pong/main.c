/*
   MyOpenGl.c
   Template by Ionildo Jose Sanches - 11/08/16
   Rest by Cauê Felchar
*/
// linker -lGL -lGLU -lglut -lm
/* Includes requeridos */
#ifdef _WIN32
#include <windows.h>
#endif
#include <GL/gl.h>
#include <GL/glut.h>
#include <math.h>
#include <string.h>
#include "itoa.h"

#define UP 1
#define STILL 0
#define DOWN -1
#define BALL_SPD 0.7

int width = 100, heigth = 100;

int PAD_W = 1, PAD_H = 5;

void myInit(void);

struct obj {
  double x, y;
  double speed;
  struct v2d{
    double x,y;
  }vec;
  char point;
  // double angle;
};

struct obj pad[2] = {{.point = 0}, {.point = 0}};
struct obj ball = {1};

double fmod(double a,double mod){
  while(a >= mod){
    a-=mod;
  }
  return a;
}

void pad_move(int pad_n) {
  if (pad[pad_n].vec.x == UP) {
    pad[pad_n].y -= pad[pad_n].speed;
  } else if (pad[pad_n].vec.x == DOWN) {
    pad[pad_n].y += pad[pad_n].speed;
  }
  if(pad[pad_n].y+PAD_H > heigth){
    pad[pad_n].y = heigth-PAD_H;
  } else if (pad[pad_n].y-PAD_H < 0){
    pad[pad_n].y = PAD_H;
  }
}

void pad_show(int pad_n) {
  glColor3f(1.0, 1.0, 1.0);
  glRectf(pad[pad_n].x + PAD_W, pad[pad_n].y + PAD_H, pad[pad_n].x - PAD_W,
          pad[pad_n].y - PAD_H);
}

void ball_show() {
  glColor3f(1.0, 1.0, 1.0);
  glRectf(ball.x + PAD_W, ball.y + PAD_W, ball.x - PAD_W, ball.y - PAD_W);
}

void ball_move(){
  ball.x += ball.vec.x*ball.speed;
  ball.y += ball.vec.y*ball.speed;
  // ball.x+= cos(ball.angle)*ball.speed;
  // ball.y+= sin(ball.angle)*ball.speed;
  if(ball.y+PAD_W > heigth || ball.y-PAD_W < 0){ //bola bateu em cima ou em baixo
    ball.vec.y = - ball.vec.y;
    if(ball.y+PAD_W > heigth){
      ball.y -= PAD_W;
    }else if(ball.y-PAD_W < 0){
      ball.y += PAD_W;
    }

    // ball.angle = fmod(ball.angle,M_PI) + M_PI;
    // ball.angle = fmod(ball.angle,2*M_PI);
  }
}

int touch(){
  if(fabs(ball.x-pad[0].x) < 1.2*PAD_W && fabs(ball.y-pad[0].y) < 1.2*PAD_H){
    return 0;
  } else if(fabs(ball.x-pad[1].x) < 1.2*PAD_W && fabs(ball.y-pad[1].y) < 1.2*PAD_H){
    return 1;
  } else if(ball.x+PAD_W > width){
    return 2; //point for player 1
  } else if(ball.x-PAD_W < 0){
    return 3; //point for player 2
  }
  return -1;
}

void write_int(int x, int y, int val){
  char buf[32] = {0};
  itoa(val,buf,10);
  glColor3f(1.0, 1.0, 1.0);
  glRasterPos2f(x,y);
  for(int i = 0; i < strlen(buf); i++){
    // printf("%c",buf[i]);
    glutBitmapCharacter(GLUT_BITMAP_TIMES_ROMAN_24, buf[i]);
  }
}

/* Executada sempre que é necessario re-exibir a imagem */
void display(void) {
  glClear(GL_COLOR_BUFFER_BIT);
  // glRectf(-0.5, 0.5, 0.5, -0.5); /* coordenadas (x0,y0) e (x1,y1) */
  pad_show(0);
  pad_show(1);
  ball_show();
  glColor3f(1.0, 1.0, 1.0);

  //linha de centro
  glPushAttrib(GL_ENABLE_BIT); 
  glLineStipple(1, 0x0F00);
  glEnable(GL_LINE_STIPPLE);

  glBegin(GL_LINES);
  glVertex2f(width/2, 0);
  glVertex2f(width/2, heigth);
  glEnd();
  
  glPopAttrib();
  //
  //score
  write_int(  width/4,10,pad[0].point);
  write_int(3*width/4,10,pad[1].point);
  // glRasterPos2f(width/4,10);
  // glutBitmapCharacter(GLUT_BITMAP_TIMES_ROMAN_24, (char) pad[0].point+'0');
  // glRasterPos2f(3*width/4,10);
  // glutBitmapCharacter(GLUT_BITMAP_TIMES_ROMAN_24, (char) pad[1].point+'0');
  //

  // glFlush();
  glutSwapBuffers();
}

/* Função ativada qdo a janela é aberta pela primeira vez e toda vez
   que a janela é reconfigurada (movida ou modificado o tamanho) */

void myReshape(int w, int h) {
  glViewport(0, 0, w, h);
  glMatrixMode(GL_PROJECTION);
  glLoadIdentity();
  glOrtho(0.0, 100.0, 100.0, 0.0, -10.0, 10.0);
  // width = w;
  // heigth = h;
  // PAD_W = (5.0 / 250.0) * w;
  // PAD_H = (20.0 / 250.0) * h;
  glMatrixMode(GL_MODELVIEW);
  glLoadIdentity();
}

/* Função ativada qdo alguma tecla é pressionada */

void Key(unsigned char key, int x, int y) {
  switch (key) {
  case 27:
    exit(1);
    break; /* Esc finaliza o programa */
  case (unsigned char)'r':
    myInit();
    break;
  case (unsigned char)'q'://player 1
    pad[0].vec.x = UP;
    // pad_move(0,UP);
    break;
  case (unsigned char)'a':
    pad[0].vec.x = DOWN;
    // pad_move(0,DOWN);
    break;
  case (unsigned char)'o'://player 2
    pad[1].vec.x = UP;
    // pad_move(1,UP);
    break;
  case (unsigned char)'l':
    pad[1].vec.x = DOWN;
    // pad_move(1,DOWN);
    break;
  default:
    break;
  }
}

void Keyup(unsigned char key, int x, int y){
  switch (key) {
  case 27:
    exit(1);
    break; /* Esc finaliza o programa */
  case (unsigned char)'q'://player 1
    pad[0].vec.x = STILL;
    // pad_move(0,UP);
    break;
  case (unsigned char)'a':
    pad[0].vec.x = STILL;
    // pad_move(0,DOWN);
    break;
  case (unsigned char)'o'://player 2
    pad[1].vec.x = STILL;
    // pad_move(1,UP);
    break;
  case (unsigned char)'l':
    pad[1].vec.x = STILL;
    // pad_move(1,DOWN);
    break;
  default:
    break;
  }}

/* Inicializações do programa */

void myInit(void) {
  glClearColor(0.0, 0.0, 0.0, 0.0);
  pad[0].x = 2*PAD_W;
  pad[0].y = heigth/2;
  pad[1].x = width - 2*PAD_W;
  pad[1].y = heigth/2;
  pad[0].speed = pad[1].speed = 1.2;
  pad[0].point = pad[1].point = 0;
  // printf("%f,%f\n", pad[1].x, pad[1].y);
  ball.x = width/2;
  ball.y = heigth/2;
  // ball.angle = M_PI;
  ball.vec.x = 1.0;
  ball.vec.y = 0.0;

  ball.speed = BALL_SPD;
}

void tick(int in) {
  ball_move();
  pad_move(0);
  pad_move(1);
  int colision = touch();
  switch(colision){
    case 0:
      // ball.angle = atan(fabs(pad[0].y-ball.y)/fabs(pad[0].x-ball.x+10));
      ball.vec.x = -cos(atan2(pad[0].y-ball.y,pad[0].x-(ball.x+2)));
      ball.vec.y = -sin(atan2(pad[0].y-ball.y,pad[0].x-(ball.x+2)));
      ball.speed = fabs(ball.speed)*1.05;
    break;
    case 1:
      ball.vec.x = -cos(atan2(pad[1].y-ball.y,pad[1].x-(ball.x-2)));
      ball.vec.y = -sin(atan2(pad[1].y-ball.y,pad[1].x-(ball.x-2)));
      // ball.angle = atan(fabs(pad[1].y-ball.y)/fabs(pad[1].x-ball.x-10));
      ball.speed = fabs(ball.speed)*1.05;
    break;
    case 2:
      ball.x = pad[1].x-5;
      ball.y = pad[1].y;
      ball.vec.x = -1;
      ball.vec.y = 0;
      pad[0].point+=1;
      ball.speed = BALL_SPD;
    break;
    case 3:
      ball.x = pad[0].x+5;
      ball.y = pad[0].y;
      ball.vec.x = 1;
      ball.vec.y = 0;
      pad[1].point+=1;
      ball.speed = BALL_SPD;
    break;
    default:
    break;
  }
  glutPostRedisplay();
  glutTimerFunc(1,tick,0);
  // return tick(in);
}

/* Parte principal - ponto de início de execução */

int main(int argc, char **argv) {
  glutInit(&argc, argv);
  glutInitDisplayMode(GLUT_DOUBLE | GLUT_RGB);
  glutInitWindowSize(640, 480);
  glutInitWindowPosition(0, 0);
  glutCreateWindow("PONG");
  glutReshapeFunc(myReshape);
  glutDisplayFunc(display);
  myInit();
  glutKeyboardFunc(Key);
  glutKeyboardUpFunc(Keyup);
  glutTimerFunc(1000,tick,0);
  glutMainLoop();
  return (0);
}
