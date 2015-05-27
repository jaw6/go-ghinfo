package main

import (
  "flag"
  "fmt"
  "github.com/octokit/go-octokit/octokit"
)

var userFlag = flag.String("user", "", "User name")

func init() {
  flag.StringVar(userFlag, "u", "", "User name")
}

func usageDetails() {
  fmt.Println("Hello - usage details go here")
}

func userDetails(userName string) {
  fmt.Println("User details for ", userName)

  client := octokit.NewClient(nil)

  url, err := octokit.UserURL.Expand(octokit.M{"user": userName})
  if err != nil  {
    fmt.Println("Error: ", err)
    return
  }

  user, result := client.Users(url).One()
  if result.HasError() {
    fmt.Println("Error: ", err)
    return
  }

  fmt.Printf("%v has shared %v public git repositories and %v gists", userName, user.PublicRepos, user.PublicGists)
}

func main() {
  flag.Parse()
  // fmt.Println(*userFlag)
  switch {
    case *userFlag != "": userDetails(*userFlag)
    default: usageDetails()
  }
}
