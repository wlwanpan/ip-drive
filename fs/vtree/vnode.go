package vtree

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"sync"

	"github.com/orbit-drive/orbit-drive/common"
	"github.com/orbit-drive/orbit-drive/fs/api"
	"github.com/orbit-drive/orbit-drive/fs/db"
	"github.com/orbit-drive/orbit-drive/fs/pb"
)

const (
	// FileCode represents a file
	FileCode = iota
	// DirCode represents a dir
	DirCode = iota
)

var (
	// ErrVNodeNotFound is returned when a vnode is missing a Link.
	ErrVNodeNotFound = errors.New("vnode does not exist")

	// ErrNotADir is returned when accessing the links of a file type vnode.
	ErrNotADir = errors.New("file does not have any links")
)

// VNode represents a file structure where each node can be (i) a dir (ii) a file.
// If is a file, Source links to the ipfs hash saved on the network.
type VNode struct {
	// Id is generated from the absolute path and refers to the key used to save to leveldb.
	ID []byte `json:"_id"`

	// Path holds the absolute path in the os file system <- Need to compress to relative path.
	Path string `json:"path"`

	// Type represents if the vnode is a file or dir.
	Type int `json:"type"`

	// Links refers all children vnode in the dir.
	Links []*VNode `json:"links"`

	// Source refers to the ipfs hash generated by the network.error
	Source *db.Source `json:"source"`
}

// PPrint pretty print vnode as json to console.
func (vn *VNode) PPrint() {
	data, _ := json.MarshalIndent(vn, "", "	")
	log.Println(common.ToStr(data))
}

// GetID parse the vtree id to string and returns.
func (vn *VNode) GetID() string {
	return common.ToStr(vn.ID)
}

// SetAsDir sets the vnode type to a dir.
func (vn *VNode) SetAsDir() {
	vn.Type = DirCode
}

// IsDir returns true of vnode is of type dircode.
func (vn *VNode) IsDir() bool {
	return vn.Type == DirCode
}

// SetAsFile sets the vnode type to a file.
func (vn *VNode) SetAsFile() {
	vn.Type = FileCode
}

// SetSource sets the vnode source to the provided source.
func (vn *VNode) SetSource(s *db.Source) {
	vn.Source = s
}

// SaveSource upload a file path to the ipfs network and
// save the return hash as the source of the vnode.
func (vn *VNode) SaveSource() error {
	// If ipfs hash empty, then upload to network.
	if !vn.Source.IsUploaded() {
		s, err := api.UploadFile(vn.Path)
		if err != nil {
			return err
		}
		vn.Source.SetSrc(s)
	}
	return vn.Source.Save(vn.ID)
}

// UpdateSource validates and updates source if given source file differ from current source.
func (vn *VNode) UpdateSource(source *db.Source) error {
	if vn.Source.IsSame(source) {
		return nil
	}
	vn.Source = source
	return vn.SaveSource()
}

// GenChildID returns a hash from the current vnode id and the given path.
func (vn *VNode) GenChildID(p string) []byte {
	i := append(vn.ID, p...)
	return common.HashStr(common.ToStr(i))
}

// NewVNode initialize and returns a new VNode under current vnode.
func (vn *VNode) NewVNode(path string) *VNode {
	i := append(vn.ID, path...)
	n := &VNode{
		ID:     common.HashStr(common.ToStr(i)),
		Path:   path,
		Links:  []*VNode{},
		Source: db.NewSource(path),
	}
	vn.Links = append(vn.Links, n)
	return n
}

// PopulateNodes read a path and populate the its links given
// the path is a directory else creates a file node.RemoveFromWatchList
func (vn *VNode) PopulateNodes(s db.Sources) error {
	files, err := ioutil.ReadDir(vn.Path)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	for _, f := range files {
		abspath := filepath.Join(vn.Path, f.Name())
		if common.IsHidden(abspath) {
			continue
		}
		nn := vn.NewVNode(abspath)
		if f.IsDir() {
			nn.SetAsDir()
			nn.PopulateNodes(s)
			continue
		}

		source := s.ExtractSource(nn.GetID())
		if nn.Source.IsSame(source) {
			nn.SetSource(source)
			continue
		}
		wg.Add(1)
		go func(vn *VNode) {
			vn.SaveSource()
			wg.Done()
		}(nn)
	}

	wg.Wait()
	return nil
}

// FindChildAt perform a full traversal to look a vnode from a given path.
func (vn *VNode) FindChildAt(path string) (*VNode, error) {
	rel, err := filepath.Rel(vn.Path, path)
	if err != nil || rel == "." {
		return vn, err
	}
	steps := strings.Split(rel, "/")
	return vn.traverse(steps)
}

// FindChild look for a given id from its Links (1 level).
func (vn *VNode) FindChild(i []byte) (*VNode, error) {
	if vn.Type == FileCode {
		return vn, ErrNotADir
	}

	for _, n := range vn.Links {
		if bytes.Equal(n.ID, i) {
			return n, nil
		}
	}
	return vn, ErrVNodeNotFound
}

// UnlinkChild traverse a VTree and remove the VNode at the given path.
func (vn *VNode) UnlinkChild(path string) error {
	return nil
}

// traverse traverse a VNode 1 level at a time down the tree.
func (vn *VNode) traverse(steps []string) (*VNode, error) {
	if len(steps) == 0 {
		return vn, nil
	}

	for _, step := range steps {
		p := filepath.Join(vn.Path, step)
		i := vn.GenChildID(p)

		link, err := vn.FindChild(i)
		if err != nil {
			return vn, err
		}
		vn = link
	}
	return vn, nil
}

// ToProto parse a vtree to protobuf.
func (vn *VNode) ToProto() *pb.FSNode {
	var wg sync.WaitGroup
	pbNode := &pb.FSNode{
		ID:     vn.ID,
		Path:   vn.Path,
		Source: vn.Source.Src,
		Links:  []*pb.FSNode{},
	}

	for _, vnode := range vn.Links {
		wg.Add(1)
		go func(vnode *VNode) {
			pbNode.Links = append(pbNode.Links, vnode.ToProto())
			wg.Done()
		}(vnode)
	}
	wg.Wait()
	return pbNode
}

// AllDirPaths traverse the vnode links and returns a slice of all the child dirpath.
func (vn *VNode) AllDirPaths() []string {
	if !vn.IsDir() {
		return []string{}
	}
	var wg sync.WaitGroup
	dirPaths := []string{vn.Path}
	for _, vnode := range vn.Links {
		wg.Add(1)
		go func(vnode *VNode) {
			paths := vnode.AllDirPaths()
			dirPaths = append(dirPaths, paths...)
			wg.Done()
		}(vnode)
	}
	wg.Wait()
	return dirPaths
}
