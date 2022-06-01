//./program  4.97s user 0.03s system 94% cpu 5.271 total
#include <stdio.h>
#include <stdlib.h>

int **getMatrix()
{
    int **matrix = (int **)malloc(sizeof(int *) * 9);
    for (int i = 0; i < 9; ++i)
    {
        matrix[i] = (int *)malloc(sizeof(int) * 9);
        for (int j = 0; j < 9; ++j)
        {
            matrix[i][j] = i + j;
        }
    }
    return matrix;
}

void sum(int **m1, int **m2)
{
    for (int i = 0; i < 9; ++i)
    {
        for (int j = 0; j < 9; ++j)
        {
            m1[i][j] = m1[i][j] + m2[i][j];
        }
    }
}

void freeMatrix(int **m1)
{
    for (int i = 0; i < 9; ++i)
    {
        free(m1[i]);
    }

    free(m1);
}

int main()
{
    for (int i = 0; i < 10000000; ++i)
    {
        int **m1 = getMatrix();
        int **m2 = getMatrix();
        sum(m1, m2);
        freeMatrix(m1);
        freeMatrix(m2);
    }
    return 0;
}