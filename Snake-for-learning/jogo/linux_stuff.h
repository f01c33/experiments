// #include <termios.h>
// #include <sys/time.h>
// #include <unistd.h>
// void move_cursor(int x, int y){
//    printf("\033[%d;%dH",x,y);
// }

int kbhit() {
    struct timeval tv;
    fd_set fds;
    tv.tv_sec = 0;
    tv.tv_usec = 0;
    FD_ZERO(&fds);
    FD_SET(STDIN_FILENO, &fds);  // STDIN_FILENO is 0
    select(STDIN_FILENO + 1, &fds, NULL, NULL, &tv);
    return FD_ISSET(STDIN_FILENO, &fds);
}

void nonblock(int state) {  // set terminal to and from non-blocking
    struct termios ttystate;

    // get the terminal state
    tcgetattr(STDIN_FILENO, &ttystate);

    if (state == 1) {
        // turn off canonical mode
        ttystate.c_lflag &= ~ICANON;
        // minimum of number input read.
        ttystate.c_cc[VMIN] = 1;
    } else if (state == 0) {
        // turn on canonical mode
        ttystate.c_lflag |= ICANON;
    }
    // set the terminal attributes.
    tcsetattr(STDIN_FILENO, TCSANOW, &ttystate);
}

unsigned int input() {
    if (kbhit()) {
        return (unsigned int)fgetc(stdin);
    } else {
        return 0;
    }
}

struct timespec wt[] = {{0, 0}};

void std_sleep(int waitTime) {
    wt[0].tv_nsec = waitTime * 1000;
    nanosleep(wt, NULL);
}

struct winsize w;

void start() {
    nonblock(1);  // STDOUT_FILENO is 0, usually
    ioctl(STDOUT_FILENO, TIOCGWINSZ, &w);
    return;
}
void end() {
    nonblock(0);
    return;
}

int width() {
    return w.ws_col;
}
int height() {
    // ioctl(STDOUT_FILENO,TIOCGWINSZ,&w);
    return w.ws_row;
}
