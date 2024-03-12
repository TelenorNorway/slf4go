package caller

import (
	"fmt"
	"runtime"
	"strings"
)

type FuncSig struct {
	IsModule          bool
	Package           string
	HasReceiver       bool
	IsReceiverRef     bool
	Receiver          string
	IsReceiverGeneric bool
	Function          string
	IsFunctionGeneric bool
}

type Info struct {
	File string
	Line int
	Sig  FuncSig
}

func (i Info) String() string {
	return fmt.Sprintf(`type Info struct {
    Sig  type FuncSig struct {
             IsModule          %t
             Package           %s
             HasReceiver       %t
             IsReceiverRef     %t
             Receiver          %s
             IsReceiverGeneric %t
             Function          %s
             IsFunctionGeneric %t
         }
    File %s
    Line %d
}`,
		i.Sig.IsModule,
		i.Sig.Package,
		i.Sig.HasReceiver,
		i.Sig.IsReceiverRef,
		i.Sig.Receiver,
		i.Sig.IsReceiverGeneric,
		i.Sig.Function,
		i.Sig.IsFunctionGeneric,
		i.File,
		i.Line)
}

func isGenericEnd(line string) (bool, string) {
	l := len(line) - 5
	if l < 0 {
		l = 0
	}
	if line[l:] == "[...]" {
		return true, line[:l]
	}
	return false, line
}

func isReceiverRefF(line string) (bool, string) {
	l := len(line)
	if l > 3 && line[0] == '(' {
		return true, line[2 : len(line)-1]
	}
	return false, line
}

func parseName(line string) (sig FuncSig) {

	if lastSlash := strings.LastIndex(line, "/"); lastSlash > -1 {
		sig.IsModule = true
		l := lastSlash + strings.Index(line[lastSlash:], ".")
		sig.Package = line[:l]
		line = line[l+1:]
	} else {
		l := strings.Index(line, ".")
		sig.Package = line[:l]
		line = line[l+1:]
	}

	sig.IsFunctionGeneric, line = isGenericEnd(line)

	if index := strings.LastIndex(line, "."); index >= 0 {
		sig.Function = line[index+1:]
		line = line[:index]
	} else {
		sig.Function = line
		return
	}

	sig.IsReceiverRef, line = isReceiverRefF(line)
	sig.IsReceiverGeneric, sig.Receiver = isGenericEnd(line)
	sig.HasReceiver = true

	return
}

func Get(n int) *Info {
	pc, file, line, ok := runtime.Caller(n + 1)
	if !ok {
		panic("could not get caller info")
	}
	fn := runtime.FuncForPC(pc)
	return &Info{
		File: file,
		Line: line,
		Sig:  parseName(fn.Name()),
	}
}
