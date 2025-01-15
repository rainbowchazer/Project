package enums

type FilePathPrefix string

const (
	NewsImages                 FilePathPrefix = "/news/"
	ProductFilesInstructions   FilePathPrefix = "/product/files/instructions/"
	ProductFilesCertifications FilePathPrefix = "/product/files/certifications/"
	ProductImages              FilePathPrefix = "/product/images/"
	CategoryImage              FilePathPrefix = "/category/"
	PhotoGallery               FilePathPrefix = "/photos/"
	SEO                        FilePathPrefix = "/seo/"
	TmpFiles                   FilePathPrefix = "/tmp/"
)

func (f FilePathPrefix) String() string {
	return string(f)
}
