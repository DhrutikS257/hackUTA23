#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#define MAX_BUFFER_SIZE 128

typedef struct{
    char *user;
    char *sensor;
    char *time;
    char *threat_level;
}data_t;

void load_data(data_t *threat, char *sensor_triggered)
{
    if(!strcmp(sensor_triggered, "Vibration Sensor triggered"))
        threat->threat_level = "threat";
    else if(!strcmp(sensor_triggered, "Ultra Sonic triggered"))
        threat->threat_level = "alert";
    else
        threat->threat_level = "warning";
    threat->sensor = sensor_triggered;

    time_t current_time;
    struct tm *time_info;

    time(&current_time);
    time_info = localtime(&current_time);
    threat->time = asctime(time_info);
    strtok(threat->time, "\n");
    threat->user = "josego0716@gmail.com";
}

void getAlert(char* data)
{
    char raw_data[MAX_BUFFER_SIZE] = {0};
    FILE *serial = fopen("COM3", "r");
    if (!serial) {
        fprintf(stderr, "Failed to open serial port.\n");
        return;
    }

    data_t *threat = calloc(1, sizeof(data_t));


    while (1) {
        if (fgets(raw_data, MAX_BUFFER_SIZE, serial) != NULL) {
            strtok(raw_data, "\n");
            load_data(threat, raw_data);
            strcat(data, "sensor'");
            strcat(data, threat->sensor);
            strcat(data, "'time'");
            strcat(data, threat->time);
            strcat(data, "'alert'");
            strcat(data, threat->threat_level);
            strcat(data, "'email'");
            strcat(data, threat->user);
            strcat(data, "'");
            fclose(serial);
            return;
        }
    }

    return;
}