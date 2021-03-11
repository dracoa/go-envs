package envs

import (
	"fmt"
	"os"
	"testing"
)

func TestFromUrl(t *testing.T) {
	url := "https://jsonplaceholder.typicode.com/todos/1"
	FromUrl(url)
	if os.Getenv("title") == "" {
		t.Error("cannot get env from url")
	}
	fmt.Println(os.Getenv("title"))
}

func TestFromDir(t *testing.T) {
	FromDir("../test_env")
	FromDir("../test_env2")
	envs := []string{"env1", "env2", "env3"}
	for _, k := range envs {
		if os.Getenv(k) != k {
			t.Errorf("cannot get env from dir %s", k)
		}
	}
	fmt.Println(os.Getenv("env1"))
}

func TestGetDefault(t *testing.T) {
	FromDir("../test_env")
	if GetDefault("env1", "env3") != "env1" {
		t.Errorf("cannot get from env")
	}
	if GetDefault("env4", "env3") != "env3" {
		t.Errorf("cannot get from default")
	}
}
