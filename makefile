# Compiler and flags
CC = gcc
CFLAGS = -Wall -Wextra -g
TARGET = ./output/vague

# Find all .c source files recursively
SRCS = $(shell find ./source -name "*.c")

# Find all directories for header inclusion (-I)
INCLUDES = $(shell find ./source -type d | sed 's/^/-I/')

all: $(TARGET)

# Automatically create the output directory and compile
$(TARGET): $(SRCS)
	mkdir -p ./output
	$(CC) $(SRCS) $(INCLUDES) $(CFLAGS) -o $(TARGET)

# Build and execute the program
run: all
	$(TARGET)

clean:
	rm -f $(TARGET)

# Prevent conflicts with files named 'all', 'run', or 'clean'
.PHONY: all run clean
