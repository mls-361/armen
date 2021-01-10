/*
------------------------------------------------------------------------------------------------------------------------
####### metrics ####### (c) 2020-2021 mls-361 ###################################################### MIT License #######
------------------------------------------------------------------------------------------------------------------------
*/

package metrics

import (
	"encoding/json"
	"net/http"
	"sync"
)

type (
	leaf map[string]interface{}

	tree struct {
		data  leaf
		mutex sync.RWMutex
	}
)

func newTree() *tree {
	return &tree{
		data: make(leaf),
	}
}

func (t *tree) getLeaf(l leaf, keys []string) leaf {
	if len(keys) == 1 {
		return l
	}

	key := keys[0]

	v, ok := l[key]
	if ok {
		ml, ok := v.(leaf)
		if ok {
			return t.getLeaf(ml, keys[1:])
		}
	}

	nl := make(leaf)

	// Sans aucun scrupule, on remplace l'éventuelle valeur précédente !
	l[key] = nl

	return t.getLeaf(nl, keys[1:])
}

// AddInt AFAIRE.
func (t *tree) AddInt(value int64, keys ...string) {
	lk := len(keys)

	if lk == 0 {
		return
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	l := t.getLeaf(t.data, keys)

	key := keys[lk-1]

	v, ok := l[key]
	if ok {
		i, ok := v.(int64)
		if ok {
			l[key] = i + value
			return
		}
	}

	// Sans aucun scrupule, on remplace l'éventuelle valeur précédente !
	l[key] = value
}

// SetInt AFAIRE.
func (t *tree) SetInt(value int64, keys ...string) {
	lk := len(keys)

	if lk == 0 {
		return
	}

	t.mutex.Lock()
	defer t.mutex.Unlock()

	l := t.getLeaf(t.data, keys)

	// Sans aucun scrupule, on remplace l'éventuelle valeur précédente !
	l[keys[lk-1]] = value
}

func (t *tree) handler() http.HandlerFunc {
	return func(rw http.ResponseWriter, _ *http.Request) {
		rw.Header().Set("Content-Type", "application/json; charset=utf-8")
		t.mutex.Lock()
		_ = json.NewEncoder(rw).Encode(t.data)
		t.mutex.Unlock()
	}
}

/*
######################################################################################################## @(°_°)@ #######
*/
