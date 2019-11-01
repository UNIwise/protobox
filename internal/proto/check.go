package proto

import "os/exec"

func HasProtoc() bool {
	return cmdCheck("protoc")
}

func HasProtocGo() bool {
	return cmdCheck("protoc-gen-go")

}

func HasProtocTypescript() bool {
	return cmdCheck("protoc-gen-ts")
}

func cmdCheck(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
