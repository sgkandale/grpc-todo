package item

import (
	"context"
	"log"

	"withMazeYard/database/primaryDB"

	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (item *Item) InsertInDB() error {
	tx, err := primaryDB.DevportDB.Begin(context.TODO())
	if err != nil {
		log.Println("error begining transaction : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	b := &pgx.Batch{}

	b.Queue(`insert into `+primaryDB.OrganisationTable+`
			(orgid, orgName, createdon, lastupadtedon, 
			orgstatus) 
			values($1, $2, $3, $4, $5)`,
		org.OrgID, org.Name, org.CreatedOn, org.LastUpdatedOn, org.Status)

	b.Queue(`insert into `+primaryDB.MembersTable+`
			(orgid, userid, useraccesslevel, joinedon) 
			values($1, $2, $3, $4)`,
		org.OrgID, org.CreatorID, UserAccessLevel_Creator, org.LastUpdatedOn)

	batchResults := tx.SendBatch(context.TODO(), b)

	var qerr error
	var rows pgx.Rows
	for qerr == nil {
		rows, qerr = batchResults.Query()
		rows.Close()
	}

	err = tx.Commit(context.TODO())
	if err != nil {
		log.Println("error inserting in db : ", err)
		return status.Error(codes.Internal, "something went wrong")
	}

	return nil
}