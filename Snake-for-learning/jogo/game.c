#include "game_types.h"
#include "queue.h"
#include "menus.h"

typedef struct snake {
    queue* q;
    char direction;
} snake;


block snake_move(field* f, snake* snake) {
    char dir = snake->direction;
    block next, out;
    block* tmp;
    tmp = queue_last(snake->q);
    bool grow = false;
    if (dir == 'w') {
        if (f cell(tmp->i - 1, tmp->j).code == IS_FOOD) {
            grow = true;
        }
        next =
            block_new(IS_SNAKE, tmp->i - 1, tmp->j, snake_up[0], snake_up[1]);
    } else if (dir == 's') {
        if (f cell(tmp->i + 1, tmp->j).code == IS_FOOD) {
            grow = true;
        }
        next = block_new(IS_SNAKE, tmp->i + 1, tmp->j, snake_down[0],
                         snake_down[1]);
    } else if (dir == 'a') {
        if (f cell(tmp->i, tmp->j - 1).code == IS_FOOD) {
            grow = true;
        }
        next = block_new(IS_SNAKE, tmp->i, tmp->j - 1, snake_left[0],
                         snake_left[1]);
    } else {  // if (dir == 'd') {
        if (f cell(tmp->i, tmp->j + 1).code == IS_FOOD) {
            grow = true;
        }
        next = block_new(IS_SNAKE, tmp->i, tmp->j + 1, snake_right[0],
                         snake_right[1]);
    }
    f cell(tmp->i, tmp->j).v[0] =
        snake_body[0];  // many segfaults were felt here
    f cell(tmp->i, tmp->j).v[1] = snake_body[1];
    tmp->v[0] = snake_body[0];
    tmp->v[1] = snake_body[1];
    queue_push(f, snake->q, next);
    out = f cell(next.i, next.j);  // what was the block before it became a
                                   // snake
    f cell(next.i, next.j) = next;
    if (!grow) {
        block tail = queue_pop(f, snake->q);
        f cell(tail.i, tail.j) = empty_block; // draws
    }
    return out;
}
bool correct_anwser(question* q, block* head){
    switch(head->v[0]){
        case L'A':
        if(q->answer == 0) return true;
        break;
        case L'B':
        if(q->answer == 1) return true;
        break;
        case L'C':
        if(q->answer == 2) return true;
        break;
        case L'D':
        if(q->answer == 3) return true;
        break;
        case L'E':
        if(q->answer == 4) return true;
        break;
    }
    return false;
}

bool snake_alive(field* f, snake* snake, block* head,question* q) {
    if (head->code == IS_SNAKE) {
        return true;
    }
    if(head->code == IS_FOOD){
        if(correct_anwser(q,head)){
            return false;//quit = false
        } else {
            return true;
        }
    }
    block* tmp = queue_last(snake->q);
    if (tmp->i == 0 || tmp->i == f->i - 1) {
        return true;
    } else if (tmp->j == 0 || tmp->j == f->j - 1) {
        return true;
    } else {
        return false;
    }
}

void new_fruit(field* f, int code) {
    int i = 1 + rand() % (f->i - 2);
    int j = 1 + rand() % (f->j - 2);
    if (f cell(i, j).code == IS_BG) {
        f cell(i, j) = block_new(2, i, j, food[code][0], food[code][1]);
    } else {
        new_fruit(f,code);
        return;
    }
}

void remove_block(field* f,int code){
    for(int i = f->i-1, j; i >= 0; i--){
        for(j = f->j-1; j >= 0; j--){
            if(f cell(i,j).code == code){
                f cell(i,j) = empty_block;
            }
        }
    }
}

// board size, width,height
int board[] = {23, 40};

int main() {
    srand(time(NULL));
    setlocale(LC_ALL, "");
    start();

    const char* filenames[] = {"pergunta0","pergunta1","pergunta2"};
    int filenamesize = sizeof(filenames)/sizeof(filenames[0]);
    bool chosen[filenamesize];
    for(int i = 0; i < filenamesize; i++){
        chosen[i] = false;
    }
    int sleep_time = 48000;
    field* f;
    board[1] = width()/2;//each block has 2 chars of width
    board[0] = height()-1;
    int difficulty = dif_chooser(sleep_time, board);

    switch (difficulty) {
        case 0:
            sleep_time += 16000;
            break;
        case 2:
            sleep_time -= 16000;
            break;
    }

    f = field_new(board[0], board[1]);
    // wchar_t scores[100] = {0};
    // get_highscore(scores,100);
    // field_write(f,scores,5,5);
    // free(scores);
    snake* snake = malloc(sizeof(snake));
    snake->q = queue_new(4);
    snake->direction = 'w';
    block head;
    queue_push(f, snake->q, block_new(1, (f->i / 2) + 1, f->j / 2,
                                      snake_body[0], snake_body[1]));
    queue_push(f, snake->q,
               block_new(1, f->i / 2, f->j / 2, snake_up[0], snake_up[1]));
    question q;
    bool new_question = true,quit = false;
    char dir;
    int in;
    while (!quit) {
        if(new_question){
            remove_block(f,IS_TEXT);
            int i;
            for(i = 0; i < filenamesize; i++){
                if(!chosen[i]) {
                    chosen[i] = true;
                    break;
                }
            }
            if(i == filenamesize){
                i = 0;//turn on congratulations
            }
            q = read_question(filenames[i]);
        }
        field_write(f,q.text,1,1);
        field_print(f);
        std_sleep(sleep_time);
        in = input();
        dir = snake->direction;
        switch (in) {
            case 'w':
                if (snake->direction != 's')
                    snake->direction = 'w';
                break;
            case 's':
                if (snake->direction != 'w')
                    snake->direction = 's';
                break;
            case 'a':
                if (snake->direction != 'd')
                    snake->direction = 'a';
                break;
            case 'd':
                if (snake->direction != 'a')
                    snake->direction = 'd';
                break;
            default:
                break;
        }
        head = snake_move(f, snake);
        quit = snake_alive(f, snake, &head, &q);
        if(new_question){
            for(int i = 0; i < 5; i++){
                new_fruit(f,i);
            }
            new_question = false;
        }
        if(!quit && head.code == IS_FOOD){
            new_question = true;
            remove_block(f,IS_FOOD);
        }
        // if (!(rand() % 60)) {
        //     new_fruit(f);
        // }
    }
    int snake_size = queue_len(snake->q) - 2;
    int retry = game_over(sleep_time, board, snake_size);
    // wprintf(L"you lost the game with %d rats!\n", );

    field_free(f);
    queue_free(snake->q);
    free(snake);
    end();

    if (retry == 1)
        return main();
    
    // save_highscore(snake_size);
    return 0;
}
