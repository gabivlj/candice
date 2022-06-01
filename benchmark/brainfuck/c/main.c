#include <stdio.h>
// 1.92s user 0.02s system 89% cpu 2.180 total
#define OP_END 0
#define OP_INC_DP 1
#define OP_DEC_DP 2
#define OP_INC_VAL 3
#define OP_DEC_VAL 4
#define OP_OUT 5
#define OP_IN 6
#define OP_JMP_FWD 7
#define OP_JMP_BCK 8

#define SUCCESS 0
#define FAILURE 1

#define PROGRAM_SIZE 4096
#define STACK_SIZE 512
#define DATA_SIZE 65535

#define STACK_PUSH(A) (STACK[SP++] = A)
#define STACK_POP() (STACK[--SP])
#define STACK_EMPTY() (SP == 0)
#define STACK_FULL() (SP == STACK_SIZE)

struct instruction_t
{
    unsigned short operator;
    unsigned short operand;
};

static struct instruction_t PROGRAM[PROGRAM_SIZE];
static unsigned short STACK[STACK_SIZE];
static unsigned int SP = 0;

int compile_bf(FILE *fp)
{
    unsigned short pc = 0, jmp_pc;
    int c;
    while ((c = getc(fp)) != EOF && pc < PROGRAM_SIZE)
    {
        switch (c)
        {
        case '>':
            PROGRAM[pc].operator= OP_INC_DP;
            break;
        case '<':
            PROGRAM[pc].operator= OP_DEC_DP;
            break;
        case '+':
            PROGRAM[pc].operator= OP_INC_VAL;
            break;
        case '-':
            PROGRAM[pc].operator= OP_DEC_VAL;
            break;
        case '.':
            PROGRAM[pc].operator= OP_OUT;
            break;
        case ',':
            PROGRAM[pc].operator= OP_IN;
            break;
        case '[':
            PROGRAM[pc].operator= OP_JMP_FWD;
            if (STACK_FULL())
            {
                return FAILURE;
            }
            STACK_PUSH(pc);
            break;
        case ']':
            if (STACK_EMPTY())
            {
                return FAILURE;
            }
            jmp_pc = STACK_POP();
            PROGRAM[pc].operator= OP_JMP_BCK;
            PROGRAM[pc].operand = jmp_pc;
            PROGRAM[jmp_pc].operand = pc;
            break;
        default:
            pc--;
            break;
        }
        pc++;
    }
    if (!STACK_EMPTY() || pc == PROGRAM_SIZE)
    {
        return FAILURE;
    }
    PROGRAM[pc].operator= OP_END;
    return SUCCESS;
}

int execute_bf()
{
    unsigned short data[DATA_SIZE], pc = 0;
    unsigned int ptr = DATA_SIZE;
    while (--ptr)
    {
        data[ptr] = 0;
    }
    while (PROGRAM[pc].operator!= OP_END)
    {
        switch (PROGRAM[pc].operator)
        {
        case OP_INC_DP:
            ptr++;
            break;
        case OP_DEC_DP:
            ptr--;
            break;
        case OP_INC_VAL:
            data[ptr]++;
            break;
        case OP_DEC_VAL:
            data[ptr]--;
            break;
        case OP_OUT:
            // putchar(data[ptr]);
            break;
        case OP_IN:
            data[ptr] = (unsigned int)getchar();
            break;
        case OP_JMP_FWD:
            if (!data[ptr])
            {
                pc = PROGRAM[pc].operand;
            }
            break;
        case OP_JMP_BCK:
            if (data[ptr])
            {
                pc = PROGRAM[pc].operand;
            }
            break;
        default:
            return FAILURE;
        }
        pc++;
    }
    return ptr != DATA_SIZE ? SUCCESS : FAILURE;
}

int main(int argc, const char *argv[])
{
    int status;
    FILE *fp;
    if (argc != 2 || (fp = fopen("program.b", "r")) == NULL)
    {
        fprintf(stderr, "Usage: %s filename\n", argv[0]);
        return FAILURE;
    }
    status = compile_bf(fp);
    fclose(fp);
    if (status == SUCCESS)
    {
        for (int i = 0; i < 10000; i++)
        {
            status = execute_bf();
            if (status != 0)
                return status;
        }
        printf("exiting...");
        return 0;
    }
    if (status == FAILURE)
    {
        fprintf(stderr, "Error!\n");
    }
    return status;
}