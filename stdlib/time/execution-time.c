#include <stdio.h>
#include <time.h>

int main () {
    clock_t start, end;
    double duration;
    double i = 0;
    start = clock();

    for(i=0; i< 10000000000; i++) {

    }

    end = clock();
    duration = ((double) (end - start));

    printf(" %f",duration);

}