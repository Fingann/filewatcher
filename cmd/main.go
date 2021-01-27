package main

import (
	"os"
	"path/filepath"
	"time"

	"github.com/Fingann/filewatcher/models"
	"github.com/Fingann/filewatcher/notify"
)

var db models.EntityList

func main() {

	db, err := buildList("/home/sf/projects/filewatcher/testFolder")

	if err != nil {
		return
	}

	for range time.Tick(time.Second * 5) {
		newFiles, err := buildList("/home/sf/projects/filewatcher/testFolder")
		if err != nil {
			return
		}
		i, j := 0, 0
		for ; i < len(db); i++ {
			// remaining files are deleted
			if j >= len(newFiles) {
				notify.Notify("FileWatcher", "Deleted", db[i].Path(), "")
				continue
			}

			//Check if equal
			if db[i].Path() == newFiles[j].Path() {
				//notify.Notify("FileWatcher", "Same", db[i].Path(), "")
				j++
				continue
			}

			// Check if added
			tmpj := j
			added := false
			for ; j < len(newFiles); j++ {
				if db[i].Path() != newFiles[j].Path() {
					continue
				}
				added = true
				for k := i; k < j; k++ {
					notify.Notify("FileWatcher", "Added", newFiles[k].Path(), "")
				}
				j++
				break
			}
			if !added {
				notify.Notify("FileWatcher", "Deleted", db[i].Path(), "")
				j = tmpj
			}

		}
		// set last read as db
		db = newFiles
	}

}

func checkDirectory() {

}

func buildList(path string) (models.EntityList, error) {
	list := models.EntityList{}

	err := filepath.Walk(path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() {
				return err
			}
			list = append(list, models.NewEntity(info.Name(), info.Size(), info.IsDir(), path))
			//fmt.Println(path, info.Size(), info.IsDir())
			//notify.Notify("FileWatcher", "Entity found", path, "")
			return err
		})

	return list, err
}
