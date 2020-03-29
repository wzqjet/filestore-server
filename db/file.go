package db

import (
	"database/sql"
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

type TableFile struct {
	FileHash string
	FileName sql.NullString
	FileSize sql.NullInt64
	FileAddr sql.NullString
}

func GetFileMeta(filehash string) (*TableFile, error) {
	stmt, err := mydb.DBConn().Prepare(
		"select file_sha1,file_name,file_size,file_addr from tbl_file where file_sha1 = ? and status = 1 limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	defer stmt.Close()
	file := TableFile{}
	err = stmt.QueryRow(filehash).Scan(&file.FileHash, &file.FileName, &file.FileSize, &file.FileAddr)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &file, nil
}
