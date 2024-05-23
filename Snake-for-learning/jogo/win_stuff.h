//#include <windows.h>

CONSOLE_SCREEN_BUFFER_INFO csbi;
void start() {
    GetConsoleScreenBufferInfo(GetStdHandle(STD_OUTPUT_HANDLE), &csbi);
    return;
}

void end() {
    return;
}
int width() {
    return csbi.srWindow.Right - csbi.srWindow.Left + 1;
}
int height() {
    return csbi.srWindow.Bottom - csbi.srWindow.Top + 1;
}
unsigned int input() {
    unsigned int input_data;
    if (_kbhit()) {
        input_data = (unsigned int)_getch();
    } else {
        input_data = 0;
    }
    return input_data;
}

void std_sleep(int waitTime) {
    __int64 time1 = 0, time2 = 0, freq = 0;

    QueryPerformanceCounter((LARGE_INTEGER*)&time1);
    QueryPerformanceFrequency((LARGE_INTEGER*)&freq);

    do {
        QueryPerformanceCounter((LARGE_INTEGER*)&time2);
    } while ((time2 - time1) < waitTime);
}