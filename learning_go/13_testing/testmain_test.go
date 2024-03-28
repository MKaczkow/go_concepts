package testmain

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var testTime time.Time

func TestMain(m *testing.M) {
	fmt.Println("Tests started")
	testTime = time.Now()
	exitVal := m.Run()
	fmt.Println("Tests finished in", time.Since(testTime))
	os.Exit(exitVal)
}

func TestFirst(t *testing.T) {
	t.Log("First test", testTime)
}

func TestSecond(t *testing.T) {
	t.Log("Second test", testTime)
}
