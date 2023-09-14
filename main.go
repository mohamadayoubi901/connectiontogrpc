package connect

import (
	"context"
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/mohamadayoubi901/connectiontogrpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr  = flag.String("addr", "localhost:50051", "the address to connect to")
	addr2 = flag.String("addr2", "localhost:50052", "the address to connect to")
)

func GEtconfgis(client string, instance string) (string, string, error) {

	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	r, err := c.Getconfigs(ctx, &pb.GetEnvRequest{Instance: instance, Client: client})
	if err != nil {
		return "", "", err
	}
	return r.GetAcctname(), r.GetAcctkey(), err
}
func GetAuth(acname string, acckey string, req *http.Request) (string, error) {

	flag.Parse()
	// Set up a connection to the server.
	conn, err := grpc.Dial(*addr2, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAuhtClientt(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	// a := "GJerLIuhAIrVpCjpWDnopLDCwYe4mt6xVfbrfPj7vNOEdv4A/oQFrgqteyBu5dyKNS4hYE1gCe8S+AStp4pdWA=="
	a2 := req.Header
	res := map[string]*pb.ListOfString{}

	for name, values := range a2 {
		aa3 := pb.ListOfString{}
		for _, valuee := range values {
			aa3.Value = append(aa3.Value, valuee)
		}
		res[name] = &aa3
	}
	r, err := c.Getauthentication(ctx, &pb.GetauthRequest{Accnamee: acname, AccKey: acckey, Reqmethod: req.Method, ReqUrl: req.URL.String(), Reqheader: res})
	if err != nil {
		return "", err

	}
	return r.GetAuthenication(), err
}

const (
	headerAuthorization      = "Authorization"
	headerCacheControl       = "Cache-Control"
	headerContentEncoding    = "Content-Encoding"
	headerContentDisposition = "Content-Disposition"
	headerContentLanguage    = "Content-Language"
	headerContentLength      = "Content-Length"
	headerContentMD5         = "Content-MD5"
	headerContentType        = "Content-Type"
	headerDate               = "Date"
	headerIfMatch            = "If-Match"
	headerIfModifiedSince    = "If-Modified-Since"
	headerIfNoneMatch        = "If-None-Match"
	headerIfUnmodifiedSince  = "If-Unmodified-Since"
	headerRange              = "Range"
	headerUserAgent          = "User-Agent"
	headerXmsDate            = "x-ms-date"
	headerXmsVersion         = "x-ms-version"
)
