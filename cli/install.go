package cli

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"context"
	"log"
	"time"

	pb "github.com/lmuench/plg/rpc/plg"
	"google.golang.org/grpc"
)

const (
	rpcSrvAddr = "localhost:50051"
)

func Install(plug string) error {
	absObjPath, err := compile(plug)
	if err != nil {
		fmt.Print("error: ", plug, " could not be compiled (", err, ")\n")
	}
	// TODO: read iface string from metadata
	return register("MockIFace", absObjPath)
}

func compile(name string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	obj := name + ".so"
	src := name + ".go"
	cmd := exec.Command("go", "build", "-buildmode=plugin", "-o="+obj, src)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return "", err
	}

	absObjPath := filepath.Join(wd, obj)
	return absObjPath, nil
}

func register(iface, absObjPath string) error {
	conn, err := grpc.Dial(rpcSrvAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewRegistryClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.RegisterPlugin(ctx, &pb.Plugin{
		Iface:      iface,
		AbsObjPath: absObjPath,
	})
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	log.Printf("response: %s", res.Msg)
	return nil
}
