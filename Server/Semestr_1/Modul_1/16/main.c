int strdiff(char *a, char *b)
{
    int i, j, n;
    if (strlen(a) < strlen(b))
        n = strlen(b);
        else n = strlen(a);
    for(i = 0; i < n; i++)
    {
        for(j = 31; j >= 24; j--)
        {
            if ((a[i]<<j) != (b[i]<<j)) return i*8+31-j;
        }
    }
    return -1;
}