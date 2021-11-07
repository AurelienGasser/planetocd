package server

import (
	"math/rand"
	"sort"
	"time"
)

func getArticleSuggestions(cur *article) ([]*article, error) {
	all, err := getArticles(cur.Lang)
	if err != nil {
		return nil, err
	}

	// create list from dict (remove current article)
	list := make([]*article, len(all)-1)
	idx := 0
	for _, a := range all {
		if a != cur {
			list[idx] = a
			idx++
		}
	}

	// shuffle list
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(list), func(i, j int) { list[i], list[j] = list[j], list[i] })

	// sort list by number of common tags
	sort.Slice(list, func(i, j int) bool {
		return getNumCommonTags(list[i], cur) > getNumCommonTags(list[j], cur)
	})

	res := make([]*article, 0)

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

func getNumCommonTags(a *article, b *article) int {
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
