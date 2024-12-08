package likes

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sort"

	"github.com/aureliengasser/planetocd/server/db"
	"github.com/jackc/pgx/v5"
)

func Save(articleID int, lang string, ip string) (int, int32, error) {
	conn, err := db.GetDbConnection()
	if err != nil {
		fmt.Printf("error while connecting to the DB: %v\n", err)
		return -1, -1, fmt.Errorf("error while connecting to the DB")
	}
	defer conn.Close(context.Background())
	randomNumber := rand.Int32()
	args := pgx.NamedArgs{
		"articleID":    articleID,
		"lang":         lang,
		"ip":           ip,
		"randomNumber": randomNumber,
	}
	rows, err := conn.Query(context.Background(), "INSERT INTO likes (article_id, lang, ip, random_number) values (@articleID, @lang, @ip, @randomNumber) RETURNING id", args)
	if err != nil {
		fmt.Printf("error while inserting like: %v\n", err)
		return -1, -1, fmt.Errorf("error while inserting like")
	}
	rows.Next()
	var likeID int
	if err := rows.Scan(&likeID); err != nil {
		fmt.Printf("error while getting like id: %v\n", err)
		return -1, -1, fmt.Errorf("error while getting like id")
	}
	return likeID, randomNumber, nil
}

func Update(id int, randomNumber int32, username string) error {
	conn, err := db.GetDbConnection()
	if err != nil {
		return err
	}
	defer conn.Close(context.Background())
	_, err = conn.Exec(context.Background(), "UPDATE likes SET username = $1 WHERE id = $2 AND random_number = $3", username, id, randomNumber)
	if err != nil {
		return err
	}
	return nil
}

type ArticleLikes struct {
	UserNames      []string
	AnonymousLikes int
}

func Get(articleID int) (*ArticleLikes, error) {
	conn, err := db.GetDbConnection()
	if err != nil {
		return nil, err
	}
	defer conn.Close(context.Background())
	rows, err := conn.Query(context.Background(), "SELECT username FROM likes WHERE article_id = $1", articleID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	res := ArticleLikes{
		UserNames:      make([]string, 0),
		AnonymousLikes: 0,
	}
	for rows.Next() {
		var username *string
		if err := rows.Scan(&username); err != nil {
			return nil, err
		}
		if username == nil || *username == "" {
			res.AnonymousLikes++
		} else {
			res.UserNames = append(res.UserNames, *username)
		}
	}
	sort.Strings(res.UserNames)
	return &res, nil
}
