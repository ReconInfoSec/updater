package tuf

import (
	"encoding/json"
	"fmt"
	"os"
	"path"

	"github.com/pkg/errors"
)

func (r *localRepo) root(opts ...func() interface{}) (*Root, error) {
	var root Root
	err := r.getRole(roleRoot, &root)
	if err != nil {
		return nil, errors.Wrap(err, "getting local root role")
	}
	return &root, nil
}

func (r *localRepo) timestamp() (*Timestamp, error) {
	var ts Timestamp
	err := r.getRole(roleTimestamp, &ts)
	if err != nil {
		return nil, errors.Wrap(err, "getting local timestamp role")
	}
	return &ts, nil
}

func (r *localRepo) snapshot(opts ...func() interface{}) (*Snapshot, error) {
	var ss Snapshot
	err := r.getRole(roleSnapshot, &ss)
	if err != nil {
		return nil, errors.Wrap(err, "getting local snapshot role")
	}
	return &ss, nil
}

func (r *localRepo) targets(opts ...func() interface{}) (*Targets, error) {
	var ts Targets
	err := r.getRole(roleTargets, &ts)
	if err != nil {
		return nil, errors.Wrap(err, "getting local targets role")
	}
	return &ts, nil
}

// save persists role information, r is the role type, and
// data is the class containing the role information
func (r *localRepo) save(roleName role, data interface{}) error {
	isRoleCorrect(roleName, data)
	f, err := os.OpenFile(path.Join(r.repoPath, fmt.Sprintf("%s.json", roleName)), os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return errors.Wrap(err, "opening role file for writing")
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(data)
}

func (r *localRepo) getRole(name role, val interface{}) error {
	err := validateRole(name)
	if err != nil {
		return err
	}
	f, err := os.Open(path.Join(r.repoPath, fmt.Sprintf("%s.json", name)))
	if err != nil {
		return errors.Wrap(err, "getting role")
	}
	defer f.Close()
	return json.NewDecoder(f).Decode(val)
}