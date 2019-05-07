package dbops

import "log"

//通过用户在前台 user->api service ->delete video
// api serveice ->scheduler->write video  deletion   record
//timer
//timer ->runner ->read wvdr ->exec ->delete video  from  folder
func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT  INTO  video_del_rec  (video_id ) VALUES(?)")
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error : %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}
