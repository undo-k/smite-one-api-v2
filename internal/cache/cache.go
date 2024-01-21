package cache

import (
	"database/sql"
	"github.com/undo-k/smite-one-api-v2/internal/models"
	"github.com/undo-k/smite-one-api-v2/internal/tools"
	"strings"
)

func CreateGodCache() (map[string]models.God, error) {
	db := tools.OpenConnection()

	rows, err := db.Query("select g.name, role, win_rate, pick_rate, ban_rate,\n       (select string_agg(i.name, ',')\n        from smiteinfo_god_top_items hi\n        join public.smiteinfo_item i on hi.item_id = i.id\n        where hi.god_id = g.name) as hot_items,\n       (select string_agg(i.name, ',')\n        from smiteinfo_god_lr_top_items ti\n        join public.smiteinfo_item i on ti.item_id = i.id\n        where ti.god_id = g.name) as top_items\n    from smiteinfo_god g\ngroup by g.name")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	defer db.Close()

	var godDetails = make(map[string]models.God)

	for rows.Next() {
		var god models.God
		var hotItems sql.NullString
		var topItems sql.NullString
		err := rows.Scan(&god.Name, &god.Role, &god.WinRate, &god.PickRate, &god.BanRate, &hotItems, &topItems)
		if err != nil {
			return nil, err
		}
		if hotItems.Valid {
			god.HotItems = strings.Split(hotItems.String, ",")
		} else {
			god.HotItems = []string{}
		}

		if topItems.Valid {
			god.TopItems = strings.Split(topItems.String, ",")
		} else {
			god.TopItems = []string{}
		}

		godDetails[god.Name] = god
	}

	return godDetails, nil
}
