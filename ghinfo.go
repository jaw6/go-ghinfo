package main

import (
  "flag"
  "fmt"
  "strings"
  "github.com/octokit/go-octokit/octokit"
)

var userFlag = flag.String("user", "", "User name")
var repoFlag = flag.String("repo", "", "Repo name or URL")

func init() {
  flag.StringVar(userFlag, "u", "", "User name")
  flag.StringVar(repoFlag, "r", "", "Repo name or URL")
}

func makeHeader(head string) (header string) {
  hr   := ""
  for i:=0; i<len(head); i++ { hr += "=" }
  header = fmt.Sprintf(head + "\n" + hr + "\n")
  return
}

func getRepo(input string) (owner, repo string) {
  parts := strings.Split(input, "/")
  l := len(parts)
  // repoName = fmt.Sprintf("%s/%s", parts[l-2], parts[l-1])
  owner, repo = parts[l-2], parts[l-1]
  return
}

func repoDetails(repoInput string) {
  owner, repo := getRepo(repoInput)
  fmt.Println(makeHeader(fmt.Sprintf("Details for repository %v/%v", owner, repo)))
}

func usageDetails() {
  fmt.Println("Hello - usage details go here")
}

func userDetails(userName string) {
  fmt.Println(makeHeader(fmt.Sprintf("Details for GitHub user %v", userName)))

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
  publicGists := "'?'"

  fmt.Printf("     Name: %v\n Location: %v\n    Email: %v\n\n", user.Name, user.Location, user.Email)
  fmt.Printf("%v has shared %v public git repositories and %v gists.\n", userName, user.PublicRepos, publicGists)
  fmt.Printf("%v is followed by %v GitHub users and follows %v users.\n", spaces, user.Followers, user.Following)
  fmt.Printf("%v has been a happy GitHub user since %v.\n", spaces, user.CreatedAt.Format("2006-01-02"))
}

func main() {
  flag.Parse()
  // fmt.Println(*userFlag)
  switch {
    case *userFlag != "": userDetails(*userFlag)
    case *repoFlag != "": repoDetails(*repoFlag)
    default: usageDetails()
  }
}
