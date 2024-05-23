struct score_t {
    wchar_t name[100];
    int sc;
};

void get_highscore(wchar_t* scores, size_t size) {
    FILE* score = fopen("scores.txt", "r+");
    unsigned int flag = 2;
    wchar_t name[100] = {0};
    wchar_t buf[size];
    for (unsigned long int i = 0; i < size; i++) {
        buf[i] = L'\0';
    }
    int scor = 0;
    while (flag != WEOF && flag >= 2) {
        flag = fwscanf(score, L"%d %ls", &scor, name);
        if (flag != WEOF) {
            swprintf(buf, size, L"%d %ls\n", scor, name);
            wcscat(scores, buf);
            // wprintf(L"%ls",scores);
        }
    }
    wcscat(scores, L" ");
    fclose(score);
}

void save_highscore(int current_score) {
    clear;
    // assumes end() has been called
    wprintf(L"type your name:\n");
    char c;
    while ((c = getchar()) != '\n' && c != EOF) {
    }  // does not fix the input error
    wchar_t name[100];
    wscanf(L"%*[^\n] %ls", &name);  // not helping
    FILE* score = fopen("scores.txt", "r+");
    struct score_t score_array[5];
    bool flag = false;
    for (int i = 0; i < 5; i++) {
        fwscanf(score, L"%d %ls", &score_array[i].sc, score_array[i].name);
        if (!flag && score_array[i].sc <= current_score) {
            if (i != 4) {
                score_array[i + 1] = score_array[i];
            }
            score_array[i].sc = current_score;
            wcscpy(score_array[i].name, name);
            flag = true;
            i++;
        }
    }
    for (int i = 0; i < 5; i++) {
        wprintf(L"%d,%ls\n", score_array[i].sc, score_array[i].name);
    }
    fclose(score);
    score = fopen("scores.txt", "w");
    for (int i = 0; i < 5; i++) {
        fwprintf(score, L"%d\n%ls", score_array[i].sc, score_array[i].name);
        if (i != 4) {
            fwprintf(score, L"\n");
        }
    }
    fclose(score);
}

void field_write(field* f,
                 wchar_t* text,
                 unsigned short int i,
                 unsigned short int j) {
    for (int d = 0, k = 0, lines = j; text[d] && k + 1 < f->j; d += 2, k++) {
        if (k >= f->j - 2 || text[d] == L'\n') {
            d++;
            k = 0;
            lines++;
        }
        if (f cell(lines, k + i).code ==
            IS_BG) {  // only write on the background
            f cell(lines, k + i).v[0] = text[d];
            f cell(lines, k + i).code = IS_TEXT;
            f cell(lines, k + i).v[1] = BG;
        }
        if (text[d + 1] == L'\0') {
            break;
        }
        if (k >= f->j - 2 || text[d + 1] == L'\n') {
            // d--;       // don't write the \n
            // d++;
            k = 0;     // start writing from the start again
            lines++;   // jump a line
            continue;  // will only start text on the first block space
        }
        if (f cell(lines, k + i).code == IS_TEXT) {
            f cell(lines, k + i).v[1] = text[d + 1];
        }
    }
}

int dif_chooser(int sleep_time, int* board) {  // good ol' event loop
    char c = 0, state = 1;
    field* f = field_new(board[0], board[1]);
    wchar_t* level;
    while (true) {
        field_clear(f);
        std_sleep(sleep_time);
        c = input();
        switch (state) {
            case 0:
                level = L"Choose your dificulty:\n \n ☒easy \n ◽medium \n ◽hard";
                break;
            case 1:
                level = L"Choose your dificulty:\n \n ◽easy \n ☒medium \n ◽hard";
                break;
            case 2:
                level = L"Choose your dificulty:\n \n ◽easy \n ◽medium \n ☒hard";
                break;
        }
        field_write(f, level, 1, 1);
        field_print(f);
        switch (c) {
            case '\n':
                return state;
            case 'w':
                if (state > 0)
                    state -= 1;
                break;
            case 's':
                if (state < 2)
                    state += 1;
                break;
        }
    }
}

int game_over(int sleep_time, int* board, int food_n) {
    char c = 0, state = 1;
    field* f = field_new(board[0], board[1]);
    wchar_t level[100];
    while (true) {
        field_clear(f);
        std_sleep(sleep_time);
        c = input();
        switch (state) {
            case 0:
                swprintf(level, 100,
                         L"You lost the game with %d rats!\n Try again?\n ☒no "
                         L"\n ◽yes",
                         food_n);
                break;
            case 1:
                swprintf(level, 100,
                         L"You lost the game with %d rats!\n Try again?\n ◽no "
                         L"\n ☒yes",
                         food_n);
                break;
        }
        field_write(f, level, 1, 1);
        field_print(f);
        switch (c) {
            case '\n':
                return state;
            case 'w':
                if (state > 0)
                    state -= 1;
                break;
            case 's':
                if (state < 1)
                    state += 1;
                break;
        }
    }
}
