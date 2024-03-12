#include <stdio.h>
#include <mosquitto.h>

int main() {
  int rc;
  mosquitto_lib_init();
  struct mosquitto *mosq = mosquitto_new("localhost", 1883, NULL);
  if(!mosq){
    fprintf(stderr, "Error: Out of memory.\n");
    return 1;
  }

  rc = mosquitto_connect(mosq, "localhost", 1883, 60);
  if(rc){
    fprintf(stderr, "Error: Connection failed.\n");
    return 1;
  }

  rc = mosquitto_publish(mosq, NULL, "hello/world", "Hello, world!", 12, 0);
  if(rc){
    fprintf(stderr, "Error: Publish failed.\n");
    return 1;
  }

  mosquitto_disconnect(mosq);
  mosquitto_destroy(mosq);
  mosquitto_lib_cleanup();
  return 0;
}
