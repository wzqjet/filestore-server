package db

import (
	mydb "filestore-server/db/mysql"
	"fmt"
)

func OnFileUploadFinished(filehash string, filename string,
	filesize int64, fileaddr string) bool {
	stmt, err := mydb.DBConn().Prepare(
		"insert ignore into tbl_file(file_sha1, file_name, file_size, file_addr, status) values (?,?,?,?,1)")
	if err != nil {
		fmt.Println("Failed to prepare statement,ERR " + err.Error())
		return false
	}
	defer stmt.Close()
	ret, err := stmt.Exec(filehash, filename, filesize, fileaddr)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	rf, err := ret.RowsAffected()
	if err != nil {
		if rf < 0 {
			fmt.Println("File has been uploaded", filehash)
		}
		return true
	}
	return false

}
