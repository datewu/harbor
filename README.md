# Example

```golang
package main

import (
	"log"
	"strconv"
	"time"

	"github.com/datewu/harbor"
)

func main() {
	ep := harbor.NewEndpoint(
		"https://harbor.changeme.please.com",
		"admin",
		"pwd",
		2*time.Hour-15*time.Minute)

	// search projects
	ps, err := ep.SearchProject("wise")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("projects:", ps)

	// search images
	pID := strconv.Itoa(ps[0].ProjectID)
	rs, err := ep.SearchImg(pID, "ant")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("repos:", rs)

	// search images
	ts, err := ep.ListImgTags(rs[0].Name)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("repo:", rs[0].Name, "tags:", ts)

	// check login
	err = ep.Login()
	log.Println("should have no err:", err)
}
```

## Result

```shell
➜  ttt go run main.go
2020/04/15 11:08:10 projects: [{wisedev 16}]
2020/04/15 11:08:11 repos: [{wisedev/ant 16 2}]
2020/04/15 11:08:11 repo: wisedev/ant tags: [{1.9.4-jdk6 269803744} {1.9 269803744}]
2020/04/15 11:08:11 should have no err: <nil>

```
