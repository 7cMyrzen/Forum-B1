package forum

import (
	"forum/lib/types"

	_ "github.com/go-sql-driver/mysql"
)

func OrderTopicsNewestToOldest(posts []types.Posts) []types.Posts {
	for i := 0; i < len(posts); i++ {
		for j := i; j < len(posts); j++ {
			if posts[i].Creation < posts[j].Creation {
				posts[i], posts[j] = posts[j], posts[i]
			}
		}
	}
	return posts
}
