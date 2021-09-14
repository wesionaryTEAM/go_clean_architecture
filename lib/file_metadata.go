package lib

// UploadMetadata metadata received after uploading file
type UploadMetadata struct {
	FieldName string
	URL       string
	FileName  string
	FileUID   string
	Size      int64
}

type UploadedFiles []UploadMetadata

func (f UploadedFiles) GetFile(fieldName string) UploadMetadata {
	for _, file := range f {
		if file.FieldName == fieldName {
			return file
		}
	}
	return UploadMetadata{}
}

func (f UploadedFiles) GetMultipleFiles(fieldName string) []UploadMetadata {
	var files []UploadMetadata
	for _, file := range f {
		if file.FieldName == fieldName {
			files = append(files, file)
		}
	}
	return files
}
