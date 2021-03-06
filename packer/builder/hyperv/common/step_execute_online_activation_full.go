package common

import (
	"fmt"
	"bytes"
	"github.com/mitchellh/multistep"
	"github.com/mitchellh/packer/packer"
	"strings"
	"log"
)

type StepExecuteOnlineActivationFull struct {
	Pk string
}

func (s *StepExecuteOnlineActivationFull) Run(state multistep.StateBag) multistep.StepAction {
	ui := state.Get("ui").(packer.Ui)
	comm := state.Get("communicator").(packer.Communicator)

	errorMsg := "Error Executing Online Activation: %s"

	var remoteCmd packer.RemoteCmd
	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)
	var err error
	var stderrString string
	var stdoutString string

	ui.Say("Executing Online Activation Full version...")

	var blockBuffer bytes.Buffer
	blockBuffer.WriteString("{ cscript \"$env:SystemRoot/system32/slmgr.vbs\" /ipk "+ s.Pk +" //nologo }")

	log.Printf("cmd: %s", blockBuffer.String())
	remoteCmd.Command = "-ScriptBlock " + blockBuffer.String()

	remoteCmd.Stdout = stdout
	remoteCmd.Stderr = stderr

	err = comm.Start(&remoteCmd)

	stderrString = strings.TrimSpace(stderr.String())
	stdoutString = strings.TrimSpace(stdout.String())

	log.Printf("stdout: %s", stdoutString)
	log.Printf("stderr: %s", stderrString)

	if len(stderrString) > 0 {
		err = fmt.Errorf(errorMsg, stderrString)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

//	ui.Say(stdoutString)

/*
	blockBuffer.Reset()
	blockBuffer.WriteString("{ cscript \"$env:SystemRoot/system32/slmgr.vbs\" -ato //nologo }")

	log.Printf("cmd: %s", blockBuffer.String())
	remoteCmd.Command = "-ScriptBlock " + blockBuffer.String()

	err = comm.Start(&remoteCmd)

	stderrString = strings.TrimSpace(stderr.String())
	stdoutString = strings.TrimSpace(stdout.String())

	log.Printf("stdout: %s", stdoutString)
	log.Printf("stderr: %s", stderrString)

	if len(stderrString) > 0 {
		err = fmt.Errorf(errorMsg, stderrString)
		state.Put("error", err)
		ui.Error(err.Error())
		return multistep.ActionHalt
	}

	ui.Say(stdoutString)
*/
	return multistep.ActionContinue
}

func (s *StepExecuteOnlineActivationFull) Cleanup(state multistep.StateBag) {
	// do nothing
}
