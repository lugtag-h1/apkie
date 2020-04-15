package apkie

import (
	"archive/zip"
	"bytes"
	"encoding/xml"
	"errors"
	"github.com/avast/apkparser"
	str "github.com/lugtag-h1/apkie/helper/strings"
	"io"
	"strings"
)

var ErrManifestNotFound = errors.New("AndroidManifest.xml not found")

func ReadAndroidManifest(filename string) (*manifest, error) {
	var (
		manifest manifest
		buff bytes.Buffer
	)

	r, _ := zip.OpenReader(filename)
	defer r.Close()

	for _, file := range r.File {
		if file.Name == "AndroidManifest.xml" {
			rc, err := file.Open()
			if err != nil {
				return nil, err
			}

			err = apkparser.ParseXml(rc, xml.NewEncoder(&buff), nil)

			if err != nil && err == apkparser.ErrPlainTextManifest {
				io.Copy(&buff, rc)
			} else if err != nil {
				return nil, err
			}

			xml.Unmarshal(buff.Bytes(), &manifest)
			return &manifest, nil
		}
	}

	return nil, ErrManifestNotFound
}

func FindExportedComponents(manifest *manifest, name string) (ret []ComponentInfo) {
	for _, components := range [][]component{
		manifest.Application.Activity,
		manifest.Application.ActivityAlias,
		manifest.Application.Provider,
		manifest.Application.Receiver,
		manifest.Application.Service,

	} {
		for _, comp := range components {
			if strings.EqualFold(comp.GetName(), name) {
				return []ComponentInfo{{
					Name:       comp.GetName(),
					IsExported: comp.IsExported(),
				}}
			}

			if strings.EqualFold(comp.GetAuthorities(), name) {
				return []ComponentInfo{{
					Name:       comp.GetAuthorities(),
					IsExported: comp.IsExported(),
				}}
			}
		}

		// Partial matches
		for _, comp := range components {
			if str.ContainsFold(comp.GetName(), name) {
				ret = append(ret, ComponentInfo{
					Name:       comp.GetName(),
					IsExported: comp.IsExported(),
				})
			}

			if str.ContainsFold(comp.GetAuthorities(), name) {
				ret = append(ret, ComponentInfo{
					Name:       comp.GetAuthorities(),
					IsExported: comp.IsExported(),
				})
			}
		}
	}

	return ret
}