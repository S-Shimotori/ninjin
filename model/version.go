package model

type version struct {
	Short shortVersion
	ProductBuild productBuildVersion
}

func NewVersion(shortVersionString string, productBuildVersionString string) (version, error) {
	var v version

	short, shortError := NewShortVersion(shortVersionString)
	if shortError != nil {
		return v, shortError
	}

	productBuild, productBuildError := NewProductBuildVersion(productBuildVersionString)
	if productBuildError != nil {
		return v, productBuildError
	}

	return version{Short: short, ProductBuild: productBuild}, nil
}

func newVersionWithoutError(shortVersionString string, productBuildVersionString string) version {
	v, _ := NewVersion(shortVersionString, productBuildVersionString)
	return v
}

func Less(v0, v1 version) bool {
	if !EqualsForShortVersion(v0.Short, v1.Short) {
		return LessForShortVersion(v0.Short, v1.Short)
	} else {
		return LessForProductBuildVersion(v0.ProductBuild, v1.ProductBuild)
	}
}
