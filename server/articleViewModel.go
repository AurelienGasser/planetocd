package server

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/aureliengasser/planetocd/server/cache"
	"github.com/aureliengasser/planetocd/server/likes"
	"github.com/aureliengasser/planetocd/server/viewModel"
)

type articleViewModel struct {
	Article          *cache.Article
	CurrentPageIndex int
	Pagination       *viewModel.Pagination
	Suggestions      []*cache.Article
	LikeURL          *url.URL
	UpdateLikeURL    *url.URL
	Likes            *likes.ArticleLikes
}

// todo: remove lang param
func (vm *articleViewModel) LikesText(lang string) string {
	// add fake likes
	fakeName := "anais"
	if lang == "es" {
		fakeName = "juan"
	}
	if lang == "zh" {
		fakeName = "张伟"
	}
	userNames := append([]string{fakeName}, vm.Likes.UserNames...)
	anonymousLikes := vm.Likes.AnonymousLikes

	return LikesText(lang, userNames, anonymousLikes)
}

func LikesText(lang string, userNames []string, anonymousLikes int) string {
	if len(userNames) > 0 {
		if anonymousLikes == 0 {
			usernamesString := strings.Join(userNames, ", ")
			if len(userNames) > 1 {
				usernamesString = strings.Join(userNames[:len(userNames)-1], ", ") + " " + Translate(lang, "and") + " " + userNames[len(userNames)-1]
			}
			return usernamesString + " " + Translate(lang, "found this article helpful", len(userNames))
		}
		return strings.Join(userNames, ", ") + " " + Translate(lang, "and") + " " + strconv.Itoa(anonymousLikes) + " " + Translate(lang, "other", anonymousLikes) + " " + Translate(lang, "found this article helpful", 2)
	}
	if anonymousLikes == 0 {
		return ""
	}
	return strconv.Itoa(anonymousLikes) + " " + Translate(lang, "people", anonymousLikes) + " " + Translate(lang, "found this article helpful", anonymousLikes)
}
