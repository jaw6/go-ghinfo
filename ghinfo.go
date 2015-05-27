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
  fmt.Println("User details for " + userName)

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

  spaces := ""
  for i:=0; i<len(userName); i++ { spaces += " " }
  publicGists := 0
  fmt.Printf("%v has shared %v public git repositories and %v gists.\n", userName, user.PublicRepos, publicGists)
  fmt.Printf("%v is followed by %v GitHub users and follows %v users.\n", spaces, user.Followers, user.Following)
  fmt.Printf("%v has been a happy GitHub user since %v.\n", spaces, user.CreatedAt.Format("2006-01-02"))
}

func main() {
  flag.Parse()
  // fmt.Println(*userFlag)
  switch {
    case *userFlag != "": userDetails(*userFlag)
    default: usageDetails()
  }
}
