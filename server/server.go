package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/junqirao/gateway/component/registry"
	"github.com/junqirao/gateway/lib/response"
	"github.com/junqirao/gateway/model"
)

// List server info
func List(ctx context.Context) (infos []*model.ServerInfo, err error) {
	infos = make([]*model.ServerInfo, 0)

	statusMap, err := registry.Instance().Get(ctx, statusRegistryKey)
	if err != nil {
		return
	}

	cfgMap, err := registry.Instance().Get(ctx, configRegistryKey)
	if err != nil {
		return nil, err
	}

	for name, cfgStr := range cfgMap {
		sc := new(model.ServerConfig)
		if err = json.Unmarshal([]byte(cfgStr), &sc); err != nil {
			g.Log().Warningf(ctx, "unmarshal server config [%s](value=%s) failed: %v", name, cfgStr, err)
			continue
		}
		st := new(model.ServerStatus)
		if statusStr, ok := statusMap[name]; ok && statusStr != "" {
			if err = json.Unmarshal([]byte(statusStr), &st); err != nil {
				g.Log().Warningf(ctx, "unmarshal server status [%s](value=%s) failed: %v", name, statusStr, err)
				continue
			}
		}

		infos = append(infos, &model.ServerInfo{
			Name:   name,
			Config: sc,
			Status: st,
		})
	}

	return
}

// Get server info
func Get(ctx context.Context, name string) (rsp interface{}, err error) {
	cfg := new(model.ServerConfig)
	if err = getRegistryOne(ctx, fmt.Sprintf("%s%s", configRegistryKey, name), &cfg); err != nil {
		return
	}
	status := new(model.ServerStatus)
	if err = getRegistryOne(ctx, fmt.Sprintf("%s%s", statusRegistryKey, name), &status); err != nil {
		return
	}
	rsp = &model.ServerInfo{
		Name:   name,
		Config: cfg,
		Status: status,
	}
	return
}

func getRegistryOne(ctx context.Context, key string, ptr interface{}) (err error) {
	vs, err := registry.Instance().Get(ctx, key)
	if err != nil {
		return
	}

	has := false
	for n, s := range vs {
		if n != key {
			continue
		}
		has = true
		if ptr != nil {
			err = json.Unmarshal([]byte(s), &ptr)
		}
		break
	}
	if !has {
		err = response.ErrorResourceNotFound
	}
	return
}

// Delete ...
func Delete(ctx context.Context, name string) (rsp interface{}, err error) {
	if err = getRegistryOne(ctx, fmt.Sprintf("%s%s", configRegistryKey, name), nil); err != nil {
		return
	}

	keys := []string{fmt.Sprintf("%s%s", configRegistryKey, name), fmt.Sprintf("%s%s", statusRegistryKey, name)}
	for _, key := range keys {
		if err = registry.Instance().Delete(ctx, key); err != nil {
			return nil, err
		}
	}

	rsp = name
	return
}

// SetConfig ...
func SetConfig(ctx context.Context, name string, config *model.ServerConfig) (rsp interface{}, err error) {
	return name, registry.Instance().Set(ctx, fmt.Sprintf("%s%s", configRegistryKey, name), config)
}

// SetStatus ...
func SetStatus(ctx context.Context, name string, status *model.ServerStatus) (rsp interface{}, err error) {
	if err = getRegistryOne(ctx, fmt.Sprintf("%s%s", configRegistryKey, name), nil); err != nil {
		return
	}

	return name, registry.Instance().Set(ctx, fmt.Sprintf("%s%s", statusRegistryKey, name), status)
}
