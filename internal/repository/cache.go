package repository

import (
	"context"
)

type Lover struct {
	Name  string  `json:"name"`
	Score float64 `json:"score"`
}

type Ranking struct {
	Rank  int    `json:"rank"`
	Lover *Lover `json:"lover"`
}

func (r *Repository) GetRank(ctx context.Context, key string) (rankings []*Ranking, err error) {
	result, err := r.rdb.ZRevRangeWithScores(ctx, key, 0, -1).Result()
	if err != nil {
		r.l.Errorf("get rank error: %s", err)
		return
	}

	rankings = make([]*Ranking, len(result))
	for i, item := range result {
		rankings[i] = &Ranking{
			Rank: i + 1,
			Lover: &Lover{
				Name:  item.Member.(string),
				Score: item.Score,
			},
		}
	}

	return
}

//func (r *Repository) UpdateRank(ctx context.Context, id int64) (err error) {
//}
