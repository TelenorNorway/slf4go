package caller

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// region test util

type parsed struct {
	t   *testing.T
	rep string
	sig FuncSig
}

func parse(t *testing.T, str string) *parsed {
	return &parsed{
		t:   t,
		rep: str,
		sig: parseName(str),
	}
}

func assertPackage(p *parsed, isModule bool, expected string) {
	if isModule {
		assert.True(p.t, p.sig.IsModule, "%s should be a module", p.rep)
	} else {
		assert.False(p.t, p.sig.IsModule, "%s should not a module", p.rep)
	}
	assert.Equal(p.t, expected, p.sig.Package, "%s package mismatch, expected '%s', but got '%s'", p.rep, expected, p.sig.Package)
}

func assertNoReceiver(p *parsed) {
	assert.False(p.t, p.sig.HasReceiver, "%s should not have a receiver", p.rep)
	assert.False(p.t, p.sig.IsReceiverRef, "%s should not be a receiver ref, expected no receiver", p.rep)
	assert.Equal(p.t, "", p.sig.Receiver, "%s should not have a receiver value", p.rep)
	assert.False(p.t, p.sig.IsReceiverGeneric, "%s should not be a generic receiver, expected no receiver", p.rep)
}

func assertReceiver(p *parsed, isRef, isGeneric bool, receiver string) {
	assert.True(p.t, p.sig.HasReceiver, "%s should be a receiver", p.rep)
	if isRef {
		assert.True(p.t, p.sig.IsReceiverRef, "%s should be a receiver ref", p.rep)
	} else {
		assert.False(p.t, p.sig.IsReceiverRef, "%s should not be a receiver ref", p.rep)
	}
	if isGeneric {
		assert.True(p.t, p.sig.IsReceiverGeneric, "%s should have a generic receiver", p.rep)
	} else {
		assert.False(p.t, p.sig.IsReceiverGeneric, "%s should not be a generic receiver", p.rep)
	}
	assert.Equal(p.t, receiver, p.sig.Receiver, "%s should have a receiver '%s', but got '%s'", p.rep, receiver, p.sig.Receiver)
}

func assertFunction(p *parsed, isGeneric bool, name string) {
	if isGeneric {
		assert.True(p.t, p.sig.IsFunctionGeneric, "%s should be a generic function", p.rep)
	} else {
		assert.False(p.t, p.sig.IsFunctionGeneric, "%s should not be a generic function", p.rep)
	}
	assert.Equal(p.t, name, p.sig.Function, "%s function should be named '%s', but got '%s'", p.rep, name, p.sig.Function)
}

// endregion

func TestParseNameMainDotMain(t *testing.T) {
	p := parse(t, "main.main")

	assertPackage(p, false, "main")
	assertNoReceiver(p)
	assertFunction(p, false, "main")
}

func TestParseNameMainDotMainGeneric(t *testing.T) {
	p := parse(t, "main.main[...]")

	assertPackage(p, false, "main")
	assertNoReceiver(p)
	assertFunction(p, true, "main")
}

func TestParseNameMainDotTestDotTest(t *testing.T) {
	p := parse(t, "main.test.test")

	assertPackage(p, false, "main")
	assertReceiver(p, false, false, "test")
	assertFunction(p, false, "test")
}

func TestParseNameMainDotTestDotTestGeneric(t *testing.T) {
	// This is not something parseName should throw about,
	// because receiver methods cannot be generic. But,
	// parseName does not check for that. So it is still
	// able to parse it.

	p := parse(t, "main.test.test[...]")

	assertPackage(p, false, "main")
	assertReceiver(p, false, false, "test")
	assertFunction(p, true, "test")
}

func TestParseNameMainDotTestGenericDotTest(t *testing.T) {
	p := parse(t, "main.test[...].test")

	assertPackage(p, false, "main")
	assertReceiver(p, false, true, "test")
	assertFunction(p, false, "test")
}

func TestParseNameMainDotTestGenericDotTestGeneric(t *testing.T) {
	// This is not something parseName should throw about,
	// because receiver methods cannot be generic. But,
	// parseName does not check for that. So it is still
	// able to parse it.

	p := parse(t, "main.test[...].test[...]")

	assertPackage(p, false, "main")
	assertReceiver(p, false, true, "test")
	assertFunction(p, true, "test")
}

func TestParseNameMainDotReceiverTestDotTest(t *testing.T) {
	p := parse(t, "main.(*test).test")

	assertPackage(p, false, "main")
	assertReceiver(p, true, false, "test")
	assertFunction(p, false, "test")
}

func TestParseNameMainDotReceiverTestDotTestGeneric(t *testing.T) {
	p := parse(t, "main.(*test).test[...]")

	assertPackage(p, false, "main")
	assertReceiver(p, true, false, "test")
	assertFunction(p, true, "test")
}

func TestParseNameMainDotReceiverTestGenericDotTest(t *testing.T) {
	p := parse(t, "main.(*test[...]).test")

	assertPackage(p, false, "main")
	assertReceiver(p, true, true, "test")
	assertFunction(p, false, "test")
}

func TestParseNameMainDotReceiverTestGenericDotTestGeneric(t *testing.T) {
	p := parse(t, "main.(*test[...]).test[...]")

	assertPackage(p, false, "main")
	assertReceiver(p, true, true, "test")
	assertFunction(p, true, "test")
}

func TestParseNameExampleDotComSlashExampleDotTest(t *testing.T) {
	p := parse(t, "example.com/example.test")

	assertPackage(p, true, "example.com/example")
	assertNoReceiver(p)
	assertFunction(p, false, "test")
}

func TestParseNameExampleDotComSlashExampleDotTestGeneric(t *testing.T) {
	p := parse(t, "example.com/example.test[...]")

	assertPackage(p, true, "example.com/example")
	assertNoReceiver(p)
	assertFunction(p, true, "test")
}

func TestParseNameExampleDotComSlashExampleDotTestDotTest(t *testing.T) {
	p := parse(t, "example.com/example.test.test")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, false, false, "test")
	assertFunction(p, false, "test")
}

func TestParseNameExampleDotComSlashExampleDotTestDotTestGeneric(t *testing.T) {
	// This is not something parseName should throw about,
	// because receiver methods cannot be generic. But,
	// parseName does not check for that. So it is still
	// able to parse it.

	p := parse(t, "example.com/example.test.test[...]")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, false, false, "test")
	assertFunction(p, true, "test")
}

func TestParseNameExampleDotComSlashExampleDotTestGenericDotTest(t *testing.T) {
	p := parse(t, "example.com/example.test[...].test")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, false, true, "test")
	assertFunction(p, false, "test")
}

func TestParseNameExampleDotComSlashExampleDotTestGenericDotTestGeneric(t *testing.T) {
	// This is not something parseName should throw about,
	// because receiver methods cannot be generic. But,
	// parseName does not check for that. So it is still
	// able to parse it.

	p := parse(t, "example.com/example.test[...].test[...]")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, false, true, "test")
	assertFunction(p, true, "test")
}

func TestParseNameExampleDotComSlashExampleDotReceiverTestDotTest(t *testing.T) {
	p := parse(t, "example.com/example.(*test).test")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, true, false, "test")
	assertFunction(p, false, "test")
}

func TestParseNameExampleDotComSlashExampleDotReceiverTestDotTestGeneric(t *testing.T) {
	// This is not something parseName should throw about,
	// because receiver methods cannot be generic. But,
	// parseName does not check for that. So it is still
	// able to parse it.

	p := parse(t, "example.com/example.(*test).test[...]")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, true, false, "test")
	assertFunction(p, true, "test")
}

func TestParseNameExampleDotComSlashExampleDotReceiverTestGenericDotTest(t *testing.T) {
	p := parse(t, "example.com/example.(*test[...]).test")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, true, true, "test")
	assertFunction(p, false, "test")
}

func TestParseNameExampleDotComSlashExampleDotReceiverTestGenericDotTestGeneric(t *testing.T) {
	p := parse(t, "example.com/example.(*test[...]).test[...]")

	assertPackage(p, true, "example.com/example")
	assertReceiver(p, true, true, "test")
	assertFunction(p, true, "test")
}

func BenchmarkParseNameMainDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("main.main")
	}
}

func BenchmarkParseNameMainDotMainGeneric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("main.main[...]")
	}
}

func BenchmarkParseNameMainDotTestDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("main.test.main")
	}
}

func BenchmarkParseNameMainDotTestGenericDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("main.test[...].main")
	}
}

func BenchmarkParseNameMainDotReceiverTestDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("main.(*test).main")
	}
}

func BenchmarkParseNameMainDotReceiverTestGenericDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("main.(*test[...]).main")
	}
}

func BenchmarkParseNameExampleDotComSlashTestDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("example.com/test.main")
	}
}

func BenchmarkParseNameExampleDotComSlashTestDotMainGeneric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("example.com/test.main[...]")
	}
}

func BenchmarkParseNameExampleDotComSlashTestDotTestDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("example.com/test.test.main")
	}
}

func BenchmarkParseNameExampleDotComSlashTestDotTestGenericDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("example.com/test.test[...].main")
	}
}

func BenchmarkParseNameExampleDotComSlashTestDotReceiverTestDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("example.com/test.(*test).main")
	}
}

func BenchmarkParseNameExampleDotComSlashTestDotReceiverTestGenericDotMain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		parseName("example.com/test.(*test[...]).main")
	}
}
