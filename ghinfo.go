package main

import (
  "flag"
  "fmt"
  "strings"
  "time"
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

func makeSpaces(str string) (spaces string) {
  for i:=0; i<len(str); i++ { spaces += " " }
  return
}

func getRepo(input string) (owner, repo string) {
  parts := strings.Split(input, "/")
  l := len(parts)
  // repoName = fmt.Sprintf("%s/%s", parts[l-2], parts[l-1])
  owner, repo = parts[l-2], parts[l-1]
  return
}

func yyyymmdd(t *time.Time) string {
  return t.Format("2006-01-02")
}

func repoDetails(repoInput string) {
  owner, repoName := getRepo(repoInput)
  fmt.Println(makeHeader(fmt.Sprintf("Details for repository %v/%v", owner, repoName)))

  client := octokit.NewClient(nil)

  // client.Repositories() is inexplicably different signature than Users()
  repo, _ := client.Repositories().One(&octokit.RepositoryURL, octokit.M{"owner": owner, "repo": repoName})

  combo := fmt.Sprintf("%v/%v", owner, repoName)
  spaces := makeSpaces(combo)
  fmt.Printf("%v by %v\n\n", repo.Name, owner)
  fmt.Println(repo.Description + "\n")
  fmt.Println("Homepage: " + repo.Homepage + "\n")
  fmt.Printf("%v has been forked %v times and starred %v times.\n", combo, repo.ForksCount, repo.StargazersCount)
  fmt.Printf("%v has %v open issues.\n", spaces, repo.OpenIssues)
  fmt.Printf("%v was created on %s and last updated %s.\n", spaces, yyyymmdd(repo.CreatedAt), yyyymmdd(repo.UpdatedAt))
  fmt.Printf("\nClone URL: %v\n", repo.CloneURL)
}

func usageDetails() {
  fmt.Println("Usage: ghinfo [options] <argv>...")
  fmt.Println("Options:")
  fmt.Println(" -u | -user <username>        Display user details")
  fmt.Println(" -r | -repo <user/repository> Display repo details")
  fmt.Println(" -h | -help                   Help")
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

  spaces := makeSpaces(userName)
  publicGists := "'?'"

  fmt.Printf("     Name: %v\n Location: %v\n    Email: %v\n\n", user.Name, user.Location, user.Email)
  fmt.Printf("%v has shared %v public git repositories and %v gists.\n", userName, user.PublicRepos, publicGists)
  fmt.Printf("%v is followed by %v GitHub users and follows %v users.\n", spaces, user.Followers, user.Following)
  fmt.Printf("%v has been a happy GitHub user since %v.\n", spaces, yyyymmdd(user.CreatedAt))
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
