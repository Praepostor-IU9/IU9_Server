unsigned long binsearch(unsigned long nel, int (*compare)(unsigned long i))
{
    unsigned long i, a = 0, b = nel-1;
    while (a <= b)
    {
        i = (a + b) / 2;
        if (compare(i) == 0) return i;
        if (compare(i) == 1) b = i - 1;
        if (compare(i) == -1) a = i + 1;
    }
    return nel;
}
