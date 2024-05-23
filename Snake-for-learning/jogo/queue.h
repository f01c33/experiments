typedef struct queue {
	block *vec; // ponteiro pra array de fato
	int first;  // indice do first valor preenchido
	int last;   // indice do last valor preenchido
	int size;   // size atual da array
} queue;

queue *queue_new(int);
void queue_grow(queue *);
int queue_len(queue *);
void queue_push(field *, queue *, block);
block queue_pop(field *, queue *);
void queue_print(queue *, bool);
void queue_free(queue*);

queue *queue_new(int size) { // aloca a a array e inicia os valores
	queue *que = (queue *)malloc(sizeof(queue));
	que->size = size;
	que->first = 0;
	que->last = 0;
	que->vec = (block *)malloc(sizeof(block) * size); // aloca a array de fato
	return que;
}

void queue_grow(queue *que) { // também arruma o vetor
	// int* new_vec = new int[que->size * 2]; // cria um vetor novo
	block *new_vec = (block *)malloc(sizeof(block) * (que->size * 2));
	if (que->last > que->first) { // se estiver em ordem
		for (int i = que->first; i <= que->last; i++) { 
		     // copia do first elemento ao last
			new_vec[i - que->first] = que->vec[i];
		}
		que->last = que->last - que->first;
		que->first = 0;
	} else if (que->last <= que->first) { // se já tiver dado a volta
		for (int i = que->first; i < que->size;
		     i++) { // copia first a parte final
			new_vec[i - que->first] = que->vec[i];
		}
		for (int i = 0; i <= que->last; i++) { // depois a do começo
			new_vec[que->size - que->first + i] = que->vec[i];
		}
		que->last += que->size - que->first;
		que->first = 0;
	}
	free(que->vec);     // desaloca a array antiga
	que->vec = new_vec; // guarda o ponteiro da nova, agora que ja a utilizamos
	que->size *= 2; 
	    // só consegue dobrar a size, parece uma utilização ok para mim
}

int queue_len(queue *que) { 
    // não precisa guardar o int tamanho na struct vetor!
	if (que->last == que->first && que->last == 0) { // se estiver vazio, tamanho 0
		return 0;
	} else if (que->last > que->first) { // se estiver na ordem esperada
		return que->last - que->first;
	} else { // que->last < que->first      caso esteja trocado
		return que->size - que->first + que->last;
	}
}

void queue_push(field *f, queue *que, block valor) {        
                // coloca no fim da queue, bem literalmente
	if (queue_len(que) < que->size) { // se tiver espaço
		if (que->last < que->size) {
			que->vec[que->last++] = valor;
		} else if (que->last == que->size && que->first != 0) { // da a volta
			que->last = 0;
			que->vec[que->last++] = valor;
		}
	} else { // sem espaço a queue cresce
		queue_grow(que);
		queue_push(f, que, valor);
		return;
	}
	// f cell(valor.i, valor.j) = valor; // draws
}

block queue_pop(field *f, queue *que) { // da a volta ao contrário
	if (que->first == que->size) {
		que->first = 0;
	}
	block val = que->vec[que->first];   // retira o primeiro valor
	// f cell(val.i, val.j) = empty_block; // draws
	que->first += 1;
	if (que->first == que->last) { // se está vazia
		que->first = 0;
		que->last = 0;
		return empty_block;
	}
	return val;
}

void queue_print(queue *que, bool dbg) {
	if (dbg) {
		// cout << "na memória:\n";
		wprintf(L"na memória:\n");
		for (int i = 0; i < que->size; i++) {
			wprintf(L"%c%c,", que->vec[i].v[0], que->vec[i].v[1]);
			// cout << que->vec[i] << ",";
		}
		wprintf(L" tamanho: %d, prim: %d, ult: %d, size: %d, len: %d Em ordem:\n",
		       queue_len(que), que->first, que->last, que->size, queue_len(que));
		// cout << " tamanho: " << queue_len(que) << ", prim: " <<
		// que->first
		//      << ", ult: " << que->last << ", size: " << que->size
		//      << ", Em ordem:\n";
	}

	// cout << "[";
	wprintf(L"[");
	int i = que->first;
	while (true) {
		if (i == que->last) {
			break;
		}
		// cout << que->vec[i] << ",";
		wprintf(L"%c%c,", que->vec[i].v[0], que->vec[i].v[1]);
		if (i >= que->size) {
			i = 0;
		} else { i++; }
	}
	// cout << "]\n";
	wprintf(L"]");
}

block* queue_last(queue *que) {
	if (que->size == 0) {
		return &empty_block;
	} else if (que->last != 0) {
        assert(que->vec[que->last-1].code == 1);
        return &que->vec[que->last-1];
	} else {
        return &que->vec[que->size-1];
	}
}

void inline queue_free(queue* que){
	free(que->vec);
	free(que);
}

// int main() {
// 	queue *queueaa = queue_new(8);
// 	queue_print(queueaa, true);
// 	for (int i = 0; i < 8; i++) {
// 		queue_push(queueaa, i);
// 	}
// 	queue_print(queueaa, true);
// 	for (int i = 0; i < 4; i++) {
// 		queue_pop(queueaa);
// 	}
// 	queue_push(queueaa, 666);
// 	queue_print(queueaa, true);
// 	return 0;
// }
