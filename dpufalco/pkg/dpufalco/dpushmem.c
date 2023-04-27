#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <sys/types.h>
#include <sys/shm.h>
#include <time.h>


#define SLEEP_IN_NANOS (10 * 1000)
#define BUFFER_SIZE 102400

struct timespec ts = {
	.tv_nsec = SLEEP_IN_NANOS,
};

char ret[256] = {0};

int wait_for_buffer() {
    key_t buf_flag_key = ftok("/", 2000);
    if (buf_flag_key == -1) {
        return -1;
    }
	int buf_flag_shmid = shmget(buf_flag_key, 1, IPC_CREAT|0666);
    if (buf_flag_shmid == -1) {
        return -1;
    }
	char *buffer_flag = (char *)shmat(buf_flag_shmid, NULL, 0);

    /* Busy-waiting for the buffer flag */
    while (buffer_flag[0]) {
        nanosleep(&ts, &ts);
    }

    shmdt(buffer_flag);
    return 0;
}

char* read_buffer(int index) {
    // char *ret;
    int i;

    key_t buf_key = ftok("/", 1000);
    if (buf_key == -1) {
        // ret = (char *)malloc(1);
        *ret = 0;
        return ret;
    }
	int buf_shmid = shmget(buf_key, BUFFER_SIZE, IPC_CREAT|0666);
	char *buffer = (char *)shmat(buf_shmid, NULL, 0);

    int start = index<<8;
    // ret = (char *)malloc(256);
    for (i=0; i<255 && buffer[start+i]; ++i) {
        ret[i] = buffer[start+i];
    }
    ret[i] = 0;

    shmdt(buffer);
    return ret;
}

int set_signal() {
    key_t buf_flag_key = ftok("/", 2000);
    if (buf_flag_key == -1) {
        return -1;
    }
	int buf_flag_shmid = shmget(buf_flag_key, 1, IPC_CREAT|0666);
    if (buf_flag_shmid == -1) {
        return -1;
    }
	char *buffer_flag = (char *)shmat(buf_flag_shmid, NULL, 0);

    buffer_flag[0] = 1;

    shmdt(buffer_flag);
    return 0;
}