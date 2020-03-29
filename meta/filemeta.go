package meta

import "filestore-server/db"

//文件元信息结构
type FileMeta struct {
	FileSha1 string
	FileName string
	FileSize int64
	Location string
	UploadAt string
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

//新增
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

//更新到Db
func UpdateFileMetaDB(fmeta FileMeta) bool {
	return db.OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}

func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	tableFile, err := db.GetFileMeta(fileSha1)
	if err != nil {
		return FileMeta{}, err
	}
	fmeta := FileMeta{
		FileSha1: tableFile.FileHash,
		FileName: tableFile.FileName.String,
		FileSize: tableFile.FileSize.Int64,
		Location: tableFile.FileAddr.String,
	}

	return fmeta, nil
}

//简单演示删除
func RemoveFileMeta(fileSha1 string) {
	delete(fileMetas, fileSha1)
}
