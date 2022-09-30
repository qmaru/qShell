package client

import (
	"context"
	"os"
	"time"

	"qShell/grpc/common"
	pb "qShell/grpc/libs"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ClientBasic struct {
	conn          *grpc.ClientConn
	ctx           context.Context
	cancel        context.CancelFunc
	Timeout       int
	ServerAddress string
	Secure        bool
}

func (c *ClientBasic) defaultTimeout() {
	if c.Timeout == 0 {
		c.Timeout = 30
	}
}

func (c *ClientBasic) connect(username, password string) (pb.ShellServiceClient, error) {
	token := common.Token{
		Username: username,
		Password: password,
	}

	c.defaultTimeout()
	cred := insecure.NewCredentials()
	c.ctx, c.cancel = context.WithTimeout(context.Background(), time.Duration(c.Timeout)*time.Second)
	conn, err := grpc.DialContext(c.ctx, c.ServerAddress, grpc.WithTransportCredentials(cred), grpc.WithBlock(), grpc.WithPerRPCCredentials(&token))
	if err != nil {
		return nil, err
	}
	c.conn = conn
	client := pb.NewShellServiceClient(c.conn)
	return client, nil
}

func (c *ClientBasic) close() {
	c.cancel()
	c.conn.Close()
}

func (c *ClientBasic) ServerCheck(username, password string) (string, error) {
	client, err := c.connect(username, password)
	if err != nil {
		return "", err
	}
	defer c.close()

	res, err := client.ServerCheck(c.ctx, &pb.Ping{State: true})
	if err != nil {
		return "", err
	}
	return res.GetMessage(), nil
}

func (c *ClientBasic) ListScripts(username, password string) ([]string, error) {
	client, err := c.connect(username, password)
	if err != nil {
		return nil, err
	}
	defer c.close()

	res, err := client.ListScripts(c.ctx, &pb.ListState{State: true})
	if err != nil {
		return nil, err
	}
	scripts := res.GetScripts()
	return scripts, nil
}

func (c *ClientBasic) RunScript(username, password string, name string, paras []string) (string, error) {
	client, err := c.connect(username, password)
	if err != nil {
		return "", err
	}
	defer c.close()

	res, err := client.RunScript(c.ctx, &pb.RunMeta{ScriptName: name, ScriptParas: paras})
	if err != nil {
		return "", err
	}
	results := res.GetResults()
	return results, nil
}

func (c *ClientBasic) Upload(username, password string, filepath string) error {
	client, err := c.connect(username, password)
	if err != nil {
		return err
	}
	defer c.close()

	fileData, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	fileInfo, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	md5, err := common.CalcMD5(fileData)
	if err != nil {
		return err
	}

	_, err = client.UploadFile(c.ctx, &pb.FileMeta{
		Md5:  md5,
		Name: fileInfo.Name(),
		Size: fileInfo.Size(),
		File: fileData,
	})
	if err != nil {
		return err
	}
	return nil
}
