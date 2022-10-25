package purl

import (
	"errors"
	"regexp"
	"strings"

	repoMeta "github.com/qinsheng99/go-train/internal/dao/persistence/repo_meta"
)

type Purl map[string]interface{}

func PasePurl(purl string) (result Purl, err error) {
	result = make(map[string]interface{})
	var scheme string
	purls := strings.Split(purl, ":")
	if len(purls) > 1 {
		scheme = purls[0]
	}

	if !strings.EqualFold(scheme, "pkg") {
		err = errors.New("purl is missing the required `pkg` scheme component")
		return
	}

	purl1 := strings.Split(purls[1], "/")
	if len(purl1) > 1 {
		var paks []string
		paks = strings.Split(purl1[len(purl1)-1], "?")
		result["type"] = purl1[0]

		if len(paks) == 1 || len(paks) > 1 {
			var meta repoMeta.RepoMeta
			pkg := strings.Split(paks[0], "@")
			if len(pkg) > 1 {
				result["name"] = pkg[0]
				result["version"] = pkg[1]
				result["repoName"], _ = meta.GetRepo(pkg[0])
			}
			if len(paks) > 1 {
				var qualifiers = make(map[string]string)
				pkgs := strings.Split(paks[1], "&")
				for _, v := range pkgs {
					vv := strings.Split(v, "=")
					if len(vv) > 1 {
						qualifiers[vv[0]] = vv[1]
					}
				}
				result["qualifiers"] = qualifiers
			}
		}
	}

	return
}

func (p Purl) GetVersions() (string, bool) {
	if v, ok := p["version"]; ok {
		return strings.Split(v.(string), "-")[0], ok
	}
	return "", false
}

func (p Purl) GetVersion() (string, bool) {
	if v, ok := p["version"]; ok {
		version := strings.Split(v.(string), "-")[0]
		var index int
		index = strings.Index(version, ".Final")
		if index > 0 {
			version = version[:index]
			return version, true
		}

		index = strings.Index(version, ".SNAPSHOT")
		if index > 0 {
			version = version[:index]
			return version, true
		}
		return version, true
	}
	return "", false
}

func (p Purl) GetRealse() (string, bool) {
	if v, ok := p["version"]; ok {
		if len(strings.Split(v.(string), "-")) > 1 {
			return strings.Split(v.(string), "-")[1], ok
		}
		return "", ok
	}
	return "", false
}

func (p Purl) GetName() (string, bool) {
	if q, ok := p["qualifiers"]; ok {
		if pk, qok := q.(map[string]string)["pkgName"]; qok {
			return pk, qok
		}
	}
	if pkg, ok := p["name"].(string); ok {
		return pkg, ok
	}
	return "", false
}

var rawString = regexp.MustCompile(`texlive-[A-z]+-[A-z]*`)

func (p Purl) GetRepo() (string, bool) {
	upstream, ok := p.GetUpstream()
	if !ok {
		return p.GetName()
	}

	version, _ := p.GetVersions()

	repos := strings.Split(upstream, "-"+version)
	if len(repos) > 1 {
		return repos[0], true
	}

	if len(repos) == 1 {
		if strings.Contains(upstream, "bpg-fonts") {
			return "bpg-fonts", true
		}

		if strings.Contains(upstream, "eclipse-license") {
			return "eclipse-license", true
		}

		res := rawString.FindAllStringSubmatch(upstream, -1)
		if len(res) > 0 && len(res[0]) > 0 {
			u := res[0][0]
			if u[len(u)-1] == '-' {
				u = u[:len(u)-1]
			}

			return u, true
		}
		return strings.Split(upstream, "-")[0], true
	}

	return "", false
}

func (p Purl) GetUpstream() (string, bool) {
	if q, ok := p["qualifiers"]; ok {
		if pk, ok := q.(map[string]string)["upstream"]; ok {
			return pk, ok
		}
	}
	return "", false
}

func (p Purl) GetArch() (string, bool) {
	if q, ok := p["qualifiers"]; ok {
		if pk, ok := q.(map[string]string)["arch"]; ok {
			return pk, ok
		}
	}
	return "", false
}
