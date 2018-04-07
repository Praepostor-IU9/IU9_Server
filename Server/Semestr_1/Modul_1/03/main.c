
#include <stdio.h>

int main()
{
    unsigned long long x, a, b, c;
    scanf("%llu", &x);
    a = 1;
    b = 1;
    c = a + b;
    while (b + c <= x)
    {
        c += b;
        a = b;
        b = c - b;
    }
    if (x > 1)
    {
        while (c > 1)
        {
            if (x >= c)
            {
                x -= c;
                printf("%c", '1');
            }
            else printf("%c", '0');
            c = b;
            b = a;
            a = c - b;
        }
    }
    printf("%i", x > 0);
    return 0;
}