package ci

import (
	"os"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type Workspace interface {
    Branch() string
    Commit() string
    Dir() string
    Env() []string
}

type workspaceImpl struct {
    branch string
    commit string
    dir string
    env []string
}

func (ws *workspaceImpl) Branch() string {
    return ws.branch
}

func (ws *workspaceImpl) Commit() string {
    return ws.commit
}

func (ws *workspaceImpl) Dir() string {
    return ws.dir
} 

func (ws *workspaceImpl) Env() []string {
    return ws.env
} 

func NewWorkspaceFromGit(root string, url string, branch string) (*workspaceImpl, error) {
	dir, err := os.MkdirTemp(root, "workspace")
	if err != nil {
		return nil, err
	}

	repo, err := git.PlainClone(dir, false, &git.CloneOptions{
		URL:               url,
		ReferenceName:     plumbing.NewBranchReferenceName(branch),
		RecurseSubmodules: git.DefaultSubmoduleRecursionDepth,
		Depth:             1,
	})
	if err != nil {
		return nil, err
	}

	ref, err := repo.Head()
	if err != nil {
		return nil, err
	}

	return &workspaceImpl{
		dir:    dir,
		branch: branch,
		commit: ref.Hash().String(),
		env:    []string{},
	}, nil
}
