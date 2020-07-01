#include <stdio.h>
#include <sys/types.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <string.h>

#define SERVER_PORT 30001
#define BUFF_LEN 512
#define SERVER_IP "127.0.0.1"
 
/*
	{"ModuleName":"goclient","OccureTime":"2019-11-03T03:43:28.806255332Z",
	"LogType":0,"LogLevel":1,"Position":"0/1/0","ErrCode":12,
	"LogDesc":"go client test","Context":""}
*/
void udp_msg_sender(int fd, struct sockaddr* dst)
{
    int i;
	int sendlen = 0;
    socklen_t len;
    struct sockaddr_in src;
    char buf[BUFF_LEN] = {0};

	strcpy(buf,"{\"ModuleName\":\"c-clent\",\"OccureTime\":\"2019-11-03T03:43:28.806255332Z\",\"LogType\":0,\"LogLevel\":1,\"Position\":\"0-0\",\"ErrCode\":102,\"LogDesc\":\"c test\",\"Context\":\" \"}");
	
    for(i=0;i<3;i++)
    {
        len = sizeof(*dst);
        
        sendlen = sendto(fd, buf, strlen(buf), 0, dst, len);
		fprintf(stderr,"family:%d,server:%x.%x.%x.%x,port:%x.%x",dst->sa_family,dst->sa_data[2], \
		        dst->sa_data[3],dst->sa_data[4],dst->sa_data[5],dst->sa_data[0],dst->sa_data[1]);
		fprintf(stderr,"send len %d\n",sendlen);
        /*sleep(1);  һ�뷢��һ����Ϣ*/
    }

    /*发告警*/
	strcpy(buf,"{\"ModuleName\":\"acl\",\"OccureTime\":\"2019-11-03T03:43:28.806255332Z\",\"LogType\":3,\"LogLevel\":1,\"Position\":\"0-0\",\"ErrCode\":12,\"LogDesc\":\"[alarm:occur]-c test\",\"Context\":\" \"}");
    for(i=0;i<3;i++)
    {
        len = sizeof(*dst);
        
        sendlen = sendto(fd, buf, strlen(buf), 0, dst, len);
		fprintf(stderr,"family:%d,server:%x.%x.%x.%x,port:%x.%x",dst->sa_family,dst->sa_data[2], \
		        dst->sa_data[3],dst->sa_data[4],dst->sa_data[5],dst->sa_data[0],dst->sa_data[1]);
		fprintf(stderr,"send len %d\n",sendlen);
        /*sleep(1);  һ�뷢��һ����Ϣ*/
    }
	strcpy(buf,"{\"ModuleName\":\"acl\",\"OccureTime\":\"2019-11-03T03:43:28.806255332Z\",\"LogType\":3,\"LogLevel\":1,\"Position\":\"0-0\",\"ErrCode\":12,\"LogDesc\":\"[alarm:clear]-c test\",\"Context\":\" \"}");
    for(i=0;i<3;i++)
    {
        len = sizeof(*dst);
        
        sendlen = sendto(fd, buf, strlen(buf), 0, dst, len);
		fprintf(stderr,"family:%d,server:%x.%x.%x.%x,port:%x.%x",dst->sa_family,dst->sa_data[2], \
		        dst->sa_data[3],dst->sa_data[4],dst->sa_data[5],dst->sa_data[0],dst->sa_data[1]);
		fprintf(stderr,"send len %d\n",sendlen);
        /*sleep(1);  һ�뷢��һ����Ϣ*/
    }
	strcpy(buf,"{\"ModuleName\":\"qos\",\"OccureTime\":\"2019-11-03T03:43:28.0Z\",\"LogType\":3,\"LogLevel\":1,\"Position\":\"0-0\",\"ErrCode\":3,\"LogDesc\":\"[alarm:occur]-c test\",\"Context\":\" \"}");
    for(i=0;i<3;i++)
    {
        len = sizeof(*dst);
        
        sendlen = sendto(fd, buf, strlen(buf), 0, dst, len);
		fprintf(stderr,"family:%d,server:%x.%x.%x.%x,port:%x.%x",dst->sa_family,dst->sa_data[2], \
		        dst->sa_data[3],dst->sa_data[4],dst->sa_data[5],dst->sa_data[0],dst->sa_data[1]);
		fprintf(stderr,"send len %d\n",sendlen);
        /*sleep(1);  һ�뷢��һ����Ϣ*/
    }

}
 
 /*
     client:
             socket-->sendto-->revcfrom-->close
 */
 
int main(int argc, char* argv[])
{
    int client_fd;
    struct sockaddr_in ser_addr;

    client_fd = socket(AF_INET, SOCK_DGRAM, 0);
    if(client_fd < 0)
    {
        printf("create socket fail!\n");
        return -1;
    }
 
    memset(&ser_addr, 0, sizeof(ser_addr));
    ser_addr.sin_family = AF_INET;
    ser_addr.sin_addr.s_addr = htonl(0x7f000001);/*inet_addr(SERVER_IP);*/
    /*ser_addr.sin_addr.s_addr = htonl(INADDR_ANY);  ע��������ת��*/
    ser_addr.sin_port = htons(SERVER_PORT);  /*ע��������ת��*/
 
    udp_msg_sender(client_fd, (struct sockaddr*)&ser_addr);

    /*close(client_fd);*/

    return 0;
 }