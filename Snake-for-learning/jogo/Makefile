#OBJS specifies which files to compile as part of the project
OBJS = game.c

#CC specifies which compiler we're using
CC = clang-3.6
# CC = gcc

#COMPILER_FLAGS specifies the additional compilation options we're using
# -w suppresses all warnings
COMPILER_FLAGS = -std=c11 -Wall -Wpedantic -g
COMPILER_FLAGS_rls = -std=c11 -march=native -O3

#LINKER_FLAGS specifies the libraries we're linking against
LINKER_FLAGS = 

#OBJ_NAME specifies the name of our exectuable
OBJ_NAME = game

#This is the target that compiles our executable
all : $(OBJS)
	$(CC) $(OBJS) $(COMPILER_FLAGS) $(LINKER_FLAGS) -o $(OBJ_NAME)

release : $(OBJS)
	$(CC) $(OBJS) $(COMPILER_FLAGS_rls) $(LINKER_FLAGS) -o $(OBJ_NAME)