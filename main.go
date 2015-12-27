package main

import (
  "fmt"
  "flag"
  "errors"
  "os"
  "os/exec"
  "strings"
)

func append_git_suffix(str string) string {
  if strings.HasSuffix(str, ".git") {
    return str
  } else {
    return str + ".git"
  }
}

func append_git_prefix(str string) string {
  if strings.HasPrefix(str, "git@") {
    return str
  } else {
    return "git@" + str
  }
}

func create_git_url(git_url string) string {
  return append_git_prefix(append_git_suffix(strings.Replace(git_url, "/", ":", 1)))
}

func download_project(token string) {
  services := []string{"github.com", "bitbucket.com"} 

  for i := range services {
    if strings.HasPrefix(token, services[i]) {
      raw_path := strings.TrimSpace(os.Getenv("GOGETTER_PATH"))
      if raw_path == "" {
        panic(errors.New("No GOGETTER_PATH specified"))
      }

      var env_path string
      if strings.HasSuffix(raw_path, "/") {
        env_path = raw_path
      } else {
        env_path = raw_path + "/"
      }

      git_url := create_git_url(token)
      splits := strings.Split(token, "/")
  
      if len(splits) > 2 {
        fmt.Println("A")
        fmt.Println(env_path)
        fmt.Println(services[i])
        fmt.Println("B")     

        full_path := strings.Replace(env_path + services[i] + "/" + strings.Join(splits[1:], "/"), ".git", "", 1)

        err := os.MkdirAll(full_path, 0755)
        if err != nil {
          fmt.Println(full_path)
          panic(err)
        } else {
          // create dir
          fmt.Println("create " + full_path)
          os.Chdir(full_path)
          args := []string{"clone", git_url, full_path}
          fmt.Println(args)
          if err := exec.Command("git", args...).Run(); err != nil {
            panic(err)
          }
        }
      }
    }
  } 
}

func main() {
  flag.Parse()

  tokens := flag.Args()
  if len(tokens) > 0 {
    for i := range tokens {
      download_project(tokens[i])
    }
  }
}
