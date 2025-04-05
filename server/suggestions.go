package server

import (
	"math/rand"
	"sort"

	"github.com/aureliengasser/planetocd/server/cache"
)

func getArticleSuggestions(cur *cache.Article) ([]*cache.Article, error) {
	all, err := getArticles(cur.Lang)
	if err != nil {
		return nil, err
	}

	// create list from dict (remove current article)
	list := make([]*cache.Article, len(all)-1)
	idx := 0
	for _, a := range all {
		if a != cur {
			list[idx] = a
			idx++
		}
	}

	// shuffle list
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	// sort list by number of common tags
	sort.Slice(list, func(i, j int) bool {
		return getNumCommonTags(list[i], cur) > getNumCommonTags(list[j], cur)
	})

	res := make([]*cache.Article, 0)

	if len(list) > 0 && getNumCommonTags(list[0], cur) == 0 {
		// if no other article has tags in common, prioritize "recovery"
		sort.Slice(list, func(i, j int) bool {
			return hasRecoveryTag(list[i]) && !hasRecoveryTag(list[j])
		})
	}

	count := 0
	for _, sugg := range list {
		print(sugg.URL)
		res = append(res, sugg)
		count += 1
		if count >= 3 {
			break
		}
	}

	return res, nil
}

func hasRecoveryTag(a *cache.Article) bool {
	if a == nil {
		return false
	}

	if a.Tags == nil {
		return false
	}

	for _, tag := range a.Tags {
		if tag == "recovery" {
			return true
		}
	}

	return false
}

func getNumCommonTags(a *cache.Article, b *cache.Article) int {
	if a == nil || b == nil {
		return 0
	}

	if a.Tags == nil || b.Tags == nil {
		return 0
	}

	count := 0
	for _, tag := range a.Tags {
		for _, tag2 := range b.Tags {
			if tag == tag2 {
				count++
				break
			}
		}
	}
	return count
}
