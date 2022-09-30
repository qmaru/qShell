package server

import (
	"context"
	"log"
	"net"
	"os"
	"path/filepath"

	"qShell/grpc/common"
	pb "qShell/grpc/libs"
	"qShell/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

type ServerBasic struct {
	ListenAddress string
	Username      string
	Password      string
}

type ShellService struct {
	pb.UnimplementedShellServiceServer
}

// Interceptor 日志拦截器
func Interceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	// 客户端基本信息
	var clientIP string
	pr, ok := peer.FromContext(ctx)
	if ok {
		clientIP = pr.Addr.String()
	} else {
		clientIP = ""
	}
	// 认证
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "metadata error")
	}

	var username string
	var password string
	if value, ok := md["username"]; ok {
		username = value[0]
	}
	if value, ok := md["password"]; ok {
		password = value[0]
	}

	log.Printf("[%s] %s - %s\n", username, clientIP, info.FullMethod)

	err := common.CheckUser(username, password)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "auth invalid: username=%s, password=%s", username, password)
	}

	// 继续处理请求
	return handler(ctx, req)
}

func (s *ServerBasic) Run() error {
	listener, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(Interceptor),
	}

	server := grpc.NewServer(opts...)
	pb.RegisterShellServiceServer(server, &ShellService{})
	return server.Serve(listener)
}

func (shell *ShellService) ServerCheck(ctx context.Context, in *pb.Ping) (*pb.Pong, error) {
	ip, err := common.GetIP()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if in.GetState() {
		return &pb.Pong{State: true, Message: ip}, nil
	}
	return nil, status.Errorf(codes.InvalidArgument, "invalid request")
}

func (shell *ShellService) ListScripts(ctx context.Context, in *pb.ListState) (*pb.ListResults, error) {
	scripts, err := services.ListScripts()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.ListResults{Scripts: scripts}, nil
}

func (shell *ShellService) RunScript(ctx context.Context, in *pb.RunMeta) (*pb.RunResults, error) {
	name := in.GetScriptName()
	paras := in.GetScriptParas()

	results, err := services.RunScript(name, paras)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.RunResults{Results: results}, nil
}

func (shell *ShellService) UploadFile(ctx context.Context, in *pb.FileMeta) (*pb.FileResults, error) {
	remoteMD5 := in.GetMd5()
	fileData := in.GetFile()

	localMD5, err := common.CalcMD5(fileData)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	if remoteMD5 != localMD5 {
		return nil, status.Errorf(codes.Unavailable, "file hash mismatch")
	}

	root, err := services.LoadFile()
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	name := in.GetName()
	localFile := filepath.Join(root, name)
	err = os.WriteFile(localFile, fileData, 0666)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &pb.FileResults{State: true}, nil
}
