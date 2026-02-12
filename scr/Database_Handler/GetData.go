package databasehandler

import (
	"context"
	"database/sql"
	"fmt"

	variables "github.com/anicse27/Website_Down_Checker/scr/Variables"
)

func GetAllURLs(DB *sql.DB, ctx context.Context) ([]variables.URL_Data, error) {
	Data := variables.URL_Data{}
	All_Data := []variables.URL_Data{}
	Rows, err := DB.QueryContext(ctx, "SELECT * FROM Websites")
	for Rows.Next() {
		Rows.Scan(&Data.Id, &Data.SiteName, &Data.SiteURL)
		All_Data = append(All_Data, Data)
	}
	if err != nil {
		fmt.Println(err)
		return All_Data, err
	}

	return All_Data, nil
}
