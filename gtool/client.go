package gtool

import (
	"context"
	"gitee.com/kelvins-io/common/conf"
	"gitee.com/kelvins-io/common/configcenter"
	"gitee.com/kelvins-io/common/file"
	"gitee.com/kelvins-io/common/gtool/grpc_interceptor"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"strconv"
	"strings"
)

type Conn struct {
	ServerName string
	ServerPort uint64

	CertFile       string
	CertServerName string
}

func (c *Conn) GetClientConn() (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile(c.CertFile, c.CertServerName)
	if err != nil {
		return nil, err
	}

	target := c.ServerName + ":" + strconv.Itoa(int(c.ServerPort))
	return grpc.Dial(target, grpc.WithTransportCredentials(creds))
}

func NewConn(serviceName string) *Conn {
	connV1 := conf.NewConfService(serviceName)
	serverName := connV1.GetServerName()
	serverPort := connV1.GetServerPort()
	certPemFullPathV1 := connV1.GetCertPemFullPath()

	conn := &Conn{
		ServerName: serverName,
		ServerPort: serverPort,
	}

	exists, _ := file.IsFileExists(certPemFullPathV1)
	if exists == true {
		conn.CertFile = certPemFullPathV1
		conn.CertServerName = connV1.GetCertServerName()
	} else {
		connV2 := configcenter.NewConfigCenterV2(serviceName)
		certPemFullPathV2, err := connV2.GetCertPemPath()
		if err != nil {
			log.Printf("NewConn.connV2.GetCertPemPath err: %v", err)
			return nil
		}

		serverNames := strings.Split(serverName, "-")
		if len(serverNames) < 1 {
			log.Printf("NewConn.strings.Split len: %d", len(serverNames))
			return nil
		}

		certServerName := serverNames[0]
		conn.CertFile = certPemFullPathV2
		conn.CertServerName = certServerName
	}

	return conn
}

func (c *Conn) GetAPMConn(ctx context.Context) (*grpc.ClientConn, error) {
	creds, err := credentials.NewClientTLSFromFile(c.CertFile, c.CertServerName)
	if err != nil {
		return nil, err
	}

	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(creds))
	opts = append(opts, grpc.WithUnaryInterceptor(
		grpcMiddleware.ChainUnaryClient(
			apmgrpc.NewUnaryClientInterceptor(),
			grpc_interceptor.UnaryCtxHandleGRPC(),
		),
	))
	opts = append(opts, grpc.WithStreamInterceptor(
		grpcMiddleware.ChainStreamClient(
			grpc_interceptor.StreamCtxHandleGRPC(),
		),
	))

	target := c.ServerName + ":" + strconv.Itoa(int(c.ServerPort))
	return grpc.DialContext(
		ctx,
		target,
		opts...,
	)
}
