package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

//读
//api->video_id-mysql
//dispatcher->mysql -video_id->datachannel
//executor->datachannel-video_id-> delete videos
func ReadVideoDeletetionRecord(count int) ([]string, error) {
	stmtOut, err := dbConn.Prepare("SELECT video_id FROM video_del_rec LIMIT ?")
	var ids [] string

	if err != nil {
		return ids, err
	}
	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("Query VideoDeletetionRecord error :%v ", err)
		return ids, err

	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	defer stmtOut.Close()
	return ids, nil
}

//写
func DelVideoDeletionRecord(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE  FROM  video_del_rel WHERE video_id =?")
	if err != nil {
		return err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("Deleting VideoDeletetionRecord error :%v ", err)
		return err
	}
	defer stmtDel.Close()
	return nil
}
