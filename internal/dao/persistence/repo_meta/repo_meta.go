package repoMeta

import (
	"fmt"

	"github.com/qinsheng99/goWeb/internal/model"
	"github.com/qinsheng99/goWeb/library/db"
)

type RepoMeta model.RepoMeta

func (r *RepoMeta) GetRepo(pack string) (data string, err error) {
	err = db.GetPostgresqlDb().
		Model(r).
		Select("repo_name").
		Where(fmt.Sprintf(`'{%s}'&&package_names`, pack)).
		Pluck("repo_name", &data).
		Error
	return
}
