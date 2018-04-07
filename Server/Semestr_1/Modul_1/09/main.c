void revarray(void *base, unsigned long nel, unsigned long width)
{
    unsigned char a;
    unsigned long i, j;
    for (i = 0; i < nel / 2; i++)
    {
        for (j = 0; j < width; j++)
        {
            a = *((char*) (base+i*width+j));
            *((char*) (base+i*width+j)) = *((char*) (base+(nel-i-1)*width+j));
            *((char*) (base+(nel-i-1)*width+j)) = a;
        }
    }
}