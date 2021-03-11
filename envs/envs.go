package envs

import (
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func GetDefault(key string, d string) string {
	val := os.Getenv(key)
	if val == "" {
		return d
	}
	return val
}

func FromDir(dir string) {
	envs := make([]string, 0)
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return
	}
	for _, f := range files {
		envs = append(envs, fmt.Sprintf("%s/%s", dir, f.Name()))
	}
	err = godotenv.Load(envs...)
	if err != nil {
		panic(err)
	}
}

func FromUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		result := make(map[string]interface{})
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		err = json.Unmarshal(data, &result)
		if err != nil {
			panic(err)
		}
		for k, s := range result {
			err = os.Setenv(strings.ToUpper(k), fmt.Sprintf("%v", s))
			if err != nil {
				panic(err)
			}
		}
	} else {
		panic(fmt.Sprintf("invalid response code: %d", resp.StatusCode))
	}

}
