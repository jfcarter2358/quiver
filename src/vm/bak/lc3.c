/**
 * LC-3 CM for running quiver programs
 *
 * See https://www.jmeiners.com/lc3-vm/ for implementation reference
 */

#include <stdio.h>
#include <stdint.h>
#include <signal.h>
/* unix only */
#include <stdlib.h>
#include <unistd.h>
#include <fcntl.h>
#include <sys/time.h>
#include <sys/types.h>
#include <sys/termios.h>
#include <sys/mman.h>

#define MEMORY_MAX (1 << 16)

/* -------- REGISTERS -------- */
enum {
    R_R0 = 0,
    R_R1,
    R_R2,
    R_R3,
    R_R4,
    R_R5,
    R_R6,
    R_R7,
    R_PC, // program counter
    R_COND,
    R_COUNT
};

/* -------- OPCODES -------- */
enum {
    OP_BR = 0, // branch
    OP_ADD,    // add
    OP_LD,     // load
    OP_ST,     // store
    OP_JSR,    // jump register
    OP_AND,    // bitwise and
    OP_LDR,    // load register
    OP_STR,    // store register
    OP_RTI,    // unused
    OP_NOT,    // bitwise not
    OP_LDI,    // load indirect
    OP_STI,    // store indirect
    OP_JMP,    // jump
    OP_RES,    // reserved (unused)
    OP_LEA,    // load effective address
    OP_TRAP    // execute trap
};

/* -------- CONDITION FLAGS -------- */
enum {
    FL_POS = 1 << 0, // positive
    FL_ZRO = 1 << 1, // zero
    FL_NEG = 1 << 2, // negative
};

/* -------- TRAP CODES -------- */

enum {
    TRAP_GETC = 0x20,  // get character from keyboard, not echoed onto the terminal
    TRAP_OUT = 0x21,   // output a character
    TRAP_PUTS = 0x22,  // output a word string
    TRAP_IN = 0x23,    // get character from keyboard, echoed onto the terminal
    TRAP_PUTSP = 0x24, // output a byte string
    TRAP_HALT = 0x25   // halt the program
};

/* -------- MEMORY MAPPED REGISTERS -------- */
enum{
    MR_KBSR = 0xFE00, // keyboard status
    MR_KBDR = 0xFE02  // keyboard data
};


uint16_t reg[R_COUNT];
uint16_t memory[MEMORY_MAX]; /* 65536 locations */

/* -------- INPUT BUFFERING -------- */

struct termios original_tio;

void disable_input_buffering() {
    tcgetattr(STDIN_FILENO, &original_tio);
    struct termios new_tio = original_tio;
    new_tio.c_lflag &= ~ICANON & ~ECHO;
    tcsetattr(STDIN_FILENO, TCSANOW, &new_tio);
}

void restore_input_buffering() {
    tcsetattr(STDIN_FILENO, TCSANOW, &original_tio);
}

uint16_t check_key() {
    fd_set readfds;
    FD_ZERO(&readfds);
    FD_SET(STDIN_FILENO, &readfds);

    struct timeval timeout;
    timeout.tv_sec = 0;
    timeout.tv_usec = 0;
    return select(1, &readfds, NULL, NULL, &timeout) != 0;
}

/* -------- MEMORY ACCESS -------- */

void mem_write(uint16_t address, uint16_t val) {
    memory[address] = val;
}

uint16_t mem_read(uint16_t address) {
    if (address == MR_KBSR) {
        if (check_key()) {
            memory[MR_KBSR] = (1 << 15);
            memory[MR_KBDR] = getchar();
        } else {
            memory[MR_KBSR] = 0;
        }
    }
    return memory[address];
}

/* -------- HANDLE INTERRUPT --------- */
void handle_interrupt(int signal)
{
    restore_input_buffering();
    printf("\n");
    exit(-2);
}

/* -------- SIGN EXTENSION -------- */
uint16_t sign_extend(uint16_t x, int bit_count) {
    if (( x >> (bit_count - 1)) & 1) {
        x |= (0xFFFF << bit_count);
    }

    return x;
}

/* -------- FLAG UPDATES -------- */
void update_flags(uint16_t r) {
    if (reg[r] == 0) {
        reg[R_COND] = FL_ZRO;
    } else if (reg[r] >> 15) { // a 1 in the left-most bit indicates negative
        reg[R_COND] = FL_NEG;
    } else {
        reg[R_COND] = FL_POS;
    }
}

/* -------- OPS -------- */

// And
void add(uint16_t instr) {
    // Destination register
    uint16_t r0 = (instr >> 9) & 0x7;
    // First operant (SR1)
    uint16_t r1 = (instr >> 6) & 0x7;
    // Whether we are in immediate mode;
    uint16_t imm_flag = (instr >> 5) & 0x1;

    if (imm_flag) {
        uint16_t imm5 = sign_extend(instr & 0x1F, 5);
        reg[r0] = reg[r1] + imm5;
    } else {
        uint16_t r2 = instr & 0x7;
        reg[r0] = reg[r1] + reg[r2];
    }

    update_flags(r0);
}

// Bitwise and
void and(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t r1 = (instr >> 6) & 0x7;
    uint16_t imm_flag = (instr >> 5) & 0x1;

    if (imm_flag) {
        uint16_t imm5 = sign_extend(instr & 0x1F, 5);
        reg[r0] = reg[r1] & imm5;
    } else {
        uint16_t r2 = instr & 0x7;
        reg[r0] = reg[r1] & reg[r2];
    }
    update_flags(r0);
} 

// Bitwise not
void not(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t r1 = (instr >> 6) & 0x7; 

    reg[r0] = ~reg[r1];
    update_flags(r0);
}

// Branch
void br(uint16_t instr) {
    uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
    uint16_t cond_flag = (instr >> 9) & 0x7;
    if (cond_flag & reg[R_COND]) {
        reg[R_PC] += pc_offset;
    }
}

// Jump
void jmp(uint16_t instr) {
    // Also handles RET
    uint16_t r1 = (instr >> 6) & 0x7;
    reg[R_PC] = reg[r1];
}

// Jump register
void jsr(uint16_t instr) {
    uint16_t long_flag = (instr >> 11) & 1;
    reg[R_R7] = reg[R_PC];
    if (long_flag) {
        uint16_t long_pc_offset = sign_extend(instr & 0x7FF, 11);
        reg[R_PC] += long_pc_offset;  // JSR
    } else {
        uint16_t r1 = (instr >> 6) & 0x7;
        reg[R_PC] = reg[r1]; // JSRR
    }
}

// Load
void ld(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
    reg[r0] = mem_read(reg[R_PC] + pc_offset);
    update_flags(r0);
}

// Load indirect
void ldi(uint16_t instr) {
    // Destination registry
    uint16_t r0 = (instr >> 9) & 0x7;
    // PC offset 9
    uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
    // Add pc_offset to the current PC, look at that memory location to get the final
    // address
    reg[r0] = mem_read(mem_read(reg[R_PC] + pc_offset));
    update_flags(r0);
}

// Load register
void ldr(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t r1 = (instr >> 6) & 0x7;
    uint16_t offset = sign_extend(instr & 0x3F, 6);
    reg[r0] = mem_read(reg[r1] + offset);
    update_flags(r0);
}

// Load effective address
void lea(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
    reg[r0] = reg[R_PC] + pc_offset;
    update_flags(r0);
}

// Store
void st(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
    mem_write(reg[R_PC] + pc_offset, reg[r0]);
}

// Store indirect
void sti(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t pc_offset = sign_extend(instr & 0x1FF, 9);
    mem_write(mem_read(reg[R_PC] + pc_offset), reg[r0]);
}

// Store register
void str(uint16_t instr) {
    uint16_t r0 = (instr >> 9) & 0x7;
    uint16_t r1 = (instr >> 6) & 0x7;
    uint16_t offset = sign_extend(instr & 0x3F, 6);
    mem_write(reg[r1] + offset, reg[r0]);
}

// Trap routines
void trap_getc() {
    // Read a single ASCII character
    reg[R_R0] = (uint16_t)getchar();
    update_flags(R_R0);
}

void trap_out() {
    putc((char)reg[R_R0], stdout);
    fflush(stdout);
}

void trap_puts() {
    // One char per word
    uint16_t* c = memory + reg[R_R0];
    while (*c)
    {
        putc((char)*c, stdout);
        ++c;
    }
    fflush(stdout);
}

void trap_in() {
    printf("Enter a character: ");
    char c = getchar();
    putc(c, stdout);
    fflush(stdout);
    reg[R_R0] = (uint16_t)c;
    update_flags(R_R0);
}

void trap_putsp() {
    // One char per byte (two bytes per word)
    // here we need to swap back to
    // big endian format */
    uint16_t* c = memory + reg[R_R0];
    while (*c)
    {
        char char1 = (*c) & 0xFF;
        putc(char1, stdout);
        char char2 = (*c) >> 8;
        if (char2) putc(char2, stdout);
        ++c;
    }
    fflush(stdout);
}

void trap_halt() {
    puts("HALT");
    fflush(stdout);
}

int trap(uint16_t instr) {
    int running = 1;
    reg[R_R7] = reg[R_PC];

    switch (instr & 0xFF) {
        case TRAP_GETC:
            trap_getc();
            break;
        case TRAP_OUT:
            trap_out();
            break;
        case TRAP_PUTS:
            trap_puts();
            break;
        case TRAP_IN:
            trap_in();
            break;
        case TRAP_PUTSP:
            trap_putsp();
            break;
        case TRAP_HALT:
            trap_halt();
            running = 0;
            break;
    }
    return running;
}

void res(uint16_t instr) {
    abort();
}

void rti(uint16_t instr) {
    abort();
}

/* -------- SWAP TO LITTLE-ENDIAN -------- */
uint16_t swap16(uint16_t x) {
    return (x << 8) | (x >> 8);
}

/* -------- READ IMAGE FILE -------- */
void read_image_file(FILE* file) {
    // The origin tells us where in memory to place the image
    uint16_t origin;
    fread(&origin, sizeof(origin), 1, file);
    origin = swap16(origin);

    // We know the maximum file size so we only need one fread
    uint16_t max_read = MEMORY_MAX - origin;
    uint16_t* p = memory + origin;
    size_t read = fread(p, sizeof(uint16_t), max_read, file);

    // Swap to little endian
    while (read-- > 0) {
        *p = swap16(*p);
        ++p;
    }
}

/* -------- READ IMAGE -------- */
int read_image(const char* image_path) {
    FILE* file = fopen(image_path, "rb");
    if (!file) { return 0; };
    read_image_file(file);
    fclose(file);
    return 1;
}

/* -------- MAIN LOOP -------- */
int main(int argc, const char *argv[])
{
    // Load arguments
    if (argc < 2)
    {
        // Show usage string
        printf("lc3 [image-file1] ...\n");
        exit(2);
    }

    for (int i = 1; i < argc; i++)
    {
        if (!read_image(argv[i]))
        {
            printf("Failed to load image: %s\n", argv[i]);
            exit(1);
        }
    }

    // Setup
    signal(SIGINT, handle_interrupt);
    disable_input_buffering();

    // Since exactly one condition flag should be set at any given time, set the Z flag
    reg[R_COND] = FL_ZRO;

    // Set the PC starting position
    enum
    {
        PC_START = 0x3000
    };
    reg[R_PC] = PC_START;

    int running = 1;
    while (running)
    {
        // Fetch
        uint16_t instr = mem_read(reg[R_PC]++);
        uint16_t op = instr >> 12;
        
        switch (op)
        {
        case OP_ADD:
            add(instr);
            break;
        case OP_AND:
            and(instr);
            break;
        case OP_NOT:
            not(instr);
            break;
        case OP_BR:
            br(instr);
            break;
        case OP_JMP:
            jmp(instr);
            break;
        case OP_JSR:
            jsr(instr);
            break;
        case OP_LD:
            ld(instr);
            break;
        case OP_LDI:
            ldi(instr);
            break;
        case OP_LDR:
            ldr(instr);
            break;
        case OP_LEA:
            lea(instr);
            break;
        case OP_ST:
            st(instr);
            break;
        case OP_STI:
            sti(instr);
            break;
        case OP_STR:
            str(instr);
            break;
        case OP_TRAP:
            running = trap(instr);
            break;
        case OP_RES:
            res(instr);
            break;
        case OP_RTI:
            rti(instr);
            break;
        default:
            break;
        }
    }

    restore_input_buffering();
}
