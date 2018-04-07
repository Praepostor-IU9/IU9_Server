int maxarray(void *base, unsigned long nel, unsigned long width, int (*compare)(void *a, void *b))
{
    unsigned long max = 0, i;
    for (i = 1; i < nel; i++)
    {
        if (compare ((base + width * i), (base + width * max)) > 0) max = i;
    }
    return max;
}
